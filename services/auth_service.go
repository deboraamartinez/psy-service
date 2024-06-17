package services

import (
	"context"
	"errors"
	"os"
	"psy-service/config"
	"psy-service/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	collection := config.DB.Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	return err
}

func (s *AuthService) Login(user *models.User) (string, error) {
	collection := config.DB.Collection("users")
	var dbUser models.User
	err := collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&dbUser)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := s.generateToken(dbUser)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) generateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.ID.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
