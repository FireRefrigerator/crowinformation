// crowInformation project main.go
package main

import (
	_ "service/routers"
	_ "service/model"
	_ "service/dao"

	"github.com/astaxie/beego"
	"fmt"
	"service/dao"
	"service/model"
)

func main() {
	fmt.Println("Hello crow news!")
	// dao.Connect(new(model.UserInfo))
	dao.Connect(new(model.News))
	beego.Run()
}
