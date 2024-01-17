package service

import (
	"context"
	"log"

	db "github.com/Reigenleif/ecomate-mobile-backend-service/internal/db"
	"github.com/Reigenleif/ecomate-mobile-backend-service/internal/token_service"
	models "github.com/Reigenleif/ecomate-mobile-backend-service/models"
	proto "github.com/Reigenleif/ecomate-mobile-backend-service/proto"

	"github.com/gofrs/uuid"
)

type MarketplaceService struct {
	proto.UnimplementedMarketplaceServer
}

func (s *MarketplaceService) GetMarketplaceCategoryList(ctx context.Context, req *proto.GetMarketplaceCategoryListRequest) (*proto.MarketplaceCategoryListResponse, error) {
	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Category\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categoryDbList []models.Category
	for rows.Next() {
		var categoryDb models.Category
		err := rows.Scan(&categoryDb.ID, &categoryDb.CreatedAt, nil, &categoryDb.Name, &categoryDb.Description)
		if err != nil {
			return nil, err
		}
		categoryDbList = append(categoryDbList, categoryDb)
	}

	var categoryList []*proto.MarketplaceCategory
	for _, category := range categoryDbList {
		categoryList = append(categoryList, &proto.MarketplaceCategory{
			Id:          category.ID.String,
			Name:        category.Name.String,
			Description: category.Description.String,
		})
	}

	return &proto.MarketplaceCategoryListResponse{
		MarketplaceCategoryList: categoryList,
	}, nil
}

func (s *MarketplaceService) GetMarketplaceItemListByCategoryId(ctx context.Context, req *proto.GetMarketplaceItemListByCategoryIdRequest) (*proto.MarketplaceItemListResponse, error) {
	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Item\" WHERE \"categoryId\" = $1", req.CategoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itemDbList []models.Item
	for rows.Next() {
		var itemDb models.Item
		err := rows.Scan(&itemDb.ID, &itemDb.CreatedAt, nil, &itemDb.Name, &itemDb.Description, &itemDb.Price, &itemDb.ImageUrl, &itemDb.IsPublished, &itemDb.CategoryId, &itemDb.UserId)
		if err != nil {
			return nil, err
		}
		itemDbList = append(itemDbList, itemDb)
	}

	var itemList []*proto.MarketplaceItem
	for _, item := range itemDbList {
		itemList = append(itemList, &proto.MarketplaceItem{
			Id:          item.ID.String,
			Name:        item.Name.String,
			Description: item.Description.String,
			ImageUrl:    item.ImageUrl.String,
			Price:       int32(item.Price.Int),
		})
	}

	return &proto.MarketplaceItemListResponse{
		MarketplaceItemList: itemList,
	}, nil

}

func (s *MarketplaceService) GetMarketplaceItemList(ctx context.Context, req *proto.GetMarketplaceItemListRequest) (*proto.MarketplaceItemListResponse, error) {
	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Item\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itemDbList []models.Item
	for rows.Next() {
		var itemDb models.Item
		err := rows.Scan(&itemDb.ID, &itemDb.CreatedAt, nil, &itemDb.Name, &itemDb.Description, &itemDb.Price, &itemDb.ImageUrl, &itemDb.IsPublished, &itemDb.CategoryId, &itemDb.UserId)
		if err != nil {
			return nil, err
		}
		itemDbList = append(itemDbList, itemDb)
	}

	var itemList []*proto.MarketplaceItem
	for _, item := range itemDbList {
		itemList = append(itemList, &proto.MarketplaceItem{
			Id:          item.ID.String,
			Name:        item.Name.String,
			Description: item.Description.String,
			ImageUrl:    item.ImageUrl.String,
			Price:       int32(item.Price.Int),
		})
	}

	return &proto.MarketplaceItemListResponse{
		MarketplaceItemList: itemList,
	}, nil
}

func (s *MarketplaceService) GetMarketplaceItemById(ctx context.Context, req *proto.GetMarketplaceItemByIdRequest) (*proto.MarketplaceItemResponse, error) {
	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Item\" WHERE id = $1", req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itemDb models.Item
	if rows.Next() {
		err := rows.Scan(&itemDb.ID, &itemDb.CreatedAt, nil, &itemDb.Name, &itemDb.Description, &itemDb.Price, &itemDb.ImageUrl, &itemDb.IsPublished, &itemDb.CategoryId, &itemDb.UserId)
		if err != nil {
			return nil, err
		}
	}

	return &proto.MarketplaceItemResponse{
		MarketplaceItem: &proto.MarketplaceItem{
			Id:          itemDb.ID.String,
			Name:        itemDb.Name.String,
			Description: itemDb.Description.String,
			ImageUrl:    itemDb.ImageUrl.String,
			Price:       int32(itemDb.Price.Int),
		},
	}, nil
}

