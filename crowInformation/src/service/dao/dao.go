package dao

import (
	"service/model"
)

type DaoAPI interface {
	CreateN(n model.News)
	GetNById(id string) model.News
	DeleteNById(id string)
	UpdateN(n model.News)
}

var daoAPI DaoAPI

func SetDaoAPI(d DaoAPI) {
	daoAPI = d
}

func GetDaoAPI() {
	return daoAPI
}

func GetCurrentDaoAPI() {
	if daoAPI == nil {
		daoAPI = &MysqlDao{}
	}
}

func CreateNews(n model.News) {
	GetCurrentDaoAPI()
	daoAPI.CreateN(n)
}

func UpdateNews(n model.News) {
	GetCurrentDaoAPI()
	daoAPI.UpdateN(n)
}

func DeleteNewsByid(id string) {
	GetCurrentDaoAPI()
	daoAPI.DeleteNById(id)
}

func getNewsById(id string)  {
	GetCurrentDaoAPI()
	daoAPI.GetNById(id)
}
