package utils

import (
	"crypto/ed25519"
	"encoding/json"
	"time"

	"github.com/cristalhq/jwt/v3"
	log "github.com/dhawton/log4g"
	"github.com/nzvirtual/go-api/lib/database/models"
)

type userClaims struct {
	jwt.RegisteredClaims
	ID   uint
	User models.User
}

type structJWTOptions struct {
	Algo       string
	PrivateKey ed25519.PrivateKey
	PublicKey  ed25519.PublicKey
	Duration   time.Duration
}

var options structJWTOptions

func SetJWTOptions(o *structJWTOptions) {
	options = *o
}

func CreateJWTToken(user *models.User) (string, error) {
	signer, err := jwt.NewSignerEdDSA(options.PrivateKey)
	if err != nil {
		return "", err
	}

	builder := jwt.NewBuilder(signer)

	claims := &userClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  []string{"nzv-api"},
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(options.Duration)),
		},
		ID:   user.ID,
		User: *user,
	}

	token, err := builder.Build(claims)
	if err != nil {
		j, _ := json.Marshal(claims)
		log.Category("jwt/CreateJWTToken").Error("Failed to build token with claims " + string(j))
		return "", err
	}

	return token.String(), nil
}
