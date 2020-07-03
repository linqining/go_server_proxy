package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"log"
)

type Cfg struct {
	Domain      string	`yaml:"domain"`
	Port	string    `yaml:"port"`
	CertPathTLS	string	`yaml:"cert_path_tls"`
	KeyPathTLS	string	`yaml:"key_path_tls"`
}


func GetConfig() Cfg {

	file, err := os.Open("cfg.yaml")

	if err != nil {
		log.Print("读取配置文件失败")
		os.Exit(1)
		//panic(err)
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Print("读取配置文件失败")
		os.Exit(1)
		//panic(err)
	}

	cfg := Cfg{}

	err = yaml.Unmarshal(bytes, &cfg)

	if err != nil {
		log.Print("读取配置文件失败")
		os.Exit(1)
		//panic(err)
	}
	return cfg
}

