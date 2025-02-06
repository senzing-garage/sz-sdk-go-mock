//go:build linux

package szconfigmanager_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-helpers/jsonutil"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/sz-sdk-go-mock/szabstractfactory"
	"github.com/senzing-garage/sz-sdk-go-mock/szconfigmanager"
	"github.com/senzing-garage/sz-sdk-go/senzing"
)

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
// Interface methods - Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSzconfigmanager_AddConfig() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfigmanager/szconfigmanager_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szConfig, err := szAbstractFactory.CreateConfig(ctx)
	if err != nil {
		handleError(err)
	}
	configHandle, err := szConfig.CreateConfig(ctx)
	if err != nil {
		handleError(err)
	}
	configDefinition, err := szConfig.ExportConfig(ctx, configHandle)
	if err != nil {
		handleError(err)
	}
	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		handleError(err)
	}
	configComment := "Example configuration"
	configID, err := szConfigManager.AddConfig(ctx, configDefinition, configComment)
	if err != nil {
		handleError(err)
	}
	fmt.Println(configID > 0) // Dummy output.
	// Output: true
}

func ExampleSzconfigmanager_GetConfig() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfigmanager/szconfigmanager_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		handleError(err)
	}
	configID, err := szConfigManager.GetDefaultConfigID(ctx)
	if err != nil {
		handleError(err)
	}
	configDefinition, err := szConfigManager.GetConfig(ctx, configID)
	if err != nil {
		handleError(err)
	}
	fmt.Println(jsonutil.Truncate(configDefinition, 10))
	// Output: {"G2_CONFIG":{"CFG_ATTR":[{"ATTR_CLASS":"ADDRESS","ATTR_CODE":"ADDR_CITY","ATTR_ID":1608,"DEFAULT_VALUE":null,"FELEM_CODE":"CITY","FELEM_REQ":"Any",...
}

func ExampleSzconfigmanager_GetConfigs() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfigmanager/szconfigmanager_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		handleError(err)
	}
	configList, err := szConfigManager.GetConfigs(ctx)
	if err != nil {
		handleError(err)
	}
	fmt.Println(jsonutil.Truncate(configList, 3))
	// Output: {"CONFIGS":[{...
}

func ExampleSzconfigmanager_GetDefaultConfigID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfigmanager/szconfigmanager_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		handleError(err)
	}
	configID, err := szConfigManager.GetDefaultConfigID(ctx)
	if err != nil {
		handleError(err)
	}
	fmt.Println(configID > 0) // Dummy output.
	// Output: true
}

func ExampleSzconfigmanager_ReplaceDefaultConfigID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfigmanager/szconfigmanager_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szConfig, err := szAbstractFactory.CreateConfig(ctx)
	if err != nil {
		handleError(err)
	}
	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		handleError(err)
	}
	currentDefaultConfigID, err := szConfigManager.GetDefaultConfigID(ctx)
	if err != nil {
		handleError(err)
	}
	configDefinition, err := szConfigManager.GetConfig(ctx, currentDefaultConfigID)
	if err != nil {
		handleError(err)
	}
	configHandle, err := szConfig.ImportConfig(ctx, configDefinition)
	if err != nil {
		handleError(err)
	}
	_, err = szConfig.AddDataSource(ctx, configHandle, "XXX")
	if err != nil {
		handleError(err)
	}
	newConfigDefinition, err := szConfig.ExportConfig(ctx, configHandle)
	if err != nil {
		handleError(err)
	}
	err = szConfig.CloseConfig(ctx, configHandle)
	if err != nil {
		handleError(err)
	}
	newConfigID, err := szConfigManager.AddConfig(ctx, newConfigDefinition, "Command")
	if err != nil {
		handleError(err)
	}
	err = szConfigManager.ReplaceDefaultConfigID(ctx, currentDefaultConfigID, newConfigID)
	if err != nil {
		handleError(err)
	}
	// Output:
}

func ExampleSzconfigmanager_SetDefaultConfigID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfigmanager/szconfigmanager_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	if err != nil {
		handleError(err)
	}
	configID, err := szConfigManager.GetDefaultConfigID(ctx) // For example purposes only. Normally would use output from GetConfigList()
	if err != nil {
		handleError(err)
	}
	err = szConfigManager.SetDefaultConfigID(ctx, configID)
	if err != nil {
		handleError(err)
	}
	// Output:
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func ExampleSzconfigmanager_SetLogLevel() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfigmanager/szconfigmanager_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	err := szConfigManager.SetLogLevel(ctx, logging.LevelInfoName)
	if err != nil {
		handleError(err)
	}
	// Output:
}

func ExampleSzconfigmanager_SetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfigmanager/szconfigmanager_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szConfigManager.SetObserverOrigin(ctx, origin)
	// Output:
}

func ExampleSzconfigmanager_GetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szconfigmanager/szconfigmanager_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szConfigManager.SetObserverOrigin(ctx, origin)
	result := szConfigManager.GetObserverOrigin(ctx)
	fmt.Println(result)
	// Output: Machine: nn; Task: UnitTest
}

// ----------------------------------------------------------------------------
// Helper functions
// ----------------------------------------------------------------------------

func getSzAbstractFactory(ctx context.Context) senzing.SzAbstractFactory {
	var result senzing.SzAbstractFactory
	_ = ctx
	result = &szabstractfactory.Szabstractfactory{
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
	return result
}

func getSzConfigManager(ctx context.Context) *szconfigmanager.Szconfigmanager {
	_ = ctx
	return &szconfigmanager.Szconfigmanager{
		AddConfigResult:          AddConfigResult,
		GetConfigResult:          GetConfigResult,
		GetConfigsResult:         GetConfigsResult,
		GetDefaultConfigIDResult: GetDefaultConfigIDResult,
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
