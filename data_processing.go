package tlgGoAPI

import (
	"fmt"

	"github.com/krokomoko/tlgGoAPI/tlg"
)

func GetCallbackQueryData(update *tlg.Update) (data string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - GetCallbackQueryData: %s", r)
		}
	}()

	if update.CallbackQuery.Data == "" {
		err = fmt.Errorf("ERROR - GetCallbackQueryData, update.CallbackQuery.Data is empty")
		return
	}

	data = update.CallbackQuery.Data

	return
}

func GetMessageId(update *tlg.Update) (messageId int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - GetMessageId: %s", r)
		}
	}()

	if update.Message.MessageId != 0 {
		messageId = update.Message.MessageId
	} else if update.EditedMessage.MessageId != 0 {
		messageId = update.Message.MessageId
	} else if update.ChannelPost.MessageId != 0 {
		messageId = update.ChannelPost.MessageId
	} else if update.CallbackQuery.Message.MessageId != 0 {
		messageId = update.CallbackQuery.Message.MessageId
	} else {
		err = fmt.Errorf("ERROR - GetMessageId, not message id")
	}

	return
}
