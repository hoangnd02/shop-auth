package usecases

import (
	"github.com/hoanggggg5/shop-pkg/infrastructure/repository"
	"github.com/hoanggggg5/shop/internal/models"
)

type userUsecase struct {
	usecase[models.User]
}

type UserUsecase interface {
	Usecase[models.User]
}

func NewUserUsecase(repo repository.Repository[models.User]) UserUsecase {
	return userUsecase{
		usecase: usecase[models.User]{
			repository: repo,
		},
	}
}
