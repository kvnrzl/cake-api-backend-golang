package repository

import (
	"backend-engineer-test-privy/model"
	"context"
	"database/sql"
)

type CakeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, cake *model.Cake) (uint, error)
	GetByID(ctx context.Context, tx *sql.Tx, id uint) (*model.Cake, error)
	GetAll(ctx context.Context, tx *sql.Tx) ([]*model.Cake, error)
	Update(ctx context.Context, tx *sql.Tx, cake *model.Cake) error
	Delete(ctx context.Context, tx *sql.Tx, id uint) error
}
