package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roihan12/assignment-2/model"
	"github.com/roihan12/assignment-2/service"
)

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *orderController {
	return &orderController{orderService}
}

// CreateOrder godoc
//
//	@Summary		Create new order
//	@Description	Create new order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.OrderCreateRequest	true	"request is required"
//	@Success		200		{object}	model.OrderResponse
//	@Failure		400		{object}	model.APIError
//	@Failure		500		{object}	model.APIError
//
//	@Router			/orders [post]
func (controller *orderController) CreateOrder(ctx *gin.Context) {
	var newOrder model.OrderCreateRequest

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}
	
	response, err := controller.orderService.CreateOrder(newOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "Order created succcessfully",
			"data":    response,
		})
}

// GetAllOrder godoc
//
//	@Summary		Get all orders
//	@Description	Get all details orders
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//
//	@Failure		500	{object}	model.APIError
//
//	@Success		200	{array}		model.OrderResponse
//	@Router			/orders [get]
func (controller *orderController) GetAllOrder(ctx *gin.Context) {
	response, err := controller.orderService.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,
		gin.H{
			"message": "Get all order succcessfully",
			"data":    response,
		})

}

// GetOrderById godoc
//
//	@Summary		Get one order by id
//	@Description	Get details order by id
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Failure		400	{object}	model.APIError	"We need ID!!"
//	@Failure		404	{object}	model.APIError	"Can not find ID"
//	@Success		200	{object}	model.OrderResponse
//	@Router			/orders/{id} [get]
func (controller *orderController) GetOrderById(ctx *gin.Context) {
	orderID := ctx.Param("id")

	uintValue, err := strconv.ParseUint(orderID, 10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response, err := controller.orderService.GetOrderByID(uintValue)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.APIError{
			ErrorCode:    http.StatusNotFound,
			ErrorMessage: "Order not found",
		})
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"message": "Get order by id succcessfully",
			"data":    response,
		})
}

// DeleteOrder godoc
//
//	@Summary		Delete order by id
//	@Description	Delete order by id
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"id"
//	@Failure		400	{object}	model.APIError	"We need ID!!"
//	@Failure		404	{object}	model.APIError	"Can not find ID"
//	@Success		200	{object}	string			"Order deleted succcessfully"
//	@Router			/orders/{id} [delete]
func (controller *orderController) DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")

	uintValue, err := strconv.ParseUint(orderID, 10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = controller.orderService.DeleteOrder(uintValue)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order deleted succcessfully",
	})
}

// UpdateOrder godoc
//
//	@Summary		Update order
//	@Description	Update new order
//	@Tags			Order
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		string						true	"id"
//
//	@Param			request	body		model.OrderUpdateRequest	true	"request is required"
//	@Success		200		{object}	model.OrderResponse
//	@Failure		400		{object}	model.APIError
//	@Failure		500		{object}	model.APIError
//
//	@Router			/orders/{id} [put]
func (controller *orderController) UpdateOrder(ctx *gin.Context) {

	orderID := ctx.Param("id")

	uintValue, err := strconv.ParseUint(orderID, 10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var updatedOrder model.OrderUpdateRequest

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	response, err := controller.orderService.UpdateOrder(uintValue, updatedOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"message": "Updated order succcessfully",
			"data":    response,
		})
}
