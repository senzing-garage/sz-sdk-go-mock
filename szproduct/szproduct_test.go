package szproduct_test

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/go-helpers/env"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/sz-sdk-go-mock/szproduct"
	"github.com/senzing-garage/sz-sdk-go-mock/testdata"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	defaultTruncation = 76
	instanceName      = "SzProduct Test"
	observerOrigin    = "SzProduct observer"
	printResults      = false
	verboseLogging    = senzing.SzNoLogging
)

// Bad parameters

const (
	badLogLevelName = "BadLogLevelName"
)

var (
	logLevel          = env.GetEnv("SENZING_LOG_LEVEL", "INFO")
	observerSingleton = &observer.NullObserver{
		ID:       "Observer 1",
		IsSilent: true,
	}
)

// ----------------------------------------------------------------------------
// Interface methods - test
// ----------------------------------------------------------------------------

func TestSzproduct_GetLicense(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(test)
	actual, err := szProduct.GetLicense(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzproduct_GetVersion(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(test)
	actual, err := szProduct.GetVersion(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzproduct_SetLogLevel_badLogLevelName(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(test)
	_ = szConfig.SetLogLevel(ctx, badLogLevelName)
}

func TestSzproduct_SetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(test)
	origin := "Machine: nn; Task: UnitTest"
	szProduct.SetObserverOrigin(ctx, origin)
}

func TestSzproduct_GetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(test)
	origin := "Machine: nn; Task: UnitTest"
	szProduct.SetObserverOrigin(ctx, origin)
	actual := szProduct.GetObserverOrigin(ctx)
	assert.Equal(test, origin, actual)
}

func TestSzproduct_UnregisterObserver(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(test)
	err := szProduct.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzproduct_AsInterface(test *testing.T) {
	ctx := context.TODO()
	szProduct := getSzProductAsInterface(ctx)
	actual, err := szProduct.GetLicense(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getSzProduct(ctx context.Context) *szproduct.Szproduct {
	testValue := &testdata.TestData{
		Int64s:   testdata.Data1_int64s,
		Strings:  testdata.Data1_strings,
		Uintptrs: testdata.Data1_uintptrs,
	}
	result := &szproduct.Szproduct{
		GetLicenseResult: testValue.String("GetLicenseResult"),
		GetVersionResult: testValue.String("GetVersionResult"),
	}
	if logLevel == "TRACE" {
		result.SetObserverOrigin(ctx, observerOrigin)
		err := result.RegisterObserver(ctx, observerSingleton)
		panicOnError(err)

		err = result.SetLogLevel(ctx, "TRACE")
		panicOnError(err)

	}
	return result
}

func getSzProductAsInterface(ctx context.Context) senzing.SzProduct {
	result := getSzProduct(ctx)
	return result
}

func getTestObject(test *testing.T) *szproduct.Szproduct {
	return getSzProduct(test.Context())
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func printResult(test *testing.T, title string, result interface{}) {
	if printResults {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result), defaultTruncation))
	}
}

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}
