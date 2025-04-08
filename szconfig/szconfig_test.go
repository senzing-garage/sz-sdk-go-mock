package szconfig_test

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/go-helpers/env"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/sz-sdk-go-mock/szconfig"
	"github.com/senzing-garage/sz-sdk-go-mock/testdata"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	baseTen           = 10
	dataSourceCode    = "GO_TEST"
	defaultTruncation = 76
	instanceName      = "SzConfig Test"
	observerOrigin    = "SzConfig observer"
	printResults      = false
	verboseLogging    = senzing.SzNoLogging
)

// Bad parameters

const (
	badConfigDefinition = "}{"
	badConfigHandle     = uintptr(0)
	badDataSourceCode   = "\n\tGO_TEST"
	badLogLevelName     = "BadLogLevelName"
	badSettings         = "{]"
)

// Nil/empty parameters

var (
	nilDataSourceCode string
)

var (
	logLevel          = env.GetEnv("SENZING_LOG_LEVEL", "INFO")
	observerSingleton = &observer.NullObserver{
		ID:       "Observer 1",
		IsSilent: true,
	}
)

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func TestSzconfig_AddDataSource(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	actual, err := szConfig.AddDataSource(ctx, dataSourceCode)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzconfig_DeleteDataSource(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	actual, err := szConfig.GetDataSources(ctx)
	require.NoError(test, err)
	printResult(test, "Original", actual)

	_, _ = szConfig.AddDataSource(ctx, dataSourceCode)
	actual, err = szConfig.GetDataSources(ctx)
	require.NoError(test, err)
	printResult(test, "     Add", actual)

	_, err = szConfig.DeleteDataSource(ctx, dataSourceCode)
	require.NoError(test, err)
	actual, err = szConfig.GetDataSources(ctx)
	require.NoError(test, err)
	printResult(test, "  Delete", actual)
}

func TestSzconfig_DeleteDataSource_nilDataSourceCode(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	_, err := szConfig.DeleteDataSource(ctx, nilDataSourceCode)
	require.NoError(test, err)
}

func TestSzconfig_Export(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	actual, err := szConfig.Export(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzconfig_GetDataSources(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	actual, err := szConfig.GetDataSources(ctx)
	require.NoError(test, err)
	printActual(test, actual)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzconfig_SetLogLevel_badLogLevelName(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	_ = szConfig.SetLogLevel(ctx, badLogLevelName)
}

func TestSzconfig_SetObserverOrigin(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	origin := "Machine: nn; Task: UnitTest"
	szConfig.SetObserverOrigin(ctx, origin)
}

func TestSzconfig_GetObserverOrigin(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	origin := "Machine: nn; Task: UnitTest"
	szConfig.SetObserverOrigin(ctx, origin)
	actual := szConfig.GetObserverOrigin(ctx)
	assert.Equal(test, origin, actual)
}

func TestSzconfig_UnregisterObserver(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	err := szConfig.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzconfig_AsInterface(test *testing.T) {
	ctx := test.Context()
	szConfig := getSzConfigAsInterface(ctx)
	actual, err := szConfig.GetDataSources(ctx)
	require.NoError(test, err)
	printActual(test, actual)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getSzConfig(ctx context.Context) *szconfig.Szconfig {
	testValue := &testdata.TestData{
		Int64s:   testdata.Data1_int64s,
		Strings:  testdata.Data1_strings,
		Uintptrs: testdata.Data1_uintptrs,
	}
	result := &szconfig.Szconfig{
		AddDataSourceResult:  testValue.String("AddDataSourceResult"),
		CreateConfigResult:   testValue.Uintptr("CreateConfigResult"),
		GetDataSourcesResult: testValue.String("GetDataSourcesResult"),
		ImportConfigResult:   testValue.Uintptr("ImportConfigResult"),
		ExportConfigResult:   testValue.String("ExportConfigResult"),
	}
	if logLevel == "TRACE" {
		result.SetObserverOrigin(ctx, observerOrigin)
		err := result.RegisterObserver(ctx, observerSingleton)
		if err != nil {
			panic(err)
		}
		err = result.SetLogLevel(ctx, "TRACE")
		if err != nil {
			panic(err)
		}
	}
	return result

}

func getSzConfigAsInterface(ctx context.Context) senzing.SzConfig {
	return getSzConfig(ctx)
}

func getTestObject(test *testing.T) *szconfig.Szconfig {
	return getSzConfig(test.Context())
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
