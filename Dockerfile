FROM golang:1.21-bullseye AS builder
WORKDIR /app
RUN apt-get update && apt-get install -y curl unzip
COPY . /app
RUN go mod download
RUN ls -lrt
RUN CGO_ENABLED=0 go build -o main /app/cmd/server/main.go

FROM alpine
RUN apk add --update --no-cache ca-certificates tzdata && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD [ "./main" ]
