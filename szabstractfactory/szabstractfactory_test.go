package szabstractfactory_test

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/sz-sdk-go-mock/szabstractfactory"
	"github.com/senzing-garage/sz-sdk-go-mock/testdata"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/stretchr/testify/require"
)

const (
	baseCallerSkip    = 4
	defaultTruncation = 76
	instanceName      = "SzAbstractFactory Test"
	printResults      = false
	verboseLogging    = senzing.SzNoLogging
)

// ----------------------------------------------------------------------------
// Interface methods - test
// ----------------------------------------------------------------------------

func TestSzAbstractFactory_CreateConfigManager(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szAbstractFactory := getTestObject(test)

	defer func() { require.NoError(test, szAbstractFactory.Destroy(ctx)) }()

	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	require.NoError(test, err)
	configList, err := szConfigManager.GetConfigRegistry(ctx)
	require.NoError(test, err)
	printActual(test, configList)
}

func TestSzAbstractFactory_CreateDiagnostic(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szAbstractFactory := getTestObject(test)

	defer func() { require.NoError(test, szAbstractFactory.Destroy(ctx)) }()

	szDiagnostic, err := szAbstractFactory.CreateDiagnostic(ctx)
	require.NoError(test, err)
	result, err := szDiagnostic.CheckDatastorePerformance(ctx, 1)
	require.NoError(test, err)
	printActual(test, result)
}

func TestSzAbstractFactory_CreateEngine(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szAbstractFactory := getTestObject(test)

	defer func() { require.NoError(test, szAbstractFactory.Destroy(ctx)) }()

	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	require.NoError(test, err)
	stats, err := szEngine.GetStats(ctx)
	require.NoError(test, err)
	printActual(test, stats)
}

func TestSzAbstractFactory_CreateProduct(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szAbstractFactory := getTestObject(test)

	defer func() { require.NoError(test, szAbstractFactory.Destroy(ctx)) }()

	szProduct, err := szAbstractFactory.CreateProduct(ctx)
	require.NoError(test, err)
	version, err := szProduct.GetVersion(ctx)
	require.NoError(test, err)
	printActual(test, version)
}

func TestSzAbstractFactory_Destroy(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szAbstractFactory := getTestObject(test)

	defer func() { require.NoError(test, szAbstractFactory.Destroy(ctx)) }()
}

func TestSzAbstractFactory_Reinitialize(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	szAbstractFactory := getTestObject(test)

	defer func() { require.NoError(test, szAbstractFactory.Destroy(ctx)) }()

	_, err := szAbstractFactory.CreateDiagnostic(ctx)
	require.NoError(test, err)
	_, err = szAbstractFactory.CreateEngine(ctx)
	require.NoError(test, err)
	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	require.NoError(test, err)
	configID, err := szConfigManager.GetDefaultConfigID(ctx)
	require.NoError(test, err)
	err = szAbstractFactory.Reinitialize(ctx, configID)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getSzAbstractFactory(ctx context.Context) *szabstractfactory.Szabstractfactory {
	_ = ctx

	testValue := &testdata.TestData{
		Int64s:   testdata.Data1_int64s,
		Strings:  testdata.Data1_strings,
		Uintptrs: testdata.Data1_uintptrs,
	}

	return &szabstractfactory.Szabstractfactory{
		AddConfigResult:                         testValue.Int64("AddConfigResult"),
		AddDataSourceResult:                     testValue.String("AddDataSourceResult"),
		AddRecordResult:                         testValue.String("AddRecordResult"),
		CheckDatastorePerformanceResult:         testValue.String("CheckDatastorePerformanceResult"),
		CountRedoRecordsResult:                  testValue.Int64("CountRedoRecordsResult"),
		CreateConfigResult:                      testValue.Uintptr("CreateConfigResult"),
		DeleteDataSourceResult:                  testValue.String("DeleteDataSourceResult"),
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
		GetConfigRegistryResult:                 testValue.String("GetConfigRegistryResult"),
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
}

func getTestObject(t *testing.T) senzing.SzAbstractFactory {
	t.Helper()
	ctx := t.Context()

	return getSzAbstractFactory(ctx)
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
