package usecase

import "goclean/repo"

type LoginUsecase interface {
	
}

type loginUsecase struct {
	userRepo repo.UserRepo
}

func NewLoginUsecase(userrepo repo.UserRepo) LoginUsecase {
	return &loginUsecase{
		userRepo: userrepo,
	}
}
