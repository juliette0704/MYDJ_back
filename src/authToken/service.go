package authtoken

import (
	"errors"
	"log"
	"mydj_server/config"
	"mydj_server/src/entity"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint64) (string, error) {
	jwtKey := []byte(config.GetConfig().Token.TokenKey)
	log.Println(jwtKey)
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Expiration après 24 heures
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	log.Println("token login = ", tokenString)
	return tokenString, nil
}

func ValidateToken(tokenString string) (uint64, error) {
	log.Println("token repris = ", tokenString)
	jwtKey := []byte(config.GetConfig().Token.TokenKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, jwt.ErrSignatureInvalid
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, jwt.ErrInvalidKey
	}
	userID, ok := claims["userID"]
	if !ok {
		return 0, jwt.ErrInvalidKey
	}
	switch userID := userID.(type) {
	case float64:
		return uint64(userID), nil
	case string:
		id, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			return 0, err
		}
		return id, nil
	default:
		return 0, errors.New("userID format not recognized")
	}
}

func GetUserByID(userID uint64) (*entity.User, error) {
	var user entity.User
	db, err := config.ReturnDB()
	if err != nil {
		return nil, err
	}

	if err := db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
