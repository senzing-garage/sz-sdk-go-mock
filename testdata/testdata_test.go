package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Interface methods - test
// ----------------------------------------------------------------------------

func TestTestdata_String(test *testing.T) {
	testValue := &TestData{
		Int64s:   Data1_int64s,
		Strings:  Data1_strings,
		Uintptrs: Data1_uintptrs,
	}
	assert.Equal(test, "", testValue.String("GetFeatureResult"))
}
