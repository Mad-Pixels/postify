package tgsend

import (
	"fmt"

	"github.com/Mad-Pixels/go-postify"
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

	if raw.GetMetadata() == nil || raw.GetMetadata().Telegram.MessageID == 0 {
		response, err := tg.Send(data.String(), telegram.ModeMarkdownV2)
		if err != nil {
			return fmt.Errorf("failed to publish post from %s: %w", getFlagFrom(ctx), err)
		}
		postify.Logger.Info(getFlagFrom(ctx), " was published")
		raw.GetMetadata().Telegram.MessageID = response.MessageID
		raw.GetMetadata().Telegram.Date = response.Date
		return raw.Sync(getFlagFrom(ctx))
	}
	_, err = tg.Edit(raw.GetMetadata().Telegram.MessageID, data.String(), telegram.ModeMarkdownV2)
	if err != nil {
		return fmt.Errorf("failed to change post %s: %w", getFlagFrom(ctx), err)
	}
	postify.Logger.Info(getFlagFrom(ctx), " post was changed")
	return nil
}
