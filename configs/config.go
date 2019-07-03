package configs

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"os"
	"path/filepath"
)

type singleton struct {
	configInfo map[string]interface{}
}

var instance *singleton

func init() {
	if instance == nil {
		instance = new(singleton)
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		err := config.LoadFile(dir + "/config.yaml")
		if err != nil {
			fmt.Println("加载配置文件错误！！请确认当前目录下包含config.yaml文件")
			fmt.Println(err)
		}
		conf := config.Map()
		instance.configInfo = conf

		fmt.Println(conf)
	}

}

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		err := config.LoadFile(dir + "/config.yaml")
		if err != nil {
			fmt.Println("加载配置文件错误！！请确认当前目录下包含config.yaml文件")
			fmt.Println(err)
		}
		conf := config.Map()
		instance.configInfo = conf
		fmt.Println(conf)

	}
	return instance
}

func Config() map[string]interface{} {

	return GetInstance().configInfo
}
