package domain

import (
	"os"
	"strconv"
)

type Config struct {
	Port       string
	RedisHost  string
	RedisPass  string
	BlockSize  int
	BlockCount int
	Seed       rune
}

func NewConfig() *Config {
	var err error
	cfg := &Config{
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPass: os.Getenv("REDIS_PASSWORD"),
	}

	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		cfg.Port = ":80"
	}

	BlockSize, err := strconv.ParseInt(os.Getenv("BLOCK_SIZE"), 10, 32)
	if err != nil {
		BlockSize = 4
	}

	BlockCount, err := strconv.ParseInt(os.Getenv("BLOCK_COUNT"), 10, 32)
	if err != nil {
		BlockCount = 3
	}

	Seed, err := strconv.ParseInt(os.Getenv("SEED"), 10, 32)
	if err != nil {
		Seed = 100
	}

	cfg.BlockSize = int(BlockSize)
	cfg.BlockCount = int(BlockCount)
	cfg.Seed = rune(Seed)
	return cfg
}
