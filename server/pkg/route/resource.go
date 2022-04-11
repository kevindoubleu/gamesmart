package route

import (
	"github.com/kevindoubleu/gamesmart/pkg/config"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/controller/helper"
	"github.com/kevindoubleu/gamesmart/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type resources struct {
	secrets	*model.Secrets

	dbConn	*mongo.Database
	
	jwtSvc	helper.JWTService
	userSvc	helper.UserService
}

func defaultResources() resources {
	db := config.ConnectDb(constants.DB)
	secrets := config.GenerateSecrets()

	return resources{
		secrets: secrets,

		dbConn: db,

		jwtSvc: helper.NewJWTService(secrets),
		userSvc: helper.NewUserService(db.Collection(constants.DB_TBL_USERS)),
	}
}
