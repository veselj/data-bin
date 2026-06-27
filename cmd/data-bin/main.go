package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/veselj/data-bin/internal/handler"
	"github.com/veselj/data-bin/internal/storage"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		slog.Error("S3_BUCKET environment variable is required")
		os.Exit(1)
	}

	ctx := context.Background()
	store, err := storage.New(ctx, bucket)
	if err != nil {
		slog.Error("init storage", "error", err)
		os.Exit(1)
	}

	slog.Info("starting", "bucket", bucket)
	h := handler.New(store)
	lambda.Start(h.Handle)
}
