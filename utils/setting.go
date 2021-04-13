package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey string
	SecretKey string
	Bucket    string
	QnSever   string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件错误，请检查路径")
	}

	LoadServer(file)
	LoadData(file)
	LoadQn(file)
}

func LoadQn(file *ini.File) {

	AccessKey = file.Section("qn").Key("AccessKey").MustString("")
	SecretKey = file.Section("qn").Key("SecretKey").MustString("")
	Bucket = file.Section("qn").Key("Bucket").MustString("")
	QnSever = file.Section("qn").Key("QnSever").MustString("")
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("HttpPort").MustString("tsUha0jwt")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("tsuhao")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("Dkhg1500")
	DbName = file.Section("database").Key("DbName").MustString("blog")

}
