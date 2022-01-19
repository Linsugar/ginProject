package adminController

import "github.com/gin-gonic/gin"

type AdminController struct {
}

func (con AdminController) AdminHome(c *gin.Context) {
	c.String(200, "我是首页")
}
