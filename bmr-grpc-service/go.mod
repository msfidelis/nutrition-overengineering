module bmr-grpc-service

go 1.21

toolchain go1.22.4

require (
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.52.0
	go.opentelemetry.io/otel v1.27.0
	go.opentelemetry.io/otel/exporters/zipkin v1.27.0
	go.opentelemetry.io/otel/sdk v1.27.0
	google.golang.org/grpc v1.64.0
)

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240520151616-dc85e6b867a5 // indirect
)

require (
	github.com/rs/zerolog v1.26.1
	golang.org/x/net v0.25.0
	google.golang.org/protobuf v1.34.1
)
