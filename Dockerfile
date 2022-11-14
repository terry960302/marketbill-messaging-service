# aws go image 
FROM golang:1.19.3-debian-11-r4	
LABEL Terry Kim <terry960302@gmail.com>


WORKDIR /app

ADD go.mod /app
ADD go.sum /app

RUN apk update && \
    apk add git && \
    go get -d ./...

COPY . /app

EXPOSE 8000
ENV PROFILE=prod

RUN go build main.go
CMD ["./main"]
