package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Telegram defines an interface for work with Telegram API.
type Telegram interface {
	Edit(id int, msg string, mode ParseMode) (*tgbotapi.Message, error)
	Send(msg string, mode ParseMode) (*tgbotapi.Message, error)
}

type telegram struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

// New initialize Telegram object.
func New(token string, chatID int64, opts ...Option) (Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	tg := &telegram{
		chatID: chatID,
		bot:    bot,
	}
	for _, opt := range opts {
		opt(tg)
	}
	return tg, nil
}

// Send message to telegram chat or chanel.
func (t *telegram) Send(msg string, mode ParseMode) (*tgbotapi.Message, error) {
	message := tgbotapi.NewMessage(t.chatID, msg)
	message.ParseMode = string(mode)

	resp, err := t.bot.Send(message)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Edit message which already published.
func (t *telegram) Edit(id int, msg string, mode ParseMode) (*tgbotapi.Message, error) {
	editMsg := tgbotapi.NewEditMessageText(t.chatID, id, msg)
	editMsg.ParseMode = string(mode)

	resp, err := t.bot.Send(editMsg)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
