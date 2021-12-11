FROM golang:1.17.5-alpine AS builder

COPY . /go/src/crud
WORKDIR /go/src/crud/

RUN go get github.com/gorilla/mux

RUN GOOS=linux GOARCH=amd64 go build -o ./crud ./main.go

FROM alpine
COPY --from=builder /go/src/crud .

EXPOSE 8000

ENTRYPOINT ["/crud"]