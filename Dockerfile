FROM golang:1.19-alpine3.16

COPY . .
RUN go get ./...
CMD ["/main.sh"]