FROM golang:1.19 AS builder

WORKDIR $GOPATH/src/recommendations-grpc-service

COPY . ./

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM cgr.dev/chainguard/wolfi-base:latest

COPY --from=builder /go/src/recommendations-grpc-service/main ./

EXPOSE 30000

ENTRYPOINT ["./main"]
