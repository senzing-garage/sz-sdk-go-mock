/*
Package testdata is not intended for public use.
It contains test case helpers.
*/
package testdata

var Data1_strings = map[string]string{
	"AddDataSourceResult":                     `{"DSRC_ID":1001}`,
	"AddRecordResult":                         "",
	"CheckDatastorePerformanceResult":         `{"numRecordsInserted":76667,"insertTime":1000}`,
	"DeleteDataSourceResult":                  "",
	"DeleteRecordResult":                      "",
	"ExportConfigResult":                      `{"G2_CONFIG":{"CFG_ATTR":[{"ATTR_CLASS":"ADDRESS","ATTR_CODE":"ADDR_CITY","ATTR_ID":1608,...`,
	"FetchNextResult":                         ``,
	"FindInterestingEntitiesByEntityIDResult": `{"INTERESTING_ENTITIES":{"ENTITIES":[]}}`,
	"FindInterestingEntitiesByRecordIDResult": `{"INTERESTING_ENTITIES":{"ENTITIES":[]}}`,
	"FindNetworkByEntityIDResult":             `{"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}`,
	"FindNetworkByRecordIDResult":             `{"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}`,
	"FindPathByEntityIDResult":                `{"ENTITY_PATHS":[{"START_ENTITY_ID":100001,"END_ENTITY_ID":100001,"ENTITIES":[100001]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}`,
	"FindPathByRecordIDResult":                `{"ENTITY_PATHS":[{"START_ENTITY_ID":100001,"END_ENTITY_ID":100001,"ENTITIES":[100001]}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}`,
	"GetConfigResult":                         `{"G2_CONFIG":{"CFG_ATTR":[{"ATTR_CLASS":"ADDRESS","ATTR_CODE":"ADDR_CITY","ATTR_ID":1608,"DEFAULT_VALUE":null,"FELEM_CODE":"CITY","FELEM_REQ":"Any",...`,
	"GetConfigsResult":                        `{"CONFIGS":[{"CONFIG_ID":41320074,"CONFIG_COMMENTS":"Example configuration","SYS_CREATE_DT":"2023-02-16 21:43:10.171"},{"CONFIG_ID":1111755672,"CONFIG_COMMENTS":"szconfigmgr_test at 2023-02-16 21:43:10.154619801 +0000 UTC","SYS_CREATE_DT":"2023-02-16 21:43:10.159"},{"CONFIG_ID":3680541328,"CONFIG_COMMENTS":"Created by szdiagnostic_test at 2023-02-16 21:43:07.294747409 +0000 UTC","SYS_CREATE_DT":"2023-02-16 21:43:07.755"}]}`,
	"GetDataSourcesResult":                    `{"DATA_SOURCES":[{"DSRC_ID":1,"DSRC_CODE":"TEST"},{"DSRC_ID":2,"DSRC_CODE":"SEARCH"}]}`,
	"GetDatastoreInfoResult":                  `{"dataStores":[{"id":"CORE","type":"sqlite3","location":"nowhere"}]}`,
	"GetEntityByEntityIDResult":               `{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}`,
	"GetEntityByRecordIDResult":               `{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}`,
	"GetFeatureResult":                        `{"LIB_FEAT_ID":1,"FTYPE_CODE":"NAME","ELEMENTS":[{"FELEM_CODE":"FULL_NAME","FELEM_VALUE":"Robert Smith"},{"FELEM_CODE":"SUR_NAME","FELEM_VALUE":"Smith"},{"FELEM_CODE":"GIVEN_NAME","FELEM_VALUE":"Robert"},{"FELEM_CODE":"CULTURE","FELEM_VALUE":"ANGLO"},{"FELEM_CODE":"CATEGORY","FELEM_VALUE":"PERSON"},{"FELEM_CODE":"TOKENIZED_NM","FELEM_VALUE":"ROBERT|SMITH"}]}`,
	"GetLicenseResult":                        `{"billing":"YEARLY","contract":"Senzing Public Test License","customer":"Senzing Public Test License",...`,
	"GetRecordResult":                         `{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}`,
	"GetRedoRecordResult":                     `{"REASON":"deferred delete","DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1003","REEVAL_ITERATION":1,"DSRC_ACTION":"X"}`,
	"GetStatsResult":                          `{"workload":{"abortedUnresolve":0,"actualAmbiguousTest":0,"addedRecords":3,...`,
	"GetVersionResult":                        `{"PRODUCT_NAME":"Senzing SDK","VERSION":"3.5.0","BUILD_VERSION":"3.5.0.23041","BUILD_DATE":"2023-02-09","BUILD_NUMBER":"2023_02_09__23_01","COMPATIBILITY_VERSION":{"CONFIG_VERSION":"10"},"SCHEMA_VERSION":{"ENGINE_SCHEMA_VERSION":"3.5","MINIMUM_REQUIRED_SCHEMA_VERSION":"3.0","MAXIMUM_REQUIRED_SCHEMA_VERSION":"3.99"}}`,
	"GetVirtualEntityByRecordIDResult":        `{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}`,
	"HowEntityByEntityIDResult":               `{"HOW_RESULTS":{"FINAL_STATE":{"NEED_REEVALUATION":0,"VIRTUAL_ENTITIES":[{"MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]},{"INTERNAL_ID":2,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]}],"VIRTUAL_ENTITY_ID":"V1-S1"}]},"RESOLUTION_STEPS":[{"INBOUND_VIRTUAL_ENTITY_ID":"V2","MATCH_INFO":{"ERRULE_CODE":"CNAME_CFF_CEXCL","MATCH_KEY":"+NAME+DOB+PHONE"},"RESULT_VIRTUAL_ENTITY_ID":"V1-S1","STEP":1,"VIRTUAL_ENTITY_1":{"MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]}],"VIRTUAL_ENTITY_ID":"V1"},"VIRTUAL_ENTITY_2":{"MEMBER_RECORDS":[{"INTERNAL_ID":2,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]}],"VIRTUAL_ENTITY_ID":"V2"}}]}}`,
	"PreprocessRecordResult":                  "{}",
	"ProcessRedoRecordResult":                 "",
	"ReevaluateEntityResult":                  "",
	"ReevaluateRecordResult":                  "",
	"SearchByAttributesResult":                `{"RESOLVED_ENTITIES":[{"ENTITY":{"RESOLVED_ENTITY":{"ENTITY_ID":100001}},"MATCH_INFO":{"ERRULE_CODE":"SF1","MATCH_KEY":"+PNAME+EMAIL","MATCH_LEVEL_CODE":"POSSIBLY_RELATED"}}]}`,
	"WhyEntitiesResult":                       `{"WHY_RESULTS":[{"ENTITY_ID":100001,"ENTITY_ID_2":100001,"MATCH_INFO":{"WHY_KEY":"+NAME+DOB+ADDRESS+PHONE+EMAIL","WHY_ERRULE_CODE":"SF1_SNAME_CFF_CSTAB","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}`,
	"WhyRecordInEntityResult":                 `{"WHY_RESULTS":[{"INTERNAL_ID":100001,"ENTITY_ID":100001,"FOCUS_RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}],"MATCH_INFO":{"WHY_KEY":"+NAME+DOB+PHONE","WHY_ERRULE_CODE":"CNAME_CFF_CEXCL","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}]}`,
	"WhyRecordsResult":                        `{"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":100001}}...`,
}

var Data1_int64s_example = map[string]int64{
	"AddConfigResult":          int64(1),
	"CountRedoRecordsResult":   int64(4),
	"GetActiveConfigIDResult":  int64(1),
	"GetDefaultConfigIDResult": int64(1),
}

var Data1_int64s = map[string]int64{
	"AddConfigResult":          int64(1),
	"CountRedoRecordsResult":   int64(0),
	"GetActiveConfigIDResult":  int64(1),
	"GetDefaultConfigIDResult": int64(1),
}

var Data1_uintptrs = map[string]uintptr{
	"CreateConfigResult":           uintptr(1),
	"ExportCsvEntityReportResult":  uintptr(1),
	"ExportJSONEntityReportResult": uintptr(1),
	"ImportConfigResult":           uintptr(1),
}
