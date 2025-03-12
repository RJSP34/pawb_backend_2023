FROM golang:alpine

ENV GIN_MODE=release
ENV PORT=3000

WORKDIR /go/src/GolangGinApi

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE $PORT

ENTRYPOINT ["./main"]