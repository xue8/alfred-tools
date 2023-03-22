package base64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	s, err := Decode("OkwqKCZAS21uc3UyZTc4")
	assert.NoError(t, err)
	assert.Equal(t, s, `:L*(&@Kmnsu2e78`)
}

func TestEncode(t *testing.T) {
	s := Encode(`:L*(&@Kmnsu2e78`)
	assert.Equal(t, s, `OkwqKCZAS21uc3UyZTc4`)
}
