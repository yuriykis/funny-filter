package main

import (
	"log"

	"github.com/yuriykis/funny-filter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
