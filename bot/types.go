package bot

import "github.com/krokomoko/tlgGoAPI/tlg"

type Client struct {
	ID    int64
	State *_State
}

type _State struct {
	Name        string
	transitions map[string]*_Transition
	loginAction func(client *Client, update *tlg.Update) error
}

type _Transition struct {
	next       *_State
	exitAction func(client *Client, update *tlg.Update) error
}
