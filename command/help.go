package command

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/nurali-techie/picopy/cli"
)

//go:embed help.txt
var help string

type helpCommand struct {
}

func NewHelpCommand() cli.Command {
	return new(helpCommand)
}

func (c *helpCommand) Execute(ctx context.Context, args []string) error {
	fmt.Println(help)
	return nil
}
