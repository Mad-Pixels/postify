package content

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/Mad-Pixels/go-postify/utils"
)

type Content interface {
	ConvWithTmpl(format fType, tmplPath string) (*bytes.Buffer, error)
	Conv(format fType) (*bytes.Buffer, error)
	Sync(path string) error
	WriteRouter(path string) error
	GetMetadata() *Metadata
}

func New(path string, contentFiles []string, opts ...Option) (Content, error) {
	if err := utils.IsDirOrError(path); err != nil {
		return nil, err
	}
	c := &content{
		path: path,
	}
	for _, opt := range opts {
		opt(c)
	}
	for _, file := range contentFiles {
		if err := utils.IsFileOrError(filepath.Join(path, filepath.Base(file))); err != nil {
			return nil, err
		}
		c.contentFiles = append(c.contentFiles, filepath.Base(file))
	}
	md, err := newMetafile(path)
	if err != nil {
		return nil, err
	}
	c.Metadata = md
	return c, nil
}

type content struct {
	*Metadata

	contentFiles []string
	metafile     string
	path         string
}

func (c *content) GetMetadata() *Metadata {
	return c.Metadata
}

func (c *content) Conv(format fType) (*bytes.Buffer, error) {
	raw, err := c.formatted(format)
	if err != nil {
		return nil, err
	}

	data := new(bytes.Buffer)
	for _, file := range c.contentFiles {
		buf, ok := raw[file]
		if !ok {
			return nil, fmt.Errorf("missing content file: %s", file)
		}
		_, err = data.Write(buf.Bytes())
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (c *content) ConvWithTmpl(format fType, tmplPath string) (*bytes.Buffer, error) {
	raw, err := c.formatted(format)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.ParseFiles(filepath.Clean(tmplPath))
	if err != nil {
		return nil, err
	}
	data := make(map[string]string, len(raw))
	for file, body := range raw {
		data[file] = body.String()
	}

	res := new(bytes.Buffer)
	if err = tmpl.Execute(res, data); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *content) raw() (map[string][]byte, error) {
	data := make(map[string][]byte)

	for _, file := range c.contentFiles {
		body, err := utils.ReadFile(filepath.Join(c.path, file))
		if err != nil {
			return nil, err
		}
		data[file] = body
	}
	return data, nil
}

func (c *content) formatted(format fType) (map[string]*bytes.Buffer, error) {
	raw, err := c.raw()
	if err != nil {
		return nil, err
	}

	data := make(map[string]*bytes.Buffer, len(raw))
	for file, body := range raw {
		formatted := new(bytes.Buffer)

		switch format {
		case HTML:
			if err = contentHTML(body, formatted); err != nil {
				return nil, err
			}
		case Telegram:
			if err = contentTelegram(body, formatted); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported format: %v", format)
		}
		data[file] = formatted
	}
	return data, nil
}
