FROM golang:1.17.5-alpine AS builder

COPY . /go/src/fiber-crud
WORKDIR /go/src/fiber-crud/

RUN go get github.com/gofiber/fiber/v2 && \
    go get gorm.io/gorm && \
    go get gorm.io/driver/mysql

RUN GOOS=linux GOARCH=amd64 go build -o ./fiber-crud ./main.go

FROM alpine
COPY --from=builder /go/src/fiber-crud .

EXPOSE 8000

ENTRYPOINT ["/fiber-crud"]