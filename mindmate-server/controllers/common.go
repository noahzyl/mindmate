/*
 * Define the common logic of the controller layer
 */

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SuccessJson struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
}

type ErrorJson struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func ReturnSuccessJson(ctx *gin.Context, code int, message interface{}, data interface{}, count int64) {
	json := &SuccessJson{
		Code:    code,
		Message: message,
		Data:    data,
		Count:   count,
	}
	ctx.JSON(http.StatusOK, json)
}

func ReturnErrorJson(ctx *gin.Context, code int, message interface{}) {
	json := &ErrorJson{
		Code:    code,
		Message: message,
	}
	ctx.JSON(http.StatusOK, json)
}
