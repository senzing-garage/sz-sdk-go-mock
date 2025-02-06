package szabstractfactory

import (
	"context"
	"fmt"
	"testing"

	truncator "github.com/aquilax/truncate"
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

// Mock variables.

var (
	AddConfigResult                         = int64(1)
	AddDataSourceResult                     = `{"DSRC_ID":1001}`
	AddRecordResult                         = "{}"
	CheckDatastorePerformanceResult         = `{"numRecordsInserted":76667,"insertTime":1000}`
	CountRedoRecordsResult                  = int64(0)
	CreateConfigResult                      = uintptr(1)
	DeleteRecordResult                      = "{}"
	ExportConfigResult                      = `{"G2_CONFIG":{"CFG_ATTR":[{"ATTR_ID":1001,"ATTR_CODE":"DATA_SOURCE","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"Yes","DEFAULT_VALUE":null,"INTERNAL":"No"},{"ATTR_ID":1003,"ATTR_CODE":"RECORD_ID","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"No","DEFAULT_VALUE":null,"INTERNAL":"No"},{"ATTR_ID":1007,"ATTR_CODE":"DSRC_ACTION","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"Yes","DEFAULT_VALUE":null,"INTERNAL":"No"},{"ATTR_ID":1101,"ATTR_CODE":"NAME_TYPE","ATTR_CLASS":"NAME","FTYPE_CODE":"NAME","FELEM_CODE":"USAGE_TYPE","FELEM_REQ":"No","DEFAULT_VALUE":null,"INTERNAL":"No"},{"ATTR_ID":1102,"ATTR_CODE":"NAME_FULL","ATTR_CLASS":"NAME","FTYPE_CODE":"NAME","FELEM_CODE":"FULL_NAME","FELEM_REQ":"Any","DEFAULT_VALUE":null,"INTERNAL":"No"},{"ATTR_ID":1103,"ATTR_CODE":"NAME_ORG","ATTR_CLASS":"NAME","FTYPE_CODE":"NAME","FELEM_CODE":"ORG_NAME","FELEM_REQ":"Any","DEFAULT_VALUE":null,"INTERNAL":"No"},{"ATTR_ID":1104,"ATTR_CODE":"NAME_LAST","ATTR_CLASS":"NAME","FTYPE_CODE":"NAME","FELEM_CODE":"SUR_NAME","FELEM_REQ":"Any","DEFAULT_VALUE":null,"INTERNAL":"No"},{"ATTR_ID":1105,"ATTR_CODE":"NAME_FIRST","ATTR_CLASS":"NAME","FTYPE_CODE":"NAME","FELEM_CODE":"GIVEN_NAME","FELEM_REQ":"Any","DEFAULT_VALUE":null,"INTERNAL":"No"},{"ATTR_ID":1106,"ATTR_CODE":"NAME_MIDDLE",`
	ExportCsvEntityReportResult             = uintptr(1)
	ExportJSONEntityReportResult            = uintptr(1)
	FetchNextResult                         = ``
	FindInterestingEntitiesByEntityIDResult = `{"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
	FindInterestingEntitiesByRecordIDResult = `{"INTERESTING_ENTITIES":{"ENTITIES":[]}}`
	FindNetworkByEntityIDResult             = `{"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`
	FindNetworkByRecordIDResult             = `{"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`
	FindPathByEntityIDResult                = `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":1,"ENTITIES":[1]}],"ENTITIES":[{"RESOLVED_ENTITY":...`
	FindPathByRecordIDResult                = `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":1,"ENTITIES":[1]}],"ENTITIES":...`
	GetActiveConfigIDResult                 = int64(1)
	GetConfigResult                         = `{"G2_CONFIG":{"CFG_ATTR":[{"ATTR_ID":1001,"ATTR_CODE":"DATA_SOURCE","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"Yes","DEFAULT_VALUE":null,"ADVANCED":"Yes","INTERNAL":"No"},{"ATTR_ID":1002,"ATTR_CODE":"ROUTE_CODE","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"No","DEFAULT_VALUE":null,"ADVANCED":"Yes","INTERNAL":"No"},{"ATTR_ID":1003,"ATTR_CODE":"RECORD_ID","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"No","DEFAULT_VALUE":null,"ADVANCED":"No","INTERNAL":"No"},{"ATTR_ID":1004,"ATTR_CODE":"ENTITY_TYPE","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,`
	GetConfigsResult                        = `{"CONFIGS":[{"CONFIG_ID":41320074,"CONFIG_COMMENTS":"Example configuration","SYS_CREATE_DT":"2023-02-16 21:43:10.171"},{"CONFIG_ID":1111755672,"CONFIG_COMMENTS":"szconfigmgr_test at 2023-02-16 21:43:10.154619801 +0000 UTC","SYS_CREATE_DT":"2023-02-16 21:43:10.159"},{"CONFIG_ID":3680541328,"CONFIG_COMMENTS":"Created by szdiagnostic_test at 2023-02-16 21:43:07.294747409 +0000 UTC","SYS_CREATE_DT":"2023-02-16 21:43:07.755"}]}`
	GetDataSourcesResult                    = `{"DATA_SOURCES":[{"DSRC_ID":1,"DSRC_CODE":"TEST"},{"DSRC_ID":2,"DSRC_CODE":"SEARCH"}]}`
	GetDatastoreInfoResult                  = `{}`
	GetDefaultConfigIDResult                = int64(1)
	GetEntityByEntityIDResult               = `{"RESOLVED_ENTITY":{"ENTITY_ID":1}}`
	GetEntityByRecordIDResult               = `{"RESOLVED_ENTITY":{"ENTITY_ID":1}}`
	GetFeatureResult                        = `{}`
	GetLicenseResult                        = `{"customer":"Senzing Public Test License","contract":"Senzing Public Test - 50K records test","issueDate":"2023-11-02","licenseType":"EVAL (Solely for non-productive use)","licenseLevel":"STANDARD","billing":"YEARLY","expireDate":"2024-11-02","recordLimit":50000}`
	GetRecordResult                         = `{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}`
	GetRedoRecordResult                     = `{"REASON":"deferred delete","DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001","DSRC_ACTION":"X"}`
	GetStatsResult                          = `{ "workload": { "loadedRecords": 5,  "addedRecords": 5,  "deletedRecords": 1,  "reevaluations": 0,  "repairedEntities": 0,  "duration":...`
	GetVersionResult                        = `{"PRODUCT_NAME":"Senzing API","VERSION":"3.5.0","BUILD_VERSION":"3.5.0.23041","BUILD_DATE":"2023-02-09","BUILD_NUMBER":"2023_02_09__23_01","COMPATIBILITY_VERSION":{"CONFIG_VERSION":"10"},"SCHEMA_VERSION":{"ENGINE_SCHEMA_VERSION":"3.5","MINIMUM_REQUIRED_SCHEMA_VERSION":"3.0","MAXIMUM_REQUIRED_SCHEMA_VERSION":"3.99"}}`
	GetVirtualEntityByRecordIDResult        = `{"RESOLVED_ENTITY":{"ENTITY_ID":1}}`
	HowEntityByEntityIDResult               = `{"HOW_RESULTS":{"FINAL_STATE":{"NEED_REEVALUATION":0,"VIRTUAL_ENTITIES":[{"MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]},{"INTERNAL_ID":2,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]}],"VIRTUAL_ENTITY_ID":"V1-S1"}]},"RESOLUTION_STEPS":[{"INBOUND_VIRTUAL_ENTITY_ID":"V2","MATCH_INFO":{"ERRULE_CODE":"CNAME_CFF_CEXCL","MATCH_KEY":"+NAME+DOB+PHONE"},"RESULT_VIRTUAL_ENTITY_ID":"V1-S1","STEP":1,"VIRTUAL_ENTITY_1":{"MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]}],"VIRTUAL_ENTITY_ID":"V1"},"VIRTUAL_ENTITY_2":{"MEMBER_RECORDS":[{"INTERNAL_ID":2,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]}],"VIRTUAL_ENTITY_ID":"V2"}}]}}`
	ImportConfigResult                      = uintptr(1)
	PreprocessRecordResult                  = "{}"
	ProcessRedoRecordResult                 = ``
	ReevaluateEntityResult                  = "{}"
	ReevaluateRecordResult                  = "{}"
	SearchByAttributesResult                = `{"RESOLVED_ENTITIES":[{"ENTITY":{"RESOLVED_ENTITY":{"ENTITY_ID":1}},"MATCH_INFO":{"ERRULE_CODE":"SF1","MATCH_KEY":"+PNAME+EMAIL","MATCH_LEVEL_CODE":"POSSIBLY_RELATED"}}]}`
	WhyEntitiesResult                       = `{"WHY_RESULTS":[{"ENTITY_ID":1,"ENTITY_ID_2":1,"MATCH_INFO":{"WHY_KEY":...`
	WhyRecordInEntityResult                 = `BOB`
	WhyRecordsResult                        = `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}],"INTERNAL_ID_2":2,"ENTITY_ID_2":1,"FOCUS_RECORDS_2":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1002"}],"MATCH_INFO":{"WHY_KEY":"+NAME+DOB+PHONE","WHY_ERRULE_CODE":"CNAME_CFF_CEXCL","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`
)

// ----------------------------------------------------------------------------
// Interface methods - test
// ----------------------------------------------------------------------------

func TestSzAbstractFactory_CreateConfig(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	defer func() { handleError(szAbstractFactory.Destroy(ctx)) }()
	szConfig, err := szAbstractFactory.CreateConfig(ctx)
	require.NoError(test, err)
	configHandle, err := szConfig.CreateConfig(ctx)
	require.NoError(test, err)
	dataSources, err := szConfig.GetDataSources(ctx, configHandle)
	require.NoError(test, err)
	printActual(test, dataSources)
}

func TestSzAbstractFactory_CreateConfigManager(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	defer func() { handleError(szAbstractFactory.Destroy(ctx)) }()
	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	require.NoError(test, err)
	configList, err := szConfigManager.GetConfigs(ctx)
	require.NoError(test, err)
	printActual(test, configList)
}

func TestSzAbstractFactory_CreateDiagnostic(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	defer func() { handleError(szAbstractFactory.Destroy(ctx)) }()
	szDiagnostic, err := szAbstractFactory.CreateDiagnostic(ctx)
	require.NoError(test, err)
	result, err := szDiagnostic.CheckDatastorePerformance(ctx, 1)
	require.NoError(test, err)
	printActual(test, result)
}

func TestSzAbstractFactory_CreateEngine(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	defer func() { handleError(szAbstractFactory.Destroy(ctx)) }()
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	require.NoError(test, err)
	stats, err := szEngine.GetStats(ctx)
	require.NoError(test, err)
	printActual(test, stats)
}

func TestSzAbstractFactory_CreateProduct(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	defer func() { handleError(szAbstractFactory.Destroy(ctx)) }()
	szProduct, err := szAbstractFactory.CreateProduct(ctx)
	require.NoError(test, err)
	version, err := szProduct.GetVersion(ctx)
	require.NoError(test, err)
	printActual(test, version)
}

func TestSzAbstractFactory_Destroy(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	defer func() { handleError(szAbstractFactory.Destroy(ctx)) }()
}

func TestSzAbstractFactory_Reinitialize(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	defer func() { handleError(szAbstractFactory.Destroy(ctx)) }()
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

func getSzAbstractFactory(ctx context.Context) *Szabstractfactory {
	_ = ctx
	return &Szabstractfactory{
		AddConfigResult:                         AddConfigResult,
		AddDataSourceResult:                     AddDataSourceResult,
		AddRecordResult:                         AddRecordResult,
		CheckDatastorePerformanceResult:         CheckDatastorePerformanceResult,
		CountRedoRecordsResult:                  CountRedoRecordsResult,
		CreateConfigResult:                      CreateConfigResult,
		DeleteRecordResult:                      DeleteRecordResult,
		ExportConfigResult:                      ExportConfigResult,
		ExportCsvEntityReportResult:             ExportCsvEntityReportResult,
		ExportJSONEntityReportResult:            ExportJSONEntityReportResult,
		FetchNextResult:                         FetchNextResult,
		FindInterestingEntitiesByEntityIDResult: FindInterestingEntitiesByEntityIDResult,
		FindInterestingEntitiesByRecordIDResult: FindInterestingEntitiesByRecordIDResult,
		FindNetworkByEntityIDResult:             FindNetworkByEntityIDResult,
		FindNetworkByRecordIDResult:             FindNetworkByRecordIDResult,
		FindPathByEntityIDResult:                FindPathByEntityIDResult,
		FindPathByRecordIDResult:                FindPathByRecordIDResult,
		GetActiveConfigIDResult:                 GetActiveConfigIDResult,
		GetConfigResult:                         GetConfigResult,
		GetConfigsResult:                        GetConfigsResult,
		GetDataSourcesResult:                    GetDataSourcesResult,
		GetDatastoreInfoResult:                  GetDatastoreInfoResult,
		GetDefaultConfigIDResult:                GetDefaultConfigIDResult,
		GetEntityByEntityIDResult:               GetEntityByEntityIDResult,
		GetEntityByRecordIDResult:               GetEntityByRecordIDResult,
		GetFeatureResult:                        GetFeatureResult,
		GetLicenseResult:                        GetLicenseResult,
		GetRecordResult:                         GetRecordResult,
		GetRedoRecordResult:                     GetRedoRecordResult,
		GetStatsResult:                          GetStatsResult,
		GetVersionResult:                        GetVersionResult,
		GetVirtualEntityByRecordIDResult:        GetVirtualEntityByRecordIDResult,
		HowEntityByEntityIDResult:               HowEntityByEntityIDResult,
		ImportConfigResult:                      ImportConfigResult,
		PreprocessRecordResult:                  PreprocessRecordResult,
		ProcessRedoRecordResult:                 ProcessRedoRecordResult,
		ReevaluateEntityResult:                  ReevaluateEntityResult,
		ReevaluateRecordResult:                  ReevaluateRecordResult,
		SearchByAttributesResult:                SearchByAttributesResult,
		WhyEntitiesResult:                       WhyEntitiesResult,
		WhyRecordInEntityResult:                 WhyRecordInEntityResult,
		WhyRecordsResult:                        WhyRecordsResult,
	}
}

func getTestObject(ctx context.Context, test *testing.T) *Szabstractfactory {
	_ = test
	return getSzAbstractFactory(ctx)
}

func handleError(err error) {
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
