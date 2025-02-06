//go:build linux

package szengine_test

import (
	"context"
	"fmt"
	"strconv"

	"github.com/senzing-garage/go-helpers/jsonutil"
	"github.com/senzing-garage/go-helpers/record"
	"github.com/senzing-garage/go-helpers/truthset"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/sz-sdk-go-mock/szabstractfactory"
	"github.com/senzing-garage/sz-sdk-go-mock/szengine"
	"github.com/senzing-garage/sz-sdk-go/senzing"
)

var (
	baseTen = 10
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

func ExampleSzengine() {}

func ExampleSzengine_AddRecord() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	dataSourceCode := "CUSTOMERS"
	recordID := "1001"
	recordDefinition := `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1001", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Robert", "DATE_OF_BIRTH": "12/11/1978", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "123 Main Street, Las Vegas NV 89132", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "702-919-1300", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "1/2/18", "STATUS": "Active", "AMOUNT": "100"}`
	flags := senzing.SzWithoutInfo
	result, err := szEngine.AddRecord(ctx, dataSourceCode, recordID, recordDefinition, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output:
}

func ExampleSzengine_AddRecord_secondRecord() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	dataSourceCode := "CUSTOMERS"
	recordID := "1002"
	recordDefinition := `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1002", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Bob", "DATE_OF_BIRTH": "11/12/1978", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1515 Adela Lane", "ADDR_CITY": "Las Vegas", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89111", "PHONE_TYPE": "MOBILE", "PHONE_NUMBER": "702-919-1300", "DATE": "3/10/17", "STATUS": "Inactive", "AMOUNT": "200"}`
	flags := senzing.SzWithoutInfo
	result, err := szEngine.AddRecord(ctx, dataSourceCode, recordID, recordDefinition, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output:
}

func ExampleSzengine_CloseExport() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	flags := senzing.SzNoFlags
	exportHandle, err := szEngine.ExportJSONEntityReport(ctx, flags)
	if err != nil {
		handleError(err)
	}
	err = szEngine.CloseExport(ctx, exportHandle)
	if err != nil {
		handleError(err)
	}
	// Output:
}

func ExampleSzengine_CountRedoRecords() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	result, err := szEngine.CountRedoRecords(ctx)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: 4
}

func ExampleSzengine_DeleteRecord() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	dataSourceCode := "CUSTOMERS"
	recordID := "1003"
	flags := senzing.SzWithoutInfo
	result, err := szEngine.DeleteRecord(ctx, dataSourceCode, recordID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output:
}

func ExampleSzengine_ExportCsvEntityReport() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	csvColumnList := ""
	flags := senzing.SzNoFlags
	exportHandle, err := szEngine.ExportCsvEntityReport(ctx, csvColumnList, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(exportHandle > 0) // Dummy output.
	// Output: true
}

func ExampleSzengine_ExportCsvEntityReportIterator() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	csvColumnList := ""
	flags := senzing.SzNoFlags
	for result := range szEngine.ExportCsvEntityReportIterator(ctx, csvColumnList, flags) {
		if result.Error != nil {
			handleError(err)
			break
		}
		fmt.Println(result.Value)
	}
	// Output: RESOLVED_ENTITY_ID,RELATED_ENTITY_ID,MATCH_LEVEL_CODE,MATCH_KEY,DATA_SOURCE,RECORD_ID
}

func ExampleSzengine_ExportJSONEntityReport() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	flags := senzing.SzNoFlags
	exportHandle, err := szEngine.ExportJSONEntityReport(ctx, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(exportHandle > 0) // Dummy output.
	// Output: true
}

func ExampleSzengine_ExportJSONEntityReportIterator() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	flags := senzing.SzNoFlags
	for result := range szEngine.ExportJSONEntityReportIterator(ctx, flags) {
		if result.Error != nil {
			handleError(err)
			break
		}
		fmt.Println(result.Value)
	}
	// Output:
}

func ExampleSzengine_FetchNext() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	flags := senzing.SzNoFlags
	exportHandle, err := szEngine.ExportJSONEntityReport(ctx, flags)
	if err != nil {
		handleError(err)
	}
	defer func() {
		err = szEngine.CloseExport(ctx, exportHandle)
	}()
	jsonEntityReport := ""
	for {
		jsonEntityReportFragment, err := szEngine.FetchNext(ctx, exportHandle)
		if err != nil {
			handleError(err)
		}
		if len(jsonEntityReportFragment) == 0 {
			break
		}
		jsonEntityReport += jsonEntityReportFragment
	}
}

