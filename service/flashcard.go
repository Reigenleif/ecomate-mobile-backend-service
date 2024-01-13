package service

import (
	"context"

	db "github.com/Reigenleif/ecomate-mobile-backend-service/internal/db"
	models "github.com/Reigenleif/ecomate-mobile-backend-service/models"
	proto "github.com/Reigenleif/ecomate-mobile-backend-service/proto"
)

type FlashcardService struct {
	proto.UnimplementedFlashcardServiceServer
}

func (s *FlashcardService) GetFlashcardList(ctx context.Context, req *proto.GetFlashcardListRequest) (*proto.FlashcardListResponse, error) {
	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Flashcard\"")
	if err != nil {
		return nil, err
	}

	var flashcardDbList []models.Flashcard
	for rows.Next() {
		var flashcardDb models.Flashcard
		err := rows.Scan(&flashcardDb.ID, &flashcardDb.CreatedAt, nil, &flashcardDb.Title, &flashcardDb.ImageUrl, &flashcardDb.Description, &flashcardDb.IsPublished)
		if err != nil {
			return nil, err
		}
		flashcardDbList = append(flashcardDbList, flashcardDb)
	}

	var flashcardList []*proto.Flashcard
	for _, item := range flashcardDbList {
		flashcardList = append(flashcardList, &proto.Flashcard{
			Id:          item.ID.String,
			Title:       item.Title.String,
			ImageUrl:    item.ImageUrl.String,
			Description: item.Description.String,
		})
	}

	return &proto.FlashcardListResponse{
		FlashcardList: flashcardList,
	}, nil
}

func (s *FlashcardService) GetFlashcardById(ctx context.Context, req *proto.GetFlashcardByIdRequest) (*proto.FlashcardResponse, error) {
	var flashcardDb models.Flashcard
	err := db.GetDB().QueryRow(ctx, "SELECT * FROM public.\"Flashcard\" WHERE id = '"+req.Id+"';").Scan(&flashcardDb.ID, &flashcardDb.CreatedAt, nil, &flashcardDb.Title, &flashcardDb.ImageUrl, &flashcardDb.Description, &flashcardDb.IsPublished)
	if err != nil {
		return nil, err
	}

	return &proto.FlashcardResponse{
		Flashcard: &proto.Flashcard{
			Id:          flashcardDb.ID.String,
			Title:       flashcardDb.Title.String,
			ImageUrl:    flashcardDb.ImageUrl.String,
			Description: flashcardDb.Description.String,
		},
	}, nil
}
