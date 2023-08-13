package tg

import (
	"fmt"
	"log"
	"os"

	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	isBotRunning  bool
	creatorChatID int64
)

func Start() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("[ERROR] error loading .env file")
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	isBotRunning = false

	generalKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/help"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/send_newsletter"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/support"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/stop"),
		),
	)

	startKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/start"),
		),
	)

	log.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			isBotRunning = true
			okEmoji := emoji.Sprintf("%v", emoji.GreenCircle)
			if isBotRunning {
				msg.Text = okEmoji + " newsletterman is already running"
			}
			msg.Text = okEmoji + " newsletterman has been started"
			msg.ReplyMarkup = generalKeyboard
		case "help":
			if isBotRunning {
				infoEmoji := emoji.Sprintf("%v", emoji.Information)
				msg.Text = infoEmoji + " newsletterman hints\n\n /help - to get all commands\n /start - to start newsletterman\n /stop - to stop newsletterman\n /send_newsletter - to send newsletter\n /support - to tell about bugs you found"
				msg.ReplyMarkup = generalKeyboard
			}
		case "send_newsletter":
			if isBotRunning {
        msg.Text = "send_newsletter" // TODO: implement send_newsletter
			}
		case "stop":
			if isBotRunning {
				stopEmoji := emoji.Sprintf("%v", emoji.RedCircle)
				msg.Text = stopEmoji + " newsletterman has been stopped"
				msg.ReplyMarkup = startKeyboard
				isBotRunning = false
			}
			isBotRunning = false
		case "support":
			if isBotRunning {
				cactusEmoji := emoji.Sprintf("%v", emoji.Cactus)
				creatorChatID = 5785150199
				msg.Text = cactusEmoji + " please describe the problem:"
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
					break
				}
				continue
			}
			isBotRunning = false

		default:
			if isBotRunning {
				idkEmoji := emoji.Sprintf("%v", emoji.OpenHands)
				msg.Text = idkEmoji + " i don't understand you\n/help"
			}
		}
		if _, err := bot.Send(msg); err != nil {
			fmt.Printf("[ERROR] error sending message")
		}
	}
}
