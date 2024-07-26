package main

type Config struct {
	LocalServer struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"localServer"`
	ProxyServer struct {
		Server1 string `yaml:"server1"`
		Server2 string `yaml:"server2"`
		Server3 string `yaml:"server3"`
	} `yaml:"ProxyServer"`
	Balancer struct {
		Method string `yaml:"method"`
	} `yaml:"balancer"`
}
