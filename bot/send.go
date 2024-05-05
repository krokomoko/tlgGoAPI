package bot

import (
	"fmt"

	"github.com/krokomoko/tlgGoAPI/constructor"
	"github.com/krokomoko/tlgGoAPI/tlg"
)

func SendKeyboard(client *Client, text string, keyboard *constructor.Keyboard) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - bot, SendKeyboard: %s", r)
		}
	}()

	keyboardMarkup, err := keyboard.Compile()
	if err != nil {
		err = fmt.Errorf("ERROR - SendKeyboard, compile keyboard: %s", err)
		return
	}

	_, err = Call(tlg.SendMessage{
		ChatId:      client.ID,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: *keyboardMarkup,
	})

	return
}
