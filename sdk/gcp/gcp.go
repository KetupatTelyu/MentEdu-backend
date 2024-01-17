package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type GoogleCloudStorage struct {
	client *storage.Client
	bucket *storage.BucketHandle
}

func NewGoogleCloudStorage(projectID, bucketName, apiKey string) (*GoogleCloudStorage, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Google Cloud Storage client: %v", err)
	}

	bucket := client.Bucket(bucketName)

	return &GoogleCloudStorage{
		client: client,
		bucket: bucket,
	}, nil
}

func (gcs *GoogleCloudStorage) createFolderIfNotExists(folder string) error {
	obj := gcs.bucket.Object(folder + "/")

	wc := obj.NewWriter(context.Background())
	if err := wc.Close(); err != nil {
		return fmt.Errorf("failed to create folder: %v", err)
	}

	return nil
}

func (gcs *GoogleCloudStorage) UploadFile(file *multipart.FileHeader, folder string) (string, error) {
	if err := gcs.createFolderIfNotExists(folder); err != nil {
		return "", err
	}

	objectName := folder + "/" + file.Filename
	wc := gcs.bucket.Object(objectName).NewWriter(context.Background())

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	_, err = io.Copy(wc, src)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	return objectName, nil
}

func (gcs *GoogleCloudStorage) UploadSavedFile(filePath string, folder string) (string, error) {
	if err := gcs.createFolderIfNotExists(folder); err != nil {
		return "", err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	objectName := folder + "/" + filepath.Base(filePath)
	wc := gcs.bucket.Object(objectName).NewWriter(context.Background())

	_, err = io.Copy(wc, file)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	return objectName, nil
}

func (gcs *GoogleCloudStorage) DeleteFile(filePath string) error {
	obj := gcs.bucket.Object(filePath)
	return obj.Delete(context.Background())
}

func (gcs *GoogleCloudStorage) Close() {
	gcs.client.Close()
}
