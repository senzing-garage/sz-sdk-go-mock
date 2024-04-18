package szengine

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/go-helpers/record"
	"github.com/senzing-garage/go-helpers/testfixtures"
	"github.com/senzing-garage/go-helpers/truthset"
	"github.com/senzing-garage/sz-sdk-go/sz"
	"github.com/senzing-garage/sz-sdk-go/szerror"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
	instanceName      = "Engine Test"
	printResults      = false
	verboseLogging    = sz.SZ_NO_LOGGING
)

type GetEntityByRecordIdResponse struct {
	ResolvedEntity struct {
		EntityId int64 `json:"ENTITY_ID"`
	} `json:"RESOLVED_ENTITY"`
}

var (
	szengineSingleton *Szengine
)

// ----------------------------------------------------------------------------
// Interface functions - test
// ----------------------------------------------------------------------------

func TestSzengine_AddRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	flags := sz.SZ_WITHOUT_INFO
	records := []record.Record{
		truthset.CustomerRecords["1004"],
		truthset.CustomerRecords["1005"],
	}
	for _, record := range records {
		actual, err := szEngine.AddRecord(ctx, record.DataSource, record.Id, record.Json, flags)
		testError(test, err)
		printActual(test, actual)
		defer szEngine.DeleteRecord(ctx, record.DataSource, record.Id, flags)
	}
}

func TestSzengine_AddRecord_szBadInput(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record1, err := record.NewRecord(`{"DATA_SOURCE": "TEST", "RECORD_ID": "ADD_TEST_ERR_1", "NAME_FULL": "NOBODY NOMATCH"}`)
	testError(test, err)
	record2, err := record.NewRecord(`{"DATA_SOURCE": "BOB", "RECORD_ID": "ADD_TEST_ERR_2", "NAME_FULL": "ERR BAD SOURCE"}`)
	testError(test, err)
	flags := sz.SZ_WITHOUT_INFO

	// This one should succeed.

	actual, err := szEngine.AddRecord(ctx, record1.DataSource, record1.Id, record1.Json, flags)
	testError(test, err)
	defer szEngine.DeleteRecord(ctx, record1.DataSource, record1.Id, sz.SZ_WITHOUT_INFO)
	printActual(test, actual)

	// This one should fail.

	actual, err = szEngine.AddRecord(ctx, "CUSTOMERS", record2.Id, record2.Json, flags)
	assert.True(test, szerror.Is(err, szerror.SzBadInput))
	printActual(test, actual)

	// Clean-up the records we inserted.

	actual, err = szEngine.DeleteRecord(ctx, record1.DataSource, record1.Id, flags)
	testError(test, err)
	defer szEngine.DeleteRecord(ctx, record2.DataSource, record2.Id, sz.SZ_NO_FLAGS)
	printActual(test, actual)
}

func TestSzengine_AddRecord_withInfo(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	flags := sz.SZ_WITH_INFO
	records := []record.Record{
		truthset.CustomerRecords["1004"],
		truthset.CustomerRecords["1005"],
	}
	for _, record := range records {
		actual, err := szEngine.AddRecord(ctx, record.DataSource, record.Id, record.Json, flags)
		testError(test, err)
		printActual(test, actual)
		defer szEngine.DeleteRecord(ctx, record.DataSource, record.Id, flags)
	}
}

func TestSzengine_CloseExport(test *testing.T) {
	// Tested in:
	//  - TestSzengine_ExportCsvEntityReport
	//  - TestSzengine_ExportJsonEntityReport
}

func TestSzengine_CountRedoRecords(test *testing.T) {
	expected := int64(1)
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.CountRedoRecords(ctx)
	testError(test, err)
	printActual(test, actual)
	assert.Equal(test, expected, actual)
}

