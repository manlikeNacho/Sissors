package sliceRepo

import (
	"github.com/google/uuid"
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDb_SaveUrl(t *testing.T) {
	db := New()
	testUrl := models.Url{
		ID:        uuid.New().String(),
		Url:       "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html",
		ShortUrl:  "jTa4L57P",
		CreatedAt: time.Time{},
	}

	err := db.SaveUrl(testUrl)
	assert.Nil(t, err)
}

func TestDb_GetUrl(t *testing.T) {
	db := New()
	testUrl := models.Url{
		ID:        uuid.New().String(),
		Url:       "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html",
		ShortUrl:  "jTa4L57P",
		CreatedAt: time.Time{},
	}
	err := db.SaveUrl(testUrl)
	val, err := db.GetUrl(testUrl)
	assert.Nil(t, err)
	assert.Equal(t, testUrl.ShortUrl, val)
}
