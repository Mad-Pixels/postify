package telegram

// ParseMode custom type.
type ParseMode string

// define parseMode rules.
const (
	ModeMarkdownV2 ParseMode = "MarkdownV2"
	ModeMarkdown   ParseMode = "Markdown"
	ModeHTML       ParseMode = "HTML"
)

// Option ...
type Option func(Telegram)
