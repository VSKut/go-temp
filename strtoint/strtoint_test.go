package strtoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MyStrToInt(t *testing.T) {
	r, err := MyStrToInt("150")

	assert.Equal(t, r, 150, "should be equal")
	assert.Nil(t, err, "should be nil")
}

func Test_MyStrToInt_EmptyString(t *testing.T) {
	r, err := MyStrToInt("")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be an error")
}

func Test_MyStrToInt_MixedString(t *testing.T) {
	r, err := MyStrToInt("Mixed123string456")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be an error")
}

func Test_MyStrToInt_OutOfRange(t *testing.T) {
	r, err := MyStrToInt("999999999999999999999999999999999999")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be an error")
}
