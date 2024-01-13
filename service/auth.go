package service

import (
	"context"

	"github.com/Reigenleif/ecomate-mobile-backend-service/internal/db"
	"github.com/Reigenleif/ecomate-mobile-backend-service/internal/token_service"
	"github.com/Reigenleif/ecomate-mobile-backend-service/models"
	"github.com/Reigenleif/ecomate-mobile-backend-service/proto"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.UserResponse, error) {
	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"User\" WHERE email = $1", req.Email)
	if err != nil {
		return nil, err
	}

	var user models.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, nil, &user.HashPassword, &user.ImageUrl, &user.Role)
		if err != nil {
			return nil, err
		}
	}

	errWrongPassword := bcrypt.CompareHashAndPassword([]byte(user.HashPassword.String), []byte(req.Password))
	if errWrongPassword != nil {
		return nil, errWrongPassword
	}

	uc := token_service.UserClaims{
		Id:    user.ID.String,
		Email: user.Email.String,
		Role:  user.Role.String,
	}

	accessToken, err := token_service.NewAccessToken(uc)
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Id:          user.ID.String,
		Name:        user.Name.String,
		Email:       user.Email.String,
		ImageUrl:    user.ImageUrl.String,
		Role:        user.Role.String,
		AccessToken: accessToken,
	}, nil
}

func (s *AuthService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.UserResponse, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID: nulls.String{
			String: uuid.Must(uuid.NewV4()).String(),
			Valid:  true,
		},
		Name: nulls.String{
			String: req.Name,
			Valid:  true,
		},
		Email: nulls.String{
			String: req.Email,
			Valid:  true,
		},
		HashPassword: nulls.String{
			String: string(hashPassword),
			Valid:  true,
		},
		Role: nulls.String{
			String: "COSTUMER",
			Valid:  true,
		},
	}

	_, err = db.GetDB().Exec(ctx, "INSERT INTO public.\"User\" (id, name, email, \"hashPassword\") VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.Email, user.HashPassword)
	if err != nil {
		return nil, err
	}

	uc := token_service.UserClaims{
		Id:    user.ID.String,
		Email: user.Email.String,
		Role:  user.Role.String,
	}

	accessToken, err := token_service.NewAccessToken(uc)
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Name:        user.Name.String,
		Email:       user.Email.String,
		Role:        user.Role.String,
		AccessToken: accessToken,
	}, nil
}

func (s *AuthService) GoogleLogin(ctx context.Context, req *proto.LoginRequest) (*proto.UserResponse, error) {
	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"User\" WHERE email = $1", req.Email)
	if err != nil {
		return nil, err
	}

	var user models.User
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, nil, &user.HashPassword, &user.ImageUrl, &user.Role)
		if err != nil {
			return nil, err
		}
	}

	errWrongPassword := bcrypt.CompareHashAndPassword([]byte(user.HashPassword.String), []byte(req.Password))
	if errWrongPassword != nil {
		return nil, errWrongPassword
	}

	uc := token_service.UserClaims{
		Id:    user.ID.String,
		Email: user.Email.String,
	}

	accessToken, err := token_service.NewAccessToken(uc)
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Id:          user.ID.String,
		Name:        user.Name.String,
		Email:       user.Email.String,
		ImageUrl:    user.ImageUrl.String,
		Role:        user.Role.String,
		AccessToken: accessToken,
	}, nil
}
