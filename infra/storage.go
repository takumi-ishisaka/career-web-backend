package infra

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"github.com/nokin-all-of-career/career-web-backend/configs"
)

// Bucket : storage bucket connection
var Bucket *storage.BucketHandle

// CtxStorage : storage context
var CtxStorage context.Context

// NewStorageConnection : create storage connection
func NewStorageConnection() error {

	CtxStorage = context.Background()
	client, err := storage.NewClient(CtxStorage)
	defer client.Close()

	bucketPath := configs.Config.BucketPath
	if bucketPath == "" {
		bucketPath = os.Getenv("STORAGE_BUCKET_PATH")
	}
	Bucket = client.Bucket(bucketPath)

	return err
}
