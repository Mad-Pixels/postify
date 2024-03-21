package content

import (
	"encoding/json"
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
	if err := utils.IsFileOrError(filepath.Join(path, metafileName)); err != nil {
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
	if err := json.Unmarshal(body, md); err != nil {
		return nil, err
	}
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
	if err := utils.IsFileOrCreate(path); err != nil {
		return err
	}
	prevBody, err := utils.ReadFile(path)
	if err != nil {
		return err
	}

	var mdList []Metadata
	if len(prevBody) > 0 {
		if err = json.Unmarshal(prevBody, &mdList); err != nil {
			return err
		}
	}
	mdList = append(mdList, *m)

	newBody, err := json.MarshalIndent(mdList, "", "  ")
	if err != nil {
		return err
	}
	return utils.WriteToFile(path, newBody)
}
