package proteins

import (
	context "context"
	"fmt"
	"os"
	"time"

	"recommendations-grpc-service/pkg/logger"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func Call(ctx context.Context, weight float64, tracer trace.Tracer) (*Response, error) {

	var backoffSchedule = []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
	}

	var resGrpc *Response
	var err error

	log := logger.Instance()
	proteinsEndpoint := os.Getenv("PROTEINS_SERVICE_ENDPOINT")

	for i, backoff := range backoffSchedule {

		ctxCall, spanCall := tracer.Start(ctx, fmt.Sprintf("proteins call attempt %v", i+1))
		defer spanCall.End()

		var conn *grpc.ClientConn
		conn, err = grpc.Dial(
			proteinsEndpoint,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		)

		if err != nil {
			log.Error().
				Str("Service", "proteins").
				Str("Error", err.Error()).
				Msg("Failed to create gRPC Connection with proteins Service")

			spanCall.SetAttributes(
				attribute.String("Service", "proteins"),
				attribute.String("gRPC connection error", err.Error()),
			)
		}
		defer conn.Close()

		grpcClient := NewProteinsServiceClient(conn)

		resGrpc, err = grpcClient.SayHello(ctxCall, &Message{
			Weight: weight,
		})

		if err != nil {
			log.Error().
				Str("Service", "proteins").
				Str("Error", err.Error()).
				Msg("Failed to communicate with proteins Service")

			spanCall.SetAttributes(
				attribute.String("Service", "proteins"),
				attribute.String("gRPC call error", err.Error()),
			)
		}

		if err == nil {
			break
		}

		log.Info().
			Str("Service", "proteins").
			Int("Retry", i+1).
			Str("Backoff", fmt.Sprintf("%s", backoff)).
			Msg("Failed to communicate with proteins Service")

		spanCall.SetAttributes(
			attribute.String("Service", "proteins"),
			attribute.Int("Attempts", i+1),
		)

		time.Sleep(backoff)
	}

	if err != nil {
		return nil, err
	}

	return resGrpc, nil

}
