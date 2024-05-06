package tlgGoAPI

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/krokomoko/tlgGoAPI/bot"
	"github.com/krokomoko/tlgGoAPI/tlg"
)

type _WaitGroupCount struct {
	wg    sync.WaitGroup
	count uint16
	limit uint16
}

func (wg *_WaitGroupCount) Add(delta int) bool {
	if wg.limit == 0 {
		return false
	}

	if wg.count >= wg.limit {
		return false
	}

	wg.count += 1
	wg.wg.Add(delta)
	return true
}

func (wg *_WaitGroupCount) Done() {
	wg.count -= 1
	wg.wg.Done()
}

func (wg *_WaitGroupCount) Count() uint16 {
	return wg.count
}

func (wg *_WaitGroupCount) Wait() {
	wg.wg.Wait()
}

func getUpdates() (tlgUpdate *tlg.TelegramUpdate, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - getUpdates: %s", r)
		}
	}()

	var gUpd = tlg.GetUpdates{}

	if _OFFSET_TO_UPDATE != 0 {
		gUpd.Offset = _OFFSET_TO_UPDATE
	}

	result, err := bot.Call(gUpd)
	if err != nil {
		return nil, err
	}

	var update_struct tlg.TelegramUpdate

	err = json.Unmarshal([]byte(result), &update_struct)
	if err != nil {
		return nil, err
	}

	return &update_struct, nil
}

func getClient(update *tlg.Update) (client *bot.Client, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - getClient: %s", r)
		}
	}()

	var clientId int64

	if update.Message.MessageId != 0 {
		clientId = update.Message.Chat.Id
	} else if update.EditedMessage.MessageId != 0 {
		clientId = update.EditedMessage.Chat.Id
	} else if update.ChannelPost.MessageId != 0 {
		clientId = update.ChannelPost.Chat.Id
	} else if update.InlineQuery.Id != "" {
		clientId = update.InlineQuery.From.Id
	} else if update.ChosenInlineResult.ResultId != "" {
		clientId = update.ChosenInlineResult.From.Id
	} else if update.CallbackQuery.Id != "" {
		clientId = update.CallbackQuery.From.Id
	} else if update.ShippingQuery.Id != "" {
		clientId = update.ShippingQuery.From.Id
	} else if update.PreCheckoutQuery.Id != "" {
		clientId = update.PreCheckoutQuery.From.Id
	} else {
		err = fmt.Errorf("ERROR - getClient, no client id in update")
		return
	}

	client, err = bot.GetClient(clientId)

	return
}

func getMessageType(update *tlg.Update) (messageType string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - getMessageType: %s", r)
		}
	}()

	//M_COMMAND
	if update.Message.Entities != nil &&
		len(update.Message.Entities) > 0 &&
		func() bool {
			for _, entitie := range update.Message.Entities {
				if entitie.Type == "bot_command" {
					return true
				}
			}
			return false
		}() {
		messageType += fmt.Sprintf("%d.", M_COMMAND)
	}

	// M_IMAGE
	if update.Message.Photo != nil {
		messageType += fmt.Sprintf("%d.", M_IMAGE)
	}

	// M_VIDEO
	if len(update.Message.Video.FileId) > 0 {
		messageType += fmt.Sprintf("%d.", M_VIDEO)
	}

	// M_VOICE
	if len(update.Message.Voice.FileId) > 0 {
		messageType += fmt.Sprintf("%d.", M_VOICE)
	}

	// M_AUDIO
	if len(update.Message.Audio.FileId) > 0 {
		messageType += fmt.Sprintf("%d.", M_AUDIO)
	}

	// M_GIF
	if len(update.Message.Animation.FileId) > 0 {
		messageType += fmt.Sprintf("%d.", M_GIF)
	}

	// M_DOCUMENT
	if len(update.Message.Document.FileId) > 0 {
		messageType += fmt.Sprintf("%d.", M_DOCUMENT)
	}

	// M_STICKER
	if len(update.Message.Sticker.FileId) > 0 {
		messageType += fmt.Sprintf("%d.", M_STICKER)
	}

	// M_URL
	if update.Message.Entities != nil &&
		len(update.Message.Entities) > 0 &&
		func() bool {
			for _, entitie := range update.Message.Entities {
				if entitie.Type == "url" {
					return true
				}
			}
			return false
		}() {
		messageType += fmt.Sprintf("%d.", M_URL)
	}

	// M_CALLBACK
	if update.CallbackQuery.Id != "" {
		messageType += fmt.Sprintf("%d.", M_CALLBACK)
	}

	// M_TEXT
	if 0 == len(messageType) {
		messageType += fmt.Sprintf("%d.", M_TEXT)
	}

	return
}

func getSignal(client *bot.Client, update *tlg.Update, messageType string) (signalName string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - bot, getSignal: %s", r)
		}
	}()

	callback, err := bot.GetMessageTypeCallback(messageType)

	signalName, err = callback(client, update)

	return
}
