package controllers

import (
	"assignment2/database"
	"assignment2/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderMethod interface {
	CreateOrder(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type OrderController struct{}

type ItemParam struct {
	ItemCode    int    `json:"item_code" example:"1"`
	Description string `json:"description" example:"Sabun"`
	Quantity    int    `json:"quantity" example:"12"`
}

type UpdateItemParam struct {
	LineItemId  int    `json:"lineItemId" example:"1"`
	ItemCode    int    `json:"item_code" example:"1"`
	Description string `json:"description" example:"Sabun"`
	Quantity    int    `json:"quantity" example:"12"`
}

type CreateParam struct {
	OrderedAt    time.Time `json:"ordered_at" example:"2024-03-11T12:34:56Z"`
	CustomerName string    `json:"customer_name" example:"Irvan"`
	Items        []ItemParam
}

type UpdateParam struct {
	OrderedAt    time.Time `json:"ordered_at" example:"2024-03-11T12:34:56Z"`
	CustomerName string    `json:"customer_name" example:"Irvan"`
	Items        []UpdateItemParam
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

// CreateOrder godoc
// @Summary Creating Order
// @Description  Creating Order based on input
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param order body CreateParam true "Order"
// @Success      201  {object}  model.Order{}
// @Failure      400  {object}  ErrorResponse
// @Router       /orders [post]
func (c OrderController) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order
	msg := ErrorResponse{}
	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		msg.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	db := database.GetDB()
	errs := db.Create(&newOrder).Error

	if errs != nil {
		msg.Message = errs.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
		fmt.Println("Error creating product data:", err)
		return
	}
	ctx.JSON(http.StatusCreated, newOrder)
}

// GetOrders godoc
// @Summary Get list of Order
// @Description  Fetching all order data
// @Tags         orders
// @Produce      json
// @Success      200  {object}  []model.Order{}
// @Failure      404  {object}  ErrorResponse
// @Router       /orders [get]
func (c OrderController) GetOrders(ctx *gin.Context) {
	db := database.GetDB()
	msg := ErrorResponse{}
	orders := []model.Order{}
	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		msg.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusNotFound, msg)
		fmt.Println("Error get data ", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, orders)

}

// UpdateOrder godoc
// @Summary Update Order
// @Description  Update a specific order data
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param orderId path int true "Order ID"
// @Param order body UpdateParam true "Order"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /orders/{orderId} [put]
func (c OrderController) UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	msg := ErrorResponse{}
	order := UpdateParam{}
	itemModel := model.Item{}
	orderModel := model.Order{}

	// Get order Id
	var ID = ctx.Param("orderId")
	orderID, errs := strconv.Atoi(ID)
	if errs != nil {
		msg.Message = errs.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	//Bind update order param
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		msg.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	// Update all item in order.Items
	for i, items := range order.Items {
		_ = i

		// Validation to ensure LineItemId is filled
		if items.LineItemId == 0 {
			msg.Message = ("LineItemId must be filled !")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
			return
		}

		//Create new Item based on item param
		newItem := model.Item{
			ItemID:      uint(items.LineItemId),
			ItemCode:    items.ItemCode,
			Description: items.Description,
			Quantity:    items.Quantity,
			OrderID:     uint(orderID),
		}

		//Update the item
		errUpdateItem := db.Model(&itemModel).Where("item_id=?", items.LineItemId).Updates(newItem).Error
		if errUpdateItem != nil {
			if errors.Is(errUpdateItem, gorm.ErrRecordNotFound) {
				msg.Message = fmt.Sprintf("Data Item with association order_id = %d Not Found", orderID)
				ctx.AbortWithStatusJSON(http.StatusNotFound, msg)
				return
			} else {
				fmt.Println("Error Delete : ", errUpdateItem)
				msg.Message = errUpdateItem.Error()
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, msg)
				return
			}
		}
	}
	//Create new order based on order param
	newOrder := model.Order{
		OrderID:      uint(orderID),
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
	}

	// Update order with order_id = orderID
	errUpdate := db.Model(&orderModel).Where("order_id=?", orderID).Updates(newOrder).Error
	if errUpdate != nil {
		if errors.Is(errUpdate, gorm.ErrRecordNotFound) {
			msg.Message = fmt.Sprintf("Data with id = %d Not Found", orderID)
			ctx.AbortWithStatusJSON(http.StatusNotFound, msg)
			return
		} else {
			fmt.Println("Error update : ", errUpdate)
			msg.Message = errUpdate.Error()
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, msg)
			return
		}
	}

	//Display success
	success := SuccessResponse{
		Message: fmt.Sprintf("Data order with id %d updated successfully", orderID),
	}

	ctx.JSON(http.StatusOK, success)

}

// DeleteOrder godoc
// @Summary Delete Order
// @Description  Delete a specific order data
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param orderId path int true "Order ID"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /orders/{orderId} [delete]
func (c OrderController) DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	msg := ErrorResponse{}
	order := model.Order{}
	item := model.Item{}
	_ = item
	var ID = ctx.Param("orderId")
	orderID, errs := strconv.Atoi(ID)

	if errs != nil {
		msg.Message = errs.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	// Delete all item with order_id = orderID
	errDelItem := db.Where("order_id=?", orderID).Delete(&item).Error
	if errDelItem != nil {
		if errors.Is(errDelItem, gorm.ErrRecordNotFound) {
			msg.Message = fmt.Sprintf("Data Item with association order_id = %d Not Found", orderID)
			ctx.AbortWithStatusJSON(http.StatusNotFound, msg)
			return
		} else {
			fmt.Println("Error Delete : ", errDelItem)
			msg.Message = errDelItem.Error()
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, msg)
			return
		}
	}

	// Delete all order with order_id = orderID
	errDel := db.Where("order_id=?", orderID).Delete(&order).Error
	if errDel != nil {
		if errors.Is(errDel, gorm.ErrRecordNotFound) {
			msg.Message = fmt.Sprintf("Data Order with order_id = %d Not Found", orderID)
			ctx.AbortWithStatusJSON(http.StatusNotFound, msg)
			return
		} else {
			fmt.Println("Error Delete : ", errDel)
			msg.Message = errDel.Error()
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, msg)
			return
		}
	}
	delResp := SuccessResponse{}
	delResp.Message = fmt.Sprintf("car with id %v has been succesfully deleted", orderID)
	ctx.JSON(http.StatusOK, delResp)

}
