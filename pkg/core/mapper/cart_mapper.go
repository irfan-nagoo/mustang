package mapper

import (
	
	"github.com/mustang/pkg/core/dto"
	"github.com/mustang/pkg/core/entity"
)



func ToCart(cDTO *dto.CartDTO) *entity.Cart {

	var ia []*entity.Item
	if len(cDTO.Items) != 0 {

		for _, item := range cDTO.Items {
			ia	= append(ia, ToItem(item))
		}
	}
	return &entity.Cart{
		Id: cDTO.Id,
		UserId: cDTO.UserId,
		ExpiryDate: cDTO.ExpiryDate,
		Items: ia,
	}
}

func ToCartDTO(c *entity.Cart) *dto.CartDTO {

	var ia []*dto.ItemDTO
	if len(c.Items) != 0 {

		for _, item := range c.Items {
			ia	= append(ia, ToItemDTO(item))
		}
	}
	return &dto.CartDTO{
		Id: c.Id,
		UserId: c.UserId,
		ExpiryDate: c.ExpiryDate,
		Items: ia,
	}
}


func ToItem(iDTO *dto.ItemDTO) *entity.Item {
	return &entity.Item{
		Id: iDTO.Id,
		Title: iDTO.Title,
		ImageUrl: iDTO.ImageUrl,
		Price: iDTO.Price,
		Quantity: iDTO.Quantity,
	}
}

func ToItemDTO(i *entity.Item) *dto.ItemDTO {
	return &dto.ItemDTO{
		Id: i.Id,
		Title: i.Title,
		ImageUrl: i.ImageUrl,
		Price: i.Price,
		Quantity: i.Quantity,
	}
}
