package bot

import (
	"fmt"

	"github.com/krokomoko/tlgGoAPI/constructor"
	"github.com/krokomoko/tlgGoAPI/tlg"
)

/*
args:
1 - keyboard  *constructor.Keyboard
2 - ProtectContent bool
3 - DisableNotification bool
4 - DisableWebPagePreview bool
*/
func SendMessage(client *Client, text string, args ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - bot, SendText: %s", r)
		}
	}()

	tlgSendMessage := tlg.SendMessage{
		ChatId:    client.ID,
		Text:      text,
		ParseMode: "HTML",
	}

	for argInd, arg := range args {
		switch argInd {
		case 0:
			if keyboard, ok := arg.(*constructor.Keyboard); ok {
				keyboardMarkup, err := keyboard.Compile()
				if err != nil {
					err = fmt.Errorf("ERROR - SendKeyboard, compile keyboard: %s", err)
					return err
				}
				tlgSendMessage.ReplyMarkup = *keyboardMarkup
			}

		case 1:
			if protectContent, ok := arg.(bool); !ok {
				err = fmt.Errorf(
					"ERROR - SendText, ProtectContent parameter wrong data type: %T",
					args[1],
				)
				return
			} else {
				tlgSendMessage.ProtectContent = protectContent
			}

		case 2:
			if disableNotification, ok := arg.(bool); !ok {
				err = fmt.Errorf(
					"ERROR - SendText, DisableNotification parameter wrong data type: %T",
					args[1],
				)
				return
			} else {
				tlgSendMessage.DisableNotification = disableNotification
			}

		case 3:
			if disableWebPagePreview, ok := arg.(bool); !ok {
				err = fmt.Errorf(
					"ERROR - SendText, DisableWebPagePreview parameter wrong data type: %T",
					args[1],
				)
				return
			} else {
				tlgSendMessage.DisableWebPagePreview = disableWebPagePreview
			}
		}

	}

	_, err = Call(tlgSendMessage)

	return
}

/*
args:
1 - keyboard *constructor.Keyboard
*/
func EditMessageText(client *Client, messageId int64, text string, args ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - EditMessageText: %s", r)
		}
	}()

	tlgEditMessageText := tlg.EditMessageText{
		ChatId:    client.ID,
		MessageId: messageId,
		Text:      text,
	}

	if len(args) > 0 {
		if keyboard, ok := args[0].(*constructor.Keyboard); ok {
			keyboardMarkup, err := keyboard.Compile()
			if err != nil {
				err = fmt.Errorf(
					"ERROR - EditMessageText, compile keyboard: %s",
					err,
				)
				return err
			}
			tlgEditMessageText.ReplyMarkup = *keyboardMarkup
		}
	}

	_, err = Call(tlgEditMessageText)

	return
}

/*
args:
1 - keyboard *constructor.Keyboard
*/
func EditMessageCaption(client *Client, messageId int64, caption string, args ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - EditMessageCaption: %s", r)
		}
	}()

	tlgEditMessageCaption := tlg.EditMessageCaption{
		ChatId:    client.ID,
		MessageId: messageId,
		Caption:   caption,
	}

	if len(args) > 0 {
		if keyboard, ok := args[0].(*constructor.Keyboard); ok {
			keyboardMarkup, err := keyboard.Compile()
			if err != nil {
				err = fmt.Errorf(
					"ERROR - EditMessageCaption, compile keyboard: %s",
					err,
				)
				return err
			}
			tlgEditMessageCaption.ReplyMarkup = *keyboardMarkup
		}
	}

	_, err = Call(tlgEditMessageCaption)

	return
}

/*
add - additional parameters:
1. caption  string
2. keyboard
3. has_spoiler bool
4. disable_notifiaction bool
5. protect_content bool
*/
func SendPhotoByURL(client *Client, photoURL tlg.InputFile, args ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - SendPhotoByURL: %s", r)
		}
	}()

	tlgSendPhoto := tlg.SendPhoto{
		ChatId:    client.ID,
		Photo:     photoURL,
		ParseMode: "HTML",
	}

	for argInd, arg := range args {
		switch argInd {
		case 0:
			if caption, ok := arg.(string); ok && caption != "" {
				tlgSendPhoto.Caption = caption
			}
		case 1:
			if keyboard, ok := arg.(*constructor.Keyboard); ok {
				keyboardMarkup, err := keyboard.Compile()
				if err != nil {
					err = fmt.Errorf("ERROR - SendPhotoByURL, compile keyboard: %s", err)
					return err
				}
				tlgSendPhoto.ReplyMarkup = *keyboardMarkup
			}

		case 2:
			if hasSpoiler, ok := arg.(bool); ok {
				tlgSendPhoto.HasSpoiler = hasSpoiler
			}
		case 3:
			if disableNotification, ok := arg.(bool); ok {
				tlgSendPhoto.DisableNotification = disableNotification
			}
		case 4:
			if protectContent, ok := arg.(bool); ok {
				tlgSendPhoto.ProtectContent = protectContent
			}
		}
	}

	_, err = Call(tlgSendPhoto)

	return
}
