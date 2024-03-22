package content

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewContent(t *testing.T) {
	tmpDir := t.TempDir()
	contentFiles := []string{"file1.txt", "file2.txt"}

	fullPaths := make([]string, 0, len(contentFiles))
	for _, f := range contentFiles {
		filePath := filepath.Join(tmpDir, f)
		require.NoError(t, os.WriteFile(filePath, []byte("test"), 0644))
		fullPaths = append(fullPaths, filePath)
	}

	metaData := Metadata{
		Telegram: telegramData{
			MessageID: 1,
			Date:      20230101,
		},
		Static: staticData{
			Title: "Test Title",
			Url:   "http://example.com",
		},
	}
	metaDataBytes, err := json.Marshal(metaData)
	require.NoError(t, err)

	metaFilePath := filepath.Join(tmpDir, metafileName)
	require.NoError(t, os.WriteFile(metaFilePath, metaDataBytes, 0644))

	newContent, err := New(tmpDir, fullPaths)
	require.NoError(t, err, "Creating new content should not fail")

	contentObj, ok := newContent.(*content)
	require.True(t, ok, "Expected newContent to be of type *content")

	require.Equal(t, len(contentFiles), len(contentObj.contentFiles), "All content files should be added")
	require.NotNil(t, contentObj.Metadata, "Metadata should be loaded")
}