func ExampleSzengine_FindInterestingEntitiesByEntityID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	entityID, err := getEntityIDForRecord("CUSTOMERS", "1001")
	if err != nil {
		handleError(err)
	}
	flags := senzing.SzNoFlags
	result, err := szEngine.FindInterestingEntitiesByEntityID(ctx, entityID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"INTERESTING_ENTITIES":{"ENTITIES":[]}}
}

func ExampleSzengine_FindInterestingEntitiesByRecordID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	dataSourceCode := "CUSTOMERS"
	recordID := "1001"
	flags := senzing.SzNoFlags
	result, err := szEngine.FindInterestingEntitiesByRecordID(ctx, dataSourceCode, recordID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"INTERESTING_ENTITIES":{"ENTITIES":[]}}
}

func ExampleSzengine_FindNetworkByEntityID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	entityID1, err := getEntityIDStringForRecord("CUSTOMERS", "1001")
	if err != nil {
		handleError(err)
	}
	entityID2, err := getEntityIDStringForRecord("CUSTOMERS", "1002")
	if err != nil {
		handleError(err)
	}
	entityList := `{"ENTITIES": [{"ENTITY_ID": ` + entityID1 + `}, {"ENTITY_ID": ` + entityID2 + `}]}`
	maxDegrees := int64(2)
	buildOutDegrees := int64(1)
	maxEntities := int64(10)
	flags := senzing.SzNoFlags
	result, err := szEngine.FindNetworkByEntityID(ctx, entityList, maxDegrees, buildOutDegrees, maxEntities, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}
}

func ExampleSzengine_FindNetworkByRecordID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	recordList := `{"RECORDS": [{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1001"}, {"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1002"}]}`
	maxDegrees := int64(1)
	buildOutDegrees := int64(2)
	maxEntities := int64(10)
	flags := senzing.SzNoFlags
	result, err := szEngine.FindNetworkByRecordID(ctx, recordList, maxDegrees, buildOutDegrees, maxEntities, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}
}

func ExampleSzengine_FindPathByEntityID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	startEntityID, err := getEntityIDForRecord("CUSTOMERS", "1001")
	if err != nil {
		handleError(err)
	}
	endEntityID, err := getEntityIDForRecord("CUSTOMERS", "1002")
	if err != nil {
		handleError(err)
	}
	maxDegrees := int64(1)
	avoidEntityIDs := ""
	requiredDataSources := ""
	flags := senzing.SzNoFlags
	result, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, maxDegrees, avoidEntityIDs, requiredDataSources, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":100001,"END_ENTITY_ID":100001,"ENTITIES":[100001]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}
}

func ExampleSzengine_FindPathByEntityID_avoiding() {
	// TODO: Implement ExampleSzEngine_FindPathByEntityID_avoiding
	// // For more information, visit https://github.com/senzing-garage/sz-sdk-go-core/blob/main/szengine/szengine_examples_test.go
	// ctx := context.TODO()
	// szEngine := getSzEngineExample(ctx)
	// startEntityID := getEntityIDForRecord("CUSTOMERS", "1001")
	// endEntityID := getEntityIDForRecord("CUSTOMERS", "1002")
	// maxDegrees := int64(1)
	// avoidEntityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + getEntityIDStringForRecord("CUSTOMERS", "1003") + `}]}`
	// requiredDataSources := ""
	// flags := senzing.SzNoFlags
	// result, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, maxDegrees, avoidEntityIDs, requiredDataSources, flags)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(truncate(result, 107))
	// // Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":1,"ENTITIES":[1]}],"ENTITIES":[{"RESOLVED_ENTITY":...
}

func ExampleSzEngine_FindPathByEntityID_avoidingAndIncluding() {
	// TODO: Implement ExampleSzEngine_FindPathByEntityID_avoidingAndIncluding
}

