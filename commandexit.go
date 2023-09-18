package main

import (
	"os"
	"errors"
)

func commandExit(cfg *config, i string) error {
	if len(i) > 0{
		return errors.New("no such commands")
	}
	os.Exit(0)
	return nil
}