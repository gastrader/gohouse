package config

import "github.com/gastrader/gohouse/utils"

type JwtConfig struct {
	Secret []byte
}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		Secret: []byte(utils.GetIni("jwt_secret", "JWT_SECRET", "awesome")),
	}
}