func ExampleSzengine_FindPathByEntityID_including() {
	// TODO: Implement ExampleSzEngine_FindPathByEntityID_including
	// // For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	// ctx := context.TODO()
	// szEngine := getSzEngineExample(ctx)
	// startEntityID := getEntityIDForRecord("CUSTOMERS", "1001")
	// endEntityID := getEntityIDForRecord("CUSTOMERS", "1002")
	// maxDegree := int64(1)
	// avoidEntityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + getEntityIDStringForRecord("CUSTOMERS", "1003") + `}]}`
	// requiredDataSources := `{"DATA_SOURCES": ["CUSTOMERS"]}`
	// flags := senzing.SzNoFlags
	// result, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, maxDegree, avoidEntityIDs, requiredDataSources, flags)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func ExampleSzengine_FindPathByRecordID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	startDataSourceCode := "CUSTOMERS"
	startRecordID := "1001"
	endDataSourceCode := "CUSTOMERS"
	endRecordID := "1002"
	maxDegrees := int64(1)
	avoidRecordKeys := ""
	requiredDataSources := ""
	flags := senzing.SzNoFlags
	result, err := szEngine.FindPathByRecordID(ctx, startDataSourceCode, startRecordID, endDataSourceCode, endRecordID, maxDegrees, avoidRecordKeys, requiredDataSources, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":100001,"END_ENTITY_ID":100001,"ENTITIES":[100001]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}
}

func ExampleSzengine_FindPathByRecordID_avoiding() {
	// TODO: Implement ExampleSzEngine_FindPathByRecordID_avoiding
	// // For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	// ctx := context.TODO()
	// szEngine := getSzEngineExample(ctx)
	// startDataSourceCode := "CUSTOMERS"
	// startRecordID := "1001"
	// endDataSourceCode := "CUSTOMERS"
	// endRecordID := "1002"
	// maxDegree := int64(1)
	// avoidRecordKeys := `{"RECORDS": [{ "DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1003"}]}`
	// requiredDataSources := ""
	// flags := senzing.SzNoFlags
	// result, err := szEngine.FindPathByRecordID(ctx, startDataSourceCode, startRecordID, endDataSourceCode, endRecordID, maxDegree, avoidRecordKeys, requiredDataSources, flags)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(truncate(result, 107))
	// // Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":1,"ENTITIES":[1]}],"ENTITIES":[{"RESOLVED_ENTITY":...
}

func ExampleSzEngine_FindPathByRecordID_excludingAndIncluding() {
	// TODO: Implement ExampleSzEngine_FindPathByRecordID_excludingAndIncluding
}

func ExampleSzengine_FindPathByRecordID_including() {
	// TODO: Implement ExampleSzEngine_FindPathByRecordID_including
	// // For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	// ctx := context.TODO()
	// szEngine := getSzEngineExample(ctx)
	// startDataSourceCode := "CUSTOMERS"
	// startRecordID := "1001"
	// endDataSourceCode := "CUSTOMERS"
	// endRecordID := "1002"
	// maxDegrees := int64(1)
	// avoidRecordKeys := `{"ENTITIES": [{"ENTITY_ID": ` + getEntityIDStringForRecord("CUSTOMERS", "1003") + `}]}`
	// requiredDataSources := `{"DATA_SOURCES": ["CUSTOMERS"]}`
	// flags := senzing.SzNoFlags
	// result, err := szEngine.FindPathByRecordID(ctx, startDataSourceCode, startRecordID, endDataSourceCode, endRecordID, maxDegrees, avoidRecordKeys, requiredDataSources, flags)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(truncate(result, 119))
	// // Output: {"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":1,"ENTITIES":[]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":...
}

func ExampleSzengine_GetActiveConfigID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	result, err := szEngine.GetActiveConfigID(ctx)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result > 0) // Dummy output.
	// Output: true
}

func ExampleSzengine_GetEntityByEntityID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	entityID, err := getEntityIDForRecord("CUSTOMERS", "1001")
	if err != nil {
		handleError(err)
	}
	flags := senzing.SzNoFlags
	result, err := szEngine.GetEntityByEntityID(ctx, entityID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"RESOLVED_ENTITY":{"ENTITY_ID":100001}}
}

func ExampleSzengine_GetEntityByRecordID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	dataSourceCode := "CUSTOMERS"
	recordID := "1001"
	flags := senzing.SzNoFlags
	result, err := szEngine.GetEntityByRecordID(ctx, dataSourceCode, recordID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"RESOLVED_ENTITY":{"ENTITY_ID":100001}}
}

func ExampleSzengine_GetRecord() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	dataSourceCode := "CUSTOMERS"
	recordID := "1001"
	flags := senzing.SzNoFlags
	result, err := szEngine.GetRecord(ctx, dataSourceCode, recordID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(jsonutil.Flatten(jsonutil.Normalize(result)))
	// Output: {"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}
}

