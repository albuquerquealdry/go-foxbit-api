
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -x
COPY src ./src
RUN ls /app
RUN go build -o foxbit-calc-api /app/src/cmd/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/foxbit-calc-api ./
EXPOSE 8000
CMD ["./foxbit-calc-api"]