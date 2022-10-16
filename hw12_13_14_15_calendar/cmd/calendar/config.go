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
	StorageType string `toml:"storage_type"` // sql|in_memory
	URI         string `toml:"uri"`
}
type Server struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

const (
	SQLStorage      = "sql"
	InMemoryStorage = "in_memory"
)
