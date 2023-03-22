package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlDecode(t *testing.T) {
	s, err := Decode("%26%5E%28JIH%5B.%2C%40%24%2509012%5D%5B%3D-~")
	assert.NoError(t, err)
	assert.Equal(t, s, "&^(JIH[.,@$%09012][=-~")
}

func TestUrlEncode(t *testing.T) {
	s := Encode(`&^(JIH[.,@$%09012][=-~`)
	assert.Equal(t, s, "%26%5E%28JIH%5B.%2C%40%24%2509012%5D%5B%3D-~")
}
