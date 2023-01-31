FROM golang:latest

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
RUN chmod 777 main.sh
CMD ["/bin/sh", "./main.sh"]