package sliceRepo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/stretchr/testify/assert"
)

func TestDb_SaveUrl(t *testing.T) {
	db := New()
	testUrl := &models.Url{
		ID:       uuid.New().String(),
		Url:      "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html",
		ShortUrl: "jTa4L57P",
	}

	err := db.SaveUrl(testUrl)
	assert.Nil(t, err)
}

func TestDb_GetUrl(t *testing.T) {
	db := New()
	testUrl := &models.Url{
		ID:       uuid.New().String(),
		Url:      "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html",
		ShortUrl: "jTa4L57P",
	}
	_ = db.SaveUrl(testUrl)
	val, err := db.GetUrl(testUrl.ShortUrl)
	assert.Nil(t, err)
	assert.Equal(t, testUrl.Url, val)

	testUrl2 := &models.Url{
		ID:       uuid.New().String(),
		Url:      "https://www.eddywm.com/lets-build-a-url-shortener-in-go-part-3-short-link-generatio",
		ShortUrl: "",
	}

	val2, err2 := db.GetUrl(testUrl2.ShortUrl)
	assert.NotNil(t, err2)
	assert.Equal(t, "", val2)

}
