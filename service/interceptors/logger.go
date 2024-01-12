package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log.Print("Request: ", req)
	return handler(ctx, req)
}