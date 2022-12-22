package mysql

import (
	"backend-engineer-test-privy/cake/repository"
	"backend-engineer-test-privy/model"
	"context"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

type CakeRepositoryImpl struct{}

func NewCakeRepositoryImpl() repository.CakeRepository {
	return &CakeRepositoryImpl{}
}

func (c *CakeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, cake *model.Cake) (uint, error) {
	sql := `INSERT INTO cakes (title, description, rating, image, created_at, updated_at) 
			VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	defer stmt.Close()

	now := time.Now()
	res, err := stmt.ExecContext(ctx, cake.Title, cake.Description, cake.Rating, cake.Image, now, now)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return uint(id), err
}

func (c *CakeRepositoryImpl) GetByID(ctx context.Context, tx *sql.Tx, id uint) (*model.Cake, error) {
	var cake model.Cake
	sql := `SELECT id, title, description, rating, image, created_at, updated_at FROM cakes 
			WHERE id = ? AND deleted_at IS NULL`
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, id).Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &cake, nil
}

func (c *CakeRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) ([]*model.Cake, error) {
	var cakes []*model.Cake
	sql := `SELECT id, title, description, rating, image, created_at, updated_at FROM cakes 
			WHERE deleted_at IS NULL 
			ORDER BY rating DESC, title ASC`
	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cake model.Cake
		err := rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		cakes = append(cakes, &cake)
	}

	return cakes, nil
}

func (c *CakeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, cake *model.Cake) error {
	sql := `UPDATE cakes SET title = ?, description = ?, rating = ?, image = ?, updated_at = ?
			WHERE id = ?`
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.ExecContext(ctx, cake.Title, cake.Description, cake.Rating, cake.Image, now, cake.ID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (c *CakeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id uint) error {
	sql := `UPDATE cakes SET deleted_at = ? WHERE id = ?`
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.ExecContext(ctx, now, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
