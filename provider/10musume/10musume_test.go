package tenmusume

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTenMusume_GetMovieInfoByID(t *testing.T) {
	provider := New()
	for _, item := range []string{
		"042922_01",
		"041607_01",
		"010906_04",
		"120409_01",
	} {
		info, err := provider.GetMovieInfoByID(item)
		data, _ := json.MarshalIndent(info, "", "\t")
		assert.True(t, assert.NoError(t, err) && assert.True(t, info.Valid()))
		t.Logf("%s", data)
	}
}

func TestTenMusume_GetReviewInfo(t *testing.T) {
	provider := New()
	for _, item := range []string{
		"042922_01",
	} {
		reviews, err := provider.GetMovieReviewInfoByID(item)
		data, _ := json.MarshalIndent(reviews, "", "\t")
		if assert.NoError(t, err) {
			for _, review := range reviews {
				assert.True(t, review.Valid())
			}
		}
		t.Logf("%s", data)
	}
}
