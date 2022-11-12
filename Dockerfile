FROM golang:alpine3.16	
LABEL Terry Kim <terry960302@gmail.com>


WORKDIR /app

ADD go.mod /app
ADD go.sum /app

RUN apk update && \
    apk add git && \
    go get -d ./...

COPY . /app

EXPOSE 80

RUN go build main.go
CMD ["./main"]
