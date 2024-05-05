package tlgGoAPI

import (
	"fmt"
	"log"

	"github.com/krokomoko/tlgGoAPI/tlg"
)

func GetCallbackQueryData(update *tlg.Update) (data string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - GetCallbackQueryData: %s", r)
			log.Println(err)
		}
	}()

	if update.CallbackQuery.Data == "" {
		err = fmt.Errorf("ERROR - GetCallbackQueryData, update.CallbackQuery.Data is empty")
		return
	}

	data = update.CallbackQuery.Data

	return
}
