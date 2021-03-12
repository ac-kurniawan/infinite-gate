package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"time"
)

type IAuthService interface {
	CreateToken(name string, email string, userId int) (string, error)
	VerifyToken(token string) (*Auth, error)
}

type AuthService struct {
	secret string
}

func (a *AuthService) CreateToken(name string, email string, userId int) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["email"] = email
	claims["userId"] = userId
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Encode token
	t, err := token.SignedString([]byte(a.secret))

	return t, err
}

func (a *AuthService) VerifyToken(token string) (*Auth, error) {
	payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			log.Errorf("[Token] - Token methode not valid")
			return nil, errors.New("[Token] - methode error")
		}

		return a.secret, nil
	})

	if err != nil || !payload.Valid {
		log.Errorf("[Token] - Invalid bearer token")
		return nil, errors.New("[Token] - Invalid bearer token")
	}

	claims, ok := payload.Claims.(jwt.MapClaims)
	exp := claims["exp"].(int64)

	if !ok {
		log.Errorf("[Token] - exp not set")
		return nil, errors.New("[Token] - exp not set")
	}

	deltaExp := time.Now().Unix() - exp

	if deltaExp < 0 {
		log.Errorf("[Token] - token expired")
		return nil, errors.New("[Token] - token expired")
	}

	return &Auth{
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}, nil
}

func NewAuthService(secret string) IAuthService {
	return &AuthService{secret}
}
