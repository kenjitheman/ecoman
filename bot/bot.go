package bot

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	api "main.go/api"
	"main.go/db"
)

var isBotRunning bool

func Datafetch(cityname, stationName string) string {
	cityData, err := api.GetCityDataFromMongoDB(cityname, stationName)
	if err != nil {
		fmt.Printf("[ERROR] error fetching city data from MongoDB: %v", err)
		return "error fetching city data! \ncity is incorrect!"
	} else {
		result := fmt.Sprintf("%v City: %s\n", emoji.Cityscape, cityData.CityName)
		result += fmt.Sprintf("%v Station: %s\n", emoji.House, cityData.StationName)
		result += fmt.Sprintf("%v Latitude: %s\n%v Longitude: %s\n%v Timezone: %s\n",
			emoji.Compass, cityData.Latitude,
			emoji.Compass, cityData.Longitude,
			emoji.ThreeOClock, cityData.Timezone)
		for _, pollutant := range cityData.Pollutants {
			pollutantInfo := fmt.Sprintf("%v Pollutant: %s\n   + Units: %s\n   + Time: %v\n   + Value: %v\n   + Average: %v\n",
				emoji.GemStone, pollutant.Pol, pollutant.Unit, pollutant.Time, pollutant.Value, pollutant.Averaging)
			result += pollutantInfo
		}
		return result
	}
}

func StartBot() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		fmt.Printf("[ERROR] error starting bot: %v", err)
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	ticker := time.NewTicker(3 * time.Hour)

	go api.FetchAndSaveData()

	generalKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("help"),
			tgbotapi.NewKeyboardButton("getdata"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("info"),
			tgbotapi.NewKeyboardButton("status"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("support"),
			tgbotapi.NewKeyboardButton("support_creator"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("stop"),
		),
	)

	startKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("start"),
		),
	)

	isBotRunning = false

	chatStates := make(map[int64]string)
	creatorChatID := int64(5785150199)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		command := strings.TrimPrefix(update.Message.Text, "/")

		switch command {
		case "start":
			isBotRunning = true
			if isBotRunning {
				StartOkEmoji := emoji.Sprintf("%v", emoji.GreenCircle)
				msg.Text = StartOkEmoji + " ecoman is already running\nuse /stop to stop the bot\nuse /help for more information"
				msg.ReplyMarkup = generalKeyboard
			} else {
				leafGreenEmoji := emoji.Sprintf("%v", emoji.LeafyGreen)
				msg.Text = leafGreenEmoji + " hey, let's start\nrun /help for usage info!"
				msg.ReplyMarkup = startKeyboard
			}

		case "help":
			helpEmoji := emoji.Sprintf("%v", emoji.Information)
			msg.Text = helpEmoji + " all available commands:\n\n/help - use if you need some help\n\n/getdata - use to get all ecoman data, specify the city name\n\n/status - use to see working status\n\n/info - use to see more info about creator and bot\n\n/support - use if you found a bug etc.\n\n/support_creator - It's an open-source free-to-use product, so I don't get any money from it\n\n/stop - use stop command to stop the bot"

		case "getdata":
			getdataEmoji := emoji.Sprintf("%v", emoji.Cityscape)
			fetchingMessage := getdataEmoji + " please select a city from the list:"
			msg.Text = fetchingMessage
			cityNames, err := db.FetchAllCityNamesFromMongoDB()
			if err != nil {
				fmt.Printf("[ERROR] failed to fetch city names: %v\n", err)
				return
			}
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
			citiesKeyboard.OneTimeKeyboard = true
			msg.ReplyMarkup = citiesKeyboard



		case "status":
			if err == nil {
				statusEmoji := emoji.Sprintf("%v", emoji.GreenCircle)
				msg.Text = statusEmoji + " ecoman is ok, working fine! ^_^"
			} else {
				statusEmoji := emoji.Sprintf("%v", emoji.RedCircle)
				msg.Text = statusEmoji + " ecoman is not ok, something isn't fine -_-\ntry again later -_-"
			}

		case "info":
			infoEmoji := emoji.Sprintf("%v", emoji.Information)
			msg.Text = infoEmoji + " hey, I'm amodotomi, the creator of ecoman.\n\necoman is a telegram bot that allows you to get the latest information about ecology in Ukraine.\ndata updates every 15 minutes.\n\nenjoy! ^_^"

		case "support":
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
				GreenHeartEmoji := emoji.Sprintf("%v", emoji.GreenHeart)
				msg.Text = GreenHeartEmoji + " thanks for your bug report!"
				bot.Send(msg)

				supportMsg := tgbotapi.NewMessage(
					creatorChatID,
					fmt.Sprintf(
						" bug report from user %s:\n%s",
						update.Message.From.UserName,
						description,
					),
				)
				bot.Send(supportMsg)

				delete(chatStates, update.Message.Chat.ID)
				break
			}

			continue

		case "support_creator":
			GreenHeartEmoji := emoji.Sprintf("%v", emoji.GreenHeart)
			msg.Text = GreenHeartEmoji + " my website: amodotomi.com\n" + GreenHeartEmoji + " my GitHub: github.com/amodotomi\n" + GreenHeartEmoji + " thanks for your support!"
		case "stop":
			isBotRunning = false
			stopEmoji := emoji.Sprintf("%v", emoji.StopSign)
			msg.Text = stopEmoji + " ecoman has been stopped"
			msg.ReplyMarkup = startKeyboard

		default:
			defaultEmoji := emoji.Sprintf("%v", emoji.OkHand)
			msg.Text = defaultEmoji + " sorry, i don't know that command\nuse /help for a list of all commands or help ^_^"
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
