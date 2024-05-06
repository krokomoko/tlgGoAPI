# tlgGoAPI

Telegram BOT API with Golang

## How to use

1. Set token
2. Set possible states of client
3. Set states transition by clients signal
4. Set functions for detect signal by message types
5. Run the bot

#### Set token

	import "github.com/krokomoko/tlgGoAPI/bot"
	//...
	bot.SetToken("some token")

#### Set possible states of client

	bot.SetStates("state 1", "state 2", "state 3" [, ...other])

#### Set states transition by clients signal

	bot.SetTransition("init", "state 1", "some signal", callback)

Callbacks function type be like:

	import "github.com/krokomoko/tlgGoAPI/tlg"
	//...
	func(client *bot.Client, update *tlg.Update) error

#### Set functions for detect signal by message

	bot.SetMessageType(detectFunc, M_TYPE_1, M_TYPE_2, M_TYPE_3 [, ...other])

Message types:

* M_COMMAND
* M_TEXT
* M_IMAGE
* M_VIDEO
* M_VOICE
* M_AUDIO
* M_GIF
* M_DOCUMENT
* M_STICKER
* M_URL

M_TEXT - message type without any other message types

Detect function type be like:

	func(client *bot.Client, update *tlg.Update) (string, error)

#### Run the bot

	bot.Run()