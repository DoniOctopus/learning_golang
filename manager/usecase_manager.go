package manager

import (
	"goclean/usecase"
	"sync"
)

type UsecaseManager interface {
	GetServiceUsecase() usecase.ServiceUsecase
	GetUserUsecase() usecase.UserUsecase
	GetLoginUsecase() usecase.LoginUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	svcUsecase usecase.ServiceUsecase
	usrUsecase usecase.UserUsecase
	lgusecase usecase.LoginUsecase
}

var onceLoadServiceUsecase sync.Once
var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once

func (um *usecaseManager) GetServiceUsecase() usecase.ServiceUsecase {
	onceLoadServiceUsecase.Do(func() {
		um.svcUsecase = usecase.NewServiceUseCase(um.repoManager.GetServiceRepo())
	})
	return um.svcUsecase
}

func (um *usecaseManager) GetUserUsecase() usecase.UserUsecase {
	onceLoadUserUsecase.Do(func() {
		um.usrUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.usrUsecase
}

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUsecase {
	onceLoadLoginUsecase.Do(func() {
		um.lgusecase = usecase.NewLoginUsecase(um.repoManager.GetUserRepo())
	})
	return um.lgusecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
