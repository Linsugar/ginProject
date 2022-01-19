package adminController

import (
	mysql "TestGo/domysql"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func (con AdminController) AdminHome(c *gin.Context) {
	c.String(200, "我是首页")
}

func (con AdminController) UpdatePage(c *gin.Context) {
	name := c.Param("name")
	result := mysql.UpdateData(name)
	c.JSON(200, gin.H{
		"结果": result,
	})

}
