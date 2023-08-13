<h1 align="center">newsletter telegram bot</h1>

# in dev stage rn | isn't deployed yet | isn't working yet

###

<img align="right" height="250" src="https://media.tenor.com/k_h_hLhzhW4AAAAd/news.gif"  />

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="30" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

### you need to add this bot to your channel or group to begin using it

## project structure:

```
.
├── cmd
│   └── main.go
├── core
│   └── core.go
├── Dockerfile
├── go.mod
├── README.md
└── tg
    └── tg.go
```

## installation

use git clone:

```
git clone https://github.com/kenjitheman/animun
```

## usage

- create .env file and inside you should create env variable with your api key
  ->

```
TELEGRAM_API_TOKEN=YOUR_TOKEN
```

- then you should uncomment commented lines in tg/tg.go ( ! you need uncomment
  commented lines only if you using this way !) ->

```
package tg

import (
	"fmt"
	"log"
	"os"

	"github.com/darenliang/jikan-go"
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	// "github.com/joho/godotenv"

	"main.go/api"
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

### you can also run it using docker ->

- you need to paste your api key in dockerfile ->

```
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
```

- then run it ->

```
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

- or you can run it with the following command:

```
cd cmd
go run main.go
```

## contributing

- pull requests are welcome. For major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
