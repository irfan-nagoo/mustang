package service

import (
	"context"
	"slices"
	"strconv"
	"time"

	"github.com/mustang/pkg/core/dto"
	"github.com/mustang/pkg/core/entity"
	"github.com/mustang/pkg/core/mapper"
	"github.com/mustang/pkg/core/request"
	"github.com/mustang/pkg/core/response"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type CartService struct {
	Colln *mongo.Collection
}

func (cs *CartService) AddToCart(cr *request.CartRequest, ctx context.Context) (*dto.CartDTO, error) {

	filter := bson.D{{Key: "userId", Value: cr.UserId}}
	c, err := findByUserId(&filter, cs, ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			nc := entity.Cart{
				UserId:     cr.UserId,
				ExpiryDate: time.Now().AddDate(0, 0, 5),
				Items:      []*entity.Item{mapper.ToItem(cr.Item)},
			}

			ic, err := cs.Colln.InsertOne(ctx, nc)
			if err != nil {
				return nil, err
			}

			nc.Id = ic.InsertedID.(bson.ObjectID)
			return mapper.ToCartDTO(&nc), nil

		} else {
			return nil, err
		}
	}

	it := mapper.ToItem(cr.Item)
	update := bson.M{
		"$push": bson.M{"items": it},
	}
	_, err = cs.Colln.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	c.Items = append(c.Items, it)

	return mapper.ToCartDTO(c), nil
}

func (cs *CartService) GetCart(userId string, ctx context.Context) (*dto.CartDTO, error) {

	filter := bson.D{{Key: "userId", Value: userId}}
	c, err := findByUserId(&filter, cs, ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ToCartDTO(c), nil
}

func (cs *CartService) DeleteFromCart(userId string, itemId string, ctx context.Context) (*response.BaseResponse, error) {

	filter := bson.D{{Key: "userId", Value: userId}}
	c, err := findByUserId(&filter, cs, ctx)
	if err != nil {
		return nil, err
	}

	iId, _ := strconv.ParseInt(itemId, 10, 64)
	update := bson.M{
		"$pull": bson.M{
			"items": bson.M{"id": iId}},
	}
	_, err = cs.Colln.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	c.Items = slices.DeleteFunc(c.Items, func(i *entity.Item) bool {
		return i.Id == iId
	})
	// if cart empty, delete it
	message := "Item successfully deleted"
	if len(c.Items) == 0 {
		if _, err := cs.Colln.DeleteOne(ctx, filter); err != nil {
			return nil, err
		}
		message = "Cart is empty now"
	}
	return &response.BaseResponse{
		Code:    "SUCCESS",
		Message: message,
	}, nil
}

func findByUserId(uf *bson.D, cs *CartService, ctx context.Context) (*entity.Cart, error) {

	opts := options.FindOne().SetSkip(0)

	var c entity.Cart
	err := cs.Colln.FindOne(ctx, uf, opts).Decode(&c)
	return &c, err
}
