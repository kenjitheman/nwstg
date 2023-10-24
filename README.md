## newsletter tg bot in golang

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

## project structure:

```go
.
├── bot
│   ├── bot.go
│   ├── keyboards.go
│   └── vars.go
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
└── users.json
```

## installation

```shell
git clone https://github.com/kenjitheman/newsman
```

## usage

- create .env file and inside you should create env variable with your api key:

```.env
TELEGRAM_API_TOKEN=YOUR_TOKEN
```

- then you should uncomment commented lines in tg/tg.go \
	- **( ! you need uncomment commented lines only if you using this way !)**

```go
//"github.com/joho/godotenv"
```

```go
// err := godotenv.Load("../.env")
// if err != nil {
// 	fmt.Println("[ERROR] error loading .env file")
// 	log.Panic(err)
// }
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

- run it

```shell
go run main.go
```

- or build and run

```shell
go build
```

```shell
./newsman
```

### you can also run it using docker

```dockerfile
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
```

```shell
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```


## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
