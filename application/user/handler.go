package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sikab-biz-test/domain"
	"sikab-biz-test/infrastructure"
)

func SaveToDB(result <-chan domain.User) {
	user := <-result
	//fmt.Printf("user: %+v", user.Addresses)
	//os.Exit(1)
	userRepo := infrastructure.InstantiateUserRepo()
	err := userRepo.SaveUserToDb(user)
	if err != nil {
		log.Printf("Error saving user data to db with following structure %+v, err is %s", user, err)
	}
}

func GetUserById(ctx echo.Context) error {
	id := ctx.Param("id")
	idUuid, err := uuid.Parse(id)
	if err != nil {
		log.Printf("There was a problem parsing uuid %s", id)
	}

	userRepo := infrastructure.InstantiateUserRepo()
	user, exists, err := userRepo.GetUserById(idUuid)
	if err != nil {
		log.Printf("There was a problem retrieving data for user with id %s", id)
	}
	if !exists {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	return ctx.JSON(http.StatusOK, user)
}
