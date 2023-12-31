package tlgGoAPI

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	//"time"
)

type WaitGroupCount struct {
	wg    sync.WaitGroup
	count uint16
	limit uint16
}

type Bot struct {
	token            string
	url              string
	offset_to_update int64
	callbacks        map[string]func(bot *Bot, update *Update) error
	not_found        string
	stateMachine     *stateMachine
	//client           http.Client
}

const PROCESSING_COUNT_LIMIT = 1

func NewBot(token string) *Bot {
	return &Bot{
		token:        token,
		url:          fmt.Sprintf("https://api.telegram.org/bot%s/", token),
		callbacks:    make(map[string]func(bot *Bot, update *Update) error),
		not_found:    "",
		stateMachine: newStateMachine(),
		/*
			client: http.Client{
				Timeout: 500*time.Millisecond,
			},
		*/
	}
}

func (wg *WaitGroupCount) Add(delta int) bool {
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

func (wg *WaitGroupCount) Done() {
	wg.count -= 1
	wg.wg.Done()
}

func (wg *WaitGroupCount) Count() uint16 {
	return wg.count
}

func (wg *WaitGroupCount) Wait() {
	wg.wg.Wait()
}

func (bot *Bot) SetUserState(userId int64, state string) {
	bot.stateMachine.Set(userId, state)
}

func (bot *Bot) GetUserState(userId int64) string {
	state, ok := bot.stateMachine.Get(userId)
	if !ok {
		bot.stateMachine.Add(userId)
		return "start"
	}
	return state
}

func (bot *Bot) Call(fn_struct interface{}) (string, error) {
	if len(bot.token) == 0 {
		return "", errors.New("Token of bot is not exist")
	}

	var method_name, err = getFuncName(fn_struct)
	if err != nil {
		return "", err
	}

	var fn_struct_json, _ = json.Marshal(fn_struct)

	postBody := bytes.NewBuffer(fn_struct_json)

	resp, err := http.Post(bot.url+method_name, "application/json", postBody)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (bot *Bot) update() (*TelegramUpdate, error) {
	var gUpd = getUpdates{}
	if bot.offset_to_update != 0 {
		gUpd.Offset = bot.offset_to_update
	}

	result, err := bot.Call(gUpd)
	if err != nil {
		return nil, err
	}

	var update_struct TelegramUpdate

	err = json.Unmarshal([]byte(result), &update_struct)
	if err != nil {
		return nil, err
	}

	return &update_struct, nil
}

func (bot *Bot) processing(update *Update, wgc *WaitGroupCount, fn func(bot *Bot, update *Update) error) error {
	defer wgc.Done()

	return fn(bot, update)
}

func getHash(some_ints *[]uint8) string {
	var result string

	for _, int_ := range *some_ints {
		result += fmt.Sprintf("%d.", int_)
	}

	return result
}

func (bot *Bot) L(message_type *[]uint8, state string, fn func(bot *Bot, update *Update) error) error {
	var hash_key = getHash(message_type) + state

	bot.callbacks[hash_key] = fn

	// errors may be added later
	return nil
}

func (bot *Bot) Set404(text string) {
	bot.not_found = text
}

// customization will be added later
func (bot *Bot) send404(update *Update) {
	if len(bot.not_found) == 0 {
		return
	}
	
	_, err := bot.Call(SendMessage{
		ChatId: update.Message.Chat.Id,
		Text:   bot.not_found,
	})
	if err != nil {
		fmt.Println("ERROR: send404")
	}

	//fmt.Println("send404 output:", output)
}

func (bot *Bot) messageTypeCheck(update *Update, state string) (func(bot *Bot, update *Update) error, string) {
	var messageType = ""

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

	// M_CALLBACK
	if update.CallbackQuery.Id != ""  {
		messageType += fmt.Sprintf("%d.", M_CALLBACK)
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

	// M_TEXT
	if 0 == len(messageType) {
		messageType += fmt.Sprintf("%d.", M_TEXT)
	}

	messageType += state

	return bot.callbacks[messageType], messageType
}

func (bot *Bot) GetFromId(update *Update) int64 {
	var result int64

	if update.Message.MessageId != 0 {
		result = update.Message.Chat.Id
	} else if update.EditedMessage.MessageId != 0 {
		result = update.EditedMessage.Chat.Id
	} else if update.ChannelPost.MessageId != 0 {
		result = update.ChannelPost.Chat.Id
	} else if update.InlineQuery.Id != "" {
		result = update.InlineQuery.From.Id
	} else if update.ChosenInlineResult.ResultId != "" {
		result = update.ChosenInlineResult.From.Id
	} else if update.CallbackQuery.Id != "" {
		result = update.CallbackQuery.From.Id
	} else if update.ShippingQuery.Id != "" {
		result = update.ShippingQuery.From.Id
	} else if update.PreCheckoutQuery.Id != "" {
		result = update.PreCheckoutQuery.From.Id
	}

	return result
}

func (bot *Bot) Run() {
	var wgc = WaitGroupCount{
		limit: PROCESSING_COUNT_LIMIT,
	}

	for {
		updates, err := bot.update()
		if err != nil {
			fmt.Println(err)
			return
		}

		if !updates.Ok {
			fmt.Println("ERROR: while updating for new messages")
			return
		}

		if 0 == len(updates.Result) {
			continue
		}

		//fmt.Println("RESULTS:")
		bot.stateMachine.Load()
		for _, update := range updates.Result {
			bot.offset_to_update = update.UpdateId + 1
			fmt.Printf("--> Update id: %d \"%s\"\n", update.UpdateId, update.Message.Text)

			for !wgc.Add(1) {
			}

			// get current state for this user
			userID := bot.GetFromId(&update)
			if 0 == userID {
				fmt.Println("From Id is 0, continue...")
				wgc.Done()
				continue
			}
			state := bot.GetUserState(userID)

			// message type check
			// (need to make available to use multiple states for one callback)
			fn, message_type_hash := bot.messageTypeCheck(&update, state)
			fmt.Println("Message type:", message_type_hash)
			if fn == nil {
				fmt.Println("ERROR: have not callback to this message type:", message_type_hash)
				bot.send404(&update)
				wgc.Done()
				continue
			}

			bot.processing(&update, &wgc, fn)
			//go bot.processing(&update, wgc, fn)
		}
		bot.stateMachine.Save()
	}
}

/* EXAMPLE
token := ""

bot := NewBot(token)

testMTypes := []uint8{M_COMMAND}
bot.L(&testMTypes, "start", func(update *Update) error {
	fmt.Println("get testMTypes message only")
	return nil
})
bot.Run()
*/
