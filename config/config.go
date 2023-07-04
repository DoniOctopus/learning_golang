package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Config struct {
	TokenConfig
}

func (c *Config) ReadConfigFile() error {

	tokenExpire, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))
	accessTokenLifeTime := time.Duration(tokenExpire) * time.Minute
	if err != nil {
		return errors.New("failed to convert token expire")
	}
	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("TOKEN_APP_NAME"),
		JwtSignatureKey:     os.Getenv("TOKEN_SECRET"),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfigFile()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
