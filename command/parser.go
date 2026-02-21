package command

import (
	"errors"
	"strings"	
)


func ParseCommand(input string) (string, []string, error) {

	fields := strings.Fields(input)

	if len(fields) == 0 {
		return "", nil, errors.New("empty command")
	}

	cmd := fields[0]
	args := fields[1:]

	return cmd, args, nil
}