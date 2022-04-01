package config

type ClientConfig struct {
	ServerHost string
	ServerPort string
	TlsEnable  bool
}

var ClientConfiguration ClientConfig

func Init(enableTls bool, host string, port string) {
	ClientConfiguration = ClientConfig{
		ServerHost: "127.0.0.1",
		ServerPort: "8972",
		TlsEnable:  enableTls,
	}
	if len(host) != 0 {
		ClientConfiguration.ServerHost = host
	}
	if len(port) != 0 {
		ClientConfiguration.ServerPort = port
	}
}
