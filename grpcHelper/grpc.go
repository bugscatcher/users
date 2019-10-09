package grpcHelper

import "fmt"

type ServerConf struct {
	Port    int    `mapstructure:"port"`
	Address string `mapstructure:"address"`
}

func (s ServerConf) Addr() string {
	return fmt.Sprintf("%s:%d", s.Address, s.Port)
}