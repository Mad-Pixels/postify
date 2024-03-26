package content

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/Mad-Pixels/go-postify/utils"
)

// Metadata ...
type Metadata struct {
	Telegram telegramData `json:"telegram"`
	Static   staticData   `json:"static"`
}

type telegramData struct {
	MessageID int `json:"message_id"`
	Date      int `json:"date"`
}

type staticData struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func newMetafile(path string) (*Metadata, error) {
	if err := utils.IsFileOrCreate(filepath.Join(path, metafileName)); err != nil {
		return nil, err
	}
	body, err := utils.ReadFile(filepath.Join(path, metafileName))
	if err != nil {
		return nil, err
	}
	md := &Metadata{
		Static: staticData{
			Title: filepath.Base(path),
			Url:   filepath.Join(urlPrefix, filepath.Base(path)),
		},
	}
	_ = json.Unmarshal(body, md)
	return md, nil
}

// Sync updates the metadata file at the specified path with the information from the provided Metadata object.
func (m *Metadata) Sync(path string) (err error) {
	body, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return utils.WriteToFile(filepath.Join(path, metafileName), body)
}

// WriteRouter appends the current Metadata object to a list of metadata stored in a file at the given path.
func (m *Metadata) WriteRouter(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	var mdList map[string]Metadata
	prevBody, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(prevBody) > 0 {
		if err = json.Unmarshal(prevBody, &mdList); err != nil {
			return err
		}
	} else {
		mdList = make(map[string]Metadata)
	}
	mdList[m.Static.Url] = *m

	file.Seek(0, io.SeekStart)
	file.Truncate(0)
	newBody, err := json.Marshal(mdList)
	if err != nil {
		return err
	}
	_, err = file.Write(newBody)
	return err
}
