package routers

import (
	"TestGo/adminController"
	mysql "TestGo/domysql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	Admin := r.Group("/admin")
	{
		Admin.GET("/", adminController.AdminController{}.AdminHome)
		Admin.GET("/page/:name", adminController.AdminController{}.UpdatePage)
		Admin.GET("/admin/:age", func(c *gin.Context) {
			// var b strings.Builder
			mysql.DoTest()
			val := c.Param("age")
			fmt.Println("==========================")
			fmt.Println(val)
			fmt.Println("==========================")
			va, err := strconv.Atoi(val)
			if err != nil {
				c.JSON(200, gin.H{"status": err})
			} else {
				rs := mysql.FindData(va)

				if rs.Age == 100 {
					fmt.Printf("age大于100")
					c.JSON(200, gin.H{"status": fmt.Sprint("小于", rs.Name)})
				} else {
					c.JSON(200, gin.H{"status": fmt.Sprint("小于", rs.Name)})
				}
			}

		})
	}
}
