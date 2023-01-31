FROM golang:1.19-alpine3.16

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build
RUN chmod 777 script.sh
ENTRYPOINT ["/bin/sh", "./app/script.sh"]