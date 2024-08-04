package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Nama      string `json:"nama"`
	Userid    string `json:"userid"`
	EmailAddr string `json:"email_addr"`
	Role      string `json:"role"`
	UnikId    string `json:"unik_id"`
}

type jwtClaimsPendamping struct {
	jwt.RegisteredClaims
	Userid      string `json:"userid"`
	Nama        string `json:"nama"`
	EmailAddr   string `json:"email_addr"`
	Role        string `json:"role"`
	UnikId      string `json:"unik_id"`
	IdLembaga   string `json:"id_lembaga"`
	NamaLembaga string `json:"nama_lembaga"`
	NoHp        string `json:"no_hp"`
}

func (w *JwtWrapper) GenerateAccessToken() (signedToken string, err error) {
	// expirationTime := time.Now().Add(5 * time.Minute)

	expirationTime := time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix()

	claims := &jwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("err signature invalid")
		}

		return nil, errors.New("err bad request")
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	return claims, nil
}
