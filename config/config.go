package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config ConfigList

// main関数より先に実行される
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

// コンフィグファイルの読み込み
func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
		Static: cfg.Section("app").Key("static").String(),
	}
}
