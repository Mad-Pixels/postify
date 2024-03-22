package main

import (
	"os"

	"github.com/Mad-Pixels/go-postify"
	"github.com/Mad-Pixels/go-postify/commands"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// define globals.
func init() {
	postify.Logger.SetOutput(os.Stdout)
	postify.Logger.SetLevel(logrus.InfoLevel)
	postify.Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func main() {
	app := &cli.App{
		Name:     postify.Name,
		Usage:    postify.Usage,
		Commands: commands.Commands(),
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
