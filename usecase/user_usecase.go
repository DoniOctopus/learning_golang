package usecase

import (
	"goclean/model"
	"goclean/repo"
)

type UserUsecase interface {
	// GetUserById(int) (*model.UserModel, error)
	InsertUser(*model.UserModel) error
}

type userUsecaseImpl struct {
	usrRepo repo.UserRepo
}

// func (svcUsecase *serviceUsecaseImpl) GetServiceById(id int) (*model.ServiceModel, error) {
// 	return svcUsecase.svcRepo.GetServiceById(id)
// }

func (usrUsecase *userUsecaseImpl) InsertUser(usr *model.UserModel) error {
	return usrUsecase.usrRepo.InsertUser(usr)
}

func NewUserUseCase(usrRepo repo.UserRepo) UserUsecase {
	return &userUsecaseImpl{
		usrRepo: usrRepo,
	}
}
