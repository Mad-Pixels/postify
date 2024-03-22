package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// IsFileOrError checks if the given path exists and is a file.
func IsFileOrError(path string) error {
	exist, isDir, err := statPath(path)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("file %s doesn't exist", path)
	}
	if isDir {
		return fmt.Errorf("%s is not a file", path)
	}
	return nil
}

// IsFileOrCreate checks if the given path exists and is a file. If the file does not exist, it attempts to create it.
func IsFileOrCreate(path string) error {
	exist, isDir, err := statPath(path)
	if err != nil {
		return err
	}
	if exist && isDir {
		return fmt.Errorf("%s exists and is a directory", path)
	}
	if exist {
		return nil
	}
	if err := IsDirOrCreate(filepath.Dir(path)); err != nil {
		return fmt.Errorf("failed to create directories for %s: %w", path, err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer file.Close()
	return nil
}

// IsDirOrError checks if the given path exists and is a directory.
func IsDirOrError(path string) error {
	exist, isDir, err := statPath(path)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("directory %s doesn't exist", path)
	}
	if !isDir {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}

// IsDirOrCreate checks if the given path exists and is a directory. If the directory does not exist, it attempts to create it.
func IsDirOrCreate(path string) error {
	exist, isDir, err := statPath(path)
	if err != nil {
		return err
	}
	if exist && !isDir {
		return fmt.Errorf("%s exists and is not a directory", path)
	}
	if exist {
		return nil
	}
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

// ReadFile reads the content of a file specified by the path.
func ReadFile(path string) ([]byte, error) {
	err := IsFileOrError(path)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s for reading: %w", path, err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read from file %s: %w", path, err)
	}
	return content, nil
}

// WriteToFile writes the given data to a file at the specified path. If the file does not exist, it is created.
func WriteToFile(path string, data []byte) error {
	if err := IsFileOrCreate(path); err != nil {
		return fmt.Errorf("could not create file if not exists %s: %w", path, err)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("could not open file %s: %w", path, err)
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("could not write to file %s: %w", path, err)
	}
	return nil
}

// Copy files and directories from source to destination.
func Copy(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if srcInfo.IsDir() {
		return copyDir(src, dst)
	}
	return copyFile(src, dst, srcInfo)
}

func copyFile(src, dst string, srcInfo os.FileInfo) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, srcInfo.Mode())
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	if err = os.Chtimes(dst, srcInfo.ModTime(), srcInfo.ModTime()); err != nil {
		return err
	}
	return nil
}

func copyDir(srcDir, dstDir string) error {
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(dstDir, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			info, err := entry.Info()
			if err != nil {
				return err
			}
			if err := copyFile(srcPath, dstPath, info); err != nil {
				return err
			}
		}
	}
	return nil
}

func statPath(path string) (exist bool, isDir bool, err error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, false, nil
		}
		return false, false, err
	}
	return true, info.IsDir(), nil
}
