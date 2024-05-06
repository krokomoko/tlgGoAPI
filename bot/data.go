package bot

import (
	"time"

	"github.com/krokomoko/tlgGoAPI/tlg"
)

var (
	states              = map[string]*_State{}
	messageTypes        = map[string]func(client *Client, update *tlg.Update) (string, error){}
	botToken            string
	methodsURL          string
	redisPrefixKey      string        = "tlg_client"
	redisClientDuration time.Duration = time.Minute * 5
)
