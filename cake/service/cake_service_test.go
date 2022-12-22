package service_test

import (
	"backend-engineer-test-privy/cake/repository/mocks"
	"backend-engineer-test-privy/cake/service"
	"backend-engineer-test-privy/model"
	"context"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAllCakes(t *testing.T) {
	db, smock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockListCakes := []*model.Cake{
		{
			ID:          1,
			Title:       "test",
			Description: "test",
			Image:       "test",
			Rating:      10,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Title:       "test 2",
			Description: "test 2",
			Image:       "test 2",
			Rating:      9,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	t.Run("Success", func(t *testing.T) {
		mockCakeRepository := new(mocks.CakeRepository)

		smock.ExpectBegin()
		mockCakeRepository.On("GetAll", mock.Anything, mock.Anything).Return(mockListCakes, nil)
		smock.ExpectCommit()

		ctx := context.Background()
		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepository, db, validate)
		cakes, err := s.GetAllCakes(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, cakes)
		assert.Len(t, cakes, len(mockListCakes))
		assert.Equal(t, cakes, mockListCakes)

		mockCakeRepository.AssertExpectations(t)
	})

	t.Run("Error-db", func(t *testing.T) {
		mockCakeRepository := new(mocks.CakeRepository)

		smock.ExpectBegin()
		mockCakeRepository.On("GetAll", mock.Anything, mock.Anything).Return(nil, model.ErrInternalServer)
		smock.ExpectCommit()

		ctx := context.Background()
		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepository, db, validate)
		cakes, err := s.GetAllCakes(ctx)
		assert.Error(t, err)
		assert.Nil(t, cakes)

		mockCakeRepository.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	db, smock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Success", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("GetByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(&model.Cake{
			ID:          1,
			Title:       "test",
			Description: "test",
			Image:       "test",
			Rating:      10,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}, nil)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		cake, err := s.GetCakeByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, cake)
		assert.Equal(t, cake.ID, uint(1))

		mockCakeRepo.AssertExpectations(t)
	})

	t.Run("Error-id-not-exist", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("GetByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(nil, model.ErrRecordNotFound)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		cake, err := s.GetCakeByID(context.Background(), 1)
		assert.Error(t, err)
		assert.Nil(t, cake)

		mockCakeRepo.AssertExpectations(t)
	})
}

func TestCreateCake(t *testing.T) {
	db, smock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockCake := &model.Cake{
		ID:          1,
		Title:       "test",
		Description: "test",
		Image:       "test",
		Rating:      10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("Create", mock.Anything, mock.Anything, mock.AnythingOfType("*model.Cake")).Return(uint(1), nil)
		mockCakeRepo.On("GetByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(mockCake, nil)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		res, err := s.CreateCake(context.Background(), mockCake)
		assert.NoError(t, err)
		assert.NotNil(t, res)

		mockCakeRepo.AssertExpectations(t)
	})

	t.Run("Error-db", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("Create", mock.Anything, mock.Anything, mock.AnythingOfType("*model.Cake")).Return(uint(0), model.ErrInternalServer)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		_, err := s.CreateCake(context.Background(), mockCake)
		assert.Error(t, err)

		mockCakeRepo.AssertExpectations(t)
	})
}

func TestUpdateCake(t *testing.T) {
	db, smock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockCake := &model.Cake{
		ID:          1,
		Title:       "test",
		Description: "test",
		Image:       "test",
		Rating:      10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("*model.Cake")).Return(nil)
		mockCakeRepo.On("GetByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(mockCake, nil)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		res, err := s.UpdateCake(context.Background(), mockCake)
		assert.NoError(t, err)
		assert.NotNil(t, res)

		mockCakeRepo.AssertExpectations(t)
	})

	t.Run("Error-id-not-exist", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("GetByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(nil, model.ErrRecordNotFound)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		res, err := s.UpdateCake(context.Background(), mockCake)
		assert.Error(t, err)
		assert.Nil(t, res)

		mockCakeRepo.AssertExpectations(t)
	})

	t.Run("Error-db", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("GetByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(mockCake, nil)
		mockCakeRepo.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("*model.Cake")).Return(model.ErrInternalServer)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		res, err := s.UpdateCake(context.Background(), mockCake)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestDeleteCake(t *testing.T) {
	db, smock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockCake := &model.Cake{
		ID:          1,
		Title:       "test",
		Description: "test",
		Image:       "test",
		Rating:      10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("GetByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(mockCake, nil)
		mockCakeRepo.On("Delete", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(nil)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		err := s.DeleteCake(context.Background(), 1)
		assert.NoError(t, err)

		mockCakeRepo.AssertExpectations(t)
	})

	t.Run("Error-id-not-exist", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("GetByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(nil, model.ErrRecordNotFound)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		err := s.DeleteCake(context.Background(), 1)
		assert.Error(t, err)

		mockCakeRepo.AssertExpectations(t)
	})

	t.Run("Error-id-not-exist", func(t *testing.T) {
		mockCakeRepo := new(mocks.CakeRepository)
		smock.ExpectBegin()
		mockCakeRepo.On("GetByID", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(mockCake, nil)
		mockCakeRepo.On("Delete", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(model.ErrCannotDelete)
		smock.ExpectCommit()

		validate := validator.New()
		s := service.NewCakeServiceImpl(mockCakeRepo, db, validate)
		err := s.DeleteCake(context.Background(), 1)
		assert.Error(t, err)
	})
}
