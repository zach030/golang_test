package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type ProductInfoTree struct {
	ID      int `json:"id"`
	TypeId   string `json:"type_id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
	Children []ProductInfoTree `json:"children"`
}
type ProductInfo struct {
	ID int `json:"id"`
	TypeId string `json:"type_id"`
	Name string `json:"name"`
	ParentId string `json:"parent_id"`
}
var(
	db *gorm.DB
	err error
	//存放商品的切片
	product []ProductInfo
)

//递归实现(返回树状结果得数据)
func productTree(allCate []ProductInfo, pid string) []ProductInfoTree {
	var arr []ProductInfoTree
	for _, v := range allCate {
		//循环遍历
		if pid == v.ParentId {
			//找到子类
			ctree := ProductInfoTree{}
			//赋值
			ctree.TypeId = v.TypeId
			ctree.ParentId = v.ParentId
			ctree.Name = v.Name
			//以此子类作为父类进行递归
			sonCate := productTree(allCate, v.TypeId)
			ctree.Children = sonCate
			arr = append(arr, ctree)
		}
	}
	return arr
}

func Search(c *gin.Context){
	typeid := c.Param("typeid")
	arr:= productTree(product, typeid)
	fmt.Println(arr)
	c.JSON(http.StatusOK, arr)
}

func Add(c *gin.Context){
	//新增商品
	var newPro ProductInfoTree
	typeid:=c.PostForm("typeid")
	name:=c.PostForm("name")
	parentid:=c.PostForm("parentid")
	newPro = ProductInfoTree{TypeId: typeid, Name: name, ParentId: parentid}
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
	//创建数据表
	//initTable()
	router:=gin.Default()
	//查找，参数为typeid
	//product为寻找范围
	db.Find(&product)
	router.GET("/search/:typeid",Search)
	//添加，参数为typeid,name,parentid
	router.POST("/add", Add)
	router.Run(":8080")
}