// crowInformation project main.go
package main

import (
	_ "crowInformation/routers"
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("Hello crow news!")
	beego.Run()
}
