package commands

import (
	htmlcontent "github.com/Mad-Pixels/go-postify/commands/html-content"
	tgsend "github.com/Mad-Pixels/go-postify/commands/tg-send"

	"github.com/urfave/cli/v2"
)

// Commands define cli-commands.
func Commands() []*cli.Command {
	return []*cli.Command{
		htmlcontent.Command(),
		tgsend.Command(),
	}
}
