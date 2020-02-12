FROM golang:1.13-alpine

RUN apk add git

RUN go get github.com/cespare/reflex
RUN export PATH=$PATH:$GOPATH/bin

WORKDIR /src/app

RUN ls -la /src/app

CMD reflex -r '\.go$' -s -- sh -c "go run main.go"
# RUN reflex -r '\.go$' -s -- sh -c "go build -o main . && /app/main"
# RUN reflex -r '\.go$' -s -- sh -c "go build -o main . && /app/main"
# RUN go run main.go


# CMD ["go", "run", "main.go"]
# CMD ["reflex", "-r", "'\.go$'", "-s", "--", "sh", "c", "'go run main.go'"]
# CMD ["reflex", "-r", "\.go$", "-s", "--", "sh", "c", "go run main.go"]