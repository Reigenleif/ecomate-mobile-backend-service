package service

import (
	"context"
	"log"

	"github.com/Reigenleif/go-grpc-gauth-pg/api"
	"github.com/Reigenleif/go-grpc-gauth-pg/internal/token_service"
)

type EmojiService struct {
	go_grpc_gauth_pg.UnimplementedEmojiServer
}

func (s *EmojiService) GetEmoji(ctx context.Context, req *go_grpc_gauth_pg.GetEmojiRequest) (*go_grpc_gauth_pg.EmojiResponse, error) {
	userClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	log.Print(userClaims)

	return &go_grpc_gauth_pg.EmojiResponse{
		Id: "ax",
	}, nil
}

func (s *EmojiService) CreateEmoji(ctx context.Context, req *go_grpc_gauth_pg.CreateEmojiRequest) (*go_grpc_gauth_pg.EmojiResponse, error) {
	return &go_grpc_gauth_pg.EmojiResponse{
		Id: "ax",
	}, nil
}

func (s *EmojiService) UpdateEmoji(ctx context.Context, req *go_grpc_gauth_pg.UpdateEmojiRequest) (*go_grpc_gauth_pg.EmojiResponse, error) {

	return &go_grpc_gauth_pg.EmojiResponse{
		Id: "ax",
	}, nil
}

func (s *EmojiService) DeleteEmoji(ctx context.Context, req *go_grpc_gauth_pg.DeleteEmojiRequest) (*go_grpc_gauth_pg.EmojiResponse, error) {

	return &go_grpc_gauth_pg.EmojiResponse{
		Id: "ax",
	}, nil
}
