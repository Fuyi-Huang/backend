package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Mechine_id int
	Db_master  map[string]string
	Db_slave   map[string]string
	Redis      map[string]string
)

func ReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config failed")
		fmt.Println(err)
	}
	Db_master = viper.GetStringMapString("database.master")
	Db_slave = viper.GetStringMapString("database.slave")
	Mechine_id = viper.GetInt("mechine_id")
	Redis = viper.GetStringMapString("redis")
}
