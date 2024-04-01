package model

import "reflect"

type Config struct {
	Github Github
	Proxy  Proxy
}

type Proxy struct {
	Address string
	Port    string
}

type Github struct {
	Url   string
	Start int
	End   int
	Agent string
}

var config Config

func SetConfig(tmpConfig Config) {

	if isEmpty := reflect.DeepEqual(
		reflect.ValueOf(tmpConfig),
		reflect.Zero(reflect.TypeOf(tmpConfig))); isEmpty {

		panic("Setting system configuration failed, Github info is null!")
	}

	config = tmpConfig
}

func GetConfig() Config {

	return config
}
