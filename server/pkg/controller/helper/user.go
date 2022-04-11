package helper

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	// returns empty string and an error if token is invalid
	GetUsernameFromSession(*jwt.Token) (string, error)

	// returns nil if username doesn't exist
	GetUserByUsername(context.Context, string) *model.User
}

func NewUserService(usersTable *mongo.Collection) UserService {
	return myUserService{
		usersTable: usersTable,
	}
}

type myUserService struct {
	usersTable	*mongo.Collection
}

func (svc myUserService) GetUsernameFromSession(token *jwt.Token) (string, error) {
	if token == nil {
		return "", constants.ErrJWTNil
	}

	session, ok := token.Claims.(*model.SessionJWT)
	if !ok {
		return "", constants.ErrJWTInvalid
	}

	return session.Username, nil
}

func (svc myUserService) GetUserByUsername(c context.Context, username string) *model.User {
	// find user in db
	row := svc.usersTable.FindOne(c, bson.M{"username": username})
	if row.Err() == mongo.ErrNoDocuments {
		return nil
	}

	// return the user
	var user model.User
	row.Decode(&user)
	return &user
}
