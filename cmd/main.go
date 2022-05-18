package main

import (
	"log"

	"github.com/CodingForFunAndProfit/apigateway/pkg/auth"
	"github.com/CodingForFunAndProfit/apigateway/pkg/config"
	"github.com/CodingForFunAndProfit/apigateway/pkg/order"
	"github.com/CodingForFunAndProfit/apigateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
