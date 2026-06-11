package dto

import (
	"github.com/mustang/pkg/core/vo"
)


type ItemDTO struct {
	Id int64
	Title string
	ImageUrl string
	Price vo.Money
	Quantity int64
}