package tlgGoAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// first make stateMachine in json file
// Postgresql and other will be later

type stateMachine map[int64]string

func newStateMachine() *stateMachine {
	result := make(stateMachine)
	return &result
}

func (st stateMachine) Add(userId int64) {
	st[userId] = "start"
}

func (st stateMachine) Set(userId int64, state string) {
	st[userId] = state
}

func (st stateMachine) Get(userId int64) (string, bool) {
	state, ok := st[userId]
	return state, ok
}

func (st *stateMachine) Load() bool {
	input, err := ioutil.ReadFile("stateMachine.json")
	if err != nil {
		fmt.Println("ERROR: while reading stateMachine.json:", err)
		return false
	}

	err = json.Unmarshal(input, st)
	if err != nil {
		fmt.Println("ERROR: while unmarshal stateMachine string to object:", err)
		return false
	}

	return true
}

func (st *stateMachine) Save() bool {
	output, err := json.Marshal(st)
	if err != nil {
		fmt.Println("ERROR: while marshal stateMachine object to string:", err)
		return false
	}

	f, err := os.Create("stateMachine.json")
	if err != nil {
		fmt.Println("ERROR: while open stateMachine.json:", err)
		return false
	}
	defer f.Close()

	_, err = f.Write(output)
	if err != nil {
		fmt.Println("ERROR: while writing to stateMachine.json:", err)
		return false
	}

	return true
}
