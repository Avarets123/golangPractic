package config

type Config struct {
	Env        string           `yaml:"env" env-default:"dev"`
	Server     ServerConfig     `yaml:"server" env-required:"true"`
	LogLevel   string           `yaml:"log_level" env-default:"INFO"`
	DbPostgres DbPostgresConfig `yaml:"db_postgres"`
}

type DbPostgresConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"localhost"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"db_name"`
}

type ServerConfig struct {
	Host string `yaml:"host" env-required:"true"`
	Port int    `yaml:"port" env-default:"4002"`
}

func NewConfig() *Config {
	return &Config{}
}
