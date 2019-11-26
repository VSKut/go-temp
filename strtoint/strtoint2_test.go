package strtoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MyStrToInt2(t *testing.T) {
	r, err := MyStrToInt2("150")

	assert.Equal(t, r, 150, "should be equal")
	assert.Nil(t, err, "should be nil")
}

func Test_MyStrToInt2_EmptyString(t *testing.T) {
	r, err := MyStrToInt2("")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be an error")
}

func Test_MyStrToInt2_OutOfRange(t *testing.T) {
	r, err := MyStrToInt2("999999999999999999999999999999999999")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be an error")
}

func Test_MyStrToInt2_MixedString(t *testing.T) {
	r, err := MyStrToInt2("Mixed123string456")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be an error")
}
