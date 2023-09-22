## newsletter tg bot in golang

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="15" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

## project structure:

```go
.
├── cmd
│   └── main.go
├── Dockerfile
├── go.mod
├── go.sum
├── README.md
├── tg
│   └── tg.go
└── users.json
```

## installation

```shell
git clone https://github.com/kenjitheman/newsletter_tgbot
```

## usage

- create .env file and inside you should create env variable with your api key:

```.env
TELEGRAM_API_TOKEN=YOUR_TOKEN
```

- then you should uncomment commented lines in tg/tg.go \
	- **( ! you need uncomment commented lines only if you using this way !)**

```go
package tg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	isBotRunning      bool
	creatorChatID     int64
	newsletterContent string
)

func Start() {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	fmt.Println("[ERROR] error loading .env file")
	// 	log.Panic(err)
	// }
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
```

- you need to add usernames(optional) and chatIDs to users.json file like this:

```json
{
  "users": {
    "kenjitheman": 5785150199,
    "username": chatID,
    "optional_not_real_username": 3942049232
  }
}
```

### you can also run it using docker:

- you need to paste your api key in dockerfile ->

```dockerfile
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
```

- then run it:

```shell
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

- or you can run it with the following command:

```shell
cd cmd
go run main.go
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
