FROM golang:1.10.2

WORKDIR /go/src
COPY . .

EXPOSE 6000

RUN go get -d -v ./...
RUN go build -o /go/src/app

CMD ["/go/src/app"]