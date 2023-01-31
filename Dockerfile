FROM golang:1.19-alpine3.16

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build
RUN chmod +x main.sh
RUN ls
ENTRYPOINT ["/bin/sh", "main.sh"]