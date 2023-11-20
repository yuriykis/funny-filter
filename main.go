package main

import (
	"github.com/yuriykis/funny-filter/cmd"
	"github.com/yuriykis/funny-filter/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
