package service

import (
	"backend-engineer-test-privy/model"
	"context"
)

type CakeService interface {
	CreateCake(ctx context.Context, cake *model.Cake) (*model.Cake, error)
	GetCakeByID(ctx context.Context, id uint) (*model.Cake, error)
	GetAllCakes(ctx context.Context) ([]*model.Cake, error)
	UpdateCake(ctx context.Context, cake *model.Cake) (*model.Cake, error)
	DeleteCake(ctx context.Context, id uint) error
}
