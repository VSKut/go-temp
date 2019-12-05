package strtoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MyStrToInt3(t *testing.T) {
	r, err := MyStrToInt3("150")

	assert.Equal(t, r, 150, "should be equal")
	assert.NoError(t, err, "shouldn't be an error")
}

func Test_MyStrToInt3_EmptyString(t *testing.T) {
	r, err := MyStrToInt3("")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be an error")
}

func Test_MyStrToInt3_OutOfRange(t *testing.T) {
	_, err := MyStrToInt3("999999999999999999999999999999999999")

	assert.Error(t, err, "should be an error")
}

func Test_MyStrToInt3_MixedString(t *testing.T) {
	_, err := MyStrToInt3("Mixed123string456")

	assert.Error(t, err, "should be an error")
}

func Benchmark_MyStrToInt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = MyStrToInt3("150")
	}
}
