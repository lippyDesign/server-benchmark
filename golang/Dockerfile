FROM golang:1.11.0-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get github.com/gin-gonic/gin

ENV SOURCES /Users/volodymyrlipunov/Desktop/server-benchmark/golang/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}golang
EXPOSE 5000
