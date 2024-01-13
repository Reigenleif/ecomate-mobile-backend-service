package service

import (
	"context"

	"github.com/Reigenleif/ecomate-mobile-backend-service/internal/db"
	"github.com/Reigenleif/ecomate-mobile-backend-service/internal/token_service"
	"github.com/Reigenleif/ecomate-mobile-backend-service/models"
	"github.com/Reigenleif/ecomate-mobile-backend-service/proto"

	"github.com/gobuffalo/nulls"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	proto.UnimplementedUserServiceServer
}

func (s *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	var userDb models.User
	err := db.GetDB().QueryRow(ctx, "SELECT * FROM public.\"User\" WHERE id = '"+req.Id+"';").Scan(&userDb.ID, &userDb.Name, &userDb.Email, nil, nil, &userDb.ImageUrl, &userDb.Role)
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Id:       userDb.ID.String,
		Name:     userDb.Name.String,
		ImageUrl: userDb.ImageUrl.String,
		Role:     userDb.Role.String,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UserResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	if UserClaims.Role != "ADMIN" && UserClaims.Id != req.Id {
		// Case when user is not admin and is trying to update another user
		return nil, models.ErrUnauthorized
	}

	var userDb models.User
	err = db.GetDB().QueryRow(ctx, "SELECT * FROM public.\"User\" WHERE id = '"+req.Id+"';").Scan(&userDb.ID, &userDb.Name, &userDb.Email, nil, nil, &userDb.ImageUrl, &userDb.Role)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		userDb.Name = nulls.NewString(*req.Name)
	}

	if req.ImageUrl != nil {
		userDb.ImageUrl = nulls.NewString(*req.ImageUrl)
	}

	if req.Password != nil {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), 12)
		if err != nil {
			return nil, err
		}
		userDb.HashPassword = nulls.NewString(string(hashPassword))
	}

	_, err = db.GetDB().Exec(ctx, "UPDATE public.\"User\" SET name = $1, \"imageUrl\" = $2, \"hashPassword\" = $3 WHERE id = $4", userDb.Name, userDb.ImageUrl, userDb.HashPassword, userDb.ID)
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Id:       userDb.ID.String,
		Name:     userDb.Name.String,
		ImageUrl: userDb.ImageUrl.String,
		Role:     userDb.Role.String,
	}, nil
}
