package main

type Config struct {
	Logger Logger `toml:"logger"`
	DB     DB     `toml:"db"`
	Server Server `toml:"server"`
}

func NewConfig() *Config {
	return &Config{}
}

type Logger struct {
	Level string `toml:"level"`
}
type DB struct {
	InMemoryStorage bool   `toml:"in_memory_storage"`
	URI             string `toml:"uri"`
}
type Server struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}
