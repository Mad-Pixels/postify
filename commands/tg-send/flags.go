package tgsend

import (
	"fmt"
	"strings"

	"github.com/Mad-Pixels/go-postify"
	"github.com/urfave/cli/v2"
)

func getFlagFrom(ctx *cli.Context) string {
	return ctx.String(flagFromPath)
}

func getFlagToken(ctx *cli.Context) string {
	return ctx.String(flagTgToken)
}

func getFlagChat(ctx *cli.Context) int64 {
	return ctx.Int64(flagTgChat)
}

func getFlagBlocks(ctx *cli.Context) []string {
	return strings.Split(strings.ReplaceAll(ctx.String(flagBlocks), " ", ""), ",")
}

func flags() []cli.Flag {
	return []cli.Flag{
		// required flags:
		&cli.StringFlag{
			Name:     flagFromPath,
			Usage:    "the source directory containig content",
			EnvVars:  []string{fmt.Sprintf("%s_TG_FROM", postify.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagTgToken,
			Usage:    "the telegram bo token",
			EnvVars:  []string{fmt.Sprintf("%s_TG_TOKEN", postify.EnvName)},
			Required: true,
		},
		&cli.Int64Flag{
			Name:     flagTgChat,
			Usage:    "the telegram chat or channel ID",
			EnvVars:  []string{fmt.Sprintf("%s_TG_CHAT", postify.EnvName)},
			Required: true,
		},
		// feature flags:
		&cli.StringFlag{
			Name:     flagBlocks,
			Usage:    "add a few markdown files by comma",
			EnvVars:  []string{fmt.Sprintf("%s_HTML_BLOCKS", postify.EnvName)},
			Required: false,
			Value:    defaultFlagBlocks,
		},
	}
}
