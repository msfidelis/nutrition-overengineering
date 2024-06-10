module github.com/msfidelis/health-api

go 1.15

require (
	github.com/Depado/ginprom v1.7.3
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-gonic/gin v1.7.7
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.9 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/google/uuid v1.3.0
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/msfidelis/gin-chaos-monkey v0.0.5
	github.com/nats-io/nats.go v1.35.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/rs/zerolog v1.22.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.16.1
	github.com/tkanos/gonfig v0.0.0-20181112185242-896f3d81fadf
	github.com/zcalusic/sysinfo v0.0.0-20200228145645-a159d7cc708b
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.31.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.31.0
	go.opentelemetry.io/otel v1.6.3
	go.opentelemetry.io/otel/exporters/jaeger v1.6.3
	go.opentelemetry.io/otel/sdk v1.6.3
	go.opentelemetry.io/otel/trace v1.6.3
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/tools v0.10.0 // indirect
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.26.0
)
