FROM golang:1.11
COPY . /tmp/

WORKDIR /go/src/app
COPY ./docker_test/main.go .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
ENTRYPOINT ["app", "-f=7", "-s=9"]