# tlgGoAPI

Simple Golang API for use Telegram Bot API

## How to use

### Creating bot instance

First create new Bot instance:

	token := "" // your telegram bot api token
    bot := tlgGoAPI.NewBot(token)
	
### Setting up a callback
	
Next add your callbacks for special input messages:

    messageTypes := []uint8{tlgGoAPI.M_COMMAND, tlgGoAPI.M_IMAGE}
    bot.L(&messageTypes, "start", func(update *tlgGoAPI.Update) error {
		_, err := bot.Call(tlgGoAPI.SendMessage {
			ChatId: update.Message.Chat.Id,
			Text: "hello, my friend!"
		})
		
		if err != nil {
			return err
		}
		
		return nil
	})

M_COMMAND and M_IMAGE - types of input messagges like:
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

"start" is the state the user must be in for the message to be processed.

### Change user state

To change user state use **SetUserState** method of bot instance:

	userId := update.Message.From.Id
	state := "some new state"
	bot.SetUserState(userId, state)

## Other

The api contains all the same data types and methods that are presented on the official [telegram bot API page](https://core.telegram.org/bots/api).


