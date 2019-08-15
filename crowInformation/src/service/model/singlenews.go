package model

import (
	"service/dao"
	"fmt"
)

// Id 设置成int64才行，不然一直报错
type News struct {
	Id      int64
	Title   string
	Autor   string
	Text    string
	Publishdate string
}

type UserInfo struct {
	Id       int64
	Username string
	Password string
}

func AddNews(n *News) {
	fmt.Println("create news")
	dao.CreateNews(n)
}

func AddUserInfo(u *UserInfo) {
	fmt.Println("create user_info")
	dao.CreateNews(u)
}

