package helper_test

import (
	"testing"

	"github.com/senzing-garage/sz-sdk-go-mock/helper"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Interface methods - test
// ----------------------------------------------------------------------------

func TestHelpers_GetMessenger(test *testing.T) {
	test.Parallel()

	x := helper.GetMessenger(1, map[int]string{}, 4)
	assert.NotEmpty(test, x)
}
