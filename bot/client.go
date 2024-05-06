package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	RedisClient *redis.Client
	clientsDB   *gorm.DB
)

func init() {
	//init client sqlite db
	var err error
	clientsDB, err = gorm.Open(
		sqlite.Open("clients.db"),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("ERROR - init clients.db: ", err)
	}

	if err = clientsDB.AutoMigrate(&Client{}); err != nil {
		log.Fatal("ERROR - migrate Client to clients.db: ", err)
	}
}

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

	//data reconciliation
	var clients []*Client

	if result := clientsDB.Find(&clients); result.Error != nil {
		err = result.Error
		return
	}

	redisConn := RedisClient.Conn()
	defer redisConn.Close()

	var clientsId struct {
		ClientsId []int64 `json:"clients_id"`
	}
	clientsId.ClientsId = []int64{}

	for _, client := range clients {
		if err = redisConn.Set(
			context.Background(),
			fmt.Sprintf("%s:%d", redisPrefixKey, client.ID),
			client.State.Name,
			time.Duration(redisClientDuration),
		).Err(); err != nil {
			err = fmt.Errorf(
				"ERROR - SetupRedisSettings, setting client to redis: %s",
				err,
			)
			return
		}

		clientsId.ClientsId = append(clientsId.ClientsId, client.ID)
	}

	output, err := json.Marshal(&clientsId)
	if err != nil {
		err = fmt.Errorf(
			"ERROR - SetupRedisSettings, marshal clients idto json: %s",
			err,
		)
		return err
	}

	if err = redisConn.Set(
		context.Background(),
		fmt.Sprintf("%s:clients_id", redisPrefixKey),
		output,
		0,
	).Err(); err != nil {
		err = fmt.Errorf(
			"ERROR - SetupRedisSettings, set clients id to redis: %s",
			err,
		)
		return err
	}

	return
}

func getClientFromFile(id int64) (client *Client, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - getClientFromFile: %s", r)
		}
	}()

	client = &Client{ID: id}

	if result := clientsDB.First(client); result.Error != nil {
		err = result.Error
	}

	return
}

func SaveClients() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - SaveClients: %s", r)
		}
	}()

	if RedisClient == nil {
		return
	}

	redisConn := RedisClient.Conn()
	defer redisConn.Close()

	var clientsId struct {
		ClientsId []int64 `json:"clients_id"`
	}
	if clientsId_str, err := redisConn.Get(
		context.Background(),
		fmt.Sprintf("%s:clients_id", redisPrefixKey),
	).Result(); err != nil {
		return fmt.Errorf(
			"ERROR - SaveClients, get clients id from redis: %s",
			err,
		)
	} else {
		if err = json.Unmarshal([]byte(clientsId_str), &clientsId); err != nil {
			return fmt.Errorf(
				"ERROR - SaveClients, unmarshal list of clients id from redis to object: %s",
				err,
			)
		}

		for _, clientId := range clientsId.ClientsId {
			clientState, err := redisConn.Get(
				context.Background(),
				fmt.Sprintf("%s:%d", redisPrefixKey, clientId),
			).Result()

			if err != nil {
				return fmt.Errorf(
					"ERROR - SaveClients, get client state from redis: %s",
					err,
				)
			}

			client := &Client{
				ID: clientId,
				State: &_State{
					Name: clientState,
				},
			}
			if result := clientsDB.Save(client); result.Error != nil {
				return fmt.Errorf(
					"ERROR - SaveClients, save client to file (or db): %s",
					result.Error,
				)
			}
		}
	}

	return
}

func GetClient(id int64) (client *Client, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - GetClient: %s", r)
		}
	}()

	var clientState string

	clientFromFile := func() error {
		if client_, err := getClientFromFile(id); err == gorm.ErrRecordNotFound {
			client_.ID = id
			client_.State = &_State{
				Name: "init",
			}
			if result := clientsDB.Save(client_); result.Error != nil {
				return fmt.Errorf(
					"ERROR - GetClient, save new client to file (or db): %s",
					result.Error,
				)
			}
			clientState = "init"
		} else if err != nil {
			return fmt.Errorf(
				"ERROR - GetClient, getting client from file (or db): %s",
				err,
			)
		} else {
			clientState = client_.State.Name
		}

		return nil
	}

	if RedisClient != nil {
		redisConn := RedisClient.Conn()
		defer redisConn.Close()

		if clientState, err = redisConn.Get(
			context.Background(),
			fmt.Sprintf("%s:%d", redisPrefixKey, id),
		).Result(); err != nil {
			if err = clientFromFile(); err != nil {
				return
			}

			if err = redisConn.Set(
				context.Background(),
				fmt.Sprintf("%s:%d", redisPrefixKey, id),
				clientState,
				time.Duration(redisClientDuration),
			).Err(); err != nil {
				err = fmt.Errorf("ERROR - GetClient, set client to redis: %s", err)
				return
			}

			var clientsId struct {
				ClientsId []int64 `json:"clients_id"`
			}
			if clientsId_str, err := redisConn.Get(
				context.Background(),
				fmt.Sprintf("%s:clients_id", redisPrefixKey),
			).Result(); err != nil {
				return nil, fmt.Errorf(
					"ERROR - GetClient, get clients id from redis: %s",
					err,
				)
			} else {
				if err = json.Unmarshal([]byte(clientsId_str), &clientsId); err != nil {
					return nil, fmt.Errorf(
						"ERROR - GetClient, unmarshal list of clients id from redis to object: %s",
						err,
					)
				}

				if !slices.Contains(clientsId.ClientsId, id) {
					clientsId.ClientsId = append(clientsId.ClientsId, id)

					output, err := json.Marshal(&clientsId)
					if err != nil {
						return nil, fmt.Errorf(
							"ERROR - GetClient, marshal updated list of clients id to json: %s",
							err,
						)
					}

					if err = redisConn.Set(
						context.Background(),
						fmt.Sprintf("%s:clients_id", redisPrefixKey),
						output,
						0,
					).Err(); err != nil {
						return nil, fmt.Errorf(
							"ERROR - GetClient, setting updated list of clients id to redis: %s",
							err,
						)
					}
				}
			}
		}
	} else {
		if err = clientFromFile(); err != nil {
			return
		}
	}

	if _, ok := states[clientState]; !ok {
		err = fmt.Errorf(
			"ERROR - GetClient, wrong client state \"%s\"",
			clientState,
		)
		return nil, err
	}
	client = &Client{
		ID:    id,
		State: states[clientState],
	}

	return
}

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
		fmt.Sprintf("%s:%d", redisPrefixKey, client.ID),
		client.State.Name,
		time.Duration(redisClientDuration),
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
