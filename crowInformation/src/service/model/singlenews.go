package model

import (
	"service/dao"
	"fmt"
)

type News struct {
	id string
	title string
	autor string
	text string
	publishdate string
}

func AddNews(n News) {
	fmt.Println("create news")
	dao.CreateNews(n)
}

