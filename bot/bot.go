package bot

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/kenjitheman/ecoman/api"
	"github.com/kenjitheman/ecoman/db"
	"github.com/kenjitheman/ecoman/openai"
)

func Start() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file (bot.go): %v", err)
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		fmt.Printf("Error starting bot: %v", err)
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	ticker := time.NewTicker(3 * time.Hour)

	go api.FetchAndSaveData()

	isBotRunning = false
	chatStates := make(map[int64]string)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		command := strings.TrimPrefix(update.Message.Text, "/")

		switch command {
		case "start", "/start":
			isBotRunning = true
			if isBotRunning {
				msg.Text = AlreadyRunningMsg
				msg.ReplyMarkup = GeneralKeyboard
			} else {
				msg.Text = StartMsg
				msg.ReplyMarkup = StartKeyboard
			}

		case "help", "/help":
			msg.Text = HelpMsg

		case "getdata", "/getdata", "get data":
			msg.Text = FetchingMsg
			cityNames, err := db.FetchAllCityNamesFromMongoDB()
			if err != nil {
				fmt.Printf("[ERROR] failed to fetch city names: %v\n", err)
				return
			}
			sort.Strings(cityNames)
			var keyboardRows [][]tgbotapi.KeyboardButton
			row := []tgbotapi.KeyboardButton{}
			for _, cityName := range cityNames {
				button := tgbotapi.NewKeyboardButton(cityName)
				row = append(row, button)
				if len(row) == 3 {
					keyboardRows = append(keyboardRows, row)
					row = []tgbotapi.KeyboardButton{}
				}
			}
			if len(row) > 0 {
				keyboardRows = append(keyboardRows, row)
			}
			citiesKeyboard := tgbotapi.NewReplyKeyboard(keyboardRows...)
			msg.ReplyMarkup = citiesKeyboard
			bot.Send(msg)

			chatStates[update.Message.Chat.ID] = "select_station"

			var selectedCity, selectedStation bool
			var selectedCityName, selectedStationName string
			var result string

			for update := range updates {
				if update.Message == nil {
					continue
				}
				if update.Message.Text != "" {
					if !selectedCity {
						selectedCityName = strings.TrimSpace(update.Message.Text)
						cityData, err := db.FetchDataFromMongoDB(selectedCityName)
						if err != nil {
							errorMessage := fmt.Sprintf(
								"[ERROR] failed to fetch city data for %s: %v",
								selectedCityName,
								err,
							)
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, errorMessage)
							bot.Send(msg)
							break
						}
						if len(cityData) == 0 {
							noDataMessage := fmt.Sprintf(
								NoDataMsg,
								selectedCityName,
							)
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, noDataMessage)
							bot.Send(msg)
							break
						}

						msg.Text = FetchingMsg

						var stationButtons [][]tgbotapi.KeyboardButton
						row := []tgbotapi.KeyboardButton{}
						for _, data := range cityData {
							button := tgbotapi.NewKeyboardButton(data.StationName)
							row = append(row, button)
							if len(row) == 2 {
								stationButtons = append(stationButtons, row)
								row = []tgbotapi.KeyboardButton{}
							}
						}
						if len(row) > 0 {
							stationButtons = append(stationButtons, row)
						}
						stationsKeyboard := tgbotapi.NewReplyKeyboard(stationButtons...)
						msg.ReplyMarkup = stationsKeyboard
						bot.Send(msg)
						selectedCity = true
					} else if !selectedStation {
						selectedStationName = strings.TrimSpace(update.Message.Text)
						result = Datafetch(selectedCityName, selectedStationName)
						responseMessage := tgbotapi.NewMessage(update.Message.Chat.ID, result)
						bot.Send(responseMessage)
						selectedStation = true
					}
				}
				if selectedCity && selectedStation {
					msg.Text = AdviceMsg
					msg.ReplyMarkup = YesOrNoKeyboard
					bot.Send(msg)

					for {
						response := <-updates

						if response.Message.Text == "yes" {
							msg.Text = YesMsg
							advice := openai.GenerateAdvice(result)
							msg.Text = advice
						} else {
							msg.Text = NoMsg
						}
						break
					}

					msg.ReplyMarkup = GeneralKeyboard
					bot.Send(msg)
					break
				}
			}

		case "status", "/status":
			if err == nil {
				msg.Text = StatusOkMsg
			} else {
				msg.Text = StatusNotOkMsg
			}

		case "info", "/info":
			msg.Text = InfoMsg

		case "bug_report", "/bug_report", "bug report":
			chatStates[update.Message.Chat.ID] = "support"
			msg.Text = "please describe the problem:"
			bot.Send(msg)

			for {
				response := <-updates

				if response.Message == nil {
					continue
				}

				if response.Message.Chat.ID != update.Message.Chat.ID {
					continue
				}

				description := response.Message.Text
				msg.Text = ThxMsg
				bot.Send(msg)

				supportMsg := tgbotapi.NewMessage(
					creatorChatID,
					fmt.Sprintf(
						BugReportMsg,
						update.Message.From.UserName,
						description,
					),
				)
				bot.Send(supportMsg)

				delete(chatStates, update.Message.Chat.ID)
				break
			}

			continue

		case "stop", "/stop":
			isBotRunning = false
			msg.Text = StopMsg
			msg.ReplyMarkup = StartKeyboard

		default:
			msg.Text = IdkMsg
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		select {
		case <-ticker.C:
			go api.FetchAndSaveData()
		default:
		}
	}
}
