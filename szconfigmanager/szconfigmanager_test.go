package szconfigmanager_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/go-helpers/env"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/sz-sdk-go-mock/szabstractfactory"
	"github.com/senzing-garage/sz-sdk-go-mock/szconfigmanager"
	"github.com/senzing-garage/sz-sdk-go-mock/testdata"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	defaultTruncation = 76
	instanceName      = "SzConfigManager Test"
	observerOrigin    = "SzConfigManager observer"
	originMessage     = "Machine: nn; Task: UnitTest"
	printResults      = false
	verboseLogging    = senzing.SzNoLogging
)

// Bad parameters

const (
	badConfigDefinition       = "\n\t"
	badConfigID               = int64(0)
	badCurrentDefaultConfigID = int64(0)
	badLogLevelName           = "BadLogLevelName"
	badNewDefaultConfigID     = int64(0)
	baseTen                   = 10
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

func TestSzconfigmanager_CreateConfigFromConfigID(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	configID, err1 := szConfigManager.GetDefaultConfigID(ctx)
	require.NoError(test, err1)

	actual, err := szConfigManager.CreateConfigFromConfigID(ctx, configID)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzconfigmanager_CreateConfigFromString(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	actual, err := szConfigManager.CreateConfigFromString(ctx, "")
	require.NoError(test, err)
	assert.NotEmpty(test, actual)
}

func TestSzconfigmanager_CreateConfigFromTemplate(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	actual, err := szConfigManager.CreateConfigFromTemplate(ctx)
	require.NoError(test, err)
	assert.NotEmpty(test, actual)
}

func TestSzconfigmanager_GetConfigRegistry(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	actual, err := szConfigManager.GetConfigRegistry(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzconfigmanager_GetDefaultConfigID(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	actual, err := szConfigManager.GetDefaultConfigID(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzconfigmanager_RegisterConfig(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	now := time.Now()
	szConfig, err := szConfigManager.CreateConfigFromTemplate(ctx)
	require.NoError(test, err)

	dataSourceCode := "GO_TEST_" + strconv.FormatInt(now.Unix(), baseTen)
	_, err = szConfig.RegisterDataSource(ctx, dataSourceCode)
	require.NoError(test, err)
	configDefinition, err := szConfig.Export(ctx)
	require.NoError(test, err)

	configComment := fmt.Sprintf("szconfigmanager_test at %s", now.UTC())
	actual, err := szConfigManager.RegisterConfig(ctx, configDefinition, configComment)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzconfigmanager_ReplaceDefaultConfigID(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	currentDefaultConfigID, err1 := szConfigManager.GetDefaultConfigID(ctx)
	require.NoError(test, err1)

	// IMPROVE: This is kind of a cheater.

	newDefaultConfigID, err2 := szConfigManager.GetDefaultConfigID(ctx)
	require.NoError(test, err2)

	err := szConfigManager.ReplaceDefaultConfigID(ctx, currentDefaultConfigID, newDefaultConfigID)
	require.NoError(test, err)
}

func TestSzconfigmanager_SetDefaultConfigID(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	configID, err := szConfigManager.GetDefaultConfigID(ctx)
	require.NoError(test, err)
	err = szConfigManager.SetDefaultConfigID(ctx, configID)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzconfigmanager_SetLogLevel_badLogLevelName(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	_ = szConfigManager.SetLogLevel(ctx, badLogLevelName)
}

func TestSzconfigmanager_SetObserverOrigin(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	szConfigManager.SetObserverOrigin(ctx, originMessage)
}

func TestSzconfigmanager_GetObserverOrigin(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	szConfigManager.SetObserverOrigin(ctx, originMessage)
	actual := szConfigManager.GetObserverOrigin(ctx)
	assert.Equal(test, originMessage, actual)
}

func TestSzconfigmanager_UnregisterObserver(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getTestObject(test)
	err := szConfigManager.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzconfigmanager_AsInterface(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szConfigManager := getSzConfigManagerAsInterface(ctx)
	actual, err := szConfigManager.GetConfigRegistry(ctx)
	require.NoError(test, err)
	printActual(test, actual)
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

// func getSzConfig(ctx context.Context) (senzing.SzConfig, error) {
// 	_ = ctx

// 	testValue := &testdata.TestData{
// 		Int64s:   testdata.Data1_int64s,
// 		Strings:  testdata.Data1_strings,
// 		Uintptrs: testdata.Data1_uintptrs,
// 	}

// 	return &szconfig.Szconfig{
// 		RegisterDataSourceResult:  testValue.String("RegisterDataSourceResult"),
// 		CreateConfigResult:   testValue.Uintptr("CreateConfigResult"),
// 		GetDataSourceRegistryResult: testValue.String("GetDataSourceRegistryResult"),
// 		ImportConfigResult:   testValue.Uintptr("ImportConfigResult"),
// 		ExportConfigResult:   testValue.String("ExportConfigResult"),
// 	}, nil
// }

func getSzConfigManager(ctx context.Context) *szconfigmanager.Szconfigmanager {
	testValue := &testdata.TestData{
		Int64s:   testdata.Data1_int64s,
		Strings:  testdata.Data1_strings,
		Uintptrs: testdata.Data1_uintptrs,
	}

	result := &szconfigmanager.Szconfigmanager{
		RegisterConfigResult:     testValue.Int64("AddConfigResult"),
		GetConfigResult:          testValue.String("GetConfigResult"),
		GetConfigRegistryResult:  testValue.String("GetConfigRegistryResult"),
		GetDefaultConfigIDResult: testValue.Int64("GetDefaultConfigIDResult"),
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

func getSzConfigManagerAsInterface(ctx context.Context) senzing.SzConfigManager {
	return getSzConfigManager(ctx)
}

func getTestObject(t *testing.T) *szconfigmanager.Szconfigmanager {
	t.Helper()

	return getSzConfigManager(t.Context())
}

func handleError(err error) {
	if err != nil {
		outputln("Error:", err)
	}
}

func outputln(message ...any) {
	fmt.Println(message...) //nolint
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
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
