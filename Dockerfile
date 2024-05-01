FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o server-bot ./cmd/main.go

CMD ["./cmd"]
