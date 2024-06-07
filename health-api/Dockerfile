FROM golang:1.19 AS builder

WORKDIR $GOPATH/src/health-api

COPY . ./

# RUN go get -u github.com/swaggo/swag/cmd/swag@v1.6.7
# RUN swag init

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM cgr.dev/chainguard/wolfi-base:latest

COPY --from=builder /go/src/health-api/main ./
COPY --from=builder /go/src/health-api/configs ./configs

EXPOSE 8080

ENTRYPOINT ["./main"]
