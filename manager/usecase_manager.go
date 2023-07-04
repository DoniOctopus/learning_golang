package manager

import (
	"goclean/usecase"
	"sync"
)

type UsecaseManager interface {
	GetServiceUsecase() usecase.ServiceUsecase
	GetUserUsecase() usecase.UserUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	svcUsecase usecase.ServiceUsecase
	usrUsecase usecase.UserUsecase
}

var onceLoadServiceUsecase sync.Once
var onceLoadUserUsecase sync.Once

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

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
