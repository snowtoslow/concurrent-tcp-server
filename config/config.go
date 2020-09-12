package config

import (
	"os"
)

type ServersConfig struct {
	RemoteServerPort string
	TcpServerPort string
	Host string
}


func New() *ServersConfig{
	return &ServersConfig{
		RemoteServerPort : getEnv("REMOTE_SERVER",""),
		TcpServerPort : getEnv("TCP_SERVER",""),
		Host: getEnv("HOST",""),
	}
}



func getEnv(key string, defaultValue string) (envVarValue string){
	if envVarValue,exist := os.LookupEnv(key);exist {
		return envVarValue
	}

	return defaultValue
}
