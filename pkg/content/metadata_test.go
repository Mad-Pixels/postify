package content

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMetafile(t *testing.T) {
	tmpDir := t.TempDir()
	metafilePath := filepath.Join(tmpDir, metafileName)

	testMetadata := &Metadata{
		Telegram: telegramData{
			MessageID: 123,
			Date:      456,
		},
		Static: staticData{
			Title: "Example Title",
			Path:  "/content/article/",
		},
	}
	data, err := json.Marshal(testMetadata)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(metafilePath, data, 0644))

	md, err := newMetafile(tmpDir)
	require.NoError(t, err)
	assert.Equal(t, testMetadata.Telegram.MessageID, md.Telegram.MessageID)
	assert.Equal(t, testMetadata.Static.Title, md.Static.Title)
}

func TestMetadata_WriteRouter(t *testing.T) {
	tmpDir := t.TempDir()
	routerFilePath := filepath.Join(tmpDir, "router.json")

	md := &Metadata{
		Telegram: telegramData{MessageID: 123, Date: 456},
		Static:   staticData{Title: "Title", Path: "/content/article/"},
	}
	err := md.WriteRouter(routerFilePath)
	require.NoError(t, err)

	var mdMap map[string]Metadata
	data, err := os.ReadFile(routerFilePath)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &mdMap))

	require.Len(t, mdMap, 1, "The map should contain exactly one entry.")
	assert.Equal(t, md.Telegram.MessageID, mdMap[md.Static.Path].Telegram.MessageID, "Telegram MessageID does not match.")
	assert.Equal(t, md.Static.Title, mdMap[md.Static.Path].Static.Title, "Static Title does not match.")
}

func TestMetadata_Sync(t *testing.T) {
	tmpDir := t.TempDir()

	md := &Metadata{
		Telegram: telegramData{MessageID: 123, Date: 456},
		Static:   staticData{Title: "New Title", Path: "/content/my_article/"},
	}
	err := md.Sync(tmpDir)
	require.NoError(t, err)

	result, err := newMetafile(tmpDir)
	require.NoError(t, err)

	assert.Equal(t, md.Telegram.MessageID, result.Telegram.MessageID)
	assert.Equal(t, md.Static.Title, result.Static.Title)
	assert.Equal(t, md.Static.Path, result.Static.Path)
}
