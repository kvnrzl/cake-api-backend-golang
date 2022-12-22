package service

import (
	"backend-engineer-test-privy/cake/repository"
	"backend-engineer-test-privy/model"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type CakeServiceImpl struct {
	cakeRepository repository.CakeRepository
	db             *sql.DB
	validate       *validator.Validate
}

func NewCakeServiceImpl(cakeRepository repository.CakeRepository, db *sql.DB, validate *validator.Validate) CakeService {
	return &CakeServiceImpl{cakeRepository: cakeRepository, db: db, validate: validate}
}

func (c *CakeServiceImpl) CreateCake(ctx context.Context, cake *model.Cake) (*model.Cake, error) {
	// validasi struct cake
	err := c.validate.Struct(cake)
	if err != nil {
		logrus.Error(err)
		return nil, model.ErrInputFieldInvalid
	}

	// buka koneksi ke database
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer tx.Rollback()

	// simpan cake ke database
	id, err := c.cakeRepository.Create(ctx, tx, cake)
	if err != nil {
		logrus.Error(err)
		return nil, model.ErrCannotCreate
	}

	// ambil cake dari database
	cake, err = c.cakeRepository.GetByID(ctx, tx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// commit transaksi
	err = tx.Commit()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return cake, nil
}

func (c *CakeServiceImpl) GetCakeByID(ctx context.Context, id uint) (*model.Cake, error) {
	// buka koneksi ke database
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer tx.Rollback()

	// ambil cake dari database
	cake, err := c.cakeRepository.GetByID(ctx, tx, id)
	if err != nil {
		logrus.Error(err)
		return nil, model.ErrRecordNotFound
	}

	// commit transaksi
	err = tx.Commit()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return cake, nil
}

func (c *CakeServiceImpl) GetAllCakes(ctx context.Context) ([]*model.Cake, error) {
	// buka koneksi ke database
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer tx.Rollback()

	// ambil semua cake dari database
	cakes, err := c.cakeRepository.GetAll(ctx, tx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// commit transaksi
	err = tx.Commit()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return cakes, nil
}

func (c *CakeServiceImpl) UpdateCake(ctx context.Context, cake *model.Cake) (*model.Cake, error) {
	// validasi struct cake
	err := c.validate.Struct(cake)
	if err != nil {
		logrus.Error(err)
		return nil, model.ErrInputFieldInvalid
	}

	// buka koneksi ke database
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer tx.Rollback()

	// cek id nya exist atau ga dulu
	_, err = c.cakeRepository.GetByID(ctx, tx, cake.ID)
	if err != nil {
		logrus.Error(err)
		return nil, model.ErrRecordNotFound
	}

	// update cake di database
	err = c.cakeRepository.Update(ctx, tx, cake)
	if err != nil {
		logrus.Error(err)
		return nil, model.ErrCannotUpdate
	}

	// ambil lagi data cake terbaru
	cake, err = c.cakeRepository.GetByID(ctx, tx, cake.ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// commit transaksi
	err = tx.Commit()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return cake, nil
}

func (c *CakeServiceImpl) DeleteCake(ctx context.Context, id uint) error {
	// buka koneksi ke database
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer tx.Rollback()

	// cek id nya exist atau ga dulu
	_, err = c.cakeRepository.GetByID(ctx, tx, id)
	if err != nil {
		logrus.Error(err)
		return model.ErrRecordNotFound
	}

	// hapus cake dari database
	err = c.cakeRepository.Delete(ctx, tx, id)
	if err != nil {
		logrus.Error(err)
		return model.ErrCannotDelete
	}

	// commit transaksi
	err = tx.Commit()
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
