package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsFileOrError(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "testfile.txt")

	require.NoError(t, os.WriteFile(tmpFile, []byte("test content"), 0644))
	err := IsFileOrError(tmpFile)
	assert.NoError(t, err)

	nonExistentFile := filepath.Join(tmpDir, "nonexistent.txt")
	err = IsFileOrError(nonExistentFile)
	assert.Error(t, err)

	err = IsFileOrError(tmpDir)
	assert.Error(t, err)
}

func TestIsDirOrError(t *testing.T) {
	tmpDir := t.TempDir()

	err := IsDirOrError(tmpDir)
	assert.NoError(t, err)

	nonExistentDir := filepath.Join(tmpDir, "nonexistentdir")
	err = IsDirOrError(nonExistentDir)
	assert.Error(t, err)

	tmpFile := filepath.Join(tmpDir, "testfile.txt")
	require.NoError(t, os.WriteFile(tmpFile, []byte("test content"), 0644))
	err = IsDirOrError(tmpFile)
	assert.Error(t, err)
}

func TestIsFileOrCreate(t *testing.T) {
	tmpDir := t.TempDir()

	tmpFile := filepath.Join(tmpDir, "testfile.txt")
	require.NoError(t, os.WriteFile(tmpFile, []byte("test content"), 0644))
	err := IsFileOrCreate(tmpFile)
	assert.NoError(t, err)

	nonExistentFile := filepath.Join(tmpDir, "newfile.txt")
	err = IsFileOrCreate(nonExistentFile)
	assert.NoError(t, err)
	_, err = os.Stat(nonExistentFile)
	assert.NoError(t, err)

	err = IsFileOrCreate(tmpDir)
	assert.Error(t, err)
}

func TestIsDirOrCreate(t *testing.T) {
	tmpDir := t.TempDir()

	err := IsDirOrCreate(tmpDir)
	assert.NoError(t, err)

	nonExistentDir := filepath.Join(tmpDir, "newdir")
	err = IsDirOrCreate(nonExistentDir)
	assert.NoError(t, err)
	_, err = os.Stat(nonExistentDir)
	assert.NoError(t, err)

	tmpFile := filepath.Join(tmpDir, "testfile.txt")
	require.NoError(t, os.WriteFile(tmpFile, []byte("test content"), 0644))
	err = IsDirOrCreate(tmpFile)
	assert.Error(t, err)
}

func TestReadFile(t *testing.T) {
	tmpDir := t.TempDir()

	tmpFile := filepath.Join(tmpDir, "testfile.txt")
	expectedContent := "test content"
	require.NoError(t, os.WriteFile(tmpFile, []byte(expectedContent), 0644))

	content, err := ReadFile(tmpFile)
	assert.NoError(t, err)
	assert.Equal(t, expectedContent, string(content))

	nonExistentFile := filepath.Join(tmpDir, "nonexistent.txt")
	_, err = ReadFile(nonExistentFile)
	assert.Error(t, err)

	_, err = ReadFile(tmpDir)
	assert.Error(t, err)
}

func TestWriteToFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "testfile.txt")

	err := WriteToFile(tmpFile, []byte("test content"))
	assert.NoError(t, err)

	content, err := os.ReadFile(tmpFile)
	require.NoError(t, err)
	assert.Equal(t, "test content", string(content))
}

func TestCopyFile(t *testing.T) {
	srcDir := t.TempDir()
	dstDir := t.TempDir()

	srcFile := filepath.Join(srcDir, "testfile.txt")
	require.NoError(t, os.WriteFile(srcFile, []byte("file content"), 0644))
	srcFileInfo, err := os.Stat(srcFile)
	require.NoError(t, err)

	dstFile := filepath.Join(dstDir, "copiedfile.txt")
	err = copyFile(srcFile, dstFile, srcFileInfo)
	assert.NoError(t, err)

	content, err := os.ReadFile(dstFile)
	require.NoError(t, err)
	assert.Equal(t, "file content", string(content))
}

func TestCopyDir(t *testing.T) {
	srcDir := t.TempDir()
	dstDir := filepath.Join(t.TempDir(), "dst")

	srcSubDir := filepath.Join(srcDir, "subdir")
	require.NoError(t, os.Mkdir(srcSubDir, 0755))
	srcFile := filepath.Join(srcSubDir, "testfile.txt")
	require.NoError(t, os.WriteFile(srcFile, []byte("hello subdir"), 0644))

	err := copyDir(srcDir, dstDir)
	assert.NoError(t, err)

	dstFile := filepath.Join(dstDir, "subdir", "testfile.txt")
	content, err := os.ReadFile(dstFile)
	require.NoError(t, err)
	assert.Equal(t, "hello subdir", string(content))
}

func TestCopy(t *testing.T) {
	srcDir := t.TempDir()
	dstDir := t.TempDir()

	srcFile := filepath.Join(srcDir, "testfile.txt")
	require.NoError(t, os.WriteFile(srcFile, []byte("hello world"), 0644))
	dstFile := filepath.Join(dstDir, "testfile.txt")
	err := Copy(srcFile, dstFile)
	assert.NoError(t, err)

	content, err := os.ReadFile(dstFile)
	require.NoError(t, err)
	assert.Equal(t, "hello world", string(content))

	srcSubDir := filepath.Join(srcDir, "subdir")
	require.NoError(t, os.Mkdir(srcSubDir, 0755))
	srcSubFile := filepath.Join(srcSubDir, "subfile.txt")
	require.NoError(t, os.WriteFile(srcSubFile, []byte("subcontent"), 0644))
	dstSubDir := filepath.Join(dstDir, "subdir")
	err = Copy(srcDir, dstDir)
	assert.NoError(t, err)

	dstSubFile := filepath.Join(dstSubDir, "subfile.txt")
	content, err = os.ReadFile(dstSubFile)
	require.NoError(t, err)
	assert.Equal(t, "subcontent", string(content))
}
