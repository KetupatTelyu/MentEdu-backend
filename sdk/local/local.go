package local

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		BasePath: basePath,
	}
}

func (ls *LocalStorage) UploadFile(file *multipart.FileHeader, folder string) (string, error) {
	// create if not exist
	if _, err := os.Stat(filepath.Join(ls.BasePath, folder)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Join(ls.BasePath, folder), 0755); err != nil {
			return "", err
		}
	}

	src, err := file.Open()

	defer src.Close()

	destPath := filepath.Join(ls.BasePath, folder, file.Filename)
	dest, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer dest.Close()

	if _, err := io.Copy(dest, src); err != nil {
		return "", err
	}

	return destPath, nil
}

func (ls *LocalStorage) UploadSavedFile(filePath string, folder string) (string, error) {
	sourceFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer sourceFile.Close()

	destPath := filepath.Join(ls.BasePath, folder, filepath.Base(filePath))
	destFile, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return "", err
	}

	return destPath, nil
}

func (ls *LocalStorage) DeleteFile(filePath string) error {
	return os.Remove(filePath)
}
