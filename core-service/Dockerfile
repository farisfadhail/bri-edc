
FROM golang:1.25.1-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o core ./

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/core .

ENV CORE_PORT=9000
EXPOSE 9000

CMD ["./core"]
