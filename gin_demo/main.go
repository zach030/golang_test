package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type ProductInfo struct {
	ID      int `json:"id"`
	TypeId   string `json:"type_id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

var(
	db *gorm.DB
	err error
)

func Search(c *gin.Context){
	typeid := c.Param("typeid")
	var product []ProductInfo
	// Get all matched records
	db.Where(&ProductInfo{ParentId: typeid}).Find(&product)
	if err:=db.Find(&ProductInfo{}).Error;err!=nil{
		c.AbortWithStatus(404)
		fmt.Println(err)
	}else{
		c.JSON(http.StatusOK, product)
	}
}

func Add(c *gin.Context){
	//新增商品
	var newPro ProductInfo
	typeid:=c.PostForm("typeid")
	name:=c.PostForm("name")
	parentid:=c.PostForm("parentid")
	newPro = ProductInfo{TypeId: typeid, Name: name, ParentId: parentid}
	db.Create(&newPro)
	err := db.Debug().Save(&newPro).Error
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"typeid": typeid,
		"name": name,
		"parentid":parentid,
	})
}
func initTable(){
	//创建数据行
	u1:= ProductInfo{TypeId: "101", Name: "家电", ParentId: "100"}
	db.Create(&u1)
	u2:=ProductInfo{TypeId: "102", Name: "数码", ParentId: "100"}
	db.Create(&u2)
	u3:=ProductInfo{TypeId: "103", Name: "食品", ParentId: "100"}
	db.Create(&u3)
	u4:=ProductInfo{TypeId: "104", Name: "电视", ParentId: "101"}
	db.Create(&u4)
	u5:=ProductInfo{TypeId: "105", Name: "笔记本", ParentId: "102"}
	db.Create(&u5)
	u6:=ProductInfo{TypeId: "106", Name: "可乐", ParentId: "103"}
	db.Create(&u6)
	u7:=ProductInfo{TypeId: "107", Name: "全面屏电视", ParentId: "104"}
	db.Create(&u7)
	u8:=ProductInfo{TypeId: "108", Name: "商务笔记本", ParentId: "105"}
	db.Create(&u8)
	u9:=ProductInfo{TypeId: "109", Name: "牛肉干", ParentId: "103"}
	db.Create(&u9)
	u10:=ProductInfo{TypeId: "110", Name: "手机", ParentId: "102"}
	db.Create(&u10)
	return
}
func main(){
	//连接数据库
	db, err = gorm.Open("mysql", "root:991030@(localhost)/golang?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	defer db.Close()
	//创建表，自动迁移（把结构体和数据表对应）
	db.AutoMigrate(&ProductInfo{})
	//initTable()
	router:=gin.Default()
	router.GET("/search/:typeid",Search)
	router.POST("/add", Add)
	router.Run(":8080")
}