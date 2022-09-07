package imc

import (
	context "context"
	"fmt"
	"os"
	"time"

	"github.com/msfidelis/health-api/pkg/logger"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func Call(ctx context.Context, weight float64, height float64, tracer trace.Tracer) (*Response, error) {

	var backoffSchedule = []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
	}

	var resGrpc *Response
	var err error

	log := logger.Instance()
	imcEndpoint := os.Getenv("IMC_SERVICE_ENDPOINT")

	for i, backoff := range backoffSchedule {

		ctxCall, spanCall := tracer.Start(ctx, fmt.Sprintf("IMC call attempt %v", i+1))
		defer spanCall.End()

		var conn *grpc.ClientConn
		conn, err = grpc.Dial(
			imcEndpoint,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		)

		if err != nil {
			log.Error().
				Str("Service", "imc").
				Str("Error", err.Error()).
				Msg("Failed to create gRPC Connection with imc Service")

			spanCall.SetAttributes(
				attribute.String("Service", "imc"),
				attribute.String("gRPC connection error", err.Error()),
			)
		}
		defer conn.Close()

		grpcClient := NewIMCServiceClient(conn)

		resGrpc, err = grpcClient.SayHello(ctxCall, &Message{
			Weight: weight,
			Height: height,
		})

		if err != nil {
			log.Error().
				Str("Service", "imc").
				Str("Error", err.Error()).
				Msg("Failed to communicate with imc Service")

			spanCall.SetAttributes(
				attribute.String("Service", "imc"),
				attribute.String("gRPC call error", err.Error()),
			)
		}

		if err == nil {
			break
		}

		log.Info().
			Str("Service", "imc").
			Int("Retry", i+1).
			Str("Backoff", fmt.Sprintf("%s", backoff)).
			Msg("Failed to communicate with imc Service")

		spanCall.SetAttributes(
			attribute.String("Service", "imc"),
			attribute.Int("Attempts", i+1),
		)

		time.Sleep(backoff)
	}

	if err != nil {
		return nil, err
	}

	return resGrpc, nil

}
