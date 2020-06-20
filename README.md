# 基于gin框架实现商品的无限级分类管理
基于golang的gin编程框架和gorm基础库，实现一个商品的无限级分类管理功能

## 一、整体项目时间安排计划表

| 时间 |                       规划                        |
| :--: | :-----------------------------------------------: |
| 6.17 |      通过看文档，看视频，学习golang基本语法       |
| 6.18 |              学习golang复杂数据类型               |
| 6.19 | 使用gin框架，建立数据表，连接数据库，完成查找功能 |
| 6.20 |        完成增加功能，对于接口功能进行测试         |



## 二、golang语法自学记录

| 时间 |                内容                 |
| :--: | :---------------------------------: |
| 6.16 |      goland软件安装，环境配置       |
| 6.17 |    学习golang基础语法、数据类型     |
| 6.18 |    学习函数、结构体、切片、接口     |
| 6.19 | 阅读gin文档，看视频学习，设计数据库 |
| 6.20 | 学习golang的post与get请求，完善程序 |



## 三、无限级别分类功能数据库设计

### 1、思路：

表中设置三个字段，**id**为自增主键，**type_id**是商品标识号，**parent_id**是商品的上一级标识号，**name**是商品名称，通过**type_id**和**parent_id**实现树状结构，达到商品的无限级分类功能

### 2、增加商品：

需要提供的参数是新增商品的**type_id**，**name**，**parent_id**

注意：新增商品的**parent_id**和直接父类的**type_id**保持一致

### 3、查找直接子类商品：

需要提供的参数是父类的**type_id**，程序返回其所有直接子级商品

### 4、建表：

```go
type ProductInfo struct {
	ID      int `json:"id"`
	TypeId   string `json:"type_id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}
```

### 5、存储过程：

```go
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
```

### 6、测试数据表：

|  ID  | Type_Id |    Name    | Parent_Id |
| :--: | :-----: | :--------: | :-------: |
|  1   |   101   |    家电    |    100    |
|  2   |   102   |    数码    |    100    |
|  3   |   103   |    食品    |    100    |
|  4   |   104   |    电视    |    101    |
|  5   |   105   |   笔记本   |    102    |
|  6   |   106   |    可乐    |    103    |
|  7   |   107   | 全面屏电视 |    104    |
|  8   |   108   | 商务笔记本 |    105    |
|  9   |   109   |   牛肉干   |    103    |
|  10  |   110   |    手机    |    102    |

### 7、数据库信息表：

|   Field   |     Type     | Null | Key  | Default |     Extra      |
| :-------: | :----------: | :--: | :--: | :-----: | :------------: |
|    ID     |   int(11)    |  NO  | PRI  |  NULL   | auto_increment |
|  Type_Id  | varchar(255) | YES  |      |  NULL   |                |
|   Name    | varchar(255) | YES  |      |  NULL   |                |
| Parent_Id | varchar(255) | YES  |      |  NULL   |                |



## 四、基于curl命令的测试用例脚本

### 增加：

```
curl --location --request POST 'http://127.0.0.1:8080/add' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'typeid=111' \
--data-urlencode 'name=百事可乐' \
--data-urlencode 'parentid=106'
```

### 查找：

```
curl --location --request GET 'http://127.0.0.1:8080/search/100' \
--header 'Content-Type: application/x-www-form-urlencoded'
```

