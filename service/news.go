package service

import (
	"context"

	db "ecomate-be/internal/db"
	"ecomate-be/internal/token_service"
	models "ecomate-be/models"
	proto "ecomate-be/proto"

	"github.com/gofrs/uuid"
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
			Id:       item.ID.String,
			Title:    item.Title.String,
			ImageUrl: item.ImageUrl.String,
			Content:  item.Content.String,
		})
	}

	return &proto.NewsListResponse{
		NewsList: newsList,
	}, nil
}

func (s *NewsService) GetNewsCommentList(ctx context.Context, req *proto.GetNewsCommentListRequest) (*proto.NewsCommentListResponse, error) {
	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"NewsComment\" WHERE \"newsId\" = '"+req.NewsId+"';")
	if err != nil {
		return nil, err
	}

	newsCommentDbList := []models.NewsComment{}
	for rows.Next() {
		var newsCommentDb models.NewsComment
		err := rows.Scan(&newsCommentDb.ID, &newsCommentDb.CreatedAt, nil, &newsCommentDb.Content, &newsCommentDb.NewsID, &newsCommentDb.UserID)
		if err != nil {
			return nil, err
		}
		newsCommentDbList = append(newsCommentDbList, newsCommentDb)
	}

	var newsCommentList []*proto.NewsComment
	for _, item := range newsCommentDbList {
		newsCommentList = append(newsCommentList, &proto.NewsComment{
			Id:        item.ID.String,
			CreatedAt: item.CreatedAt.String(),
			Content:   item.Content.String,
			NewsId:    item.NewsID.String,
			UserId:    item.UserID.String,
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

	randId := uuid.Must(uuid.NewV4()).String()
	_, errExec := db.GetDB().Exec(ctx, "INSERT INTO public.\"NewsComment\" (id, content, \"newsId\", \"userId\") VALUES ($1, $2, $3, $4)", randId, req.Content, req.NewsId, userClaims.Id)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *NewsService) UpdateNewsComment(ctx context.Context, req *proto.UpdateNewsCommentRequest) (*proto.GeneralStatusResponse, error) {
	userClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	_, errExec := db.GetDB().Exec(ctx, "UPDATE public.\"NewsComment\" SET content = $1 WHERE id = $2 AND \"userId\" = $3", req.Content, req.CommentId, userClaims.Id)
	if err != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *NewsService) DeleteNewsComment(ctx context.Context, req *proto.DeleteNewsCommentRequest) (*proto.GeneralStatusResponse, error) {
	userClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	_, errExec := db.GetDB().Exec(ctx, "DELETE FROM public.\"NewsComment\" WHERE id = $1 AND \"userId\" = $2", req.CommentId, userClaims.Id)
	if err != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}
