FROM golang:1.14 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN go build -o api .

FROM alpine:3.7

RUN adduser -S -D -H -h /app api

USER api

COPY --from=builder /build/api /app/

WORKDIR /app

CMD ["./api"]