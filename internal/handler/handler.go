package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Store interface {
	Put(ctx context.Context, contentType string, body []byte) (string, error)
}

type Handler struct {
	store Store
}

func New(store Store) *Handler {
	return &Handler{store: store}
}

type response struct {
	Key string `json:"key"`
}

func (h *Handler) Handle(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	method := req.RequestContext.HTTP.Method
	slog.Info("request", "method", method, "sourceIP", req.RequestContext.HTTP.SourceIP)

	if method != http.MethodPost {
		slog.Warn("method not allowed", "method", method)
		return apiResp(http.StatusMethodNotAllowed, `{"error":"method not allowed"}`), nil
	}

	body, err := requestBody(req)
	if err != nil {
		slog.Error("failed to decode body", "error", err)
		return apiResp(http.StatusBadRequest, `{"error":"failed to read body"}`), nil
	}

	if len(body) == 0 {
		slog.Warn("empty body")
		return apiResp(http.StatusBadRequest, `{"error":"empty body"}`), nil
	}

	ct := req.Headers["content-type"]
	if ct == "" {
		ct = "application/octet-stream"
	}
	if i := strings.Index(ct, ";"); i != -1 {
		ct = strings.TrimSpace(ct[:i])
	}

	key, err := h.store.Put(ctx, ct, body)
	if err != nil {
		slog.Error("failed to store object", "error", err)
		return apiResp(http.StatusInternalServerError, `{"error":"failed to store data"}`), fmt.Errorf("store: %w", err)
	}

	slog.Info("stored", "key", key, "bytes", len(body), "contentType", ct)
	out, _ := json.Marshal(response{Key: key})
	return apiResp(http.StatusCreated, string(out)), nil
}

func requestBody(req events.APIGatewayV2HTTPRequest) ([]byte, error) {
	if req.IsBase64Encoded {
		return base64.StdEncoding.DecodeString(req.Body)
	}
	return []byte(req.Body), nil
}

func apiResp(status int, body string) events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}
}
