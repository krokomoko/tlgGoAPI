package constructor

import (
	"fmt"

	"github.com/krokomoko/tlgGoAPI/tlg"
)

type Keyboard struct {
	rows [][]*_KeyboardButton
}

type _KeyboardButton struct {
	content      string
	callbackData string
	url          string
	webApp       *tlg.WebAppInfo
	pay          bool
}

func NewKeyboard(buttonRows ...[]string) (keyboard *Keyboard, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - NewKeyboard: %s", r)
		}
	}()

	keyboard = &Keyboard{
		rows: [][]*_KeyboardButton{},
	}

	var flag bool

	for _, buttonRow := range buttonRows {
		if len(buttonRow) == 0 {
			continue
		}

		buttons := []*_KeyboardButton{}
		for _, button := range buttonRow {
			if button != "" {
				flag = true
			}

			buttons = append(buttons, &_KeyboardButton{
				content: button,
			})
		}

		keyboard.rows = append(keyboard.rows, buttons)
	}

	if !flag {
		err = fmt.Errorf("ERROR - NewKeyboard, no buttons for new keyboard")
		return
	}

	return
}

func (k *Keyboard) SetUrl(i, j int, url string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - Keyboard, SetUrl: %s", r)
		}
	}()

	if len(k.rows) < i || len(k.rows[i]) < j {
		err = fmt.Errorf("ERROR - Keyboard, SetUrl: incorrect button indexes (x: %d | %y: %d)", i, j)

		return
	}

	k.rows[i][j].url = url

	return
}

func (k *Keyboard) SetWebApp(i, j int, url string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - Keyboard, SetWebApp: %s", r)
		}
	}()

	if len(k.rows) < i || len(k.rows[i]) < j {
		err = fmt.Errorf("ERROR - Keyboard, SetWebApp: incorrect button indexes (x: %d | %y: %d)", i, j)

		return
	}

	k.rows[i][j].webApp = &tlg.WebAppInfo{Url: url}

	return
}

func (k *Keyboard) SetPayButton(i, j int, pay bool) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - Keyboard, SetPayButton: %s", r)
		}
	}()

	if len(k.rows) < i || len(k.rows[i]) < j {
		err = fmt.Errorf("ERROR - Keyboard, SetPayButton: incorrect button indexes (x: %d | %y: %d)", i, j)

		return
	}

	k.rows[i][j].pay = pay

	return
}

func (k *Keyboard) SetCallbackData(i, j int, callbackData string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - Keyboard, SetCallbackData: %s", r)
		}
	}()

	if len(k.rows) < i || len(k.rows[i]) < j {
		err = fmt.Errorf("ERROR - Keyboard, SetCallbackData: incorrect button indexes (x: %d | %y: %d)", i, j)

		return
	}

	k.rows[i][j].callbackData = callbackData

	return
}

func (k *Keyboard) Compile() (markup *tlg.InlineKeyboardMarkup, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - Keyboard, Compile: %s", r)
		}
	}()

	markup = &tlg.InlineKeyboardMarkup{
		InlineKeyboard: [][]tlg.InlineKeyboardButton{},
	}

	var flag bool

	for _, row := range k.rows {
		if len(row) > 0 {
			flag = true
		} else {
			continue
		}

		buttonsRow := []tlg.InlineKeyboardButton{}
		for _, button := range row {
			ln := len(buttonsRow)
			buttonsRow = append(buttonsRow, tlg.InlineKeyboardButton{
				Text: button.content,
			})

			if button.callbackData != "" {
				buttonsRow[ln].CallbackData = button.callbackData
			} else if button.url != "" {
				buttonsRow[ln].Url = button.url
			} else if button.webApp != nil {
				buttonsRow[ln].WebApp = *button.webApp
			} else if button.pay {
				buttonsRow[ln].Pay = button.pay
			} else {
				continue
			}
		}

		markup.InlineKeyboard = append(markup.InlineKeyboard, buttonsRow)
	}

	if !flag {
		err = fmt.Errorf("ERROR - Keyboard Compile, no buttons for compile")
		return
	}

	return
}