func ExampleSzengine_GetRedoRecord() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	result, err := szEngine.GetRedoRecord(ctx)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"REASON":"deferred delete","DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1003","REEVAL_ITERATION":1,"DSRC_ACTION":"X"}
}

func ExampleSzengine_GetStats() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	result, err := szEngine.GetStats(ctx)
	if err != nil {
		handleError(err)
	}
	fmt.Println(jsonutil.Truncate(result, 5))
	// Output: {"workload":{"abortedUnresolve":0,"actualAmbiguousTest":0,"addedRecords":3,...
}

func ExampleSzengine_GetVirtualEntityByRecordID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	recordList := `{"RECORDS": [{"DATA_SOURCE": "CUSTOMERS","RECORD_ID": "1001"},{"DATA_SOURCE": "CUSTOMERS","RECORD_ID": "1002"}]}`
	flags := senzing.SzNoFlags
	result, err := szEngine.GetVirtualEntityByRecordID(ctx, recordList, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"RESOLVED_ENTITY":{"ENTITY_ID":100001}}
}

func ExampleSzengine_HowEntityByEntityID() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	entityID, err := getEntityIDForRecord("CUSTOMERS", "1001")
	if err != nil {
		handleError(err)
	}
	flags := senzing.SzNoFlags
	result, err := szEngine.HowEntityByEntityID(ctx, entityID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(jsonutil.Truncate(result, 5))
	// Output: {"HOW_RESULTS":{"FINAL_STATE":{"NEED_REEVALUATION":0,"VIRTUAL_ENTITIES":[...
}

func ExampleSzengine_PreprocessRecord() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	recordDefinition := `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1001", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Robert", "DATE_OF_BIRTH": "12/11/1978", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "123 Main Street, Las Vegas NV 89132", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "702-919-1300", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "1/2/18", "STATUS": "Active", "AMOUNT": "100"}`
	flags := senzing.SzNoFlags
	result, err := szEngine.PreprocessRecord(ctx, recordDefinition, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {}
}

func ExampleSzengine_PrimeEngine() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	err = szEngine.PrimeEngine(ctx)
	if err != nil {
		handleError(err)
	}
	// Output:
}

func ExampleSzEngine_ProcessRedoRecord() {
	// TODO: Uncomment after it has been implemented.
	// // For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	// ctx := context.TODO()
	// szEngine := getSzEngineExample(ctx)
	// redoRecord, err := szEngine.GetRedoRecord(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// flags := senzing.SzWithoutInfo
	// result, err := szEngine.ProcessRedoRecord(ctx, redoRecord, flags)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)
	// // Output: {}
}

func ExampleSzEngine_ProcessRedoRecord_withInfo() {
	// TODO: Uncomment after it has been implemented.
	// // For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	// ctx := context.TODO()
	// szEngine := getSzEngineExample(ctx)
	// redoRecord, err := szEngine.GetRedoRecord(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// flags := senzing.SzWithInfo
	// result, err := szEngine.ProcessRedoRecord(ctx, redoRecord, flags)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)
	// // Output: {}
}

func ExampleSzengine_ReevaluateEntity() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	entityID, err := getEntityIDForRecord("CUSTOMERS", "1001")
	if err != nil {
		handleError(err)
	}
	flags := senzing.SzWithoutInfo
	result, err := szEngine.ReevaluateEntity(ctx, entityID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output:
}

func ExampleSzengine_ReevaluateRecord() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	dataSourceCode := "CUSTOMERS"
	recordID := "1001"
	flags := senzing.SzWithoutInfo
	result, err := szEngine.ReevaluateRecord(ctx, dataSourceCode, recordID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output:
}

