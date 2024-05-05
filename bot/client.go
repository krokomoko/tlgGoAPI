package bot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetupRedisSettings(addr, password string, db int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - SetupRedisSettings: %s", r)
		}
	}()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err = RedisClient.Ping(context.Background()).Result()

	return
}

// TODO: if client is not in redis, try find him from file
func GetClient(id int64) (client *Client, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - GetClient: %s", r)
		}
	}()

	redisConn := RedisClient.Conn()
	defer redisConn.Close()

	if clientState, err := redisConn.Get(
		context.Background(),
		fmt.Sprintf("tlg_client:%d", id),
	).Result(); err != nil {
		client = &Client{
			ID:    id,
			State: states["init"],
		}
	} else {
		if _, ok := states[clientState]; !ok {
			err = fmt.Errorf("ERROR - GetClient, wrong client state \"%s\"", clientState)
			return nil, err
		}
		client = &Client{
			ID:    id,
			State: states[clientState],
		}
	}

	return
}

// TODO: save redis data to file by timer
func (client *Client) save() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - Client save: %s", r)
		}
	}()

	redisConn := RedisClient.Conn()
	defer redisConn.Close()

	err = redisConn.Set(
		context.Background(),
		fmt.Sprintf("tlg_client:%d", client.ID),
		client.State.Name,
		time.Duration(time.Minute*5),
	).Err()

	return
}

func (client *Client) SetState(stateName string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - Client SetState: %s", r)
			log.Println(err)
		}
	}()

	if _, ok := states[stateName]; !ok {
		err = fmt.Errorf("ERROR - Client SetState, state \"%s\" does not exist", stateName)
		return
	}

	client.State = states[stateName]

	client.save()

	return
}