func (s *MarketplaceService) CreateMarketplaceItem(ctx context.Context, req *proto.CreateMarketplaceItemRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	log.Print(req.Price, req.Description, req)

	id, _ := uuid.NewV4()
	_, errExec := db.GetDB().Exec(ctx, "INSERT INTO public.\"Item\" (id, name, description, \"imageUrl\", price, \"categoryId\", \"userId\") VALUES ($1, $2, $3, $4, $5, $6, $7)", id, req.Name, req.Description, req.ImageUrl, req.Price, req.CategoryId, UserClaims.Id)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *MarketplaceService) UpdateMarketplaceItem(ctx context.Context, req *proto.UpdateMarketplaceItemRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	_, errExec := db.GetDB().Exec(ctx, "UPDATE public.\"Item\" SET name = $1, description = $2, \"imageUrl\" = $3, price = $4, \"categoryId\" = $5 WHERE id = $6 AND \"userId\" = $7", req.Name, req.Description, req.ImageUrl, req.Price, req.CategoryId, req.Id, UserClaims.Id)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *MarketplaceService) DeleteMarketplaceItem(ctx context.Context, req *proto.DeleteMarketplaceItemRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	_, errExec := db.GetDB().Exec(ctx, "DELETE FROM public.\"Item\" WHERE id = $1 AND \"userId\" = $2", req.Id, UserClaims.Id)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *MarketplaceService) GetCartByUserId(ctx context.Context, req *proto.GetCartByUserIdRequest) (*proto.CartResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartDb models.Cart
	if rows.Next() {
		err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
		if err != nil {
			return nil, err
		}
	} else {
		// Case Cart not found: create 1
		id, _ := uuid.NewV4()
		_, errExec := db.GetDB().Exec(ctx, "INSERT INTO public.\"Cart\" (id, \"userId\") VALUES ($1, $2)", id, UserClaims.Id)
		if errExec != nil {
			return nil, errExec
		}

		rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
		if err != nil {
			return nil, err
		}

		if rows.Next() {
			err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
			if err != nil {
				return nil, err
			}
		}
	}

	rows, err = db.GetDB().Query(ctx, "SELECT * FROM public.\"CartItem\" WHERE \"cartId\" = $1", cartDb.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItemDbList []models.CartItem
	for rows.Next() {
		var cartItemDb models.CartItem
		err := rows.Scan(&cartItemDb.ID, &cartItemDb.CreatedAt, nil, &cartItemDb.Quantity, &cartItemDb.CartId, &cartItemDb.ItemId, &cartItemDb.IsConfirmed, &cartItemDb.IsCheckedOut)
		if err != nil {
			return nil, err
		}
		cartItemDbList = append(cartItemDbList, cartItemDb)
	}

	var cartItemList []*proto.CartItem
	for _, cartItem := range cartItemDbList {
		cartItemList = append(cartItemList, &proto.CartItem{
			Id:       cartItem.ID.String,
			CartId:   cartItem.CartId.String,
			ItemId:   cartItem.ItemId.String,
			Quantity: cartItem.Quantity.Int32,
		})
	}

	return &proto.CartResponse{
		CartItemList: cartItemList,
	}, nil
}

func (s *MarketplaceService) DeleteCartByUserId(ctx context.Context, req *proto.DeleteCartByUserIdRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	_, errExec := db.GetDB().Exec(ctx, "DELETE FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *MarketplaceService) AddCartItem(ctx context.Context, req *proto.AddCartItemRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartDb models.Cart
	if rows.Next() {
		err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
		if err != nil {
			return nil, err
		}
	} else {
		// Case Cart not found: create 1
		id, _ := uuid.NewV4()
		_, errExec := db.GetDB().Exec(ctx, "INSERT INTO public.\"Cart\" (id, \"userId\") VALUES ($1, $2)", id, UserClaims.Id)
		if errExec != nil {
			return nil, errExec
		}

		rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
		if err != nil {
			return nil, err
		}

		if rows.Next() {
			err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
			if err != nil {
				return nil, err
			}
		}
	}

	id, _ := uuid.NewV4()
	_, errExec := db.GetDB().Exec(ctx, "INSERT INTO public.\"CartItem\" (id, \"cartId\", \"itemId\", quantity) VALUES ($1, $2, $3, $4)", id, cartDb.ID, req.ItemId, req.Quantity)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *MarketplaceService) UpdateCartItem(ctx context.Context, req *proto.UpdateCartItemRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartDb models.Cart
	if rows.Next() {
		err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
		if err != nil {
			return nil, err
		}
	} else {
		// Case Cart not found: create 1
		id, _ := uuid.NewV4()
		_, errExec := db.GetDB().Exec(ctx, "INSERT INTO public.\"Cart\" (id, \"userId\") VALUES ($1, $2)", id, UserClaims.Id)
		if errExec != nil {
			return nil, errExec
		}

		rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
		if err != nil {
			return nil, err
		}

		if rows.Next() {
			err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
			if err != nil {
				return nil, err
			}
		}
	}

	_, errExec := db.GetDB().Exec(ctx, "UPDATE public.\"CartItem\" SET quantity = $1 WHERE \"cartId\" = $2 AND \"itemId\" = $3", req.Quantity, cartDb.ID, req.ItemId)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *MarketplaceService) RemoveCartItem(ctx context.Context, req *proto.RemoveCartItemRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartDb models.Cart
	if rows.Next() {
		err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
		if err != nil {
			return nil, err
		}
	} else {
		// Case Cart not found: create 1
		id, _ := uuid.NewV4()
		_, errExec := db.GetDB().Exec(ctx, "INSERT INTO public.\"Cart\" (id, \"userId\") VALUES ($1, $2)", id, UserClaims.Id)
		if errExec != nil {
			return nil, errExec
		}

		rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
		if err != nil {
			return nil, err
		}

		if rows.Next() {
			err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
			if err != nil {
				return nil, err
			}
		}
	}

	_, errExec := db.GetDB().Exec(ctx, "DELETE FROM public.\"CartItem\" WHERE \"cartId\" = $1 AND \"itemId\" = $2", cartDb.ID, req.ItemId)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *MarketplaceService) CheckOutCart(ctx context.Context, req *proto.CheckOutCartRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartDb models.Cart
	if rows.Next() {
		err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	_, errExec := db.GetDB().Exec(ctx, "UPDATE public.\"CartItem\" SET \"isCheckedOut\" = true WHERE \"cartId\" = $1", cartDb.ID)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *MarketplaceService) CheckOutCartItem(ctx context.Context, req *proto.CheckOutCartItemRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"Cart\" WHERE \"userId\" = $1", UserClaims.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartDb models.Cart
	if rows.Next() {
		err := rows.Scan(&cartDb.ID, &cartDb.CreatedAt, nil, &cartDb.UserId)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	_, errExec := db.GetDB().Exec(ctx, "UPDATE public.\"CartItem\" SET \"isCheckedOut\" = true WHERE \"id\" = $1", req.CartItemId)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}

func (s *MarketplaceService) GetCheckedOutItemList(ctx context.Context, req *proto.GetCheckedOutItemListRequest) (*proto.CartItemListResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := db.GetDB().Query(ctx, "SELECT * FROM public.\"CartItem\" WHERE \"itemId\" IN (SELECT id FROM public.\"Item\" WHERE \"userId\" = $1) AND \"isCheckedOut\" = true", UserClaims.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItemDbList []models.CartItem
	for rows.Next() {
		var cartItemDb models.CartItem
		err := rows.Scan(&cartItemDb.ID, &cartItemDb.CreatedAt, nil, &cartItemDb.Quantity, &cartItemDb.CartId, &cartItemDb.ItemId, &cartItemDb.IsConfirmed, &cartItemDb.IsCheckedOut)
		if err != nil {
			return nil, err
		}
		cartItemDbList = append(cartItemDbList, cartItemDb)
	}

	var cartItemList []*proto.CartItem
	for _, cartItem := range cartItemDbList {
		cartItemList = append(cartItemList, &proto.CartItem{
			Id:       cartItem.ID.String,
			CartId:   cartItem.CartId.String,
			ItemId:   cartItem.ItemId.String,
			Quantity: cartItem.Quantity.Int32,
		})
	}

	return &proto.CartItemListResponse{
		CartItemList: cartItemList,
	}, nil
}

func (s *MarketplaceService) ConfirmCartItem(ctx context.Context, req *proto.ConfirmCartItemRequest) (*proto.GeneralStatusResponse, error) {
	UserClaims, err := token_service.CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	_, errExec := db.GetDB().Exec(ctx, "UPDATE public.\"CartItem\" SET \"isConfirmed\" = true WHERE \"itemId\" IN (SELECT id FROM \"Item\" WHERE \"userId\" = $1) AND \"id\" = $2", UserClaims.Id, req.CartItemId)
	if errExec != nil {
		return nil, errExec
	}

	return &proto.GeneralStatusResponse{
		Success: true,
	}, nil
}
