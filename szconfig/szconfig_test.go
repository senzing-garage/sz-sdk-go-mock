package szconfig_test

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/go-helpers/env"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/sz-sdk-go-mock/szabstractfactory"
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
	originMessage     = "Machine: nn; Task: UnitTest"
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

func TestSzconfig_Export(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	actual, err := szConfig.Export(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzconfig_GetDataSourceRegistry(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	actual, err := szConfig.GetDataSourceRegistry(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzconfig_RegisterDataSource(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	actual, err := szConfig.RegisterDataSource(ctx, dataSourceCode)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzconfig_UnregisterDataSource(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	actual, err := szConfig.GetDataSourceRegistry(ctx)
	require.NoError(test, err)
	printResult(test, "Original", actual)

	_, _ = szConfig.RegisterDataSource(ctx, dataSourceCode)
	actual, err = szConfig.GetDataSourceRegistry(ctx)
	require.NoError(test, err)
	printResult(test, "     Add", actual)

	_, err = szConfig.UnregisterDataSource(ctx, dataSourceCode)
	require.NoError(test, err)
	actual, err = szConfig.GetDataSourceRegistry(ctx)
	require.NoError(test, err)
	printResult(test, "  Delete", actual)
}

// ----------------------------------------------------------------------------
// Public non-interface methods
// ----------------------------------------------------------------------------

func TestSzconfig_Import(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	configDefinition, err := szConfig.Export(ctx)
	require.NoError(test, err)
	err = szConfig.Import(ctx, configDefinition)
	require.NoError(test, err)
}

func TestSzconfig_ImportTemplate(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	err := szConfig.ImportTemplate(ctx)
	require.NoError(test, err)
}

func TestSzconfig_VerifyConfigDefinition(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	configDefinition, err := szConfig.Export(ctx)
	require.NoError(test, err)
	err = szConfig.VerifyConfigDefinition(ctx, configDefinition)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzconfig_SetLogLevel_badLogLevelName(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	_ = szConfig.SetLogLevel(ctx, badLogLevelName)
}

func TestSzconfig_SetObserverOrigin(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	origin := originMessage
	szConfig.SetObserverOrigin(ctx, origin)
}

func TestSzconfig_GetObserverOrigin(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	origin := originMessage
	szConfig.SetObserverOrigin(ctx, origin)
	actual := szConfig.GetObserverOrigin(ctx)
	assert.Equal(test, origin, actual)
}

func TestSzconfig_UnregisterObserver(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getTestObject(test)
	err := szConfig.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzconfig_AsInterface(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfig := getSzConfigAsInterface(ctx)
	actual, err := szConfig.GetDataSourceRegistry(ctx)
	require.NoError(test, err)
	printActual(test, actual)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getSzAbstractFactory(ctx context.Context) senzing.SzAbstractFactory {
	var result senzing.SzAbstractFactory

	_ = ctx

	testValue := &testdata.TestData{
		Int64s:   testdata.Data1_int64s,
		Strings:  testdata.Data1_strings,
		Uintptrs: testdata.Data1_uintptrs,
	}

	result = &szabstractfactory.Szabstractfactory{
		AddConfigResult:                         testValue.Int64("AddConfigResult"),
		AddRecordResult:                         testValue.String("AddRecordResult"),
		CheckRepositoryPerformanceResult:        testValue.String("CheckRepositoryPerformanceResult"),
		CountRedoRecordsResult:                  testValue.Int64("CountRedoRecordsResult"),
		CreateConfigResult:                      testValue.Uintptr("CreateConfigResult"),
		DeleteRecordResult:                      testValue.String("DeleteRecordResult"),
		ExportConfigResult:                      testValue.String("ExportConfigResult"),
		ExportCsvEntityReportResult:             testValue.Uintptr("ExportCsvEntityReportResult"),
		ExportJSONEntityReportResult:            testValue.Uintptr("ExportJSONEntityReportResult"),
		FetchNextResult:                         testValue.String("FetchNextResult"),
		FindInterestingEntitiesByEntityIDResult: testValue.String("FindInterestingEntitiesByEntityIDResult"),
		FindInterestingEntitiesByRecordIDResult: testValue.String("FindInterestingEntitiesByRecordIDResult"),
		FindNetworkByEntityIDResult:             testValue.String("FindNetworkByEntityIDResult"),
		FindNetworkByRecordIDResult:             testValue.String("FindNetworkByRecordIDResult"),
		FindPathByEntityIDResult:                testValue.String("FindPathByEntityIDResult"),
		FindPathByRecordIDResult:                testValue.String("FindPathByRecordIDResult"),
		GetActiveConfigIDResult:                 testValue.Int64("GetActiveConfigIDResult"),
		GetConfigRegistryResult:                 testValue.String("GetConfigRegistryResult"),
		GetConfigResult:                         testValue.String("GetConfigResult"),
		GetDataSourceRegistryResult:             testValue.String("GetDataSourceRegistryResult"),
		GetDefaultConfigIDResult:                testValue.Int64("GetDefaultConfigIDResult"),
		GetEntityByEntityIDResult:               testValue.String("GetEntityByEntityIDResult"),
		GetEntityByRecordIDResult:               testValue.String("GetEntityByRecordIDResult"),
		GetFeatureResult:                        testValue.String("GetFeatureResult"),
		GetLicenseResult:                        testValue.String("GetLicenseResult"),
		GetRecordPreviewResult:                  testValue.String("GetRecordPreviewResult"),
		GetRecordResult:                         testValue.String("GetRecordResult"),
		GetRedoRecordResult:                     testValue.String("GetRedoRecordResult"),
		GetRepositoryInfoResult:                 testValue.String("GetRepositoryInfoResult"),
		GetStatsResult:                          testValue.String("GetStatsResult"),
		GetVersionResult:                        testValue.String("GetVersionResult"),
		GetVirtualEntityByRecordIDResult:        testValue.String("GetVirtualEntityByRecordIDResult"),
		HowEntityByEntityIDResult:               testValue.String("HowEntityByEntityIDResult"),
		ImportConfigResult:                      testValue.Uintptr("ImportConfigResult"),
		ProcessRedoRecordResult:                 testValue.String("ProcessRedoRecordResult"),
		ReevaluateEntityResult:                  testValue.String("ReevaluateEntityResult"),
		ReevaluateRecordResult:                  testValue.String("ReevaluateRecordResult"),
		RegisterDataSourceResult:                testValue.String("RegisterDataSourceResult"),
		SearchByAttributesResult:                testValue.String("SearchByAttributesResult"),
		WhyEntitiesResult:                       testValue.String("WhyEntitiesResult"),
		WhyRecordInEntityResult:                 testValue.String("WhyRecordInEntityResult"),
		WhyRecordsResult:                        testValue.String("WhyRecordsResult"),
	}

	return result
}

func getSzConfig(ctx context.Context) *szconfig.Szconfig {
	testValue := &testdata.TestData{
		Int64s:   testdata.Data1_int64s,
		Strings:  testdata.Data1_strings,
		Uintptrs: testdata.Data1_uintptrs,
	}

	result := &szconfig.Szconfig{
		RegisterDataSourceResult:    testValue.String("RegisterDataSourceResult"),
		CreateConfigResult:          testValue.Uintptr("CreateConfigResult"),
		GetDataSourceRegistryResult: testValue.String("GetDataSourceRegistryResult"),
		ImportConfigResult:          testValue.Uintptr("ImportConfigResult"),
		ExportResult:                testValue.String("ExportConfigResult"),
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

func getTestObject(t *testing.T) *szconfig.Szconfig {
	t.Helper()

	return getSzConfig(t.Context())
}

func handleError(err error) {
	if err != nil {
		outputln("Error:", err)
	}
}

func outputln(message ...any) {
	fmt.Println(message...) //nolint
}

func printActual(t *testing.T, actual interface{}) {
	t.Helper()
	printResult(t, "Actual", actual)
}

func printResult(t *testing.T, title string, result interface{}) {
	t.Helper()

	if printResults {
		t.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result), defaultTruncation))
	}
}

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}
