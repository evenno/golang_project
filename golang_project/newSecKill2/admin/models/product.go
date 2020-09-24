package models

import (
	//"fmt"

	"github.com/astaxie/beego/logs"
)

type SecKillProduct struct {
	ProductId   int64  `orm:"pk;auto"`
	ProductName string `orm:"unique;size(60);default('')" form:"ProductName"  valid:"Required"`
	Total       int    `orm:"size(10);default(0)" form:"Total" valid:"Min(0)"`
	Status      int    `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
}

/*
	type Product struct {
		ProductId   int64  商品id
		ProductName string 商品名称
		Total       int    商品数量
		Status      int    商品状态（1.上架、2.下架）
	}
*/

// 创建模型
func NewProductModel() *SecKillProduct {
	return &SecKillProduct{}
}

func (this *SecKillProduct) GetProductList(slimit, elimit uint) (lists []SecKillProduct, err error) {
	if slimit == 0 && elimit == 0 {
		slimit = 0
		elimit = 20
	}
	num, err := DB.Raw("select * from sec_kill_product limit ?,?", slimit, elimit).QueryRows(&lists)
	if err != nil && num < 0 {
		logs.Warn("get Product list err : %v ", err)
		return
	}
	return
}

func (this *SecKillProduct) GetProductListByName(name string,limit uint)(num int64, result []SecKillProduct,err error) {
		if limit == 0 {
			limit=1
		}
		// 原生sql 获取数量
		num,err=DB.Raw("select* from sec_kill_product where product_name = ?  limit ?",name,limit).QueryRows(&result)
		if err !=nil {
			logs.Warn("get product by name err : %v", err)
			return
		}
		return
}

/**
*  添加产品
		o := orm.NewOrm()
		var user User
		user.Name = "slene"
		user.IsActive = true

		id, err := o.Insert(&user)
		if err == nil {
			fmt.Println(id)
		}
*/
func (this *SecKillProduct) InsertProduct(p *SecKillProduct)(num int64,err error){
	num,err=DB.Insert(p)  //num  是创建主键id值
	if err !=nil {
		logs.Warn("add product err : %v", err)
		return
	}
	return
}