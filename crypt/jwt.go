package crypt

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type (
	JWT struct {
		SecretEnv string // your ENV key name for the secret | has to be set before calling 'Sign()' or 'Validate()'
		TokenStr  string /* populated after a successfull 'Sign()' call |
		when calling 'Validate()' as a standalone,
		you should set TokenStr manually to the token string you obtained from the client*/
	}

	StandardClaims = jwt.StandardClaims

	Token = jwt.Token
)

var HMAC_HS512 = jwt.SigningMethodHS512

// Signs JWT token with given custom claims and signing method
func (c *JWT) Sign(claims jwt.Claims, method jwt.SigningMethod) error {
	t, err := jwt.NewWithClaims(method, claims).SignedString([]byte(os.Getenv(c.SecretEnv)))
	if err != nil {
		return err
	}

	c.TokenStr = t
	return nil
}

// Validates given token string with your custom claims type and keyfunc
func (c *JWT) Validate(claims jwt.Claims, keyfunc jwt.Keyfunc) (jwt.Claims, error) {
	t, err := jwt.ParseWithClaims(c.TokenStr, claims, keyfunc)
	if err != nil {
		fmt.Println("hier")
		return nil, err
	}

	cl, ok := t.Claims.(jwt.Claims)
	if ok && t.Valid {
		return cl, nil
	}

	return nil, errors.New("Given token is invalid")
}
