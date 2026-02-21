package command

import (
	"errors"
	"fmt"
	"github.com/Dhanushkumar-S-G/ToyDB/toy"
)

func ExecuteCommand(cmd string, args []string, toyDB *toy.Toy) (string, error) {

	switch cmd {
		case "SET":
			if(len(args) != 2) {
				return "", errors.New("SET command requires exactly 2 arguments")
			}
			toyDB.Store[args[0]] = args[1]
			return "OK", nil
		case "GET":
			value, exists := toyDB.Store[args[0]]
			if !exists {
				return "", errors.New("Key not found")
			}
			return value, nil
		case "DEL":
			delete(toyDB.Store, args[0])
			return "OK", nil
		default:
			return "", errors.New(fmt.Sprintf("Command not found: %s", cmd))

	}
}