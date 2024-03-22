package content

import (
	"bytes"
	"fmt"

	tgmd "github.com/Mad-Pixels/goldmark-tgmd"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// contentHTML define goldmark pkg for convert Markdown to HTML.
func contentHTML(in []byte, out *bytes.Buffer) error {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.CJK,
			extension.Typographer,
		),
		goldmark.WithParserOptions(
			parser.WithAttribute(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
			html.WithXHTML(),
		),
	)
	if err := md.Convert(in, out); err != nil {
		return fmt.Errorf("generate HTML content failed, got: %w", err)
	}
	return nil
}

// contentTelegram define goldmark pkg for conver Markdown to Telegram.
func contentTelegram(in []byte, out *bytes.Buffer) error {
	md := tgmd.TGMD()
	if err := md.Convert(in, out); err != nil {
		return fmt.Errorf("generate Telegram content failed, got: %w", err)
	}
	return nil
}
