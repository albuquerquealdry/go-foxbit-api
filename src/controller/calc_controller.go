package controller

import (
	"foxbit-calc-api/src/model"
	"foxbit-calc-api/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getParams(c *gin.Context) (float64, float64, bool) {
	termOneStr := c.Query("term_one")
	termTwoStr := c.Query("term_two")

	termOne, err1 := strconv.ParseFloat(termOneStr, 64)
	termTwo, err2 := strconv.ParseFloat(termTwoStr, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parâmetros inválidos"})
		return 0, 0, false
	}
	return termOne, termTwo, true
}

func Sum(c *gin.Context) {
	a, b, ok := getParams(c)
	if !ok {
		return
	}
	result := service.Sum(a, b)
	c.JSON(http.StatusOK, model.CalcResponse{Resultado: result})
}

func Sub(c *gin.Context) {
	a, b, ok := getParams(c)
	if !ok {
		return
	}
	result := service.Sub(a, b)
	c.JSON(http.StatusOK, model.CalcResponse{Resultado: result})
}

func Mul(c *gin.Context) {
	a, b, ok := getParams(c)
	if !ok {
		return
	}
	result := service.Mul(a, b)
	c.JSON(http.StatusOK, model.CalcResponse{Resultado: result})
}

func Div(c *gin.Context) {
	a, b, ok := getParams(c)
	if !ok {
		return
	}
	result, err := service.Div(a, b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.CalcResponse{Resultado: result})
}
