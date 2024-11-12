package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBName     string
	DBUser     string
	Passwd     string
	DBAddress  string
}

func InitConfigFunc() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnvElement("PUBLIC_HOST", "http://localhost:"),
		Port:       getEnvElement("PORT", ":8080"),
		DBUser:     getEnvElement("DBName", "root"),
		DBName:     getEnvElement("DBName", "econ"),
		Passwd:     getEnvElement("Passwd", "mypassword"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnvElement("DBHost", "127.0.0.1"), getEnvElement("DBPORT", "3306")),
	}
}

var InitConfig = InitConfigFunc()

func getEnvElement(key, elsee string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return elsee
}
