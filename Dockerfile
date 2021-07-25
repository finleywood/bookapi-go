FROM golang:1.16.6-alpine

WORKDIR /go/src/app
COPY . .
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "main.go"]