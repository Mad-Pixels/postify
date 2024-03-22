package tgsend

import (
	"fmt"

	"github.com/Mad-Pixels/go-postify/pkg/content"
	"github.com/Mad-Pixels/go-postify/pkg/telegram"
	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {
	raw, err := content.New(getFlagFrom(ctx), getFlagBlocks(ctx))
	if err != nil {
		return fmt.Errorf("failed to initialize content: %w", err)
	}

	tg, err := telegram.New(getFlagToken(ctx), getFlagChat(ctx))
	if err != nil {
		return fmt.Errorf("failed initialize tg API: %w", err)
	}
	data, err := raw.Conv(content.Telegram)
	if err != nil {
		return fmt.Errorf("failed convert data from sources: %w", err)
	}
	response, err := tg.Send(data.String(), telegram.ModeMarkdownV2)
	if err != nil {
		return fmt.Errorf("failed send message: %w", err)
	}

	raw.GetMetadata().Telegram.MessageID = response.MessageID
	raw.GetMetadata().Telegram.Date = response.Date
	return raw.Sync(getFlagFrom(ctx))
}
