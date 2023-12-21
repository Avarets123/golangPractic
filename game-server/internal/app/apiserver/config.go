package apiserver

type Config struct {
	Env      string       `yaml:"env" env-default:"dev"`
	Server   ServerConfig `yaml:"server" env-required:"true"`
	LogLevel string       `yaml:"log_level" env-default:"info"`
}

type ServerConfig struct {
	Host string `yaml:"host" env-required:"true"`
	Port int    `yaml:"port" env-default:"4002"`
}

func NewConfig() *Config {
	return &Config{}
}
