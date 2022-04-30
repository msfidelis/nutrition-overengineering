FROM golang:1.17 AS builder

WORKDIR $GOPATH/src/calories-grpc-service

COPY . ./

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:3.12.3

COPY --from=builder /go/src/calories-grpc-service/main ./

EXPOSE 30000

ENTRYPOINT ["./main"]
