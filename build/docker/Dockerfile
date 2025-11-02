FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download 

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -C cmd/app -o go-ify 

FROM alpine:latest
WORKDIR /build
COPY web ./web
COPY config ./config
COPY --from=builder build/cmd/app/go-ify .

EXPOSE 8080

CMD ["./go-ify"]
