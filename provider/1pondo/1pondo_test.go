package onepondo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnePondo_GetMovieInfoByID(t *testing.T) {
	provider := New()
	for _, item := range []string{
		"071319_870",
		"042922_001",
		"080812_401",
		"071912_387",
		"050522_001",
	} {
		info, err := provider.GetMovieInfoByID(item)
		data, _ := json.MarshalIndent(info, "", "\t")
		assert.True(t, assert.NoError(t, err) && assert.True(t, info.Valid()))
		t.Logf("%s", data)
	}
}

func TestOnePondo_GetReviewInfo(t *testing.T) {
	provider := New()
	for _, item := range []string{
		"071319_870",
	} {
		reviews, err := provider.GetMovieReviewInfo(item)
		data, _ := json.MarshalIndent(reviews, "", "\t")
		if assert.NoError(t, err) {
			for _, review := range reviews {
				assert.True(t, review.Valid())
			}
		}
		t.Logf("%s", data)
	}
}
