package controller_test

import (
	"backend-engineer-test-privy/cake/controller"
	"backend-engineer-test-privy/cake/service/mocks"
	"backend-engineer-test-privy/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCake(t *testing.T) {
	mockCake := model.Cake{
		Title:       "Test",
		Description: "Test",
		Rating:      10,
		Image:       "Test",
	}

	mockCakReturn := model.Cake{
		ID:          1,
		Title:       "Test",
		Description: "Test",
		Rating:      10,
		Image:       "Test",
		CreatedAt:   mockCake.CreatedAt,
		UpdatedAt:   mockCake.UpdatedAt,
	}

	mockCake.ID = 1

	json, err := json.Marshal(mockCake)
	assert.NoError(t, err)

	mockCakeService := new(mocks.CakeService)
	mockCakeService.On("CreateCake", mock.Anything, mock.AnythingOfType("*model.Cake")).Return(&mockCakReturn, nil)

	g := gin.New()
	req, err := http.NewRequest("POST", "/cake", strings.NewReader(string(json)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	g.POST("/cake", controller.NewCakeControllerImpl(mockCakeService).CreateCake)
	g.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "application/json; charset=utf-8", rec.Header().Get("Content-Type"))
}
