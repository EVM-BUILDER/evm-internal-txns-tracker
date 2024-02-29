package config

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	RestServer struct {
		Port string
	}
}

var conf *Configuration

func Init(confPath string) {
	conf = &Configuration{}
	viper.SetConfigFile(confPath)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Warning(err, "config file not found")
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Panic(err, "Error Unmarshal Viper Config")
	}
	logrus.SetReportCaller(true)
	spew.Dump("vdex: ", conf)
}

func GetConfig() *Configuration {
	return conf
}
