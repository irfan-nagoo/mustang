package web

import (
	"context"
	"net/http"

	"github.com/mustang/pkg/core/request"
	"github.com/mustang/pkg/core/service"

	"github.com/gin-gonic/gin"
)

type CartRoute struct {
	Service *service.CartService
}

// @BasePath /
// Mustang APIs godoc
// @Summary Add item
// @Description Add item to cart
// @Tags addcart
// @Accept json
// @Produce json
// @Param request body request.CartRequest true "Add item payload"
// @Success 200 {object} dto.CartDTO
// @Router /carts [post]
func (cr *CartRoute) AddToCart(c *gin.Context) {

	var r request.CartRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := cr.Service.AddToCart(&r, context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @BasePath /
// Mustang APIs godoc
// @Summary Get cart
// @Schemes
// @Description Get cart
// @Tags getcart
// @Accept json
// @Produce json
// @Param  id   path  int  true  "User ID"
// @Success 200 {object} dto.CartDTO
// @Router /carts/users/{id} [get]
func (cr *CartRoute) GetCart(c *gin.Context) {
	
	userId := c.Param("id")
	res, err := cr.Service.GetCart(userId, context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @BasePath /
// Mustang APIs godoc
// @Summary Delete item from cart
// @Description Delete item from cart
// @Tags deletecart
// @Accept json
// @Produce json
// @Param id      path  int  true "User ID"
// @Param itemId  path  int  true "Item ID"
// @Success 200 {object} response.BaseResponse
// @Router /carts/users/{id}/items/{itemId} [delete]
func (cr *CartRoute) DeleteFromCart(c *gin.Context) {

	userId := c.Param("id")
	itemId := c.Param("itemId")
	res, err := cr.Service.DeleteFromCart(userId, itemId, context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}