package models

import ("google.golang.org/grpc/codes"
"google.golang.org/grpc/status"
)

var (
	ErrUnauthorized = status.Errorf(codes.Unauthenticated, "UNAUTHORIZED")
)