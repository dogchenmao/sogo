package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	StartUserID uint32
	EndUserID   uint32
	Password    string
}

var GlobalConfig *Config

//判断一个文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (c *Config) Reload() {
	if confFileExists, _ := PathExists("config/config.json"); confFileExists != true {
		return
	}
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}
	//将json数据解析到struct中
	err = json.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}

}

func init() {
	GlobalConfig = &Config{
		StartUserID: 1,
		EndUserID:   1,
		Password:    "",
	}

	GlobalConfig.Reload()
}
