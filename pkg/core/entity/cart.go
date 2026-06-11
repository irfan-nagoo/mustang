package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Cart struct {
	Id bson.ObjectID		`bson:"_id,omitempty"`
	UserId string			`bson:"userId,omitempty"`
	ExpiryDate time.Time 	`bson:"expiryDate,omitempty"`
	Items []*Item			`bson:"items,omitempty"`
}