func TestSzengine_DeleteRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1009"]
	flags := sz.SZ_WITHOUT_INFO
	actual, err := szEngine.AddRecord(ctx, record.DataSource, record.Id, record.Json, flags)
	printActual(test, actual)
	testError(test, err)
	actual, err = szEngine.DeleteRecord(ctx, record.DataSource, record.Id, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_DeleteRecord_withInfo(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1010"]
	flags := sz.SZ_WITH_INFO
	actual, err := szEngine.AddRecord(ctx, record.DataSource, record.Id, record.Json, flags)
	testError(test, err)
	printActual(test, actual)
	actual, err = szEngine.DeleteRecord(ctx, record.DataSource, record.Id, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_ExportCsvEntityReport(test *testing.T) {
	expected := []string{
		`RESOLVED_ENTITY_ID,RELATED_ENTITY_ID,MATCH_LEVEL_CODE,MATCH_KEY,DATA_SOURCE,RECORD_ID`,
		`1,0,"","","CUSTOMERS","1001"`,
		`1,0,"RESOLVED","+NAME+DOB+PHONE","CUSTOMERS","1002"`,
		`1,0,"RESOLVED","+NAME+DOB+EMAIL","CUSTOMERS","1003"`,
	}
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	csvColumnList := ""
	flags := sz.SZ_EXPORT_INCLUDE_ALL_ENTITIES
	aHandle, err := szEngine.ExportCsvEntityReport(ctx, csvColumnList, flags)
	defer func() {
		err := szEngine.CloseExport(ctx, aHandle)
		testError(test, err)
	}()
	testError(test, err)
	actualCount := 0
	for actual := range szEngine.ExportCsvEntityReportIterator(ctx, csvColumnList, flags) {
		assert.Equal(test, expected[actualCount], strings.TrimSpace(actual.Value))
		actualCount += 1
	}
	assert.Equal(test, len(expected), actualCount)
}

func TestSzengine_ExportCsvEntityReportIterator(test *testing.T) {
	expected := []string{
		`RESOLVED_ENTITY_ID,RELATED_ENTITY_ID,MATCH_LEVEL_CODE,MATCH_KEY,DATA_SOURCE,RECORD_ID`,
		`1,0,"","","CUSTOMERS","1001"`,
		`1,0,"RESOLVED","+NAME+DOB+PHONE","CUSTOMERS","1002"`,
		`1,0,"RESOLVED","+NAME+DOB+EMAIL","CUSTOMERS","1003"`,
	}
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	csvColumnList := ""
	flags := sz.SZ_EXPORT_INCLUDE_ALL_ENTITIES
	actualCount := 0
	for actual := range szEngine.ExportCsvEntityReportIterator(ctx, csvColumnList, flags) {
		testError(test, actual.Error)
		assert.Equal(test, expected[actualCount], strings.TrimSpace(actual.Value))
		actualCount += 1
	}
	assert.Equal(test, len(expected), actualCount)
}

func TestSzengine_ExportJsonEntityReport(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	aRecord := testfixtures.FixtureRecords["65536-periods"]
	flags := sz.SZ_WITH_INFO
	actual, err := szEngine.AddRecord(ctx, aRecord.DataSource, aRecord.Id, aRecord.Json, flags)
	testError(test, err)
	printActual(test, actual)
	defer szEngine.DeleteRecord(ctx, aRecord.DataSource, aRecord.Id, sz.SZ_WITHOUT_INFO)
	// TODO: Figure out correct flags.
	// flags := sz.Flags(sz.SZ_EXPORT_DEFAULT_FLAGS, sz.SZ_EXPORT_INCLUDE_ALL_HAVING_RELATIONSHIPS, sz.SZ_EXPORT_INCLUDE_ALL_HAVING_RELATIONSHIPS)
	flags = int64(-1)
	aHandle, err := szEngine.ExportJsonEntityReport(ctx, flags)
	defer func() {
		err := szEngine.CloseExport(ctx, aHandle)
		testError(test, err)
	}()
	testError(test, err)
	jsonEntityReport := ""
	for {
		jsonEntityReportFragment, err := szEngine.FetchNext(ctx, aHandle)
		testError(test, err)
		if len(jsonEntityReportFragment) == 0 {
			break
		}
		jsonEntityReport += jsonEntityReportFragment
	}
	testError(test, err)
	assert.True(test, len(jsonEntityReport) > 65536)
}

func TestSzengine_ExportJsonEntityReportIterator(test *testing.T) {
	expected := 1
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	flags := sz.SZ_EXPORT_INCLUDE_ALL_ENTITIES
	actualCount := 0
	for actual := range szEngine.ExportJsonEntityReportIterator(ctx, flags) {
		testError(test, actual.Error)
		printActual(test, actual.Value)
		actualCount += 1
	}
	assert.Equal(test, expected, actualCount)
}

func TestSzengine_FetchNext(test *testing.T) {
	// Tested in:
	//  - TestSzengine_ExportJsonEntityReport
}

func TestSzengine_FindNetworkByEntityId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	entityList := `{"ENTITIES": [{"ENTITY_ID": ` + getEntityIdString(record1) + `}, {"ENTITY_ID": ` + getEntityIdString(record2) + `}]}`
	maxDegrees := int64(2)
	buildOutDegree := int64(1)
	maxEntities := int64(10)
	flags := sz.SZ_FIND_NETWORK_DEFAULT_FLAGS
	actual, err := szEngine.FindNetworkByEntityId(ctx, entityList, maxDegrees, buildOutDegree, maxEntities, flags)
	testErrorNoFail(test, err)
	printActual(test, actual)
}

func TestSzengine_FindNetworkByRecordId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	record3 := truthset.CustomerRecords["1003"]
	recordList := `{"RECORDS": [{"DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.Id + `"}, {"DATA_SOURCE": "` + record2.DataSource + `", "RECORD_ID": "` + record2.Id + `"}, {"DATA_SOURCE": "` + record3.DataSource + `", "RECORD_ID": "` + record3.Id + `"}]}`
	maxDegrees := int64(1)
	buildOutDegree := int64(2)
	maxEntities := int64(10)
	flags := sz.SZ_FIND_NETWORK_DEFAULT_FLAGS
	actual, err := szEngine.FindNetworkByRecordId(ctx, recordList, maxDegrees, buildOutDegree, maxEntities, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	startEntityId := getEntityId(truthset.CustomerRecords["1001"])
	endEntityId := getEntityId(truthset.CustomerRecords["1002"])
	maxDegrees := int64(1)
	exclusions := sz.SZ_NO_EXCLUSIONS
	requiredDataSources := sz.SZ_NO_REQUIRED_DATASOURCES
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.FindPathByEntityId(ctx, startEntityId, endEntityId, maxDegrees, exclusions, requiredDataSources, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityId_excluding(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	startRecord := truthset.CustomerRecords["1001"]
	startEntityId := getEntityId(startRecord)
	endEntityId := getEntityId(truthset.CustomerRecords["1002"])
	maxDegrees := int64(1)
	exclusions := `{"ENTITIES": [{"ENTITY_ID": ` + getEntityIdString(startRecord) + `}]}`
	requiredDataSources := sz.SZ_NO_REQUIRED_DATASOURCES
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.FindPathByEntityId(ctx, startEntityId, endEntityId, maxDegrees, exclusions, requiredDataSources, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityId_excludingAndIncluding(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	startRecord := truthset.CustomerRecords["1001"]
	startEntityId := getEntityId(startRecord)
	endEntityId := getEntityId(truthset.CustomerRecords["1002"])
	maxDegrees := int64(1)
	exclusions := `{"ENTITIES": [{"ENTITY_ID": ` + getEntityIdString(startRecord) + `}]}`
	requiredDataSources := `{"DATA_SOURCES": ["` + startRecord.DataSource + `"]}`
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.FindPathByEntityId(ctx, startEntityId, endEntityId, maxDegrees, exclusions, requiredDataSources, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityId_including(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	startRecord := truthset.CustomerRecords["1001"]
	startEntityId := getEntityId(startRecord)
	endEntityId := getEntityId(truthset.CustomerRecords["1002"])
	maxDegrees := int64(1)
	exclusions := sz.SZ_NO_EXCLUSIONS
	requiredDataSources := `{"DATA_SOURCES": ["` + startRecord.DataSource + `"]}`
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.FindPathByEntityId(ctx, startEntityId, endEntityId, maxDegrees, exclusions, requiredDataSources, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByRecordId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	maxDegree := int64(1)
	exclusions := sz.SZ_NO_EXCLUSIONS
	requiredDataSources := sz.SZ_NO_REQUIRED_DATASOURCES
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.FindPathByRecordId(ctx, record1.DataSource, record1.Id, record2.DataSource, record2.Id, maxDegree, exclusions, requiredDataSources, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByRecordId_excluding(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	maxDegree := int64(1)
	exclusions := `{"RECORDS": [{ "DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.Id + `"}]}`
	requiredDataSources := sz.SZ_NO_REQUIRED_DATASOURCES
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.FindPathByRecordId(ctx, record1.DataSource, record1.Id, record2.DataSource, record2.Id, maxDegree, exclusions, requiredDataSources, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByRecordId_excludingAndIncluding(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	maxDegree := int64(1)
	exclusions := `{"RECORDS": [{ "DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.Id + `"}]}`
	requiredDataSources := `{"DATA_SOURCES": ["` + record1.DataSource + `"]}`
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.FindPathByRecordId(ctx, record1.DataSource, record1.Id, record2.DataSource, record2.Id, maxDegree, exclusions, requiredDataSources, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByRecordId_including(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	maxDegree := int64(1)
	exclusions := `{"ENTITIES": [{"ENTITY_ID": ` + getEntityIdString(record1) + `}]}`
	requiredDataSources := `{"DATA_SOURCES": ["` + record1.DataSource + `"]}`
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.FindPathByRecordId(ctx, record1.DataSource, record1.Id, record2.DataSource, record2.Id, maxDegree, exclusions, requiredDataSources, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetActiveConfigId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.GetActiveConfigId(ctx)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetEntityByEntityId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	entityId := getEntityId(truthset.CustomerRecords["1001"])
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.GetEntityByEntityId(ctx, entityId, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetEntityByRecordId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.GetEntityByRecordId(ctx, record.DataSource, record.Id, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.GetRecord(ctx, record.DataSource, record.Id, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetRedoRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.GetRedoRecord(ctx)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetRepositoryLastModifiedTime(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.GetRepositoryLastModifiedTime(ctx)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetStats(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.GetStats(ctx)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetVirtualEntityByRecordId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	recordList := `{"RECORDS": [{"DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.Id + `"}, {"DATA_SOURCE": "` + record2.DataSource + `", "RECORD_ID": "` + record2.Id + `"}]}`
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.GetVirtualEntityByRecordId(ctx, recordList, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_HowEntityByEntityId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	entityId := getEntityId(truthset.CustomerRecords["1001"])
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.HowEntityByEntityId(ctx, entityId, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_PrimeEngine(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	err := szEngine.PrimeEngine(ctx)
	testError(test, err)
}

func TestSzengine_ProcessRedoRecord(test *testing.T) {
	// TODO: Implement TestSzengine_ProcessRedoRecord
	// ctx := context.TODO()
	// szEngine := getTestObject(ctx, test)
	// flags := sz.SZ_WITHOUT_INFO
	// actual, err := szEngine.ProcessRedoRecord(ctx, redoRecord, flags)
	// testError(test, err)
	// printActual(test, actual)
}

func TestSzengine_ProcessRedoRecord_withInfo(test *testing.T) {
	// TODO: Implement TestSzengine_ProcessRedoRecord_withInfo
	// ctx := context.TODO()
	// szEngine := getTestObject(ctx, test)
	// flags := sz.SZ_WITH_INFO
	// actual, err := szEngine.ProcessRedoRecord(ctx, redoRecord, flags)
	// testError(test, err)
	// printActual(test, actual)
}

func TestSzengine_ReevaluateEntity(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	entityId := getEntityId(truthset.CustomerRecords["1001"])
	flags := sz.SZ_WITHOUT_INFO
	actual, err := szEngine.ReevaluateEntity(ctx, entityId, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateEntity_withInfo(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	entityId := getEntityId(truthset.CustomerRecords["1001"])
	flags := sz.SZ_WITH_INFO
	actual, err := szEngine.ReevaluateEntity(ctx, entityId, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := sz.SZ_WITHOUT_INFO
	actual, err := szEngine.ReevaluateRecord(ctx, record.DataSource, record.Id, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateRecord_withInfo(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := sz.SZ_WITH_INFO
	actual, err := szEngine.ReevaluateRecord(ctx, record.DataSource, record.Id, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_SearchByAttributes(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	attributes := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "JOHNSON"}], "SSN_NUMBER": "053-39-3251"}`
	searchProfile := sz.SZ_NO_SEARCH_PROFILE
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.SearchByAttributes(ctx, attributes, searchProfile, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_StreamExportCsvEntityReport(test *testing.T) {
	// TODO: Write TestSzengine_StreamExportCsvEntityReport
}

func TestSzengine_StreamExportJsonEntityReport(test *testing.T) {
	// TODO: Write TestSzengine_StreamExportJsonEntityReport
}

func TestSzengine_SearchByAttributes_searchProfile(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	attributes := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "JOHNSON"}], "SSN_NUMBER": "053-39-3251"}`
	searchProfile := sz.SZ_NO_SEARCH_PROFILE // TODO: Figure out the search profile
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.SearchByAttributes(ctx, attributes, searchProfile, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_WhyEntities(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	entityId1 := getEntityId(truthset.CustomerRecords["1001"])
	entityId2 := getEntityId(truthset.CustomerRecords["1002"])
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.WhyEntities(ctx, entityId1, entityId2, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_WhyRecordInEntity(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.WhyRecordInEntity(ctx, record.DataSource, record.Id, flags)
	testError(test, err)
	printActual(test, actual)
}

func TestSzengine_WhyRecords(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	flags := sz.SZ_NO_FLAGS
	actual, err := szEngine.WhyRecords(ctx, record1.DataSource, record1.Id, record2.DataSource, record2.Id, flags)
	testError(test, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzengine_SetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	szEngine.SetObserverOrigin(ctx, origin)
}

func TestSzengine_GetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	szEngine.SetObserverOrigin(ctx, origin)
	actual := szEngine.GetObserverOrigin(ctx)
	assert.Equal(test, origin, actual)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzengine_AsInterface(test *testing.T) {
	expected := int64(0)
	ctx := context.TODO()
	szEngine := getSzEngineAsInterface(ctx)
	actual, err := szEngine.CountRedoRecords(ctx)
	testError(test, err)
	printActual(test, actual)
	assert.Equal(test, expected, actual)
}

func TestSzengine_Initialize(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	settings, err := getSettings()
	testError(test, err)
	configId := sz.SZ_INITIALIZE_WITH_DEFAULT_CONFIGURATION
	err = szEngine.Initialize(ctx, instanceName, settings, configId, verboseLogging)
	testError(test, err)
}

func TestSzengine_Initialize_withConfigId(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	settings, err := getSettings()
	testError(test, err)
	configId := getDefaultConfigId()
	err = szEngine.Initialize(ctx, instanceName, settings, configId, verboseLogging)
	testError(test, err)
}

func TestSzengine_Reinitialize(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	configId, err := szEngine.GetActiveConfigId(ctx)
	testError(test, err)
	err = szEngine.Reinitialize(ctx, configId)
	testError(test, err)
	printActual(test, configId)
}

func TestSzengine_Destroy(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	err := szEngine.Destroy(ctx)
	testError(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getDefaultConfigId() int64 {
	return int64(1)
}

func getSzEngine(ctx context.Context) *Szengine {
	_ = ctx
	if szengineSingleton == nil {
		szengineSingleton = &Szengine{
			AddRecordResult:                     `{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1003","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`,
			CountRedoRecordsResult:              int64(1),
			DeleteRecordResult:                  `{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1003","AFFECTED_ENTITIES":[],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`,
			ExportConfigResult:                  `{"G2_CONFIG":{"CFG_ETYPE":[{"ETYPE_ID":...`,
			ExportCsvEntityReportResult:         1,
			ExportJsonEntityReportResult:        1,
			FetchNextResult:                     ``,
			FindNetworkByEntityIdResult:         `{"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"Robert Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"CUSTOMERS","RECORD_COUNT":3,"FIRST_SEEN_DT":...`,
			FindNetworkByRecordIdResult:         `{"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"Robert Smith","RECORD_SUMMARY":[{"DATA_SOURCE":"CUSTOMERS","RECORD_COUNT":3,"FIRST_SEEN_DT":...`,
			FindPathByEntityIdResult:            `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":1,"ENTITIES":[1]}],"ENTITIES":[{"RESOLVED_ENTITY":...`,
			FindPathByRecordIdResult:            `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":1,"ENTITIES":[1]}],"ENTITIES":...`,
			GetActiveConfigIdResult:             int64(1),
			GetEntityByEntityIdResult:           `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":...`,
			GetEntityByRecordIdResult:           `{"RESOLVED_ENTITY":{"ENTITY_ID":...`,
			GetRecordResult:                     `{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001","JSON_DATA":{"RECORD_TYPE":"PERSON","PRIMARY_NAME_LAST":"Smith","PRIMARY_NAME_FIRST":"Robert","DATE_OF_BIRTH":"12/11/1978","ADDR_TYPE":"MAILING","ADDR_LINE1":"123 Main Street, Las Vegas NV 89132","PHONE_TYPE":"HOME","PHONE_NUMBER":"702-919-1300","EMAIL_ADDRESS":"bsmith@work.com","DATE":"1/2/18","STATUS":"Active","AMOUNT":"100","DATA_SOURCE":"CUSTOMERS","ENTITY_TYPE":"GENERIC","DSRC_ACTION":"A","RECORD_ID":"1001"}}`,
			GetRedoRecordResult:                 `{"REASON":"deferred delete","DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001","ENTITY_TYPE":"GENERIC","DSRC_ACTION":"X"}`,
			GetRepositoryLastModifiedTimeResult: int64(1),
			GetVirtualEntityByRecordIdResult:    `{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":...`,
			HowEntityByEntityIdResult:           `{"HOW_RESULTS":{"RESOLUTION_STEPS":[{"STEP":1,"VIRTUAL_ENTITY_1":{"VIRTUAL_ENTITY_ID":"V1","MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}]}]},"VIRTUAL_ENTITY_2":{"VIRTUAL_ENTITY_ID":"V2","MEMBER_RECORDS":[{"INTERNAL_ID":2,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1002"}]}]},"INBOUND_VIRTUAL_ENTITY_ID":"V2","RESULT_VIRTUAL_ENTITY_ID":"V1-S1","MATCH_INFO":{"MATCH_KEY":"+NAME+DOB+PHONE","ERRULE_CODE":"CNAME_CFF_CEXCL","FEATURE_SCORES":{"ADDRESS":[{"INBOUND_FEAT_ID":20,"INBOUND_FEAT":"1515 Adela Lane Las Vegas NV 89111","INBOUND_FEAT_USAGE_TYPE":"HOME","CANDIDATE_FEAT_ID":3,"CANDIDATE_FEAT":"123 Main Street, Las Vegas NV 89132","CANDIDATE_FEAT_USAGE_TYPE":"MAILING","FULL_SCORE":42,"SCORE_BUCKET":"NO_CHANCE","SCORE_BEHAVIOR":"FF"}],"DOB":[{"INBOUND_FEAT_ID":19,"INBOUND_FEAT":"11/12/1978","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":2,"CANDIDATE_FEAT":"12/11/1978","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":95,"SCORE_BUCKET":"CLOSE","SCORE_BEHAVIOR":"FMES"}],"NAME":[{"INBOUND_FEAT_ID":18,"INBOUND_FEAT":"Bob Smith","INBOUND_FEAT_USAGE_TYPE":"PRIMARY","CANDIDATE_FEAT_ID":1,"CANDIDATE_FEAT":"Robert Smith","CANDIDATE_FEAT_USAGE_TYPE":"PRIMARY","GNR_FN":97,"GNR_SN":100,"GNR_GN":95,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"CLOSE","SCORE_BEHAVIOR":"NAME"}],"PHONE":[{"INBOUND_FEAT_ID":4,"INBOUND_FEAT":"702-919-1300","INBOUND_FEAT_USAGE_TYPE":"MOBILE","CANDIDATE_FEAT_ID":4,"CANDIDATE_FEAT":"702-919-1300","CANDIDATE_FEAT_USAGE_TYPE":"HOME","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FF"}],"RECORD_TYPE":[{"INBOUND_FEAT_ID":16,"INBOUND_FEAT":"PERSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":16,"CANDIDATE_FEAT":"PERSON","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}]}}},{"STEP":2,"VIRTUAL_ENTITY_1":{"VIRTUAL_ENTITY_ID":"V1-S1","MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}]},{"INTERNAL_ID":2,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1002"}]}]},"VIRTUAL_ENTITY_2":{"VIRTUAL_ENTITY_ID":"V100001","MEMBER_RECORDS":[{"INTERNAL_ID":100001,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1003"}]}]},"INBOUND_VIRTUAL_ENTITY_ID":"V1-S1","RESULT_VIRTUAL_ENTITY_ID":"V1-S2","MATCH_INFO":{"MATCH_KEY":"+NAME+DOB+EMAIL","ERRULE_CODE":"SF1_PNAME_CSTAB","FEATURE_SCORES":{"DOB":[{"INBOUND_FEAT_ID":2,"INBOUND_FEAT":"12/11/1978","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":2,"CANDIDATE_FEAT":"12/11/1978","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FMES"}],"EMAIL":[{"INBOUND_FEAT_ID":5,"INBOUND_FEAT":"bsmith@work.com","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":5,"CANDIDATE_FEAT":"bsmith@work.com","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"F1"}],"NAME":[{"INBOUND_FEAT_ID":18,"INBOUND_FEAT":"Bob Smith","INBOUND_FEAT_USAGE_TYPE":"PRIMARY","CANDIDATE_FEAT_ID":32,"CANDIDATE_FEAT":"Bob J Smith","CANDIDATE_FEAT_USAGE_TYPE":"PRIMARY","GNR_FN":93,"GNR_SN":100,"GNR_GN":93,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"CLOSE","SCORE_BEHAVIOR":"NAME"},{"INBOUND_FEAT_ID":1,"INBOUND_FEAT":"Robert Smith","INBOUND_FEAT_USAGE_TYPE":"PRIMARY","CANDIDATE_FEAT_ID":32,"CANDIDATE_FEAT":"Bob J Smith","CANDIDATE_FEAT_USAGE_TYPE":"PRIMARY","GNR_FN":90,"GNR_SN":100,"GNR_GN":88,"GENERATION_MATCH":-1,"GNR_ON":-1,"SCORE_BUCKET":"CLOSE","SCORE_BEHAVIOR":"NAME"}],"RECORD_TYPE":[{"INBOUND_FEAT_ID":16,"INBOUND_FEAT":"PERSON","INBOUND_FEAT_USAGE_TYPE":"","CANDIDATE_FEAT_ID":16,"CANDIDATE_FEAT":"PERSON","CANDIDATE_FEAT_USAGE_TYPE":"","FULL_SCORE":100,"SCORE_BUCKET":"SAME","SCORE_BEHAVIOR":"FVME"}]}}}],"FINAL_STATE":{"NEED_REEVALUATION":0,"VIRTUAL_ENTITIES":[{"VIRTUAL_ENTITY_ID":"V1-S2","MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}]},{"INTERNAL_ID":2,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1002"}]},{"INTERNAL_ID":100001,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1003"}]}]}]}}}`,
			ProcessRedoRecordResult:             ``,
			ReevaluateEntityResult:              `{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001","AFFECTED_ENTITIES":[{"ENTITY_ID":1}],"INTERESTING_ENTITIES":{"ENTITIES":[]}}`,
			SearchByAttributesResult:            `{"RESOLVED_ENTITIES":[{"MATCH_INFO":{"MATCH_LEVEL":3,"MATCH_LEVEL_CODE":"POSSIBLY_RELATED","MATCH_KEY":"+PNAME+EMAIL","ERRULE_CODE":"SF1","FEATURE_SCORES":{"EMAIL":[{"INBOUND_FEAT":"bsmith@work.com","CANDIDATE_FEAT":"bsmith@work.com","FULL_SCORE":100}],"NAME":[{"INBOUND_FEAT":"Smith","CANDIDATE_FEAT":"Bob J Smith","GNR_FN":83,"GNR_SN":100,"GNR_GN":40,"GENERATION_MATCH":-1,"GNR_ON":-1},{"INBOUND_FEAT":"Smith","CANDIDATE_FEAT":"Robert Smith","GNR_FN":88,"GNR_SN":100,"GNR_GN":40,"GENERATION_MATCH":-1,"GNR_ON":-1}]}},"ENTITY":{"RESOLVED_ENTITY":{"ENTITY_ID":1,"ENTITY_NAME":"Robert Smith","FEATURES":{"ADDRESS":[{"FEAT_DESC":"1515 Adela Lane Las Vegas NV 89111","LIB_FEAT_ID":20,"USAGE_TYPE":"HOME","FEAT_DESC_VALUES":[{"FEAT_DESC":"1515 Adela Lane Las Vegas NV 89111","LIB_FEAT_ID":20}]},{"FEAT_DESC":"123 Main Street, Las Vegas NV 89132","LIB_FEAT_ID":3,"USAGE_TYPE":"MAILING","FEAT_DESC_VALUES":[{"FEAT_DESC":"123 Main Street, Las Vegas NV 89132","LIB_FEAT_ID":3}]}],"DOB":[{"FEAT_DESC":"12/11/1978","LIB_FEAT_ID":2,"FEAT_DESC_VALUES":[{"FEAT_DESC":"12/11/1978","LIB_FEAT_ID":2},{"FEAT_DESC":"11/12/1978","LIB_FEAT_ID":19}]}],"EMAIL":[{"FEAT_DESC":"bsmith@work.com","LIB_FEAT_ID":5,"FEAT_DESC_VALUES":[{"FEAT_DESC":"bsmith@work.com","LIB_FEAT_ID":5}]}],"NAME":[{"FEAT_DESC":"Robert Smith","LIB_FEAT_ID":1,"USAGE_TYPE":"PRIMARY","FEAT_DESC_VALUES":[{"FEAT_DESC":"Robert Smith","LIB_FEAT_ID":1},{"FEAT_DESC":"Bob J Smith","LIB_FEAT_ID":32},{"FEAT_DESC":"Bob Smith","LIB_FEAT_ID":18}]}],"PHONE":[{"FEAT_DESC":"702-919-1300","LIB_FEAT_ID":4,"USAGE_TYPE":"HOME","FEAT_DESC_VALUES":[{"FEAT_DESC":"702-919-1300","LIB_FEAT_ID":4}]},{"FEAT_DESC":"702-919-1300","LIB_FEAT_ID":4,"USAGE_TYPE":"MOBILE","FEAT_DESC_VALUES":[{"FEAT_DESC":"702-919-1300","LIB_FEAT_ID":4}]}],"RECORD_TYPE":[{"FEAT_DESC":"PERSON","LIB_FEAT_ID":16,"FEAT_DESC_VALUES":[{"FEAT_DESC":"PERSON","LIB_FEAT_ID":16}]}]},"RECORD_SUMMARY":[{"DATA_SOURCE":"CUSTOMERS","RECORD_COUNT":3,"FIRST_SEEN_DT":...`,
			GetStatsResult:                      `{ "workload": { "loadedRecords": 5,  "addedRecords": 5,  "deletedRecords": 1,  "reevaluations": 0,  "repairedEntities": 0,  "duration":...`,
			WhyEntitiesResult:                   `{"WHY_RESULTS":[{"ENTITY_ID":1,"ENTITY_ID_2":1,"MATCH_INFO":{"WHY_KEY":...`,
			WhyRecordsResult:                    `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}],"INTERNAL_ID_2":2,"ENTITY_ID_2":1,"FOCUS_RECORDS_2":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1002"}],"MATCH_INFO":{"WHY_KEY":"+NAME+DOB+PHONE","WHY_ERRULE_CODE":"CNAME_CFF_CEXCL","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`,
		}
	}
	return szengineSingleton
}

func getSzEngineAsInterface(ctx context.Context) sz.SzEngine {
	return getSzEngine(ctx)
}

func getEntityId(record record.Record) int64 {
	return getEntityIdForRecord(record.DataSource, record.Id)
}

func getEntityIdForRecord(datasource string, id string) int64 {
	ctx := context.TODO()
	var result int64 = 0
	szEngine := getSzEngine(ctx)
	response, err := szEngine.GetEntityByRecordId(ctx, datasource, id, sz.SZ_WITHOUT_INFO)
	if err != nil {
		return result
	}
	getEntityByRecordIdResponse := &GetEntityByRecordIdResponse{}
	err = json.Unmarshal([]byte(response), &getEntityByRecordIdResponse)
	if err != nil {
		return result
	}
	return getEntityByRecordIdResponse.ResolvedEntity.EntityId
}

func getEntityIdString(record record.Record) string {
	entityId := getEntityId(record)
	return strconv.FormatInt(entityId, 10)
}

func getEntityIdStringForRecord(datasource string, id string) string {
	entityId := getEntityIdForRecord(datasource, id)
	return strconv.FormatInt(entityId, 10)
}

func getSettings() (string, error) {
	return "{}", nil
}

func getTestObject(ctx context.Context, test *testing.T) *Szengine {
	_ = test
	return getSzEngine(ctx)
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func printResult(test *testing.T, title string, result interface{}) {
	if printResults {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result), defaultTruncation))
	}
}

func testError(test *testing.T, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}

func testErrorNoFail(test *testing.T, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		if szerror.Is(err, szerror.SzUnrecoverable) {
			fmt.Printf("\nUnrecoverable error detected. \n\n")
		}
		if szerror.Is(err, szerror.SzRetryable) {
			fmt.Printf("\nRetryable error detected. \n\n")
		}
		if szerror.Is(err, szerror.SzBadInput) {
			fmt.Printf("\nBad user input error detected. \n\n")
		}
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}
