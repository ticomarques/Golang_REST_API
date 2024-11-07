package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = ""

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	//Abaixo o parse recebe o token, e uma funcao anonima que checa se o signingmethod está correto, se estiver errado retorna
	//um erro, se estiver certo retorna o secretkey (tipo interface salvo na variavel parsedToken)
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	//Checa se existe algum err, caso contrário parsedToken retornou com sucesso
	if err != nil {
		return 0, errors.New("could not parse token")
	}

	//checa se o token é valido (bool)
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("token not valid")
	}

	//Example how to extract payload from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims.")
	}

	//It shows how email and userId can be extracted from the token
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil

}
