package bot

import (
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

const (
	creatorChatID int64 = 5785150199
)

var (
	isBotRunning        bool
	newsletterContent   string
	chatIDsByUser       = make(map[string]int64)
	lastUserMessageTime time.Time
	autoOff             *time.Timer
	supportMsg          tgbotapi.MessageConfig
)

var (
	infoEmoji       = emoji.Sprintf("%v", emoji.Information)
	infinityEmoji   = emoji.Sprintf("%v", emoji.Infinity)
	stopEmoji       = emoji.Sprintf("%v", emoji.RedCircle)
	okEmoji         = emoji.Sprintf("%v", emoji.GreenCircle)
	cactusEmoji     = emoji.Sprintf("%v", emoji.Cactus)
	idkEmoji        = emoji.Sprintf("%v", emoji.OpenHands)
	GreenHeartEmoji = emoji.Sprintf("%v", emoji.GreenHeart)
)

var (
	alreadyRunningMsg               = okEmoji + " newsman is already running\n[ ? ] /stop - to stop newsman"
	startMsg                        = okEmoji + " newsman is started\n[ ? ] /help - to get all commands"
	stopMsg                         = stopEmoji + " newsman is stopped\n[ ? ] /start - to start newsman"
	helpMsg                         = infoEmoji + " newsman hints\n\n+ /help - to get all commands\n+ /start - to start newsman\n+ /stop - to stop newsman\n+ /send_newsletter - to send newsletter\n+ /bug_report - to tell about bugs you found"
	thxForBugReportMsg              = GreenHeartEmoji + " thanks for your bug report!"
	describeTheProblemMsg           = cactusEmoji + " please describe the problem:"
	idkMsg                          = idkEmoji + " i don't know what you mean\n[ ? ] /help - to get all commands"
	bugReportMsg                    = "[ ! ] bug report from user @%s\n[ ! ] msg: %s"
	alreadyStoppedMsg               = stopEmoji + " newsman is already stopped\n[ ? ] /start - to start newsman"
	backToMenuMsg                   = okEmoji + " back to menu"
	autoOffMsg                      = stopEmoji + " newsman is stopped due to inactivity\n[ ? ] /start - to start newsman"
	sentNewsletterMsg               = okEmoji + " newsletter content has been sent to all users successfully"
	pleaseEnterNewsletterContentMsg = "please enter the newsletter content:"
)
