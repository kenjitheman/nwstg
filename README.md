# newsman

### Newsletter telegram bot in golang

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

## Project structure:

```go
newsman
│
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

## Installation

```sh
git clone https://github.com/kenjitheman/newsman
```

## Usage

- Create .env file and inside you should create env variable with your api key:

```.env
TELEGRAM_API_TOKEN=YOUR_TOKEN
```

- Then you should uncomment commented lines in tg/tg.go \
	- **( ! You need uncomment commented lines only if you using this way !)**

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

- You need to add usernames(optional) and chatIDs to users.json file like this:

```json
{
  "users": {
    "kenjitheman": 5785150199,
    "username": chatID,
    "optional_not_real_username": 3942049232
  }
}
```

- Run it:

```sh
go run main.go
```

- Or build and run:

```sh
go build
```

```sh
./newsman
```

### You can also run it using Docker

```dockerfile
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
```

```sh
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## Contributing

- Pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change.

## License

- [MIT](https://choosealicense.com/licenses/mit/)
