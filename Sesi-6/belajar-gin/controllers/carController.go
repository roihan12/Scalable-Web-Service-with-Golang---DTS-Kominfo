package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
var newCar Car

err := ctx.ShouldBindJSON(&newCar); err != nil {
	ctx.AbortWithError(http.statusBadRequest, err)
	return
}


newCar.id

}