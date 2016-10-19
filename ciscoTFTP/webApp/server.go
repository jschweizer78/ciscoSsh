package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	addr := ":8080"
	apiHandler := newAPIRes(r)

	r.Static("/public", "./public")

	apiHandler.setResoreces()

	fmt.Printf("running at %s\n", addr)
	r.Run(addr)
}
