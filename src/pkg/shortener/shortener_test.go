package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	type testCases struct {
		name string
		url  string
		want string
	}

	test := []testCases{
		{
			name: "test_1",
			url:  "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html",
			want: "jTa4L57P",
		},
		{
			name: "test_2",
			url:  "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/",
			want: "d66yfx7N",
		},
		{
			name: "test_3",
			url:  "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator",
			want: "dhZTayYQ",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateShortLink(tt.url, UserId)
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}
