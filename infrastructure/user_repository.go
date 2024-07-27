package infrastructure

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sikab-biz-test/domain"
)

type UserRepository struct {
	databaseConn *gorm.DB
}

func InstantiateUserRepo() *UserRepository {
	return &UserRepository{databaseConn: GetDBConn()}
}

func (repo *UserRepository) SaveUserToDb(user domain.User) error {
	if err := repo.databaseConn.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) GetUserById(id uuid.UUID) (user domain.User, exists bool, err error) {
	result := repo.databaseConn.Preload("Addresses").Where("id = ?", id).Find(&user)
	if result.Error != nil {
		err = result.Error
		return
	}

	exists = result.RowsAffected > 0
	return
}
