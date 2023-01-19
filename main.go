package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:coin", func(ctx *gin.Context) {
		coin := ctx.Param("coin")
		switch coin {
		case "bitcoin":
			response(ctx.Writer, "BTC")
		case "canadiandollar":
			response(ctx.Writer, "CAD")
		case "dogecoin":
			response(ctx.Writer, "DOGE")
		case "dollar":
			response(ctx.Writer, "USD")
		case "euro":
			response(ctx.Writer, "EUR")
		case "pound":
			response(ctx.Writer, "GBP")
		default:
			response(ctx.Writer, coin)
		}
	})

	log.Print("The is Server Running on localhost port 8080")
	router.Run(":8080")
}
