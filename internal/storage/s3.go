package storage

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type S3Store struct {
	client *s3.Client
	bucket string
}

func New(ctx context.Context, bucket string) (*S3Store, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("load aws config: %w", err)
	}
	return &S3Store{
		client: s3.NewFromConfig(cfg),
		bucket: bucket,
	}, nil
}

// Put stores body in S3 and returns the object key.
func (s *S3Store) Put(ctx context.Context, contentType string, body []byte) (string, error) {
	now := time.Now().UTC()
	key := fmt.Sprintf("%d_%02d_%02d_%s", now.Year(), now.Month(), now.Day(), uuid.New().String())

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(body),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("put object: %w", err)
	}
	return key, nil
}
