package usecase

import (
	"goclean/model"
	"goclean/repo"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	// GetUserById(int) (*model.UserModel, error)
	InsertUser(*model.UserModel) error
	GetUserByName(string) (*model.UserModel, error)
	
}

type userUsecaseImpl struct {
	usrRepo repo.UserRepo
}

// func (usrUsecase *UserUsecaseImpl) GetUserById(id int) (*model.UserModel, error) {
// 	return usrUsecase.usrRepo.GetUserById(id)
// }

func (usrUsecase *userUsecaseImpl) GetUserByName(name string) (*model.UserModel, error) {
	return usrUsecase.usrRepo.GetUserByName(name)
}


func (usrUsecase *userUsecaseImpl) InsertUser(usr *model.UserModel) error {
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.Password = string(hashedPassword)

	return usrUsecase.usrRepo.InsertUser(usr)
}

func NewUserUseCase(usrRepo repo.UserRepo) UserUsecase {
	return &userUsecaseImpl{
		usrRepo: usrRepo,
	}
}
