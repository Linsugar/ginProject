package main

import (
	// "fmt"

	"TestGo/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("haakkaka")
	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "HEllo")
	// })
	routers.AdminRouter(r)
	r.Run(":8022")

}
