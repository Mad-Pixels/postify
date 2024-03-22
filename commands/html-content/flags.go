package htmlcontent

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func getFlagFrom(ctx *cli.Context) string {
	return ctx.String(flagFromPath)
}

func getFlagTo(ctx *cli.Context) string {
	return ctx.String(flagToPath)
}

func getFlagRouter(ctx *cli.Context) string {
	return ctx.String(flagRouterPath)
}

func getFlagTmpl(ctx *cli.Context) string {
	return ctx.String(flagTmplPath)
}

func getFlagAssets(ctx *cli.Context) string {
	return ctx.String(flagAssetsPath)
}

func getFlagContentName(ctx *cli.Context) string {
	return ctx.String(flagContentName)
}

func getFlagBlocks(ctx *cli.Context) []string {
	return strings.Split(strings.ReplaceAll(ctx.String(flagBlocks), " ", ""), ",")
}

func flags() []cli.Flag {
	return []cli.Flag{
		// required flags:
		&cli.StringFlag{
			Name:     flagFromPath,
			Usage:    "the source directory containing content",
			EnvVars:  []string{"CONTENT_HTML_FROM"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagToPath,
			Usage:    "the directory where generated content will be placed",
			EnvVars:  []string{"CONTENT_HTML_TO"},
			Required: true,
		},
		// feature flags:
		&cli.StringFlag{
			Name: flagAssetsPath,
			Usage: fmt.Sprintf(
				"the directory with assests that will be added along with the content in path '--%s'",
				flagToPath,
			),
			EnvVars:  []string{"CONTENT_HTML_ASSETS"},
			Required: false,
			Value:    defaultFlagAssetsPath,
		},
		&cli.StringFlag{
			Name: flagContentName,
			Usage: fmt.Sprintf(
				"the result generetade content filename in path '--%s'",
				flagToPath,
			),
			EnvVars:  []string{"CONTENT_HTML_FILENAME"},
			Required: false,
			Value:    defaultFlagContentName,
		},
		&cli.StringFlag{
			Name:     flagBlocks,
			Usage:    "add a few markdown files by comma",
			EnvVars:  []string{"CONTENT_HTML_BLOCKS"},
			Required: false,
			Value:    defaultFlagBlocks,
		},
		&cli.StringFlag{
			Name:     flagTmplPath,
			Usage:    "specify the template path for adding generated data",
			EnvVars:  []string{"CONTENT_HTML_TMPL"},
			Required: false,
			Value:    defaultFlagTmplPath,
		},
		&cli.StringFlag{
			Name:     flagRouterPath,
			Usage:    "specify the JSON file for generate static router",
			EnvVars:  []string{"CONTENT_HTML_ROUTER"},
			Required: false,
			Value:    defaultFlagRouterPath,
		},
	}
}
