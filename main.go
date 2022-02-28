package main

import (
	"fmt"
	"os"

	"github.com/inoth/inobot/cmd"
	"github.com/inoth/inobot/proxy"
)

func init() {
	proxy.AutoSet()
}

func main() {
	defer func() {
		if exception := recover(); exception != nil {
			if err, ok := exception.(error); ok {
				fmt.Printf("%v\n", err)
			} else {
				panic(exception)
			}
			os.Exit(1)
		}
	}()
	cmd.Execute()
}
