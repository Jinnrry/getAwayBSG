package main

import (
	"fmt"
	"getAwayBSG/db"
)

func main() {

	fmt.Println("清除抓取状态（不清除状态的话爬虫会从上次停止位置继续抓取）")
	fmt.Println("请选择需要清除哪个爬虫的的状态数据：（输入数字）")
	fmt.Println("1.链家")
	fmt.Println("2.智联")
	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		db.SetLianjiaStatus(0)
		fmt.Println("Done!")
	} else if choice == 2 {
		db.SetZhilianStatus(0, 0)
		fmt.Println("Done!")
	} else {
		fmt.Println("选择错误！")
	}

}
