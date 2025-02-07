/*
Package test is not intended for public use.
It contains test case helpers.
*/
package testdata

type TestData struct {
	Int64s   map[string]int64
	Strings  map[string]string
	Uintptrs map[string]uintptr
}

func (testData *TestData) Int64(key string) int64 {
	value, ok := testData.Int64s[key]
	if !ok {
		return int64(0)
	}
	return value
}

func (testData *TestData) String(key string) string {
	value, ok := testData.Strings[key]
	if !ok {
		return ""
	}
	return value
}

func (testData *TestData) Uintptr(key string) uintptr {
	value, ok := testData.Uintptrs[key]
	if !ok {
		return uintptr(0)
	}
	return value
}
