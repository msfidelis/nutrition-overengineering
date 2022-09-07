package bmr

import (
	context "context"
	"fmt"
	"os"
	"time"

	"github.com/msfidelis/health-api/pkg/logger"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func Call(ctx context.Context, gender string, weight float64, height float64, activity_intensity string, retries int, tracer trace.Tracer) (*Response, error) {

	var backoffSchedule = []time.Duration{
		1 * time.Second,
		3 * time.Second,
		10 * time.Second,
	}

	var resBMR *Response
	var err error

	log := logger.Instance()
	bmrEndpoint := os.Getenv("BMR_SERVICE_ENDPOINT")

	for i, backoff := range backoffSchedule {

		var conn *grpc.ClientConn
		conn, err = grpc.Dial(
			bmrEndpoint,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		)

		if err != nil {
			log.Error().
				Str("Service", "bmr").
				Str("Error", err.Error()).
				Msg("Failed to create gRPC Connection with BMR Service")
		}
		defer conn.Close()

		bmrClient := NewBMRServiceClient(conn)

		resBMR, err = bmrClient.SayHello(ctx, &Message{
			Gender:   gender,
			Weight:   weight,
			Height:   height,
			Activity: activity_intensity,
		})

		if err != nil {
			log.Error().
				Str("Service", "bmr").
				Str("Error", err.Error()).
				Msg("Failed to communicate with BMR Service")
		}

		if err == nil {
			break
		}

		log.Info().
			Str("Service", "bmr").
			Int("Retry", i+1).
			Str("Backoff", fmt.Sprintf("%s", backoff)).
			Msg("Failed to communicate with BMR Service")

		time.Sleep(backoff)
	}

	if err != nil {
		return nil, err
	}

	return resBMR, nil

	// var conn *grpc.ClientConn
	// conn, err = grpc.Dial(
	// 	bmrEndpoint,
	// 	grpc.WithInsecure(),
	// 	grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	// 	grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	// )

	// if err != nil {
	// 	log.Error().
	// 		Str("Service", "bmr").
	// 		Str("Error", err.Error()).
	// 		Msg("Failed to create gRPC Connection with BMR Service")

	// 	return nil, err
	// }
	// defer conn.Close()

	// bmrClient := NewBMRServiceClient(conn)

	// resBMR, err = bmrClient.SayHello(ctx, &Message{
	// 	Gender:   gender,
	// 	Weight:   weight,
	// 	Height:   height,
	// 	Activity: activity_intensity,
	// })

	// if err != nil {
	// 	log.Error().
	// 		Str("Service", "bmr").
	// 		Str("Error", err.Error()).
	// 		Msg("Failed to communicate with BMR Service")

	// 	return nil, err
	// }

	// return resBMR, nil

}
