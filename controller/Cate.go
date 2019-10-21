package controller

import "github.com/guygubaby/hotlist-server/model"

func InitAllCates() []model.Category {
	var cates []model.Category
	c := model.Category{1,"知乎"}
	cates = append(cates, c)
	return cates
}
