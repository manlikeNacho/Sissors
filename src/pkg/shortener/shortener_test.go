package shortener

import (
	"github.com/manlikeNacho/Sissors/src/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	type testCases struct {
		name string
		url  models.Url
		want string
	}

	test := []testCases{
		{
			name: "test_1",
			url:  models.Url{ID: UserId, Url: "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/", ShortUrl: ""},
			want: "hFYku75v",
		},
		{
			name: "test_2",
			url:  models.Url{ID: UserId, Url: "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator", ShortUrl: ""},
			want: "AFZXB5Hs",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateShortLink(tt.url)
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}
