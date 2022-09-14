package usecases

import (
	"context"
	"errors"

	"github.com/hoanggggg5/shop-pkg/gpa"
	"github.com/hoanggggg5/shop-pkg/infrastructure/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Usecase[V schema.Tabler] interface {
	First(context context.Context, filters ...gpa.Filter) (*V, error)
	Find(context context.Context, filters ...gpa.Filter) (*V, error)
	Create(context context.Context, model interface{}) error
	Updates(context context.Context, model interface{}, value interface{}, filters ...gpa.Filter)
}

type usecase[V schema.Tabler] struct {
	repository repository.Repository[V]
}

func (u usecase[V]) First(context context.Context, filters ...gpa.Filter) (model *V, err error) {
	if err := u.repository.First(context, &model, filters...); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return
}

func (u usecase[V]) Find(context context.Context, filters ...gpa.Filter) (model *V, err error) {
	if err := u.repository.Find(context, &model, filters...); errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return
}

func (u usecase[V]) Create(context context.Context, model interface{}) (err error) {
	if err := u.repository.Create(context, model); errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if err != nil {
		panic(err)
	}

	return
}

func (u usecase[V]) Updates(context context.Context, model interface{}, value interface{}, filters ...gpa.Filter) {
	if err := u.repository.Update(context, model, value, filters...); errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}
}
