package bmr

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
	_ "google.golang.org/grpc/balancer/roundrobin" // Import round robin balancer
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func Call(ctx context.Context, gender string, weight float64, height float64, activity_intensity string, tracer trace.Tracer) (*Response, error) {

	var backoffSchedule = []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
	}

	var resGrpc *Response
	var err error

	log := logger.Instance()
	bmrEndpoint := os.Getenv("BMR_SERVICE_ENDPOINT")

	for i, backoff := range backoffSchedule {

		ctxCall, spanCall := tracer.Start(ctx, fmt.Sprintf("BMR call attempt %v", i+1))
		defer spanCall.End()

		var conn *grpc.ClientConn

		conn, err = grpc.NewClient(bmrEndpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
			grpc.WithDefaultServiceConfig(`{
				"loadBalancingPolicy":"round_robin",
				"methodConfig": [{
					"name": [{"service": ""}],
					"waitForReady": true,
					"retryPolicy": {
						"maxAttempts": 3,
						"initialBackoff": "0.1s",
						"maxBackoff": "1s",
						"backoffMultiplier": 2,
						"retryableStatusCodes": ["UNAVAILABLE", "DEADLINE_EXCEEDED"]
					}
				}]
			}`),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                10 * time.Second,
				Timeout:             3 * time.Second,
				PermitWithoutStream: true,
			}),
			grpc.WithInitialWindowSize(1<<20),     // 1MB
			grpc.WithInitialConnWindowSize(1<<20), // 1MB
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(4<<20), // 4MB
				grpc.MaxCallSendMsgSize(4<<20), // 4MB
			),
		)

		if err != nil {
			log.Error().
				Str("Service", "bmr").
				Str("traceID", spanCall.SpanContext().TraceID().String()).
				Str("Endpoint", bmrEndpoint).
				Str("Error", err.Error()).
				Msg("Failed to create gRPC Connection with BMR Service")

			spanCall.SetAttributes(
				attribute.String("Service", "BMR"),
				attribute.String("gRPC connection error", err.Error()),
			)
		}
		defer conn.Close()

		grpcClient := NewBMRServiceClient(conn)

		resGrpc, err = grpcClient.SayHello(ctxCall, &Message{
			Gender:   gender,
			Weight:   weight,
			Height:   height,
			Activity: activity_intensity,
		})

		if err != nil {
			log.Error().
				Str("Service", "bmr").
				Str("traceID", spanCall.SpanContext().TraceID().String()).
				Str("Endpoint", bmrEndpoint).
				Str("Error", err.Error()).
				Msg("Failed to communicate with BMR Service")

			spanCall.SetAttributes(
				attribute.String("Service", "BMR"),
				attribute.String("gRPC call error", err.Error()),
			)
		}

		if err == nil {
			break
		}

		log.Info().
			Str("Service", "bmr").
			Str("traceID", spanCall.SpanContext().TraceID().String()).
			Int("Retry", i+1).
			Str("Backoff", fmt.Sprintf("%s", backoff)).
			Msg("Failed to communicate with BMR Service")

		spanCall.SetAttributes(
			attribute.String("Service", "BMR"),
			attribute.Int("Attempts", i+1),
		)

		time.Sleep(backoff)
	}

	if err != nil {
		return nil, err
	}

	return resGrpc, nil

}
