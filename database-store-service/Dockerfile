FROM golang:1.19 AS builder

WORKDIR $GOPATH/src/database-store-service

COPY . ./

RUN go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
RUN swag init

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:3.12.3

COPY --from=builder /go/src/database-store-service/main ./
COPY --from=builder /go/src/database-store-service/configs ./configs

EXPOSE 8080

ENTRYPOINT ["./main"]
