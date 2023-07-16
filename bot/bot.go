package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	api "main.go/api"
)

func Datafetch(cityname string) string {
	cityData, err := api.GetCityDataFromMongoDB(cityname)
	if err != nil {
		fmt.Printf("[ERROR] error fetching city data from MongoDB: %v", err)
		return "Error fetching city data. Please try again later."
	} else {
		result := fmt.Sprintf("%v City: %s\n%v Street: %s\n%v Latitude: %s\n%v Longitude: %s\n%v Timezone: %s\n",
			emoji.Cityscape, cityData.CityName,
			emoji.House, cityData.StationName,
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

	// bot, err := tgbotapi.NewBotAPI("6026087255:AAHNLUWdeFRaiqhNlETTfLtL6ia1YVQsFQs")
	// 	if err != nil {
	// 		fmt.Printf("[ERROR] error starting bot: %v\n", err)
	// 		log.Panic(err)
	// 	}

	bot.Debug = true

	log.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Create a ticker that ticks every 3 hours
	ticker := time.NewTicker(3 * time.Hour)

	// Run FetchAndSaveData initially
	go api.FetchAndSaveData()

	for {
		select {
		case update := <-updates:
			if update.Message == nil {
				continue
			}

			if !update.Message.IsCommand() {
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				leafGreenEmoji := emoji.Sprintf("%v", emoji.LeafyGreen)
				msg.Text = leafGreenEmoji + " Wassup, let's start, run /help for usage info !"
			case "help":
				helpEmoji := emoji.Sprintf("%v", emoji.Information)
				msg.Text = helpEmoji + " All available commands:\n\t" + "  + /help - use if you need some help\n\t" + "  + /getdata - use to get all ecoman data, specify the cityn\n\t" + "  + /status - use to see working status\n\t" + "  + /info - use to see more info about creator and bot\n\t" + "  + /support - use if you found a bug etc.\n\t" + "  + /support_creator - it's open source free to use product,\n     so i don't get any money from it\n\t" + "  + /stop - use stop command to stop the bot"
			case "getdata":
				getdataEmoji := emoji.Sprintf("%v", emoji.GreenCircle)
				fetchingMessage := getdataEmoji + " data fetching..."
				msg.Text = fetchingMessage
				if len(update.Message.CommandArguments()) > 0 {
					cityname := update.Message.CommandArguments()
					go func() {
						time.Sleep(1 * time.Second) // simulating data fetching delay
						dataResult := Datafetch(cityname)
						msg.Text = dataResult
						if _, err := bot.Send(msg); err != nil {
							log.Panic(err)
						}
					}()
				} else {
					getdataEmoji := emoji.Sprintf("%v", emoji.Cityscape)
					msg.Text = getdataEmoji + " Please specify a city name after the /getdata command. \t Example: /getdata Kyiv"
				}
			case "status":
				if err == nil {
					statusEmoji := emoji.Sprintf("%v", emoji.GreenCircle)
					msg.Text = statusEmoji + " Ecoman is ok. Working fine ^_^"
				} else {
					statusEmoji := emoji.Sprintf("%v", emoji.RedCircle)
					msg.Text = statusEmoji + " Ecoman is not ok. Something isn't fine -_- \n" + "Try again later -_-"
				}
			case "info":
				infoEmoji := emoji.Sprintf("%v", emoji.Information)
				msg.Text = infoEmoji + " Hey, I'm amodotomi, the creator of ecoman. Ecoman is a Telegram bot that allows you to get the latest information about ecology in Ukraine. The information updates every 15 minutes. Enjoy! ^_^"
			case "stop":
				stopEmoji := emoji.Sprintf("%v", emoji.StopSign)
				msg.Text = stopEmoji + " Ecoman has been stopped"
			case "support":
				CactusEmoji := emoji.Sprintf("%v", emoji.Cactus)
				msg.Text = CactusEmoji + " You've found some errors? Describe the problem, please. Example: /support error_description"
				if len(update.Message.CommandArguments()) > 0 {
					errorDescription := update.Message.CommandArguments()
					if errorDescription == "" {
						msg.Text = "Please provide an error description after the /support command. Example: /support error_description"
					} else {
						// Send the error_description to @amodotomi
						chatID := int64(5785150199) // Replace with the correct chat ID of @amodotomi
						_, err := bot.Send(tgbotapi.NewMessage(chatID, errorDescription))
						if err != nil {
							log.Printf("[ERROR] failed to send error_description to @amodotomi: %v", err)
						}
						// Send a response indicating that the error description has been sent
						msg.Text = "Thank you for reporting the error. The description has been sent to @amodotomi."
					}
				} else {
					// Handle when the command is used without error_description
					msg.Text = "Please provide an error description after the /support command. Example: /support error_description"
				}
				_, err := bot.Send(msg)
				if err != nil {
					log.Panic(err)
				}
			case "support_creator":
				stopEmoji := emoji.Sprintf("%v", emoji.GreenHeart)
				msg.Text = stopEmoji + " My website: amodotomi.com\n" + stopEmoji + " My GitHub: github.com/amodotomi\n" + stopEmoji + " Thanks for your support!"

			default:
				defaultEmoji := emoji.Sprintf("%v", emoji.OkHand)
				msg.Text = defaultEmoji + " Sorry, I don't know that command. Use /help for a list of all commands or help ^_^"
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case <-ticker.C:
			// Run FetchAndSaveData every 3 hours
			go api.FetchAndSaveData()
		}
	}
}
