package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    string
	DBPath        string
	DeepSeekKey   string
	DeepSeekModel string
	DeepSeekBase  string
}

// Load 从 .env 文件和环境变量中加载配置。
// 查找顺序：当前目录 .env -> 项目根目录 .env -> 仅环境变量。
// 环境变量始终优先于 .env 文件中的值。
func Load() *Config {
	tryLoadEnvFile()

	return &Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		DBPath:        getEnv("DB_PATH", "data/eino-oa.db"),
		DeepSeekKey:   getEnv("DEEPSEEK_API_KEY", ""),
		DeepSeekModel: getEnv("DEEPSEEK_MODEL", "deepseek-chat"),
		DeepSeekBase:  getEnv("DEEPSEEK_BASE_URL", "https://api.deepseek.com/v1"),
	}
}

func tryLoadEnvFile() {
	candidates := []string{
		".env",
		filepath.Join("..", ".env"),
		filepath.Join("..", "..", ".env"),
	}
	for _, path := range candidates {
		if err := godotenv.Load(path); err == nil {
			log.Printf("已加载配置文件: %s", path)
			return
		}
	}
	log.Println("未找到 .env 文件，使用环境变量和默认值")
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
