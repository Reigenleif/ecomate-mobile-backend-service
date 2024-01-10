package service

import (
	"context"
	"fmt"

	db "github.com/Reigenleif/ecomate-mobile-backend-service/internal/db"
	"github.com/Reigenleif/ecomate-mobile-backend-service/internal/token_service"
	models "github.com/Reigenleif/ecomate-mobile-backend-service/models"
	proto "github.com/Reigenleif/ecomate-mobile-backend-service/proto"
)

type NewsService struct {
	proto.UnimplementedNewsServiceServer
}

func (s *NewsService) GetNewsList(ctx context.Context, req *proto.GetNewsListRequest) (*proto.NewsListResponse, error) {
	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"News\"")
	if err != nil {
		return nil, err
	}

	var newsDbList []models.News
	for rows.Next() {
		var newsDb models.News
		err := rows.Scan(&newsDb.ID, &newsDb.CreatedAt, &newsDb.UpdatedAt, &newsDb.Title, &newsDb.ImageUrl, &newsDb.Content, &newsDb.IsPublished)
		if err != nil {
			return nil, err
		}
		newsDbList = append(newsDbList, newsDb)
	}

	var newsList []*proto.News
	for _, item := range newsDbList {
		newsList = append(newsList, &proto.News{
			Id: item.ID.String,
			Title: item.Title.String,
			ImageUrl: item.ImageUrl.String,
			Content: item.Content.String,
		})
	}


	return &proto.NewsListResponse{
		NewsList: newsList,
	}, nil
}

func (s *NewsService) GetNewsCommentList(ctx context.Context, req *proto.GetNewsCommentListRequest) (*proto.NewsCommentListResponse, error) {
	newsComment, err := db.Get[models.NewsComment](ctx, "SELECT * FROM public.\"NewsComment\"")

	if err != nil {
		return nil, err
	}

	var newsCommentList []*proto.NewsComment
	for _, item := range newsComment {
		newsCommentList = append(newsCommentList, &proto.NewsComment{
			Id: item.ID.String(),
			Content: item.Content.String,
			NewsId: item.NewsID.String(),
			UserId: item.UserID.String(),
		})
	}

	return &proto.NewsCommentListResponse{
		NewsCommentList: newsCommentList,
	}, nil
}

func (s *NewsService) CreateNewsComment(ctx context.Context, req *proto.CreateNewsCommentRequest) (*proto.GeneralStatusResponse, error) {
	userClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	res , err := db.Exec(ctx, fmt.Sprintf("INSERT INTO public.\"NewsComment\" (content, newsid, userid) VALUES (%s, %s, %s)", req.Content, req.NewsId, userClaims.Id))
	if err != nil {
		return nil, err
	}

	return &proto.GeneralStatusResponse{
		Success: res,
	}, nil
}

func (s *NewsService) UpdateNewsComment(ctx context.Context, req *proto.UpdateNewsCommentRequest) (*proto.GeneralStatusResponse, error) {
	userClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	res , err := db.Exec(ctx, fmt.Sprintf("UPDATE public.\"NewsComment\" SET content = %s WHERE id = %s", req.Content, userClaims.Id))
	if err != nil {
		return nil, err
	}

	return &proto.GeneralStatusResponse{
		Success: res,
	}, nil
}

func (s *NewsService) DeleteNewsComment(ctx context.Context, req *proto.DeleteNewsCommentRequest) (*proto.GeneralStatusResponse, error) {
	userClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	res , err := db.Exec(ctx, fmt.Sprintf("DELETE FROM public.\"NewsComment\" WHERE id = %s", userClaims.Id))
	if err != nil {
		return nil, err
	}

	return &proto.GeneralStatusResponse{
		Success: res,
	}, nil
}


