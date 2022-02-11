/*================================================================
*
*  文件名称：config.go
*  创 建 者: mongia
*  创建日期：2021年12月28日
*
================================================================*/

package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type mysql struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Database string `yaml:"Database"`
	Conns    int    `yaml:"Conns"`
}

type redis struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Password string `yaml:"Password"`
	Conns    int    `yaml:"Conns"`
}

type config struct {
	Host   string `yaml:"Host"`
	Port   int    `yaml:"Port"`
	Daemon bool   `yaml:"Daemon"`
	Mysql  mysql  `yaml:"Mysql"`
	Redis  redis  `yaml:"Redis"`
}

// Cfg global config
var Cfg config

func init() {
}

// LoadConfig load pharmacyerp.yaml file
func LoadConfig() error {
	config, err := ioutil.ReadFile("runtime/config/pharmacyerp.yaml")
	if err != nil {
		return err
	}
	yaml.Unmarshal(config, &Cfg)

	return nil
}
