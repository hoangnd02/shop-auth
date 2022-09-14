package utils

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	GJWT "github.com/golang-jwt/jwt"
	"github.com/hoanggggg5/shop/models"
	"github.com/zsmartex/pkg/jwt"
)

func RandomUID() string {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000
	max := 9999999999
	return fmt.Sprintf("UID%v", rand.Intn(max-min+1)+min)
}

func RandomCode() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return fmt.Sprintf("%v", rand.Intn(max-min+1)+min)
}

func CheckJWT(token string) (*jwt.Auth, error) {
	ks := jwt.KeyStore{}
	ks.LoadPublicKeyFromFile("./public.key")
	j, err := jwt.ParseAndValidate(token, ks.PublicKey)

	if err != nil {
		return nil, err
	}

	return &j, nil
}

func GenerateJWT(user *models.User) (string, error) {
	ks := jwt.KeyStore{}
	ks.LoadPrivateKey("./private.key")
	token, err := jwt.ForgeToken(user.UID, user.Email, user.Role, sql.NullString{}, 3, ks.PrivateKey, GJWT.MapClaims{})

	if err != nil {
		return "", err
	}

	return token, nil
}
