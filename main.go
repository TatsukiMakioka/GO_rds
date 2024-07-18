package main

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"strings"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
	"my-todo-app/config"
	"my-todo-app/routers"
)

var router *gin.Engine

func init() {
	db := config.SetupDatabase()
	router = routers.SetupRouter(db)
}

// AWS Lambda handler
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	w := newResponseWriter()
	bodyReader := strings.NewReader(req.Body)

	r, err := http.NewRequest(req.HTTPMethod, req.Path, bodyReader)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	for key, value := range req.Headers {
		r.Header.Add(key, value)
	}

	q := r.URL.Query()
	for key, value := range req.QueryStringParameters {
		q.Add(key, value)
	}
	r.URL.RawQuery = q.Encode()

	router.ServeHTTP(w, r)

	return events.APIGatewayProxyResponse{
		StatusCode: w.status,
		Body:       w.body.String(),
		Headers:    w.Headers(),
	}, nil
}

func main() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		router.Run(":8080")
	} else {
		lambda.Start(Handler)
	}
}

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

func (w *responseWriter) Headers() map[string]string {
	headers := make(map[string]string)
	for key, values := range w.header {
		headers[key] = values[0]
	}
	return headers
}
