FROM amazonlinux:2

RUN yum update -y && \
    yum install -y golang gcc zlib-devel make

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bootstrap main.go
