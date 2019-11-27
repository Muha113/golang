package task

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyStrToInt1(t *testing.T) {
	var res int
	var err error

	res, err = MyStrToInt1("123abc.345efg")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
	res, err = MyStrToInt1("123")
	assert.Equal(t, []interface{}{123, nil}, []interface{}{res, err})
	res, err = MyStrToInt1("1.23")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
	res, err = MyStrToInt1("-567")
	assert.Equal(t, []interface{}{-567, nil}, []interface{}{res, err})
	res, err = MyStrToInt1("+567")
	assert.Equal(t, []interface{}{567, nil}, []interface{}{res, err})
	res, err = MyStrToInt1("+000765")
	assert.Equal(t, []interface{}{765, nil}, []interface{}{res, err})
	res, err = MyStrToInt1("10.0")
	assert.Equal(t, []interface{}{10, nil}, []interface{}{res, err})
	res, err = MyStrToInt1("")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
	res, err = MyStrToInt1("    +234.0    ")
	assert.Equal(t, []interface{}{234, nil}, []interface{}{res, err})
	res, err = MyStrToInt1(strconv.Itoa(maxInt) + "0")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
	res, err = MyStrToInt1(strconv.Itoa(minInt) + "0")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
}

func TestMyStrToInt2(t *testing.T) {
	var res int
	var err error

	res, err = MyStrToInt2("123abc.345efg")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
	res, err = MyStrToInt2("123")
	assert.Equal(t, []interface{}{123, nil}, []interface{}{res, err})
	res, err = MyStrToInt2("1.23")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
	res, err = MyStrToInt2("-567")
	assert.Equal(t, []interface{}{-567, nil}, []interface{}{res, err})
	res, err = MyStrToInt2("+567")
	assert.Equal(t, []interface{}{567, nil}, []interface{}{res, err})
	res, err = MyStrToInt2("10.0")
	assert.Equal(t, []interface{}{10, nil}, []interface{}{res, err})
	res, err = MyStrToInt2("")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
	res, err = MyStrToInt2("    +234.0    ")
	assert.Equal(t, []interface{}{234, nil}, []interface{}{res, err})
	res, err = MyStrToInt2(strconv.Itoa(maxInt) + "0")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
	res, err = MyStrToInt2(strconv.Itoa(minInt) + "0")
	assert.NotEqual(t, []interface{}{res, nil}, []interface{}{res, err})
}

// BenchmarkMyStrToInt1 : benchmark
func BenchmarkMyStrToInt1(b *testing.B) {
	str := "   +123456789   "
	for i := 0; i < b.N; i++ {
		MyStrToInt1(str)
	}
}

// BenchmarkMyStrToInt2 : benchmark
func BenchmarkMyStrToInt2(b *testing.B) {
	str := "   +123456789   "
	for i := 0; i < b.N; i++ {
		MyStrToInt2(str)
	}
}
