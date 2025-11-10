package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeAndDecodeStrings(t *testing.T) {
	s := EncodeAndDecodeStringsSolution{}
	input := []string{"we", "say", ":", "yes"}
	encoded := s.Encode(input)
	assert.Equal(t, input, s.Decode(encoded))
}

func TestEncodeAndDecodeStrings2(t *testing.T) {
	s := EncodeAndDecodeStringsSolution{}
	input := []string{"neet", "code", "love", "you"}
	encoded := s.Encode(input)
	assert.Equal(t, input, s.Decode(encoded))
}

func TestEncodeAndDecodeStrings3(t *testing.T) {
	s := EncodeAndDecodeStringsSolution{}
	input := []string{"#", "01", "3##", "#33test", "---test"}
	encoded := s.Encode(input)
	assert.Equal(t, input, s.Decode(encoded))
}
