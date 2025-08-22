package main

import (
	"foxbit-calc-api/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/sum", controller.Sum)
		api.GET("/sub", controller.Sub)
		api.GET("/mul", controller.Mul)
		api.GET("/div", controller.Div)
	}
	r.GET("/health/liveness", controller.Liveness)
	r.GET("/health/readiness", controller.Readiness)
	r.Run(":8000")
}
