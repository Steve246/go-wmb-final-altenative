package utils

import (
	"fmt"
	"livecode-wmb-2/config"
	"livecode-wmb-2/model"
	"time"

	"github.com/golang-jwt/jwt"
)

type Token interface {
	CreateAccesToken(cred *model.Credential) (string, error)
	VerrifyAccesToken(tokenString string) (jwt.MapClaims, error)
}

type token struct {
	cfg config.TokenConfig
}

// CreateAccesToken implements Token
func (t *token) CreateAccesToken(cred *model.Credential) (string, error) {
	now := time.Now().UTC()
	end := now.Add(t.cfg.AccesTokenLifeTime)
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: t.cfg.AplicationName,
		},
		Username: cred.Username,
		Email:    cred.Email,
	}
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()
	token := jwt.NewWithClaims(
		t.cfg.JwtSigningMethod,
		claims,
	)
	return token.SignedString([]byte(t.cfg.JwtSignatureKey))

}

// VerrifyAccesToken implements Token
func (t *token) VerrifyAccesToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != t.cfg.JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(t.cfg.JwtSignatureKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != t.cfg.AplicationName {
		return nil, err
	}
	return claims, nil
}

func NewTokenService(cfg config.TokenConfig) Token {
	return &token{
		cfg: cfg,
	}
}
