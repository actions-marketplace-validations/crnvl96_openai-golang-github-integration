FROM golang:1.19-alpine3.16

WORKDIR /app
COPY . .
RUN go mod download
RUN go build .
COPY ./main.sh .
RUN chmod +x main.sh
ENTRYPOINT ["/bin/sh", "/app/main.sh"]