package routers

import (
	"TestGo/adminController"
	mysql "TestGo/domysql"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	Admin := r.Group("/admin")
	{
		Admin.GET("/", adminController.AdminController{}.AdminHome)
		Admin.GET("/admin", func(c *gin.Context) {
			mysql.DoTest()
			rs := mysql.FindData()
			c.JSON(200, gin.H{"status": rs})
		})
	}
}
