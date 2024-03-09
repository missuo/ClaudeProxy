FROM golang:1.22 AS builder
WORKDIR /go/src/github.com/missuo/claude-proxy
COPY main.go ./
COPY go.mod ./
COPY go.sum ./
RUN go get -d -v ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o claude-proxy .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/github.com/missuo/claude-proxy/claude-proxy /app/claude-proxy
CMD ["/app/claude-proxy"]