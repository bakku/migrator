package commands

import (
	"errors"
	"fmt"
)

type Command interface {
	Run(args ...string) error
}

func Select(command string) (Command, error) {
	switch command {
	case "generate":
		return &Generator{}, nil
	case "migrate":
		return &Migrator{}, nil
	case "init":
		return &Initializer{}, nil
	case "rollback":
		return &Backroller{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("the command '%s' does not exist", command))
	}
}