package szengine

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/go-helpers/record"
	"github.com/senzing-garage/go-helpers/truthset"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/sz-sdk-go-mock/helper"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	avoidEntityIDs      = senzing.SzNoAvoidance
	avoidRecordKeys     = senzing.SzNoAvoidance
	buildOutDegrees     = int64(2)
	buildOutMaxEntities = int64(10)
	defaultTruncation   = 76
	instanceName        = "SzEngine Test"
	maxDegrees          = int64(2)
	observerOrigin      = "SzEngine observer"
	printResults        = false
	requiredDataSources = senzing.SzNoRequiredDatasources
	searchProfile       = senzing.SzNoSearchProfile
	verboseLogging      = senzing.SzNoLogging
)

// Bad parameters

const (
	badAttributes          = "}{"
	badAvoidEntityIDs      = "}{"
	badAvoidRecordKeys     = "}{"
	badBuildOutDegrees     = int64(-1)
	badBuildOutMaxEntities = int64(-1)
	badCsvColumnList       = "BAD, CSV, COLUMN, LIST"
	badDataSourceCode      = "BadDataSourceCode"
	badEntityID            = int64(-1)
	badExportHandle        = uintptr(0)
	badLogLevelName        = "BadLogLevelName"
	badMaxDegrees          = int64(-1)
	badRecordDefinition    = "}{"
	badRecordID            = "BadRecordID"
	badRedoRecord          = "{}"
	badRequiredDataSources = "}{"
	badSearchProfile       = "}{"
)

// Nil/empty parameters

var (
	nilAvoidEntityIDs      string
	nilBuildOutDegrees     int64
	nilBuildOutMaxEntities int64
	nilCsvColumnList       string
	nilDataSourceCode      string
	nilEntityID            int64
	nilMaxDegrees          int64
	nilRecordID            string
	nilRequiredDataSources string
	nilSearchProfile       string
)

var (
	defaultConfigID   int64
	logLevel          = helper.GetEnv("SENZING_LOG_LEVEL", "INFO")
	observerSingleton = &observer.NullObserver{
		ID:       "Observer 1",
		IsSilent: true,
	}
	szEngineSingleton *Szengine
)

type GetEntityByRecordIDResponse struct {
	ResolvedEntity struct {
		EntityID int64 `json:"ENTITY_ID"`
	} `json:"RESOLVED_ENTITY"`
}

// ----------------------------------------------------------------------------
// Interface methods - test
// ----------------------------------------------------------------------------

func TestSzengine_AddRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	flags := senzing.SzWithoutInfo
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	for _, record := range records {
		actual, err := szEngine.AddRecord(ctx, record.DataSource, record.ID, record.JSON, flags)
		require.NoError(test, err)
		printActual(test, actual)
	}
	for _, record := range records {
		actual, err := szEngine.DeleteRecord(ctx, record.DataSource, record.ID, flags)
		require.NoError(test, err)
		printActual(test, actual)
	}
}

