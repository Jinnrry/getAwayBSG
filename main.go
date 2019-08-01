package main

import (
	"flag"
	"fmt"
	"getAwayBSG/configs"
	"getAwayBSG/entrance"
)

// 申明配置变量
var (
	help           bool
	config         string
	lianjia_ershou bool
	lianjia_zufang bool
	zhilian        bool
	clean          bool
	info           bool
	infoSaveTo     string
)

func init() {
	flag.BoolVar(&help, "help", false, "显示帮助")
	flag.StringVar(&config, "config", "./config.yaml", "设置配置文件")
	flag.BoolVar(&lianjia_ershou, "lianjia_ershou", false, "抓取链家二手房数据")
	flag.BoolVar(&lianjia_zufang, "lianjia_zufang", false, "抓取链家租房数据")
	flag.BoolVar(&zhilian, "zhilian", false, "抓取智联招聘数据")
	flag.BoolVar(&clean, "clean", false, "清理缓存")
	flag.BoolVar(&info, "info", false, "保存抓取状态")
	flag.StringVar(&infoSaveTo, "info_save_to", "./numlog.txt", "输入状态文件保存位置")
}

func main() {
	flag.Parse()
	//初始化配置信息，同时输出配置信息
	if config != "" {
		configs.SetConfig(config)
	}
	fmt.Println(configs.Config())


	//进入不同入口
	if help {
		flag.Usage()
	} else if lianjia_ershou {
		entrance.Start_lianjia_ershou()
	} else if lianjia_zufang {
		entrance.Start_LianjiaZufang()
	} else if zhilian {
		entrance.Start_zhilian()
	} else if clean {
		entrance.Start_clean()
	} else if info {
		entrance.Start_info(infoSaveTo)
	} else {
		flag.Usage()
	}

}
