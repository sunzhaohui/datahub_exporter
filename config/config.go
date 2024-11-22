package config

import (
	"flag"
	"fmt"
	"log"
	"os/user"

	"github.com/go-ini/ini"
)

type ConfigStruct struct {
	SecretId  string
	SecretKey string
	Region    string
}

var Config *ConfigStruct

func init() {
	// 获取当前用户的家目录
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}
	// 构建文件路径
	filePath := fmt.Sprintf("%s/.datahub_exporter.conf", usr.HomeDir)
	var conf string
	flag.StringVar(&conf, "conf", filePath, "The conf of datahub_exporter")
	flag.Parse()
	fmt.Println("conf is:", conf)
	// 加载 INI 文件
	cfg, err := ini.Load(conf)
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}
	// 获取 [common] 部分的配置
	common := cfg.Section("common")
	secret_id := common.Key("secret_id").String()
	secret_key := common.Key("secret_key").String()
	region := common.Key("region").String()
	Config = &ConfigStruct{}
	Config.SecretId = secret_id
	Config.SecretKey = secret_key
	Config.Region = region

}
