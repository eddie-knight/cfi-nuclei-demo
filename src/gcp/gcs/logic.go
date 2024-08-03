package main

import (
	"context"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

// ref: https://medium.com/google-cloud/golang-copy-to-gcs-check-bucket-58721285788e

var ctx context.Context

// TODO: This is not yet tested.
// createBucket creates a new bucket using the Google Cloud Storage client context
func createBucket(bucketName string) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	projectID := "CFI-" + uuid.New().String()
	bucket := client.Bucket(bucketName)
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		return err
	}

	return nil
}

// TODO: This is not yet tested.
// deleteBucket deletes an existing bucket using the Google Cloud Storage client context
func deleteBucket(bucketName string) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	if err := bucket.Delete(ctx); err != nil {
		return err
	}

	return nil
}

// TODO: This is not yet tested.
// uploadFile uploads a file to an existing bucket using the Google Cloud Storage client context
func uploadFile(bucketName, objectName, filePath string) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	object := bucket.Object(objectName)

	wc := object.NewWriter(ctx)
	if err := writeFileToGCS(wc, filePath); err != nil {
		return err
	}

	return nil
}

// TODO: This is not yet tested.
// writeFileToGCS writes a file to Google Cloud Storage
func writeFileToGCS(wc *storage.Writer, filePath string) error {
	defer wc.Close()

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = io.Copy(wc, file); err != nil {
		return err
	}

	return nil
}

func connect() {
	ctx = context.Background()
	// Use Google Application Default Credentials to authorize and authenticate the client.
	// More information about Application Default Credentials and how to enable is at
	// https://developers.google.com/identity/protocols/application-default-credentials.
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic("noctx")
	}
	// Use the client.

	// Close the client when finished.
	if err := client.Close(); err != nil {
		panic("noclose")
	}
}
