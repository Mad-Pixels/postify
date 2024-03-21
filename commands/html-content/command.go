package htmlcontent

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/Mad-Pixels/go-postify"
	"github.com/urfave/cli/v2"
)

var (
	flagFromPath    = "from"
	flagToPath      = "to"
	flagRouterPath  = "with-router"
	flagContentName = "with-name"
	flagTmplPath    = "with-tmpl"
	flagAssetsPath  = "with-assets"
	flagBlocks      = "with-blocks"

	defaultFlagContentName = "index.html"
	defaultFlagBlocks      = "main.md"
	defaultFlagRouterPath  = ""
	defaultFlagAssetsPath  = ""
	defaultFlagTmplPath    = ""

	usage = "create HTML static content from Markdown text"
	name  = "html-content"
)

type tmplUsage struct {
	FlagContentName string
	FlagAssetsPath  string
	FlagRouterPath  string
	FlagFromPath    string
	FlagTmplPath    string
	FlagBlocks      string
	FlagToPath      string

	DefaultFlagContentName string
	DefaultFlagRouterPath  string
	DefaultFlagBlocks      string
	DefaultFlagAssetsPath  string
	DefaultFlagTmplPath    string
}

func Command() *cli.Command {
	tmpl, err := template.New("usage").Funcs(template.FuncMap{
		"Join": strings.Join,
	}).Parse(usageTemplate)
	if err != nil {
		postify.Logger.Fatal(err)
	}

	var usageText bytes.Buffer
	err = tmpl.Execute(&usageText, tmplUsage{
		FlagContentName: flagContentName,
		FlagAssetsPath:  flagAssetsPath,
		FlagRouterPath:  flagRouterPath,
		FlagFromPath:    flagFromPath,
		FlagTmplPath:    flagTmplPath,
		FlagBlocks:      flagBlocks,
		FlagToPath:      flagToPath,

		DefaultFlagContentName: defaultFlagContentName,
		DefaultFlagRouterPath:  defaultFlagRouterPath,
		DefaultFlagBlocks:      defaultFlagBlocks,
		DefaultFlagAssetsPath:  defaultFlagAssetsPath,
		DefaultFlagTmplPath:    defaultFlagTmplPath,
	})
	if err != nil {
		postify.Logger.Fatal(err)
	}
	return &cli.Command{
		Name:      name,
		Usage:     usage,
		UsageText: usageText.String(),
		Flags:     flags(),
		Action:    action,
	}
}
