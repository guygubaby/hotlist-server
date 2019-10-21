package controller

import "github.com/guygubaby/hotlist-server/model"


func InitAllTypes() []model.Category {
	var cates []model.Category
	c1 := model.Category{1,"知乎"}
	c2 := model.Category{2,"虎扑"}
	c3 := model.Category{3,"v2ex"}
	c4 := model.Category{4,"微博"}
	c5 := model.Category{5,"github"}
	c6 := model.Category{6,"贴吧"}
	c7 := model.Category{7,"豆瓣"}
	c8 := model.Category{8,"天涯"}
	c9 := model.Category{9,"百度"}
	c10 := model.Category{10,"36kr"}
	c11 := model.Category{11,"好奇心"}
	c12 := model.Category{12,"果壳"}
	c13 := model.Category{13,"虎嗅"}
	cates = append(cates, c1,c2,c3,c4,c5,c6,c7,c8,c9,c10,c11,c12,c13)
	return cates
}
