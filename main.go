package main

import (
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
	r, err := http.NewRequest(req.HTTPMethod, req.Path, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}
	router.ServeHTTP(w, r)
	return events.APIGatewayProxyResponse{
		StatusCode: w.status,
		Body:       w.body,
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
	gin.ResponseWriter
	header http.Header
	status int
	body   string
}

func newResponseWriter() *responseWriter {
	return &responseWriter{
		header: http.Header{},
	}
}

func (w *responseWriter) Header() http.Header {
	return w.header
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

func (w *responseWriter) Write(body []byte) (int, error) {
	w.body = string(body)
	return len(body), nil
}
