FROM golang:1.21-alpine AS builder

WORKDIR /workspace

RUN apk add --update --no-cache git && rm -rf /var/cache/apk/*
COPY go.mod go.sum /workspace/
RUN go mod download
COPY cmd /workspace/cmd
RUN go build -o server ./cmd/server

FROM alpine
RUN apk add --update --no-cache ca-certificates tzdata && rm -rf /var/cache/apk/*
COPY --from=builder /workspace/server /usr/local/bin/server
EXPOSE 8080
CMD [ "/usr/local/bin/server", "--server-stream-delay=500ms" ]
