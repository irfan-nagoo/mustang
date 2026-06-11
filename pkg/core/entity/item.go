package entity

import (
	"github.com/mustang/pkg/core/vo"
)

type Item struct {
	Id int64			`bson:"id,omitempty"`
	Title string		`bson:"title,omitempty"`
	ImageUrl string		`bson:"imageUrl,omitempty"`
	Price vo.Money		`bson:"price,omitempty"`
	Quantity int64		`bson:"quantity,omitempty"`
}