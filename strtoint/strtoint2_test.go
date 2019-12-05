package strtoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MyStrToInt2(t *testing.T) {
	r, err := MyStrToInt2("150")

	assert.Equal(t, r, 150, "should be equal")
	assert.NoError(t, err, "shouldn't be an error")
}

func Test_MyStrToInt2_EmptyString(t *testing.T) {
	r, err := MyStrToInt2("")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be an error")
}

func Test_MyStrToInt2_OutOfRange(t *testing.T) {
	_, err := MyStrToInt2("999999999999999999999999999999999999")

	assert.Error(t, err, "should be an error")
}

func Test_MyStrToInt2_MixedString(t *testing.T) {
	r, err := MyStrToInt2("Mixed123string456")

	assert.Equal(t, r, 0, "should be equal")
	assert.Error(t, err, "should be an error")
}

func Benchmark_MyStrToInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = MyStrToInt2("150")
	}
}
