package main

import (
	"bytes"
	"context"
	"my-todo-app/config"
	"my-todo-app/routers"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	db := config.SetupDatabase()
	router = routers.SetupRouter(db)
}

// AWS Lambda用のハンドラー
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	w := newResponseWriter()
	r, err := http.NewRequest(req.HTTPMethod, req.Path, bytes.NewReader([]byte(req.Body)))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	// Add headers from API Gateway request
	for key, value := range req.Headers {
		r.Header.Add(key, value)
	}

	// Set query parameters
	q := r.URL.Query()
	for key, value := range req.QueryStringParameters {
		q.Add(key, value)
	}
	r.URL.RawQuery = q.Encode()

	router.ServeHTTP(w, r)

	return events.APIGatewayProxyResponse{
		StatusCode: w.status,
		Body:       w.Body(),
		Headers:    w.header,
	}, nil
}

// ローカル環境でのエントリーポイント
func main() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		router.Run(":8080")
	} else {
		lambda.Start(Handler)
	}
}

// カスタムのResponseWriterを作成
type responseWriter struct {
	header http.Header
	status int
	body   *bytes.Buffer
}

func newResponseWriter() *responseWriter {
	return &responseWriter{
		header: http.Header{},
		body:   new(bytes.Buffer),
	}
}

func (w *responseWriter) Header() http.Header {
	return w.header
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

func (w *responseWriter) Write(body []byte) (int, error) {
	return w.body.Write(body)
}

func (w *responseWriter) Body() string {
	return w.body.String()
}
