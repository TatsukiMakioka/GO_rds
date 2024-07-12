package main

import (
    "context"
    "fmt"
    "my-todo-app/config"
    "my-todo-app/routers"
    "net/http"
    "os"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/gin-gonic/gin"
    "log"
)

var router *gin.Engine

func init() {
    db := config.SetupDatabase()
    router = routers.SetupRouter(db)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    log.Print("Handler invoked with:", req)

    w := newResponseWriter()
    r, err := http.NewRequest(req.HTTPMethod, req.Path, nil)
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       err.Error(),
        }, nil
    }
    router.ServeHTTP(w, r)

    log.Printf("Response: %d - %s", w.status, w.body)

    return events.APIGatewayProxyResponse{
        StatusCode: w.status,
        Body:       w.body,
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
