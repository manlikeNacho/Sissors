package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/manlikeNacho/Sissors/src/repository/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_CreateShortUrl(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controllers := gomock.NewController(t)
	defer controllers.Finish()
	mockRepo := mocks.NewMockRepository(controllers)
	ctrl := New(mockRepo)
	testUrl := &models.Url{
		ID:       "sum",
		Url:      "google.com",
		ShortUrl: "google",
	}
	mockRepo.EXPECT().SaveUrl(testUrl).Return(nil).AnyTimes()
	router.POST("/Url", ctrl.CreateShortUrl)
	rr := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/Url", nil)
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, request)
	assert.NotNil(t, rr.Body)
}
