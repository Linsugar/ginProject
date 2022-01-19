package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	// "mysql", "root:XXXX@tcp(127.0.0.1:3306)/test"
	host := "139.155.88.241"
	port := "3388"
	database := "jiazu"
	username := "root"
	password := "123456"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	dataDase, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("链接数据库有误%v", err)
		return
	}
	fmt.Printf("链接成功")
	Db = dataDase
	// defer Db.Close()
}

type Test3 struct {
	//model 必须大写
	gorm.Model
	Name    string
	Address string
	Age     int
}

func DoTest() {
	//没有该表就创建-该表结构发生变化
	Db.AutoMigrate(&Test3{})
	// Db.CreateTable(&Test3{})
	var t3 = &Test3{Name: "唐哟哟", Address: "日本", Age: 20}
	fmt.Printf("得到结果：%v", t3)
	Db.Create(&t3)
	// defer Db.Close()
}

func FindData() *Test3 {
	t1 := &Test3{}
	Db.First(&t1)
	fmt.Printf("得到的数据：%v", t1)
	return t1
}
