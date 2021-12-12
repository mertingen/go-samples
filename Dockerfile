FROM golang:1.17.5-alpine AS builder

COPY . /go/src/crud
WORKDIR /go/src/crud/

RUN go get github.com/gorilla/mux && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/gobuffalo/pop/... && \
    go install github.com/gobuffalo/pop/soda

RUN GOOS=linux GOARCH=amd64 go build -o ./crud ./main.go

FROM alpine
COPY --from=builder /go/src/crud .

EXPOSE 8000

ENTRYPOINT ["/crud"]