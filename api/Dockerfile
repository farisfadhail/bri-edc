
FROM golang:1.25.1-alpine AS builder

RUN apk add --no-cache git bash build-base

RUN go install github.com/google/wire/cmd/wire@latest

WORKDIR /app
COPY . .

RUN go mod tidy

RUN cd internal/injector && wire

RUN go build -o api ./

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/api .

ENV PORT=8080
EXPOSE 8080

CMD ["./api"]