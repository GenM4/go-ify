FROM golang:1.23

EXPOSE 8080

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -C cmd/app -o go-ify 

CMD ["./cmd/app/go-ify"]
