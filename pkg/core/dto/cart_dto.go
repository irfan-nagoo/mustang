package dto

import (
	"time"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CartDTO struct {
	Id     bson.ObjectID
	UserId string
	ExpiryDate time.Time 
	Items  []*ItemDTO
}
