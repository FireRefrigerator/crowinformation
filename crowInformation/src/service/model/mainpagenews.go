package model

import (
	"regexp"
	"fmt"
	"service/util"
)

const (
	URL_QUEUE    = "url_queue"
	URL_VIST_SET = "url_vist_set"
)

func ProductNewsUrl(htmlstr string) {
	fmt.Println("begin product news url")
	// 正则需要输出的字符串放在(.*)里，过滤/2019-08-12还需学习
	rule := `<a href="(http://sports.sina.com.cn/basketball/nba/.*.shtml)" target="_blank">`
	reg := regexp.MustCompile(rule)
	result := reg.FindAllStringSubmatch(htmlstr, -1)

	for i := 0; i < len(result); i++ {
		if !ISVist(URL_VIST_SET, result[i][1]) {
			fmt.Println("product news", result[i][1])
			PutinQueue(URL_QUEUE, result[i][1])
			AddToSet(URL_VIST_SET, result[i][1])
		}
	}

}

func ConsumeNewsUrl() {
	sigleUrl := PopFromQueue(URL_QUEUE)
	htmlstr := util.GetResStr(sigleUrl)
	fmt.Println(htmlstr)
	 news := News{Title:"aa",Text:"nihao",Publishdate:"2019/08/14",Autor:"bin"}
	 AddNews(&news)
	// user := UserInfo{Id: 3, Username: "wb", Password: "pp"}
	// 需要传入指针变量
	// AddUserInfo(&user)
	fmt.Println("add news success")
}