func ExampleSzengine_SearchByAttributes() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	attributes := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "Smith"}], "EMAIL_ADDRESS": "bsmith@work.com"}`
	searchProfile := ""
	flags := senzing.SzNoFlags
	result, err := szEngine.SearchByAttributes(ctx, attributes, searchProfile, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(jsonutil.Flatten(jsonutil.Redact(jsonutil.Flatten(jsonutil.NormalizeAndSort(result)), "FIRST_SEEN_DT", "LAST_SEEN_DT")))
	// Output: {"RESOLVED_ENTITIES":[{"ENTITY":{"RESOLVED_ENTITY":{"ENTITY_ID":100001}},"MATCH_INFO":{"ERRULE_CODE":"SF1","MATCH_KEY":"+PNAME+EMAIL","MATCH_LEVEL_CODE":"POSSIBLY_RELATED"}}]}
}

func ExampleSzEngine_SearchByAttributes_searchProfile() {
	// TODO: Implement ExampleSzEngine_SearchByAttributes_searchProfile
}

func ExampleSzengine_WhyEntities() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	entityID1, err := getEntityID(truthset.CustomerRecords["1001"])
	if err != nil {
		handleError(err)
	}
	entityID2, err := getEntityID(truthset.CustomerRecords["1002"])
	if err != nil {
		handleError(err)
	}
	flags := senzing.SzNoFlags
	result, err := szEngine.WhyEntities(ctx, entityID1, entityID2, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"WHY_RESULTS":[{"ENTITY_ID":100001,"ENTITY_ID_2":100001,"MATCH_INFO":{"WHY_KEY":"+NAME+DOB+ADDRESS+PHONE+EMAIL","WHY_ERRULE_CODE":"SF1_SNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}
}

func ExampleSzengine_WhyRecordInEntity() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	dataSourceCode := "CUSTOMERS"
	recordID := "1001"
	flags := senzing.SzNoFlags
	result, err := szEngine.WhyRecordInEntity(ctx, dataSourceCode, recordID, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(result)
	// Output: {"WHY_RESULTS":[{"INTERNAL_ID":100001,"ENTITY_ID":100001,"FOCUS_RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}],"MATCH_INFO":{"WHY_KEY":"+NAME+DOB+PHONE","WHY_ERRULE_CODE":"CNAME_CFF_CEXCL","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}
}

func ExampleSzengine_WhyRecords() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactory(ctx)
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	if err != nil {
		handleError(err)
	}
	dataSourceCode1 := "CUSTOMERS"
	recordID1 := "1001"
	dataSourceCode2 := "CUSTOMERS"
	recordID2 := "1002"
	flags := senzing.SzNoFlags
	result, err := szEngine.WhyRecords(ctx, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags)
	if err != nil {
		handleError(err)
	}
	fmt.Println(jsonutil.Truncate(result, 7))
	// Output: {"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}...
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func ExampleSzengine_SetLogLevel() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szEngine := getSzEngine(ctx)
	err := szEngine.SetLogLevel(ctx, logging.LevelInfoName)
	if err != nil {
		handleError(err)
	}
	// Output:
}

func ExampleSzengine_SetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szEngine := getSzEngine(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szEngine.SetObserverOrigin(ctx, origin)
	// Output:
}

func ExampleSzengine_GetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szengine/szengine_examples_test.go
	ctx := context.TODO()
	szEngine := getSzEngine(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szEngine.SetObserverOrigin(ctx, origin)
	result := szEngine.GetObserverOrigin(ctx)
	fmt.Println(result)
	// Output: Machine: nn; Task: UnitTest
}

// ----------------------------------------------------------------------------
// Helper functions
// ----------------------------------------------------------------------------

func getEntityID(record record.Record) (int64, error) {
	return getEntityIDForRecord(record.DataSource, record.ID)
}

func getEntityIDForRecord(datasource string, id string) (int64, error) {
	var err error
	_ = datasource
	_ = id
	return 1, err
}

// func getEntityIDString(record record.Record) (string, error) {
// 	entityID, err := getEntityID(record)
// 	return strconv.FormatInt(entityID, baseTen), err
// }

func getEntityIDStringForRecord(datasource string, id string) (string, error) {
	entityID, err := getEntityIDForRecord(datasource, id)
	return strconv.FormatInt(entityID, baseTen), err
}

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

func getSzEngine(ctx context.Context) *szengine.Szengine {
	_ = ctx
	return &szengine.Szengine{
		AddRecordResult:                         AddRecordResult,
		CountRedoRecordsResult:                  CountRedoRecordsResult,
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
		GetEntityByEntityIDResult:               GetEntityByEntityIDResult,
		GetEntityByRecordIDResult:               GetEntityByRecordIDResult,
		GetRecordResult:                         GetRecordResult,
		GetRedoRecordResult:                     GetRedoRecordResult,
		GetStatsResult:                          GetStatsResult,
		GetVirtualEntityByRecordIDResult:        GetVirtualEntityByRecordIDResult,
		HowEntityByEntityIDResult:               HowEntityByEntityIDResult,
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

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
