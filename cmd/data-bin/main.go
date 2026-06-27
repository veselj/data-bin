package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/veselj/data-bin/internal/handler"
	"github.com/veselj/data-bin/internal/storage"
)

func main() {
	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		log.Fatal("S3_BUCKET environment variable is required")
	}

	ctx := context.Background()
	store, err := storage.New(ctx, bucket)
	if err != nil {
		log.Fatalf("init storage: %v", err)
	}

	h := handler.New(store)
	lambda.Start(h.Handle)
}
