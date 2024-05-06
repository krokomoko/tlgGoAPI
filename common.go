package tlgGoAPI

import (
	"fmt"
	"log"
	"time"
)

// message types
const (
	M_COMMAND = iota + 1
	M_TEXT
	M_IMAGE
	M_VIDEO
	M_VOICE
	M_AUDIO
	M_GIF
	M_DOCUMENT
	M_STICKER
	M_URL
	M_CALLBACK
)

var (
	_OFFSET_TO_UPDATE       int64  = 0
	_PROCESSING_COUNT_LIMIT uint16 = 1
	_CLIENTS_SAVING_TIMER          = time.Minute * 2
)

func SetUpdateOffset(offset int64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - bot, SetUpdateOffset: %s", r)
			log.Println(err)
		}
	}()

	_OFFSET_TO_UPDATE = offset

	return
}

func SetProcessingCountLimit(count uint16) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - bot, SetProcessingCountLimit: %s", r)
			log.Println(err)
		}
	}()

	_PROCESSING_COUNT_LIMIT = count

	return
}
