package main

import (
	"errors"
	"runtime"

	"github.com/yuriykis/funny-filter/cmd"
	"github.com/yuriykis/funny-filter/log"
)

func main() {
	if err := checkOS(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func checkOS() error {
	if runtime.GOOS != "linux" {
		return errors.New("OS is not supported")
	}
	return nil
}