func TestSzengine_AddRecord_nilDataSourceCode(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.AddRecord(ctx, nilDataSourceCode, record.ID, record.JSON, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_AddRecord_nilRecordID(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.AddRecord(ctx, record.DataSource, nilRecordID, record.JSON, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_AddRecord_withInfo(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	flags := senzing.SzWithInfo
	records := []record.Record{
		truthset.CustomerRecords["1003"],
		truthset.CustomerRecords["1004"],
	}
	for _, record := range records {
		actual, err := szEngine.AddRecord(ctx, record.DataSource, record.ID, record.JSON, flags)
		require.NoError(test, err)
		printActual(test, actual)
	}
	for _, record := range records {
		actual, err := szEngine.DeleteRecord(ctx, record.DataSource, record.ID, flags)
		require.NoError(test, err)
		printActual(test, actual)
	}
}

func TestSzengine_CloseExport(test *testing.T) {
	_ = test
	// Tested in:
	//  - TestSzengine_ExportCsvEntityReport
	//  - TestSzengine_ExportJSONEntityReport
}

func TestSzengine_CountRedoRecords(test *testing.T) {
	ctx := context.TODO()
	expected := int64(0)
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.CountRedoRecords(ctx)
	require.NoError(test, err)
	printActual(test, actual)
	assert.Equal(test, expected, actual)
}

func TestSzengine_DeleteRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	records := []record.Record{
		truthset.CustomerRecords["1005"],
	}
	err := addRecords(ctx, records)
	require.NoError(test, err)
	record := truthset.CustomerRecords["1005"]
	flags := senzing.SzWithoutInfo
	require.NoError(test, err)
	actual, err := szEngine.DeleteRecord(ctx, record.DataSource, record.ID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_DeleteRecord_badRecordID(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1005"]
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.DeleteRecord(ctx, record.DataSource, badRecordID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_DeleteRecord_nilRecordID(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1005"]
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.DeleteRecord(ctx, record.DataSource, nilRecordID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_DeleteRecord_withInfo(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	records := []record.Record{
		truthset.CustomerRecords["1009"],
	}
	err := addRecords(ctx, records)
	require.NoError(test, err)
	record := truthset.CustomerRecords["1009"]
	flags := senzing.SzWithInfo
	actual, err := szEngine.DeleteRecord(ctx, record.DataSource, record.ID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_DeleteRecord_withInfo_badDataSourceCode_fix(test *testing.T) {
	_ = test
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1009"],
	}
	handleError(deleteRecords(ctx, records))
}

func TestSzengine_ExportCsvEntityReportIterator(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	csvColumnList := ""
	flags := senzing.SzNoFlags
	for result := range szEngine.ExportCsvEntityReportIterator(ctx, csvColumnList, flags) {
		if result.Error != nil {
			fmt.Println(result.Error)
			break
		}
		fmt.Println(result.Value)
	}
}

func TestSzengine_FetchNext(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.FetchNext(ctx, 0)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ExportCsvEntityReport_nilCsvColumnList(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	flags := senzing.SzExportIncludeAllEntities
	exportHandle, err := szEngine.ExportCsvEntityReport(ctx, nilCsvColumnList, flags)
	defer func() {
		err := szEngine.CloseExport(ctx, exportHandle)
		require.NoError(test, err)
	}()
	require.NoError(test, err)
}

func TestSzengine_FindInterestingEntitiesByEntityID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
		truthset.CustomerRecords["1003"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	entityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindInterestingEntitiesByEntityID(ctx, entityID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindInterestingEntitiesByRecordID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindInterestingEntitiesByRecordID(ctx, record.DataSource, record.ID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindNetworkByEntityID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	entityID1, err := getEntityIDString(record1)
	require.NoError(test, err)
	entityID2, err := getEntityIDString(record2)
	require.NoError(test, err)
	entityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + entityID1 + `}, {"ENTITY_ID": ` + entityID2 + `}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByEntityID(ctx, entityIDs, maxDegrees, buildOutDegrees, buildOutMaxEntities, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindNetworkByEntityID_badMaxDegrees(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	entityID1, err := getEntityIDString(record1)
	require.NoError(test, err)
	entityID2, err := getEntityIDString(record2)
	require.NoError(test, err)
	entityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + entityID1 + `}, {"ENTITY_ID": ` + entityID2 + `}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByEntityID(ctx, entityIDs, badMaxDegrees, buildOutDegrees, buildOutMaxEntities, flags)
	require.NoError(test, err) // TODO: TestSzengine_FindNetworkByEntityID_badMaxDegrees should fail.
	printActual(test, actual)
}

func TestSzengine_FindNetworkByEntityID_badBuildOutDegree(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	entityID1, err := getEntityIDString(record1)
	require.NoError(test, err)
	entityID2, err := getEntityIDString(record2)
	require.NoError(test, err)
	entityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + entityID1 + `}, {"ENTITY_ID": ` + entityID2 + `}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByEntityID(ctx, entityIDs, maxDegrees, badBuildOutDegrees, buildOutMaxEntities, flags)
	require.NoError(test, err) // TODO: TestSzengine_FindNetworkByEntityID_badBuildOutDegree should fail.
	printActual(test, actual)
}

func TestSzengine_FindNetworkByEntityID_badBuildOutMaxEntities(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	entityID1, err := getEntityIDString(record1)
	require.NoError(test, err)
	entityID2, err := getEntityIDString(record2)
	require.NoError(test, err)
	entityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + entityID1 + `}, {"ENTITY_ID": ` + entityID2 + `}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByEntityID(ctx, entityIDs, maxDegrees, buildOutDegrees, badBuildOutMaxEntities, flags)
	require.NoError(test, err) // TODO: TestSzengine_FindNetworkByEntityID_badBuildOutMaxEntities should fail.
	printActual(test, actual)
}

func TestSzengine_FindNetworkByEntityID_nilMaxDegrees(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	entityID1, err := getEntityIDString(record1)
	require.NoError(test, err)
	entityID2, err := getEntityIDString(record2)
	require.NoError(test, err)
	entityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + entityID1 + `}, {"ENTITY_ID": ` + entityID2 + `}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByEntityID(ctx, entityIDs, nilMaxDegrees, buildOutDegrees, buildOutMaxEntities, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindNetworkByEntityID_nilBuildOutDegrees(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	entityID1, err := getEntityIDString(record1)
	require.NoError(test, err)
	entityID2, err := getEntityIDString(record2)
	require.NoError(test, err)
	entityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + entityID1 + `}, {"ENTITY_ID": ` + entityID2 + `}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByEntityID(ctx, entityIDs, maxDegrees, nilBuildOutDegrees, buildOutMaxEntities, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindNetworkByEntityID_nilBuildOutMaxEntities(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	entityID1, err := getEntityIDString(record1)
	require.NoError(test, err)
	entityID2, err := getEntityIDString(record2)
	require.NoError(test, err)
	entityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + entityID1 + `}, {"ENTITY_ID": ` + entityID2 + `}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByEntityID(ctx, entityIDs, maxDegrees, buildOutDegrees, nilBuildOutMaxEntities, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindNetworkByRecordID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
		truthset.CustomerRecords["1003"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	record3 := truthset.CustomerRecords["1003"]
	recordKeys := `{"RECORDS": [{"DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.ID + `"}, {"DATA_SOURCE": "` + record2.DataSource + `", "RECORD_ID": "` + record2.ID + `"}, {"DATA_SOURCE": "` + record3.DataSource + `", "RECORD_ID": "` + record3.ID + `"}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByRecordID(ctx, recordKeys, maxDegrees, buildOutDegrees, buildOutMaxEntities, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindNetworkByRecordID_nilMaxDegrees(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
		truthset.CustomerRecords["1003"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	record3 := truthset.CustomerRecords["1003"]
	recordKeys := `{"RECORDS": [{"DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.ID + `"}, {"DATA_SOURCE": "` + record2.DataSource + `", "RECORD_ID": "` + record2.ID + `"}, {"DATA_SOURCE": "` + record3.DataSource + `", "RECORD_ID": "` + record3.ID + `"}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByRecordID(ctx, recordKeys, nilMaxDegrees, buildOutDegrees, buildOutMaxEntities, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindNetworkByRecordID_nilBuildOutDegrees(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
		truthset.CustomerRecords["1003"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	record3 := truthset.CustomerRecords["1003"]
	recordKeys := `{"RECORDS": [{"DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.ID + `"}, {"DATA_SOURCE": "` + record2.DataSource + `", "RECORD_ID": "` + record2.ID + `"}, {"DATA_SOURCE": "` + record3.DataSource + `", "RECORD_ID": "` + record3.ID + `"}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByRecordID(ctx, recordKeys, maxDegrees, nilBuildOutDegrees, buildOutMaxEntities, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindNetworkByRecordID_nilBuildOutMaxEntities(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
		truthset.CustomerRecords["1003"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	record3 := truthset.CustomerRecords["1003"]
	recordKeys := `{"RECORDS": [{"DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.ID + `"}, {"DATA_SOURCE": "` + record2.DataSource + `", "RECORD_ID": "` + record2.ID + `"}, {"DATA_SOURCE": "` + record3.DataSource + `", "RECORD_ID": "` + record3.ID + `"}]}`
	flags := senzing.SzFindNetworkDefaultFlags
	actual, err := szEngine.FindNetworkByRecordID(ctx, recordKeys, maxDegrees, buildOutDegrees, nilBuildOutMaxEntities, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	startEntityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	endEntityID, err := getEntityID(truthset.CustomerRecords["1002"])
	require.NoError(test, err)
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, maxDegrees, avoidEntityIDs, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityID_badMaxDegrees(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	startEntityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	endEntityID, err := getEntityID(truthset.CustomerRecords["1002"])
	require.NoError(test, err)
	exclusions := senzing.SzNoAvoidance
	requiredDataSources := senzing.SzNoRequiredDatasources
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, badMaxDegrees, exclusions, requiredDataSources, flags)
	require.NoError(test, err) // TODO: TestSzengine_FindPathByEntityID_badMaxDegrees should fail.
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityID_nilMaxDegrees(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	startEntityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	endEntityID, err := getEntityID(truthset.CustomerRecords["1002"])
	require.NoError(test, err)
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, nilMaxDegrees, avoidEntityIDs, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityID_nilAvoidEntityIDs(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	startEntityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	endEntityID, err := getEntityID(truthset.CustomerRecords["1002"])
	require.NoError(test, err)
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, maxDegrees, nilAvoidEntityIDs, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityID_nilRequiredDataSource(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	startEntityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	endEntityID, err := getEntityID(truthset.CustomerRecords["1002"])
	require.NoError(test, err)
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, maxDegrees, avoidEntityIDs, nilRequiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityID_avoiding(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	startRecord := truthset.CustomerRecords["1001"]
	startEntityID, err := getEntityID(startRecord)
	require.NoError(test, err)
	endEntityID, err := getEntityID(truthset.CustomerRecords["1002"])
	require.NoError(test, err)
	startEntityIDString, err := getEntityIDStringForRecord("CUSTOMERS", "1001")
	require.NoError(test, err)
	avoidEntityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + startEntityIDString + `}]}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, maxDegrees, avoidEntityIDs, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityID_avoidingAndIncluding(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	startRecord := truthset.CustomerRecords["1001"]
	startEntityID, err := getEntityID(startRecord)
	require.NoError(test, err)
	endEntityID, err := getEntityID(truthset.CustomerRecords["1002"])
	require.NoError(test, err)
	startRecordEntityIDString, err := getEntityIDString(startRecord)
	require.NoError(test, err)
	avoidEntityIDs := `{"ENTITIES": [{"ENTITY_ID": ` + startRecordEntityIDString + `}]}`
	requiredDataSources := `{"DATA_SOURCES": ["` + startRecord.DataSource + `"]}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, maxDegrees, avoidEntityIDs, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByEntityID_including(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	startRecord := truthset.CustomerRecords["1001"]
	startEntityID, err := getEntityID(startRecord)
	require.NoError(test, err)
	endEntityID, err := getEntityID(truthset.CustomerRecords["1002"])
	require.NoError(test, err)
	requiredDataSources := `{"DATA_SOURCES": ["` + startRecord.DataSource + `"]}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByEntityID(ctx, startEntityID, endEntityID, maxDegrees, avoidEntityIDs, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByRecordID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByRecordID(ctx, record1.DataSource, record1.ID, record2.DataSource, record2.ID, maxDegrees, avoidRecordKeys, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByRecordID_avoiding(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	avoidRecordKeys := `{"RECORDS": [{ "DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.ID + `"}]}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByRecordID(ctx, record1.DataSource, record1.ID, record2.DataSource, record2.ID, maxDegrees, avoidRecordKeys, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByRecordID_avoidingAndIncluding(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	avoidRecordKeys := `{"RECORDS": [{ "DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.ID + `"}]}`
	requiredDataSources := `{"DATA_SOURCES": ["` + record1.DataSource + `"]}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByRecordID(ctx, record1.DataSource, record1.ID, record2.DataSource, record2.ID, maxDegrees, avoidRecordKeys, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_FindPathByRecordID_including(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	record1EntityID, err := getEntityIDString(record1)
	require.NoError(test, err)
	avoidRecordKeys := `{"ENTITIES": [{"ENTITY_ID": ` + record1EntityID + `}]}`
	requiredDataSources := `{"DATA_SOURCES": ["` + record1.DataSource + `"]}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.FindPathByRecordID(ctx, record1.DataSource, record1.ID, record2.DataSource, record2.ID, maxDegrees, avoidRecordKeys, requiredDataSources, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetActiveConfigID(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.GetActiveConfigID(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetEntityByEntityID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	entityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	flags := senzing.SzNoFlags
	actual, err := szEngine.GetEntityByEntityID(ctx, entityID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetEntityByRecordID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzNoFlags
	actual, err := szEngine.GetEntityByRecordID(ctx, record.DataSource, record.ID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetRecord(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzNoFlags
	actual, err := szEngine.GetRecord(ctx, record.DataSource, record.ID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetRedoRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.GetRedoRecord(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetStats(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	actual, err := szEngine.GetStats(ctx)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_GetVirtualEntityByRecordID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	recordList := `{"RECORDS": [{"DATA_SOURCE": "` + record1.DataSource + `", "RECORD_ID": "` + record1.ID + `"}, {"DATA_SOURCE": "` + record2.DataSource + `", "RECORD_ID": "` + record2.ID + `"}]}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.GetVirtualEntityByRecordID(ctx, recordList, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_HowEntityByEntityID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	entityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	flags := senzing.SzNoFlags
	actual, err := szEngine.HowEntityByEntityID(ctx, entityID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_PreprocessRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	flags := senzing.SzWithoutInfo
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	for _, record := range records {
		actual, err := szEngine.PreprocessRecord(ctx, record.JSON, flags)
		require.NoError(test, err)
		printActual(test, actual)
	}
}

func TestSzengine_PrimeEngine(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	err := szEngine.PrimeEngine(ctx)
	require.NoError(test, err)
}

func TestSzengine_ProcessRedoRecord(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	redoRecord, err := szEngine.GetRedoRecord(ctx)
	require.NoError(test, err)
	if len(redoRecord) > 0 {
		flags := senzing.SzWithoutInfo
		actual, err := szEngine.ProcessRedoRecord(ctx, redoRecord, flags)
		require.NoError(test, err)
		printActual(test, actual)
	}
}

func TestSzengine_ProcessRedoRecord_withInfo(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
		truthset.CustomerRecords["1003"],
		truthset.CustomerRecords["1004"],
		truthset.CustomerRecords["1005"],
		truthset.CustomerRecords["1009"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	redoRecord, err := szEngine.GetRedoRecord(ctx)
	require.NoError(test, err)
	if len(redoRecord) > 0 {
		flags := senzing.SzWithInfo
		actual, err := szEngine.ProcessRedoRecord(ctx, redoRecord, flags)
		require.NoError(test, err)
		printActual(test, actual)
	}
}

func TestSzengine_ReevaluateEntity(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	entityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.ReevaluateEntity(ctx, entityID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateEntity_badEntityID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.ReevaluateEntity(ctx, badEntityID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateEntity_nilEntityID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.ReevaluateEntity(ctx, nilEntityID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateEntity_withInfo(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	entityID, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	flags := senzing.SzWithInfo
	actual, err := szEngine.ReevaluateEntity(ctx, entityID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateEntity_withInfo_badEntityID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	flags := senzing.SzWithInfo
	actual, err := szEngine.ReevaluateEntity(ctx, badEntityID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateEntity_withInfo_nilEntityID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	flags := senzing.SzWithInfo
	actual, err := szEngine.ReevaluateEntity(ctx, nilEntityID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateRecord(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.ReevaluateRecord(ctx, record.DataSource, record.ID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateRecord_badRecordID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.ReevaluateRecord(ctx, record.DataSource, badRecordID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateRecord_nilRecordID(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzWithoutInfo
	actual, err := szEngine.ReevaluateRecord(ctx, record.DataSource, nilRecordID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_ReevaluateRecord_withInfo(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzWithInfo
	actual, err := szEngine.ReevaluateRecord(ctx, record.DataSource, record.ID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_SearchByAttributes(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
		truthset.CustomerRecords["1003"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	attributes := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "JOHNSON"}], "SSN_NUMBER": "053-39-3251"}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.SearchByAttributes(ctx, attributes, searchProfile, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_SearchByAttributes_nilSearchProfile(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
		truthset.CustomerRecords["1003"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	attributes := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "JOHNSON"}], "SSN_NUMBER": "053-39-3251"}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.SearchByAttributes(ctx, attributes, nilSearchProfile, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_SearchByAttributes_withSearchProfile(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
		truthset.CustomerRecords["1003"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	attributes := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "JOHNSON"}], "SSN_NUMBER": "053-39-3251"}`
	searchProfile := "SEARCH"
	flags := senzing.SzNoFlags
	actual, err := szEngine.SearchByAttributes(ctx, attributes, searchProfile, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

// TODO: Implement TestSzengine_StreamExportCsvEntityReport
// func TestSzengine_StreamExportCsvEntityReport(test *testing.T) {}

// TODO: Implement TestSzengine_StreamExportJSONEntityReport
// func TestSzengine_StreamExportJSONEntityReport(test *testing.T) {}

func TestSzengine_SearchByAttributes_searchProfile(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	attributes := `{"NAMES": [{"NAME_TYPE": "PRIMARY", "NAME_LAST": "JOHNSON"}], "SSN_NUMBER": "053-39-3251"}`
	flags := senzing.SzNoFlags
	actual, err := szEngine.SearchByAttributes(ctx, attributes, searchProfile, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_WhyEntities(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	entityID1, err := getEntityID(truthset.CustomerRecords["1001"])
	require.NoError(test, err)
	entityID2, err := getEntityID(truthset.CustomerRecords["1002"])
	require.NoError(test, err)
	flags := senzing.SzNoFlags
	actual, err := szEngine.WhyEntities(ctx, entityID1, entityID2, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_WhyRecordInEntity(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record := truthset.CustomerRecords["1001"]
	flags := senzing.SzNoFlags
	actual, err := szEngine.WhyRecordInEntity(ctx, record.DataSource, record.ID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

func TestSzengine_WhyRecords(test *testing.T) {
	ctx := context.TODO()
	records := []record.Record{
		truthset.CustomerRecords["1001"],
		truthset.CustomerRecords["1002"],
	}
	defer func() { handleError(deleteRecords(ctx, records)) }()
	err := addRecords(ctx, records)
	require.NoError(test, err)
	szEngine := getTestObject(ctx, test)
	record1 := truthset.CustomerRecords["1001"]
	record2 := truthset.CustomerRecords["1002"]
	flags := senzing.SzNoFlags
	actual, err := szEngine.WhyRecords(ctx, record1.DataSource, record1.ID, record2.DataSource, record2.ID, flags)
	require.NoError(test, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzengine_SetLogLevel_badLogLevelName(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	_ = szConfig.SetLogLevel(ctx, badLogLevelName)
}

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

func TestSzengine_UnregisterObserver(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	err := szEngine.UnregisterObserver(ctx, observerSingleton)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzengine_AsInterface(test *testing.T) {
	expected := int64(0)
	ctx := context.TODO()
	szEngine := getSzEngineAsInterface(ctx)
	actual, err := szEngine.CountRedoRecords(ctx)
	require.NoError(test, err)
	printActual(test, actual)
	assert.Equal(test, expected, actual)
}

func TestSzengine_Initialize(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	settings, err := getSettings()
	require.NoError(test, err)
	configID := senzing.SzInitializeWithDefaultConfiguration
	err = szEngine.Initialize(ctx, instanceName, settings, configID, verboseLogging)
	require.NoError(test, err)
}

// TODO: Implement TestSzengine_Initialize_error
// func TestSzengine_Initialize_error(test *testing.T) {}

func TestSzengine_Initialize_withConfigID(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	settings, err := getSettings()
	require.NoError(test, err)
	configID := getDefaultConfigID()
	err = szEngine.Initialize(ctx, instanceName, settings, configID, verboseLogging)
	require.NoError(test, err)
}

// TODO: Implement TestSzengine_Initialize_withConfigID_error
// func TestSzengine_Initialize_withConfigID_error(test *testing.T) {}

func TestSzengine_Reinitialize(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	configID, err := szEngine.GetActiveConfigID(ctx)
	require.NoError(test, err)
	err = szEngine.Reinitialize(ctx, configID)
	require.NoError(test, err)
	printActual(test, configID)
}

// TODO: Implement TestSzengine_Reinitialize_badConfigID
// func TestSzengine_Reinitialize_badConfigID(test *testing.T) {}

func TestSzengine_Destroy(test *testing.T) {
	ctx := context.TODO()
	szEngine := getTestObject(ctx, test)
	err := szEngine.Destroy(ctx)
	require.NoError(test, err)
}

// TODO: Implement TestSzengine_Destroy_error
// func TestSzengine_Destroy_error(test *testing.T) {}

func TestSzengine_Destroy_withObserver(test *testing.T) {
	ctx := context.TODO()
	szEngineSingleton = nil
	szEngine := getTestObject(ctx, test)
	err := szEngine.Destroy(ctx)
	require.NoError(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func addRecords(ctx context.Context, records []record.Record) error {
	var err error
	_ = ctx
	_ = records
	return err
}

func deleteRecords(ctx context.Context, records []record.Record) error {
	var err error
	_ = ctx
	_ = records
	return err
}

func getDefaultConfigID() int64 {
	return defaultConfigID
}

func getEntityID(record record.Record) (int64, error) {
	return getEntityIDForRecord(record.DataSource, record.ID)
}

func getEntityIDForRecord(datasource string, id string) (int64, error) {
	ctx := context.TODO()
	var result int64
	szEngine, err := getSzEngine(ctx)
	if err != nil {
		return result, err
	}
	response, err := szEngine.GetEntityByRecordID(ctx, datasource, id, senzing.SzWithoutInfo)
	if err != nil {
		return result, err
	}
	getEntityByRecordIDResponse := &GetEntityByRecordIDResponse{}
	err = json.Unmarshal([]byte(response), &getEntityByRecordIDResponse)
	if err != nil {
		return result, err
	}
	return getEntityByRecordIDResponse.ResolvedEntity.EntityID, err
}

func getEntityIDString(record record.Record) (string, error) {
	entityID, err := getEntityID(record)
	return strconv.FormatInt(entityID, baseTen), err
}

func getEntityIDStringForRecord(datasource string, id string) (string, error) {
	entityID, err := getEntityIDForRecord(datasource, id)
	return strconv.FormatInt(entityID, baseTen), err
}

func getSettings() (string, error) {
	return "{}", nil
}

func getSzEngine(ctx context.Context) (*Szengine, error) {
	var err error
	if szEngineSingleton == nil {
		settings, err := getSettings()
		if err != nil {
			return szEngineSingleton, fmt.Errorf("getSettings() Error: %w", err)
		}
		szEngineSingleton = &Szengine{
			AddRecordResult:                         "{}",
			CountRedoRecordsResult:                  int64(0),
			DeleteRecordResult:                      "{}",
			ExportConfigResult:                      `{"G2_CONFIG":{"CFG_ETYPE":[{"ETYPE_ID":...`,
			ExportCsvEntityReportResult:             1,
			ExportJSONEntityReportResult:            1,
			FetchNextResult:                         ``,
			FindInterestingEntitiesByEntityIDResult: `{"INTERESTING_ENTITIES":{"ENTITIES":[]}}`,
			FindInterestingEntitiesByRecordIDResult: `{"INTERESTING_ENTITIES":{"ENTITIES":[]}}`,
			FindNetworkByEntityIDResult:             `{"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`,
			FindNetworkByRecordIDResult:             `{"ENTITY_PATHS":[],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`,
			FindPathByEntityIDResult:                `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":1,"ENTITIES":[1]}],"ENTITIES":[{"RESOLVED_ENTITY":...`,
			FindPathByRecordIDResult:                `{"ENTITY_PATHS":[{"START_ENTITY_ID":1,"END_ENTITY_ID":1,"ENTITIES":[1]}],"ENTITIES":...`,
			GetActiveConfigIDResult:                 int64(1),
			GetEntityByEntityIDResult:               `{"RESOLVED_ENTITY":{"ENTITY_ID":1}}`,
			GetEntityByRecordIDResult:               `{"RESOLVED_ENTITY":{"ENTITY_ID":1}}`,
			GetRecordResult:                         `{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}`,
			GetRedoRecordResult:                     `{"REASON":"deferred delete","DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001","DSRC_ACTION":"X"}`,
			GetStatsResult:                          `{ "workload": { "loadedRecords": 5,  "addedRecords": 5,  "deletedRecords": 1,  "reevaluations": 0,  "repairedEntities": 0,  "duration":...`,
			GetVirtualEntityByRecordIDResult:        `{"RESOLVED_ENTITY":{"ENTITY_ID":1}}`,
			HowEntityByEntityIDResult:               `{"HOW_RESULTS":{"FINAL_STATE":{"NEED_REEVALUATION":0,"VIRTUAL_ENTITIES":[{"MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]},{"INTERNAL_ID":2,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]}],"VIRTUAL_ENTITY_ID":"V1-S1"}]},"RESOLUTION_STEPS":[{"INBOUND_VIRTUAL_ENTITY_ID":"V2","MATCH_INFO":{"ERRULE_CODE":"CNAME_CFF_CEXCL","MATCH_KEY":"+NAME+DOB+PHONE"},"RESULT_VIRTUAL_ENTITY_ID":"V1-S1","STEP":1,"VIRTUAL_ENTITY_1":{"MEMBER_RECORDS":[{"INTERNAL_ID":1,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]}],"VIRTUAL_ENTITY_ID":"V1"},"VIRTUAL_ENTITY_2":{"MEMBER_RECORDS":[{"INTERNAL_ID":2,"RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":null}]}],"VIRTUAL_ENTITY_ID":"V2"}}]}}`,
			PreprocessRecordResult:                  "{}",
			ProcessRedoRecordResult:                 ``,
			ReevaluateEntityResult:                  "{}",
			ReevaluateRecordResult:                  "{}",
			SearchByAttributesResult:                `{"RESOLVED_ENTITIES":[{"ENTITY":{"RESOLVED_ENTITY":{"ENTITY_ID":1}},"MATCH_INFO":{"ERRULE_CODE":"SF1","MATCH_KEY":"+PNAME+EMAIL","MATCH_LEVEL_CODE":"POSSIBLY_RELATED"}}]}`,
			WhyEntitiesResult:                       `{"WHY_RESULTS":[{"ENTITY_ID":1,"ENTITY_ID_2":1,"MATCH_INFO":{"WHY_KEY":...`,
			WhyRecordsResult:                        `{"WHY_RESULTS":[{"INTERNAL_ID":1,"ENTITY_ID":1,"FOCUS_RECORDS":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1001"}],"INTERNAL_ID_2":2,"ENTITY_ID_2":1,"FOCUS_RECORDS_2":[{"DATA_SOURCE":"CUSTOMERS","RECORD_ID":"1002"}],"MATCH_INFO":{"WHY_KEY":"+NAME+DOB+PHONE","WHY_ERRULE_CODE":"CNAME_CFF_CEXCL","MATCH_LEVEL_CODE":"RESOLVED"}}],"ENTITIES":[{"RESOLVED_ENTITY":{"ENTITY_ID":1}}]}`,
		}
		err = szEngineSingleton.SetLogLevel(ctx, logLevel)
		if err != nil {
			return szEngineSingleton, fmt.Errorf("SetLogLevel() Error: %w", err)
		}
		if logLevel == "TRACE" {
			szEngineSingleton.SetObserverOrigin(ctx, observerOrigin)
			err = szEngineSingleton.RegisterObserver(ctx, observerSingleton)
			if err != nil {
				return szEngineSingleton, fmt.Errorf("RegisterObserver() Error: %w", err)
			}
			err = szEngineSingleton.SetLogLevel(ctx, logLevel) // Duplicated for coverage testing
			if err != nil {
				return szEngineSingleton, fmt.Errorf("SetLogLevel() - 2 Error: %w", err)
			}
		}
		err = szEngineSingleton.Initialize(ctx, instanceName, settings, getDefaultConfigID(), verboseLogging)
		if err != nil {
			return szEngineSingleton, fmt.Errorf("Initialize() Error: %w", err)
		}
	}
	return szEngineSingleton, err
}

func getSzEngineAsInterface(ctx context.Context) senzing.SzEngine {
	result, err := getSzEngine(ctx)
	if err != nil {
		panic(err)
	}
	return result
}

func getTestObject(ctx context.Context, test *testing.T) *Szengine {
	result, err := getSzEngine(ctx)
	require.NoError(test, err)
	return result
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
