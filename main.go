package main

import (
	"fmt"
	"github.com/naoki912/slack-tools/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
