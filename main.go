package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/nurali-techie/picopy/cli"
	"github.com/nurali-techie/picopy/command"
)

func main() {
	var commandName string
	var args []string

	if len(os.Args) <= 1 {
		commandName = "help"
		args = []string{}
	} else {
		commandName = os.Args[1]
		args = os.Args[2:]
	}

	ExecuteCommand(commandName, args)
}

func ExecuteCommand(commandName string, args []string) {
	var cmd cli.Command
	switch strings.ToLower(commandName) {
	case "help":
		cmd = command.NewHelpCommand()
	case "backup":
		cmd = command.NewBackupCommand()
	}

	if cmd == nil {
		fmt.Printf("%q command not valid\n", commandName)
		os.Exit(0)
	}

	err := cmd.Execute(context.Background(), args)
	if err != nil {
		fmt.Printf("%q command failed with error: %s\n", commandName, err.Error())
		os.Exit(1)
	}
}
