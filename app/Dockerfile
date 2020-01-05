FROM golang:latest

RUN mkdir -p /go/src/github.com/sheepover96/poep_api/
WORKDIR /go/src/github.com/sheepover96/poep_api/app
ADD . /go/src/github.com/sheepover96/poep_api/app

RUN go get -u github.com/golang/dep/cmd/dep
#  && dep init && dep ensure
# CMD ["go", "run", "main.go"]
