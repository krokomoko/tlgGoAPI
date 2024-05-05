package tlgGoAPI

import (
	"fmt"
	"log"

	"github.com/krokomoko/tlgGoAPI/bot"
	"github.com/krokomoko/tlgGoAPI/tlg"
)

/*
   main.go - основной метод для запуска бота
   bot/
       bot.go - всё, что доступно пользователю
   common/
       common.go - общие данные
       methods.go - методы для использования внутри модуля
   tlg/
       methods.go - структуры телеграм методов
       types.go - типы данных, которые использует телеграм
   constructor/nnnn
       menu.go - конструктор меню
*/

func Run() error {
	var wgc = _WaitGroupCount{
		limit: _PROCESSING_COUNT_LIMIT,
	}

	for {
		updates, err := getUpdates()
		if err != nil {
			log.Printf("ERROR - getUpdates: %s", err)
			return err
		}

		if !updates.Ok {
			err = fmt.Errorf("ERROR - while updating for new messages")
			log.Println(err)
			return err
		}

		if 0 == len(updates.Result) {
			continue
		}

		for _, update := range updates.Result {
			_OFFSET_TO_UPDATE = update.UpdateId + 1
			log.Printf("Update id: %d \"%s\"\n", update.UpdateId, update.Message.Text)

			// ожидание завершения обработки запросов до достижения лимита горутин
			for !wgc.Add(1) {
			}

			client, err := getClient(&update)
			if err != nil {
				log.Printf("ERROR - getClient: %s", err)
				wgc.Done()
				continue
			}

			messageType, err := getMessageType(&update)
			if err != nil {
				log.Printf("ERROR - getMessageType: %s", err)
				wgc.Done()
				continue
			}

			signal, err := getSignal(client, &update, messageType)
			if err != nil {
				log.Printf("ERROR - getSignal: %s", err)
				wgc.Done()
				continue
			}

			go func(update *tlg.Update) {
				err := bot.MakeTransition(client, update, signal)
				if err != nil {
					log.Printf("ERROR - MakeTransition: %s", err)
				}
				wgc.Done()
			}(&update)
		}
	}
}
