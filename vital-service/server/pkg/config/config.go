package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppName       string        `mapstructure:"APP_NAME"`
	Env           string        `mapstructure:"ENV"`
	Port          string        `mapstructure:"PORT"`
	ReadTimeout   time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout  time.Duration `mapstructure:"WRITE_TIMEOUT"`
	RedisURL      string        `mapstructure:"REDIS_URL"`
	RabbitMQURL   string        `mapstructure:"RABBITMQ_URL"`
	RiskAPIURL    string        `mapstructure:"RISK_API_URL"`
	JWTSecret     string        `mapstructure:"JWT_SECRET"`
	AllowedOrigin string        `mapstructure:"ALLOWED_ORIGIN"`
}

var Cfg *Config

func LoadConfig() {
	viper.SetEnvPrefix("VITALS")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Optional: Load from .env if present
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
	}

	viper.SetDefault("APP_NAME", "vitals-service")
	viper.SetDefault("PORT", "3000")
	viper.SetDefault("READ_TIMEOUT", 10*time.Second)
	viper.SetDefault("WRITE_TIMEOUT", 10*time.Second)
	viper.SetDefault("REDIS_URL", "redis://localhost:6379")
	viper.SetDefault("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	viper.SetDefault("RISK_API_URL", "http://localhost:8080")
	viper.SetDefault("JWT_SECRET", "your_should_never_use_this_in_production")
	viper.SetDefault("ALLOWED_ORIGIN", "*")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("ERROR: Failed to unmarshal config: %v", err)
	}

	Cfg = &config
	log.Printf("INFO: config loaded: %+v", Cfg)
}
