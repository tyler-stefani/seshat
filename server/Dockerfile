FROM golang:1.19-alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o ./server ./main.go

FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
