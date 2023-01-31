FROM golang:latest

WORKDIR /app
COPY . .
RUN go mod download
CMD ["/bin/sh", "go", "run", "./main.go"]