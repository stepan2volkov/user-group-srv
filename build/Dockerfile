FROM golang:1.18 AS builder

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.6.0

WORKDIR /app
COPY . .
RUN go mod download
RUN make build

FROM scratch
WORKDIR /app
COPY --from=builder /app/build/user-group-srv .
COPY --from=builder /go/bin/goose .
COPY --from=builder /app/migrations /migrations

EXPOSE 8000
CMD ["./user-group-srv"]