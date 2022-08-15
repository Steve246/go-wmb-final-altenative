package config

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type ApiConfig struct {
	Url string
}

type GrpcConfig struct {
	GrpcUrl string
}
type DbConfig struct {
	DataSourceName string
}

type TokenConfig struct {
	AplicationName     string
	JwtSigningMethod   *jwt.SigningMethodHMAC
	JwtSignatureKey    string
	AccesTokenLifeTime time.Duration
}

type Config struct {
	ApiConfig
	DbConfig
	GrpcConfig
	TokenConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	grpcUrl := os.Getenv("GRPC_URL")

	// token config
	c.TokenConfig = TokenConfig{
		AplicationName:     "ENIGMA",
		JwtSigningMethod:   jwt.SigningMethodHS256,
		JwtSignatureKey:    "3N1GM4",
		AccesTokenLifeTime: 2 * time.Minute,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	c.DbConfig = DbConfig{DataSourceName: dsn}
	c.ApiConfig = ApiConfig{Url: api}
	c.GrpcConfig = GrpcConfig{GrpcUrl: grpcUrl}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
