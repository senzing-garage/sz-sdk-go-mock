package szdiagnostic_test

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/go-helpers/env"
	"github.com/senzing-garage/go-helpers/record"
	"github.com/senzing-garage/go-helpers/truthset"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/sz-sdk-go-mock/szabstractfactory"
	"github.com/senzing-garage/sz-sdk-go-mock/szdiagnostic"
	"github.com/senzing-garage/sz-sdk-go-mock/testdata"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	defaultTruncation = 76
	instanceName      = "SzDiagnostic Test"
	jsonIndentation   = "    "
	observerOrigin    = "SzDiagnostic observer"
	printResults      = false
	verboseLogging    = senzing.SzNoLogging
)

// Bad parameters

const (
	badFeatureID    = int64(-1)
	badLogLevelName = "BadLogLevelName"
	badSecondsToRun = -1
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

func TestSzdiagnostic_CheckDatastorePerformance(test *testing.T) {
	ctx := test.Context()
	szDiagnostic := getTestObject(test)
	secondsToRun := 1
	actual, err := szDiagnostic.CheckDatastorePerformance(ctx, secondsToRun)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzdiagnostic_GetDatastoreInfo(test *testing.T) {
	ctx := test.Context()
	szDiagnostic := getTestObject(test)
	actual, err := szDiagnostic.GetDatastoreInfo(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzdiagnostic_GetFeature(test *testing.T) {
	ctx := test.Context()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}

	defer func() { deleteRecords(ctx, records) }()

	addRecords(ctx, records)

	szDiagnostic := getTestObject(test)
	featureID := int64(1)
	actual, err := szDiagnostic.GetFeature(ctx, featureID)
	require.NoError(test, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzdiagnostic_SetLogLevel_badLogLevelName(test *testing.T) {
	ctx := test.Context()
	szConfig := getTestObject(test)
	_ = szConfig.SetLogLevel(ctx, badLogLevelName)
}

func TestSzdiagnostic_SetObserverOrigin(test *testing.T) {
	ctx := test.Context()
	szDiagnostic := getTestObject(test)
	origin := "Machine: nn; Task: UnitTest"
	szDiagnostic.SetObserverOrigin(ctx, origin)
}

func TestSzdiagnostic_GetObserverOrigin(test *testing.T) {
	ctx := test.Context()
	szDiagnostic := getTestObject(test)
	origin := "Machine: nn; Task: UnitTest"
	szDiagnostic.SetObserverOrigin(ctx, origin)
	actual := szDiagnostic.GetObserverOrigin(ctx)
	assert.Equal(test, origin, actual)
}

func TestSzdiagnostic_UnregisterObserver(test *testing.T) {
	ctx := test.Context()
	szDiagnostic := getTestObject(test)
	err := szDiagnostic.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzdiagnostic_AsInterface(test *testing.T) {
	ctx := test.Context()
	szDiagnostic := getSzDiagnosticAsInterface(ctx)
	secondsToRun := 1
	actual, err := szDiagnostic.CheckDatastorePerformance(ctx, secondsToRun)
	require.NoError(test, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func addRecords(ctx context.Context, records []record.Record) {
	_ = ctx
	_ = records
}

func deleteRecords(ctx context.Context, records []record.Record) {
	_ = ctx
	_ = records
}

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
		AddDataSourceResult:                     testValue.String("AddDataSourceResult"),
		AddRecordResult:                         testValue.String("AddRecordResult"),
		CheckDatastorePerformanceResult:         testValue.String("CheckDatastorePerformanceResult"),
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
		GetConfigResult:                         testValue.String("GetConfigResult"),
		GetConfigsResult:                        testValue.String("GetConfigsResult"),
		GetDataSourcesResult:                    testValue.String("GetDataSourcesResult"),
		GetDatastoreInfoResult:                  testValue.String("GetDatastoreInfoResult"),
		GetDefaultConfigIDResult:                testValue.Int64("GetDefaultConfigIDResult"),
		GetEntityByEntityIDResult:               testValue.String("GetEntityByEntityIDResult"),
		GetEntityByRecordIDResult:               testValue.String("GetEntityByRecordIDResult"),
		GetFeatureResult:                        testValue.String("GetFeatureResult"),
		GetLicenseResult:                        testValue.String("GetLicenseResult"),
		GetRecordResult:                         testValue.String("GetRecordResult"),
		GetRedoRecordResult:                     testValue.String("GetRedoRecordResult"),
		GetStatsResult:                          testValue.String("GetStatsResult"),
		GetVersionResult:                        testValue.String("GetVersionResult"),
		GetVirtualEntityByRecordIDResult:        testValue.String("GetVirtualEntityByRecordIDResult"),
		HowEntityByEntityIDResult:               testValue.String("HowEntityByEntityIDResult"),
		ImportConfigResult:                      testValue.Uintptr("ImportConfigResult"),
		PreprocessRecordResult:                  testValue.String("PreprocessRecordResult"),
		ProcessRedoRecordResult:                 testValue.String("ProcessRedoRecordResult"),
		ReevaluateEntityResult:                  testValue.String("ReevaluateEntityResult"),
		ReevaluateRecordResult:                  testValue.String("ReevaluateRecordResult"),
		SearchByAttributesResult:                testValue.String("SearchByAttributesResult"),
		WhyEntitiesResult:                       testValue.String("WhyEntitiesResult"),
		WhyRecordInEntityResult:                 testValue.String("WhyRecordInEntityResult"),
		WhyRecordsResult:                        testValue.String("WhyRecordsResult"),
	}
	return result
}

func getSzDiagnostic(ctx context.Context) *szdiagnostic.Szdiagnostic {
	testValue := &testdata.TestData{
		Int64s:   testdata.Data1_int64s,
		Strings:  testdata.Data1_strings,
		Uintptrs: testdata.Data1_uintptrs,
	}
	result := &szdiagnostic.Szdiagnostic{
		CheckDatastorePerformanceResult: testValue.String("CheckDatastorePerformanceResult"),
		GetDatastoreInfoResult:          testValue.String("GetDatastoreInfoResult"),
		GetFeatureResult:                testValue.String("GetFeatureResult"),
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

func getSzDiagnosticAsInterface(ctx context.Context) senzing.SzDiagnostic {
	return getSzDiagnostic(ctx)
}

func getTestObject(t *testing.T) *szdiagnostic.Szdiagnostic {
	t.Helper()
	return getSzDiagnostic(t.Context())
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
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

// ----------------------------------------------------------------------------
