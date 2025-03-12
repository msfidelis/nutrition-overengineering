FROM golang:1.23 AS builder

WORKDIR $GOPATH/src/health-data-offload

COPY . ./

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM cgr.dev/chainguard/wolfi-base:latest

COPY --from=builder /go/src/health-data-offload/main ./

EXPOSE 8080

ENTRYPOINT ["./main"]
