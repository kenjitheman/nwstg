package bot

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	isBotRunning := false

	log.Printf("[SUCCESS] authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	lastUserMessageTime := time.Now()

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			userInput := update.Message.Text

			if autoOff != nil {
				autoOff.Stop()
			}

			if time.Since(lastUserMessageTime) > 5*time.Minute {
				if isBotRunning {
					isBotRunning = false
					msg.Text = autoOffMsg
					msg.ReplyMarkup = StartKeyboard
					bot.Send(msg)
				}
			}

			lastUserMessageTime = time.Now()

			autoOff := time.NewTimer(5 * time.Minute)
			go func() {
				<-autoOff.C
				if isBotRunning {
					isBotRunning = false
					msg.Text = autoOffMsg
					msg.ReplyMarkup = StartKeyboard
					bot.Send(msg)
				}
			}()

			switch userInput {
			case "/start", "start":
				if !isBotRunning {
					isBotRunning = true
					msg.Text = startMsg
					msg.ReplyMarkup = GeneralKeyboard
				} else {
					msg.Text = alreadyRunningMsg
				}

			case "/help", "help":
				if isBotRunning {
					msg.Text = helpMsg
					msg.ReplyMarkup = GeneralKeyboard
				}

			case "send_newsletter", "/send_newsletter", "send newsletter":
				if isBotRunning {
					msg.Text = pleaseEnterNewsletterContentMsg
					bot.Send(msg)

					response := <-updates
					if response.Message != nil {
						newsletterContent = response.Message.Text
						msg.Text = sentNewsletterMsg
						bot.Send(msg)

						usersJSON, err := os.ReadFile("../users.json")
						if err != nil {
							log.Printf("[ERROR] error reading users.json: %v", err)
							continue
						}

						var usersData struct {
							Users map[string]int64 `json:"users"`
						}

						if err := json.Unmarshal(usersJSON, &usersData); err != nil {
							log.Printf("[ERROR] error parsing users.json: %v", err)
							continue
						}

						for _, chatID := range usersData.Users {
							individualMsg := tgbotapi.NewMessage(
								chatID,
								newsletterContent,
							)
							bot.Send(individualMsg)
						}
					}
				}

			case "/stop", "stop":
				if isBotRunning {
					isBotRunning = false
					msg.Text = stopMsg
					msg.ReplyMarkup = StartKeyboard
				} else {
					msg.Text = alreadyStoppedMsg
				}

			case "/bug_report", "bug_report", "bug report":
				if isBotRunning {
					msg.ReplyMarkup = BackKeyboard
					msg.Text = describeTheProblemMsg
					bot.Send(msg)

					response := <-updates

					if response.Message != nil {
						if response.Message.Chat.ID != update.Message.Chat.ID {
							continue
						}
						description := response.Message.Text

						if description == "back to menu" {
							msg.Text = backToMenuMsg
							msg.ReplyMarkup = GeneralKeyboard
						} else {
							supportMsg = tgbotapi.NewMessage(
								creatorChatID,
								fmt.Sprintf(
									bugReportMsg,
									update.Message.From.UserName,
									description,
								),
							)
							msg.Text = thxForBugReportMsg
						}
						bot.Send(supportMsg)
						msg.ReplyMarkup = GeneralKeyboard
					}

				} else {
					msg.Text = alreadyStoppedMsg
				}
			default:
				if isBotRunning {
					msg.Text = idkMsg
				}
			}
			lastUserMessageTime = time.Now()
			if _, err := bot.Send(msg); err != nil {
				fmt.Println("[ERROR] error sending message")
			}
		}
	}

	if autoOff != nil {
		autoOff.Stop()
	}
}
