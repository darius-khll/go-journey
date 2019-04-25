package main

import (
	"fmt"
	"os"

	"go-journey/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ./t mars --lang es
// ./t traveller Boston
