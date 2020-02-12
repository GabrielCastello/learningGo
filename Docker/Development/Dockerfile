FROM golang:1.13-alpine

RUN apk add git

RUN go get github.com/cespare/reflex
RUN go get github.com/cespare/reflex
RUN export PATH=$PATH:$GOPATH/bin

RUN mkdir /app
ADD . /app
WORKDIR /app
# RUN reflex -r '\.go$' -s -- sh -c "go run main.go"
RUN go run main.go


# CMD ["/app/main"]