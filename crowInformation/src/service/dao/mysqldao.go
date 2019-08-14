package dao

import (
	"service/model"
	"github.com/astaxie/beego/orm"
)

type MysqlDao struct {
	orm.Ormer
}

var mysqlDao MysqlDao

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(model.News))
	mysqlDao = orm.NewOrm()
}

func (d *MysqlDao) CreateN(n model.News){
	mysqlDao.Insert(n)
}

func (d *MysqlDao) GetNById(id string) model.News {
	return nil
}

func (d *MysqlDao) DeleteNById(id string) {

}

func (d *MysqlDao) UpdateN(n model.News) {

}
