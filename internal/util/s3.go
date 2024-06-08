package util

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func UploadFileToS3(file *multipart.FileHeader, filename string, prefix string) (string, string, error) {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return "", "", err
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	// Get filename from filepath and add timestamp for uniqueness (optional)
	key := fmt.Sprintf("%s/%s", prefix, filename)

	fileBody, err := file.Open()
	if err != nil {
		log.Fatalf("unable to open file, %v", err)
		return "", "", err
	}
	defer fileBody.Close()

	output, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String("peluang-images"),
		Key:    aws.String(key),
		Body:   fileBody,
	})
	if err != nil {
		log.Fatalf("unable to upload file, %v", err)
		return "", "", err
	}
	return output.Location, *output.Key, nil
}

func DeleteFileFromS3(key string) error {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return err
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String("peluang-images"),
		Delete: &types.Delete{
			Objects: []types.ObjectIdentifier{
				{
					Key: aws.String(key),
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("unable to delete file: %w", err)
	}
	return nil
}
