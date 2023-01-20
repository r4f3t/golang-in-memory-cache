FROM golang:1.18

RUN mkdir -p /go/src/github.com/r4f3t/golang-in-memory-cache
#RUN CGO_ENABLED=0
#RUN GOOS=linux

ENV GOPATH /go
WORKDIR /go/src/github.com/r4f3t/golang-in-memory-cache

ADD go.mod /go/src/github.com/r4f3t/golang-in-memory-cache
ADD go.sum /go/src/github.com/r4f3t/golang-in-memory-cache
RUN go mod download
ADD . /go/src/github.com/r4f3t/golang-in-memory-cache

#RUN go get
#RUN go build


#RUN go build

RUN go build  -o /go-docker-demo


EXPOSE 8080

ENTRYPOINT [ "go","run","main.go","api", "-c ./config/config.qa.json", "-p 5558" ]

