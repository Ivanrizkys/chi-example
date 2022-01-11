package config

import "fmt"

type Config struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
}

func Get() *Config {
	conf := &Config{}
	conf.dbHost = "hahaha"
	conf.dbUser = "hahaha"
	conf.dbPassword = "hahaha"
	conf.dbPort = "hahaha"
	conf.dbName = "hahaha"
	fmt.Println(conf.dbHost)
	return conf
}

func SayHallo() {
	fmt.Println("hallo madafaka")
}
