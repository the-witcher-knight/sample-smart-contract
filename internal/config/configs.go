package config

import (
	"os"
)

type Config struct {
	Network      string
	Addr         string
	ContractAddr string
	SecretKey    string
}

func ReadConfigFromEnv() Config {
	return Config{
		Network:      os.Getenv("NETWORK"),
		Addr:         os.Getenv("ADDR"),
		ContractAddr: os.Getenv("CONTRACT_ADDR"),
		SecretKey:    os.Getenv("SK"),
	}
}
