package htmlcontent

import (
	"fmt"
	"strings"

	"github.com/Mad-Pixels/go-postify"
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
			EnvVars:  []string{fmt.Sprintf("%s_HTML_FROM", postify.EnvName)},
			Required: true,
		},
		&cli.StringFlag{
			Name:     flagToPath,
			Usage:    "the directory where generated content will be placed",
			EnvVars:  []string{fmt.Sprintf("%s_HTML_TO", postify.EnvName)},
			Required: true,
		},
		// feature flags:
		&cli.StringFlag{
			Name: flagAssetsPath,
			Usage: fmt.Sprintf(
				"the directory with assests that will be added along with the content in path '--%s'",
				flagToPath,
			),
			EnvVars:  []string{fmt.Sprintf("%s_HTML_ASSETS", postify.EnvName)},
			Required: false,
			Value:    defaultFlagAssetsPath,
		},
		&cli.StringFlag{
			Name: flagContentName,
			Usage: fmt.Sprintf(
				"the result generetade content filename in path '--%s'",
				flagToPath,
			),
			EnvVars:  []string{fmt.Sprintf("%s_HTML_FILENAME", postify.EnvName)},
			Required: false,
			Value:    defaultFlagContentName,
		},
		&cli.StringFlag{
			Name:     flagBlocks,
			Usage:    "add a few markdown files by comma",
			EnvVars:  []string{fmt.Sprintf("%s_HTML_BLOCKS", postify.EnvName)},
			Required: false,
			Value:    defaultFlagBlocks,
		},
		&cli.StringFlag{
			Name:     flagTmplPath,
			Usage:    "specify the template path for adding generated data",
			EnvVars:  []string{fmt.Sprintf("%s_HTML_TMPL", postify.EnvName)},
			Required: false,
			Value:    defaultFlagTmplPath,
		},
		&cli.StringFlag{
			Name:     flagRouterPath,
			Usage:    "specify the JSON file for generate static router",
			EnvVars:  []string{fmt.Sprintf("%s_HTML_ROUTER", postify.EnvName)},
			Required: false,
			Value:    defaultFlagRouterPath,
		},
	}
}
