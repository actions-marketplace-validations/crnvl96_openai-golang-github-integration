FROM golang:1.19-alpine3.16

WORKDIR /app
COPY . .
RUN go mod download
RUN go build .
RUN chmod +x main.sh
ENTRYPOINT ["./app/main.sh"]