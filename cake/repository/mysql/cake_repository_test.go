package mysql_test

import (
	"backend-engineer-test-privy/cake/repository/mysql"
	"backend-engineer-test-privy/model"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockCake := []*model.Cake{
		{
			ID:          1,
			Title:       "Test Cake",
			Description: "Test Description",
			Rating:      5,
			Image:       "https://www.google.com/image",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Title:       "Test Cake 2",
			Description: "Test Description 2",
			Rating:      4,
			Image:       "https://www.google.com/image2",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(
			mockCake[0].ID, mockCake[0].Title, mockCake[0].Description, mockCake[0].Rating, mockCake[0].Image, mockCake[0].CreatedAt, mockCake[0].UpdatedAt,
		).
		AddRow(
			mockCake[1].ID, mockCake[1].Title, mockCake[1].Description, mockCake[1].Rating, mockCake[1].Image, mockCake[1].CreatedAt, mockCake[1].UpdatedAt,
		)

	query := `SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE deleted_at IS NULL ORDER BY rating DESC, title ASC`

	mock.ExpectBegin()
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := mysql.NewCakeRepositoryImpl()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	cakes, err := a.GetAll(ctx, tx)
	assert.Len(t, cakes, 2)
	assert.NoError(t, err)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockCake := &model.Cake{
		ID:          1,
		Title:       "Test Cake",
		Description: "Test Description",
		Rating:      10,
		Image:       "https://www.google.com/image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(
			mockCake.ID, mockCake.Title, mockCake.Description, mockCake.Rating, mockCake.Image, mockCake.CreatedAt, mockCake.UpdatedAt,
		)

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = \\? AND deleted_at IS NULL"

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(1).WillReturnRows(rows)
	a := mysql.NewCakeRepositoryImpl()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("error begintx : '%s'", err)
	}
	cake, err := a.GetByID(ctx, tx, 1)
	if err != nil {
		t.Fatalf("error GetByID : '%s'", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, cake)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockCake := &model.Cake{
		Title:       "Test Cake",
		Description: "Test Description",
		Rating:      10,
		Image:       "https://www.google.com/image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := "INSERT INTO cakes \\(title, description, rating, image, created_at, updated_at\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(mockCake.Title, mockCake.Description, mockCake.Rating, mockCake.Image, mockCake.CreatedAt, mockCake.UpdatedAt).WillReturnResult(sqlmock.NewResult(10, 1))

	a := mysql.NewCakeRepositoryImpl()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("error begintx : '%s'", err)
	}
	id, err := a.Create(ctx, tx, mockCake)
	if err != nil {
		t.Fatalf("error Create : '%s'", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, 10, int(id))
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockCake2 := &model.Cake{
		ID:          11,
		Title:       "Test Cake",
		Description: "Test Description",
		Rating:      10,
		Image:       "https://www.google.com/image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := "UPDATE cakes SET title = \\?, description = \\?, rating = \\?, image = \\?, updated_at = \\? WHERE id = \\?"

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(mockCake2.Title, mockCake2.Description, mockCake2.Rating, mockCake2.Image, mockCake2.UpdatedAt, mockCake2.ID).WillReturnResult(sqlmock.NewErrorResult(nil))

	a := mysql.NewCakeRepositoryImpl()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("error begintx : '%s'", err)
	}
	err = a.Update(ctx, tx, mockCake2)
	if err != nil {
		t.Fatalf("error Update : '%s'", err)
	}

	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "UPDATE cakes SET deleted_at = \\? WHERE id = \\?"

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(time.Now(), 15).WillReturnResult(sqlmock.NewErrorResult(nil))

	a := mysql.NewCakeRepositoryImpl()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("error begintx : '%s'", err)
	}
	err = a.Delete(ctx, tx, 15)
	if err != nil {
		t.Fatalf("error Delete : '%s'", err)
	}

	assert.NoError(t, err)
}
