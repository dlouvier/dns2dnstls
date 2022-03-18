FROM golang:1.11
WORKDIR /go/src/github.com/dlouvier/dns2dnstls
COPY . . 
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.15
WORKDIR /root/
COPY --from=0 /go/src/github.com/dlouvier/dns2dnstls/app .
CMD ["./app"]  