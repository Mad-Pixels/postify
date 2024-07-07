package htmlcontent

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/Mad-Pixels/postify"
	"github.com/Mad-Pixels/postify/pkg/content"
	"github.com/Mad-Pixels/postify/utils"
)

func action(ctx *cli.Context) error {
	raw, err := content.New(getFlagFrom(ctx), getFlagBlocks(ctx))
	if err != nil {
		return fmt.Errorf("failed to initialize content: %w", err)
	}

	var data *bytes.Buffer
	if getFlagTmpl(ctx) != "" {
		data, err = raw.ConvWithTmpl(content.HTML, getFlagTmpl(ctx))
	} else {
		data, err = raw.Conv(content.HTML)
	}
	if err != nil {
		return fmt.Errorf("failed convert data from sources: %w", err)
	}
	err = utils.WriteToFile(filepath.Join(getFlagTo(ctx), getFlagContentName(ctx)), data.Bytes())
	if err != nil {
		return fmt.Errorf(
			"failed write data to %s: %w",
			filepath.Join(getFlagTo(ctx), getFlagContentName(ctx)),
			err,
		)
	}

	if getFlagAssets(ctx) != "" {
		if err = utils.Copy(getFlagAssets(ctx), getFlagTo(ctx)); err != nil {
			return fmt.Errorf("failed to copy asstets data: %w", err)
		}
	}
	if getFlagRouter(ctx) != "" {
		if err = raw.WriteRouter(getFlagRouter(ctx)); err != nil {
			return fmt.Errorf("failed preapre site router content: %w", err)
		}
	}
	postify.Logger.Info(getFlagFrom(ctx), " created static content")
	return raw.Sync(getFlagFrom(ctx))
}
