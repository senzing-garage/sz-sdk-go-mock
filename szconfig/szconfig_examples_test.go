//go:build linux

package szconfig_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-helpers/jsonutil"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/sz-sdk-go-mock/szabstractfactory"
	"github.com/senzing-garage/sz-sdk-go-mock/testdata"
	"github.com/senzing-garage/sz-sdk-go/senzing"
)

// ----------------------------------------------------------------------------
// Interface methods - Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSzconfig_AddDataSource() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfig/szconfig_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)

	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		panicOnError(err)
	}

	szConfig, err := szConfigManager.CreateConfigFromTemplate(ctx)
	if err != nil {
		panicOnError(err)
	}

	dataSourceCode := "GO_TEST"

	result, err := szConfig.AddDataSource(ctx, dataSourceCode)
	if err != nil {
		panicOnError(err)
	}

	fmt.Println(result)
	// Output: {"DSRC_ID":1001}
}

func ExampleSzconfig_DeleteDataSource() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfig/szconfig_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)

	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		panicOnError(err)
	}

	szConfig, err := szConfigManager.CreateConfigFromTemplate(ctx)
	if err != nil {
		panicOnError(err)
	}

	dataSourceCode := "TEST"

	result, err := szConfig.DeleteDataSource(ctx, dataSourceCode)
	if err != nil {
		panicOnError(err)
	}

	fmt.Println(result)
	// Output:
}

func ExampleSzconfig_Export() {
	// For more information, visit
	// https://github.com/senzing-garage/sz-sdk-go-core/blob/main/szconfig/szconfig_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)

	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		panicOnError(err)
	}

	szConfig, err := szConfigManager.CreateConfigFromTemplate(ctx)
	if err != nil {
		panicOnError(err)
	}

	configDefinition, err := szConfig.Export(ctx)
	if err != nil {
		panicOnError(err)
	}

	fmt.Println(jsonutil.Truncate(configDefinition, 7))
	// Output: {"G2_CONFIG":{"CFG_ATTR":[{"ATTR_CLASS":"ADDRESS","ATTR_CODE":"ADDR_CITY","ATTR_ID":1608,...
}

func ExampleSzconfig_GetDataSources() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfig/szconfig_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)

	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		panicOnError(err)
	}

	szConfig, err := szConfigManager.CreateConfigFromTemplate(ctx)
	if err != nil {
		panicOnError(err)
	}

	result, err := szConfig.GetDataSources(ctx)
	if err != nil {
		panicOnError(err)
	}

	fmt.Println(result)
	// Output: {"DATA_SOURCES":[{"DSRC_ID":1,"DSRC_CODE":"TEST"},{"DSRC_ID":2,"DSRC_CODE":"SEARCH"}]}
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func ExampleSzconfig_SetLogLevel() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfig/szconfig_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)

	err := szConfig.SetLogLevel(ctx, logging.LevelInfoName)
	if err != nil {
		panicOnError(err)
	}
	// Output:
}

func ExampleSzconfig_SetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfig/szconfig_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szConfig.SetObserverOrigin(ctx, origin)
	// Output:
}

func ExampleSzconfig_GetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfig/szconfig_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szConfig.SetObserverOrigin(ctx, origin)
	result := szConfig.GetObserverOrigin(ctx)
	fmt.Println(result)
	// Output: Machine: nn; Task: UnitTest
}

// ----------------------------------------------------------------------------
// Helper functions
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
