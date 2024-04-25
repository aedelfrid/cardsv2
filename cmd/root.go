package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type command struct {
	name string
	executor func(commandList)
}

type commandList map[string]command

var itemCommands commandList


func menu() {

	names := []string{
		"deck",
	}

	initCommands := []command{
		{
			name: "deck",
			executor: func (commandList)  {
				Cards()
			},
		},
	}

	for _, name := range names {
		for _, initCommand := range initCommands {
			itemCommands[name] = initCommand
		}
	}

	for label, items := range itemCommands {
		prompt := promptui.Select{
			Label: label,
			Items: []string{},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Print(err)
		}
	}

	


}