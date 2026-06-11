package request

import "github.com/mustang/pkg/core/dto"

type CartRequest struct {
	UserId string
	Item   *dto.ItemDTO
}