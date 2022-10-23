package config

type Config struct {
	Logger  LoggerConf
	DSN     string
	HTTPSrv string
	LogName string
}

type LoggerConf struct {
	Level   string
	LogFile string
}

func NewConfig() *Config {
	return &Config{}
}

func Default() *Config {
	return &Config{
		Logger: LoggerConf{
			Level:   "INFO",
			LogFile: "./log.txt",
		},
		DSN:     "root:masterkey@/users",
		HTTPSrv: "127.0.0.1:1122",
		LogName: "lrus",
	}
}
