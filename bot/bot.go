package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/krokomoko/tlgGoAPI/tlg"
)

func SetToken(token string) {
	botToken = token
	methodsURL = fmt.Sprintf("https://api.telegram.org/bot%s/", token)
}

func SetRedisPrefixKey(prefix string) {
	redisPrefixKey = prefix
}

func SetStates(stateNames ...string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - bot, SetStates: %s", r)
		}
	}()

	if _, ok := states["init"]; !ok {
		states["init"] = &_State{
			Name:        "init",
			transitions: make(map[string]*_Transition),
		}
	}

	for _, _stateName := range stateNames {
		states[_stateName] = &_State{
			Name:        _stateName,
			transitions: make(map[string]*_Transition),
		}
	}

	return
}

func SetTransition(
	stateFrom, stateTo, signal string,
	fn func(client *Client, update *tlg.Update) error,
) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - bot, SetTransition: %s", r)
		}
	}()

	if _, ok := states[stateFrom]; !ok {
		err = fmt.Errorf(
			"ERROR - SetTransition, state \"%s\" does not exist",
			stateFrom,
		)
		return
	}

	if _, ok := states[stateTo]; !ok {
		err = fmt.Errorf(
			"ERROR - SetTransition, state \"%s\" does not exist",
			stateTo,
		)
		return
	}

	if _, ok := states[stateFrom].transitions[signal]; ok {
		err = fmt.Errorf(
			"ERROR - SetTransition, transition from \"%s\" to \"%s\" by signal \"%s\" already exist",
			stateFrom,
			stateTo,
			signal,
		)
		return
	}

	states[stateFrom].transitions[signal] = &_Transition{
		next:       states[stateTo],
		exitAction: fn,
	}

	return
}

func SetMessageType(
	fn func(client *Client, update *tlg.Update) (string, error),
	types ...int64,
) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - bot, SetMessageType: %s", r)
		}
	}()

	if len(types) == 0 {
		err = fmt.Errorf("ERROR - SetMessageType, no types")
		return
	}

	sort.Slice(types, func(i, j int) bool {
		return types[i] < types[j]
	})

	var hash string

	for _, type_ := range types {
		hash += fmt.Sprintf("%d.", type_)
	}

	if _, ok := messageTypes[hash]; ok {
		err = fmt.Errorf("ERROR - SetMessageType, this type already exist")
		return
	}

	messageTypes[hash] = fn

	return
}

func GetMessageTypeCallback(messageTypeHash string) (fn func(client *Client, update *tlg.Update) (string, error), err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - GetMessageTypeCallback: %s", r)
		}
	}()

	var ok bool
	if fn, ok = messageTypes[messageTypeHash]; !ok {
		err = fmt.Errorf(
			"ERROR - GetMessageTypeCallback, not registered message type with hash \"%s\"",
			messageTypeHash,
		)

		return
	}

	return
}

func MakeTransition(client *Client, update *tlg.Update, signal string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - MakeTransition: %s", r)
		}
	}()

	if transition, ok := client.State.transitions[signal]; !ok {
		return
	} else if transition.next == nil {
		err = fmt.Errorf(
			"ERROR - MakeTransition: state \"%s\" have not transition by signal \"%s\"",
			client.State.Name,
			signal,
		)
		return
	}

	if client.State.transitions[signal].exitAction != nil {
		err = client.State.transitions[signal].exitAction(client, update)
		if err != nil {
			err = fmt.Errorf(
				"ERROR - MakeTransition, exitAction from state \"%s\" by signal \"%s\"",
				client.State.Name,
				signal,
			)
			return
		}
	}

	if client.State.transitions[signal].next.loginAction != nil {
		err = client.State.transitions[signal].next.loginAction(client, update)
		if err != nil {
			err = fmt.Errorf(
				"ERROR - MakeTransition, loginAction from state \"%s\" to state \"%s\" by signal \"%s\"",
				client.State.Name,
				client.State.transitions[signal].next.Name,
				signal,
			)
			return
		}
	}

	client.State = client.State.transitions[signal].next
	client.save()

	return
}

func Call(fn_struct interface{}) (responseBody string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("PANIC - Call tlg method: %s", r)
		}
	}()

	if len(botToken) == 0 {
		return "", errors.New("Token of bot is not exist")
	}

	method_name, err := tlg.GetFuncName(fn_struct)
	if err != nil {
		return "", err
	}

	fn_struct_json, _ := json.Marshal(fn_struct)

	postBody := bytes.NewBuffer(fn_struct_json)

	resp, err := http.Post(methodsURL+method_name, "application/json", postBody)
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
