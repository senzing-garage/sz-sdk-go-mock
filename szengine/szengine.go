/*
The [Szengine] implementation of the [senzing.SzEngine] interface
communicates with the Senzing native C binary, libSz.so.
*/
package szengine

import (
	"context"
	"strconv"
	"time"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-observing/notifier"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
	"github.com/senzing-garage/sz-sdk-go-mock/helper"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/senzing-garage/sz-sdk-go/szengine"
	"github.com/senzing-garage/sz-sdk-go/szerror"
)

type Szengine struct {
	AddRecordResult                         string
	CountRedoRecordsResult                  int64
	DeleteRecordResult                      string
	ExportConfigResult                      string
	ExportCsvEntityReportResult             uintptr
	ExportJSONEntityReportResult            uintptr
	FetchNextResult                         string
	FindInterestingEntitiesByEntityIDResult string
	FindInterestingEntitiesByRecordIDResult string
	FindNetworkByEntityIDResult             string
	FindNetworkByRecordIDResult             string
	FindPathByEntityIDResult                string
	FindPathByRecordIDResult                string
	GetActiveConfigIDResult                 int64
	GetEntityByEntityIDResult               string
	GetEntityByRecordIDResult               string
	GetRecordResult                         string
	GetRedoRecordResult                     string
	GetStatsResult                          string
	GetVirtualEntityByRecordIDResult        string
	HowEntityByEntityIDResult               string
	isTrace                                 bool
	logger                                  logging.Logging
	observerOrigin                          string
	observers                               subject.Subject
	PreprocessRecordResult                  string
	ProcessRedoRecordResult                 string
	ReevaluateEntityResult                  string
	ReevaluateRecordResult                  string
	SearchByAttributesResult                string
	WhyEntitiesResult                       string
	WhyRecordInEntityResult                 string
	WhyRecordsResult                        string
	WhySearchResult                         string
}

const (
	baseCallerSkip       = 4
	baseTen              = 10
	initialByteArraySize = 65535
	noError              = 0
)

// ----------------------------------------------------------------------------
// sz-sdk-go.SzEngine interface methods
// ----------------------------------------------------------------------------

/*
Method AddRecord adds a record into the Senzing repository.
The unique identifier of a record is the [dataSourceCode, recordID] compound key.
If the unique identifier does not exist in the Senzing repository, a new record definition is created in the
Senzing repository.
If the unique identifier already exists, the new record definition will replace the old record definition.
If the record definition contains JSON keys of `DATA_SOURCE` and/or `RECORD_ID`, they must match the values of `
dataSourceCode` and `recordID`.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - recordDefinition: A JSON document containing the record to be added to the Senzing repository.
  - flags: Flags used to control information returned.

Output
  - A JSON document containing metadata as specified by the flags.
*/
func (client *Szengine) AddRecord(
	ctx context.Context,
	dataSourceCode string,
	recordID string,
	recordDefinition string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(1, dataSourceCode, recordID, recordDefinition, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(2, dataSourceCode, recordID, recordDefinition, flags, result, err, time.Since(entryTime))
		}()
	}

	result = client.AddRecordResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"flags":          strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8001, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method CloseExportReport closes the exported document created by [Szengine.ExportJSONEntityReport] or
[Szengine.ExportCsvEntityReport].
It is part of the ExportXxxEntityReport(), [Szengine.FetchNext], CloseExportReport lifecycle of a list of entities
to export.
CloseExportReport is idempotent; an exportHandle may be closed multiple times.

Input
  - ctx: A context to control lifecycle.
  - exportHandle: A handle created by [Szengine.ExportJSONEntityReport] or [Szengine.ExportCsvEntityReport]
    that is to be closed.
*/
func (client *Szengine) CloseExportReport(ctx context.Context, exportHandle uintptr) error {
	var err error

	if client.isTrace {
		client.traceEntry(5, exportHandle)

		entryTime := time.Now()

		defer func() { client.traceExit(6, exportHandle, err, time.Since(entryTime)) }()
	}

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8002, err, details)
		}()
	}

	return wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method CountRedoRecords returns the number of records needing re-evaluation.
These are often called "redo records".

Input
  - ctx: A context to control lifecycle.

Output
  - The number of redo records in Senzing's redo queue.
*/
func (client *Szengine) CountRedoRecords(ctx context.Context) (int64, error) {
	var (
		err    error
		result int64
	)

	if client.isTrace {
		client.traceEntry(7)

		entryTime := time.Now()

		defer func() { client.traceExit(8, result, err, time.Since(entryTime)) }()
	}

	result = client.CountRedoRecordsResult

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8003, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method DeleteRecord deletes a record from the Senzing repository.
The unique identifier of a record is the [dataSourceCode, recordID] compound key.
DeleteRecord() is idempotent.
Multiple calls to delete the same unique identifier will all succeed,
even if the unique identifier is not present in the Senzing repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document containing metadata as specified by the flags.
*/
func (client *Szengine) DeleteRecord(
	ctx context.Context,
	dataSourceCode string,
	recordID string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(9, dataSourceCode, recordID, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(10, dataSourceCode, recordID, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.DeleteRecordResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"flags":          strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8004, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method ExportCsvEntityReport initializes a cursor over a CSV document of exported entities.
It is part of the ExportCsvEntityReport, [Szengine.FetchNext], [Szengine.CloseExportReport] lifecycle
of a list of entities to export.
The first exported line is the CSV header.
Each subsequent line contains metadata for a single entity.

Input
  - ctx: A context to control lifecycle.
  - csvColumnList: Use `*` to request all columns, an empty string to request "standard" columns,
    or a comma-separated list of column names for customized columns.
  - flags: Flags used to control information returned.

Output
  - exportHandle: A handle that identifies the document to be scrolled through using [Szengine.FetchNext].
*/
func (client *Szengine) ExportCsvEntityReport(ctx context.Context, csvColumnList string, flags int64) (uintptr, error) {
	var (
		err    error
		result uintptr
	)

	if client.isTrace {
		client.traceEntry(13, csvColumnList, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(14, csvColumnList, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.ExportCsvEntityReportResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"flags": strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8006, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method ExportCsvEntityReportIterator creates an Iterator that can be used in a for-loop
to scroll through a CSV document of exported entities.
It is a convenience method for the [Szenzine.ExportCsvEntityReport], [Szengine.FetchNext], [Szengine.CloseExportReport]
lifecycle of a list of entities to export.

Input
  - ctx: A context to control lifecycle.
  - csvColumnList: Use `*` to request all columns, an empty string to request "standard" columns,
    or a comma-separated list of column names for customized columns.
  - flags: Flags used to control information returned.

Output
  - A channel of strings that can be iterated over.
*/
func (client *Szengine) ExportCsvEntityReportIterator(
	ctx context.Context,
	csvColumnList string,
	flags int64,
) chan senzing.StringFragment {
	stringFragmentChannel := make(chan senzing.StringFragment)

	go func() {
		defer close(stringFragmentChannel)

		var err error

		if client.isTrace {
			client.traceEntry(15, csvColumnList, flags)

			entryTime := time.Now()

			defer func() { client.traceExit(16, csvColumnList, flags, err, time.Since(entryTime)) }()
		}

		if client.observers != nil {
			go func() {
				details := map[string]string{
					"flags": strconv.FormatInt(flags, baseTen),
				}
				notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8007, err, details)
			}()
		}
	}()

	return stringFragmentChannel
}

/*
Method ExportJSONEntityReport initializes a cursor over a JSON document of exported entities.
It is part of the ExportJSONEntityReport, [Szengine.FetchNext], [Szengine.CloseExportReport] lifecycle
of a list of entities to export.

Input
  - ctx: A context to control lifecycle.
  - flags: Flags used to control information returned.

Output
  - A handle that identifies the document to be scrolled through using [Szengine.FetchNext].
*/
func (client *Szengine) ExportJSONEntityReport(ctx context.Context, flags int64) (uintptr, error) {
	var (
		err    error
		result uintptr
	)

	if client.isTrace {
		client.traceEntry(17, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(18, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.ExportJSONEntityReportResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"flags": strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8008, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method ExportJSONEntityReportIterator creates an Iterator that can be used in a for-loop
to scroll through a JSON document of exported entities.
It is a convenience method for the [Szengine.ExportJSONEntityReport], [Szengine.FetchNext], [Szengine.CloseExportReport]
lifecycle of a list of entities to export.

Input
  - ctx: A context to control lifecycle.
  - flags: Flags used to control information returned.

Output
  - A channel of strings that can be iterated over.
*/
func (client *Szengine) ExportJSONEntityReportIterator(ctx context.Context, flags int64) chan senzing.StringFragment {
	stringFragmentChannel := make(chan senzing.StringFragment)

	go func() {
		defer close(stringFragmentChannel)

		var err error

		if client.isTrace {
			client.traceEntry(19, flags)

			entryTime := time.Now()

			defer func() { client.traceExit(20, flags, err, time.Since(entryTime)) }()
		}

		if client.observers != nil {
			go func() {
				details := map[string]string{}
				notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8009, err, details)
			}()
		}
	}()

	return stringFragmentChannel
}

/*
Method FetchNext is used to scroll through an exported JSON or CSV document.
It is part of the [Szengine.ExportJSONEntityReport] or [Szengine.ExportCsvEntityReport], FetchNext,
[Szengine.CloseExportReport] lifecycle of a list of exported entities.

Input
  - ctx: A context to control lifecycle.
  - exportHandle: A handle created by [Szengine.ExportJSONEntityReport] or [Szengine.ExportCsvEntityReport].

Output
  - The next chunk of exported data. An empty string signifies end of data.
*/
func (client *Szengine) FetchNext(ctx context.Context, exportHandle uintptr) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(21, exportHandle)

		entryTime := time.Now()

		defer func() { client.traceExit(22, exportHandle, result, err, time.Since(entryTime)) }()
	}

	result = client.FetchNextResult

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8010, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method FindInterestingEntitiesByEntityID is an experimental method.
Not recommended for use.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) FindInterestingEntitiesByEntityID(
	ctx context.Context,
	entityID int64,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(23, entityID, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(24, entityID, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.FindInterestingEntitiesByEntityIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": formatEntityID(entityID),
				"flags":    strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8011, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method FindInterestingEntitiesByRecordID is an experimental method.
Not recommended for use.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) FindInterestingEntitiesByRecordID(
	ctx context.Context,
	dataSourceCode string,
	recordID string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(25, dataSourceCode, recordID, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(26, dataSourceCode, recordID, flags, result, err, time.Since(entryTime))
		}()
	}

	result = client.FindInterestingEntitiesByRecordIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"flags":          strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8012, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method FindNetworkByEntityID finds a network of entities surrounding a requested set of entities.
This includes the requested entities, paths between them, and relations to other nearby entities.
The size and character of the returned network can be modified by input parameters.

Input
  - ctx: A context to control lifecycle.
  - entityIDs: A JSON document listing entities.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - maxDegrees: The maximum number of degrees in paths between entityIDs.
  - buildOutDegrees: The number of degrees of relationships to show around each search entity. Zero (0)
    prevents buildout.
  - buildOutMaxEntities: The maximum number of entities to build out in the returned network.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) FindNetworkByEntityID(
	ctx context.Context,
	entityIDs string,
	maxDegrees int64,
	buildOutDegrees int64,
	buildOutMaxEntities int64,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(27, entityIDs, maxDegrees, buildOutDegrees, buildOutMaxEntities, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(
				28,
				entityIDs,
				maxDegrees,
				buildOutDegrees,
				buildOutMaxEntities,
				flags,
				result,
				err,
				time.Since(entryTime),
			)
		}()
	}

	result = client.FindNetworkByEntityIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityIDs": entityIDs,
				"flags":     strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8013, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method FindNetworkByRecordID finds a network of entities surrounding a requested set of entities identified
by record keys.
This includes the requested entities, paths between them, and relations to other nearby entities.
The size and character of the returned network can be modified by input parameters.

Input
  - ctx: A context to control lifecycle.
  - recordKeys: A JSON document listing records.
    Example: `{"RECORDS": [{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1001"}]}`
  - maxDegrees: The maximum number of degrees in paths between entities identified by the recordKeys.
  - buildOutDegrees: The number of degrees of relationships to show around each search entity.
    Zero (0) prevents buildout.
  - buildOutMaxEntities: The maximum number of entities to build out in the returned network.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) FindNetworkByRecordID(
	ctx context.Context,
	recordKeys string,
	maxDegrees int64,
	buildOutDegrees int64,
	buildOutMaxEntities int64,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(29, recordKeys, maxDegrees, buildOutDegrees, buildOutMaxEntities, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(
				30,
				recordKeys,
				maxDegrees,
				buildOutDegrees,
				buildOutMaxEntities,
				flags,
				result,
				err,
				time.Since(entryTime),
			)
		}()
	}

	result = client.FindNetworkByRecordIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"recordKeys": recordKeys,
				"flags":      strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8014, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method FindPathByEntityID finds a relationship path between two entities.
Paths are found using known relationships with other entities.
The path can be modified by input parameters.

Input
  - ctx: A context to control lifecycle.
  - startEntityID: The entity ID for the starting entity of the search path.
  - endEntityID: The entity ID for the ending entity of the search path.
  - maxDegrees: The maximum number of degrees in paths between search entities.
  - avoidEntityIDs: A JSON document listing entities that should be avoided on the path.
    An empty string disables this capability.
    Example: `{"ENTITIES": [{"ENTITY_ID": 1}, {"ENTITY_ID": 2}, {"ENTITY_ID": 3}]}`
  - requiredDataSources: A JSON document listing data sources that should be included on the path.
    An empty string disables this capability.
    Example: `{"DATA_SOURCES": ["MY_DATASOURCE_1", "MY_DATASOURCE_2", "MY_DATASOURCE_3"]}`
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) FindPathByEntityID(
	ctx context.Context,
	startEntityID int64,
	endEntityID int64,
	maxDegrees int64,
	avoidEntityIDs string,
	requiredDataSources string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(31, startEntityID, endEntityID, maxDegrees, avoidEntityIDs, requiredDataSources, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(32, startEntityID, endEntityID, maxDegrees, avoidEntityIDs, requiredDataSources,
				flags, result, err, time.Since(entryTime))
		}()
	}

	result = client.FindPathByEntityIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"startEntityID":       formatEntityID(startEntityID),
				"endEntityID":         formatEntityID(endEntityID),
				"avoidEntityIDs":      avoidEntityIDs,
				"requiredDataSources": requiredDataSources,
				"flags":               strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8015, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method FindPathByRecordID finds a relationship path between two entities identified by record keys.
Paths are found using known relationships with other entities.
The path can be modified by input parameters.

Input
  - ctx: A context to control lifecycle.
  - startDataSourceCode: Identifies the provenance of the record for the starting
    entity of the search path.
  - startRecordID: The unique identifier within the records of the same data source
    for the starting entity of the search path.
  - endDataSourceCode: Identifies the provenance of the record for the ending entity
    of the search path.
  - endRecordID: The unique identifier within the records of the same data source for
    the ending entity of the search path.
  - maxDegrees: The maximum number of degrees in paths between search entities.
  - avoidRecordKeys: A JSON document listing entities that should be avoided on the path.
    An empty string disables this capability.
    Example: `{"RECORDS": [
    {"DATA_SOURCE": "MY_DATASOURCE", "RECORD_ID": "1"},
    {"DATA_SOURCE": "MY_DATASOURCE", "RECORD_ID": "2"},
    {"DATA_SOURCE": "MY_DATASOURCE", "RECORD_ID": "3"}
    ]}`
  - requiredDataSources: A JSON document listing data sources that should be included on the path.
    An empty string disables this capability.
    Example: `{"DATA_SOURCES": ["MY_DATASOURCE_1", "MY_DATASOURCE_2", "MY_DATASOURCE_3"]}`
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) FindPathByRecordID(
	ctx context.Context,
	startDataSourceCode string,
	startRecordID string,
	endDataSourceCode string,
	endRecordID string,
	maxDegrees int64,
	avoidRecordKeys string,
	requiredDataSources string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(33, startDataSourceCode, startRecordID, endDataSourceCode, endRecordID, maxDegrees,
			avoidRecordKeys, requiredDataSources, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(
				34,
				startDataSourceCode,
				startRecordID,
				endDataSourceCode,
				endRecordID,
				maxDegrees,
				avoidRecordKeys,
				requiredDataSources,
				flags,
				result,
				err,
				time.Since(entryTime),
			)
		}()
	}

	result = client.FindPathByRecordIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"startDataSourceCode": startDataSourceCode,
				"startRecordID":       startRecordID,
				"endDataSourceCode":   endDataSourceCode,
				"endRecordID":         endRecordID,
				"avoidRecordKeys":     avoidRecordKeys,
				"requiredDataSources": requiredDataSources,
				"flags":               strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8016, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method GetActiveConfigID returns the Senzing configuration JSON document identifier.

Input
  - ctx: A context to control lifecycle.

Output
  - configID: The Senzing configuration JSON document identifier that is currently in use by the Senzing engine.
*/
func (client *Szengine) GetActiveConfigID(ctx context.Context) (int64, error) {
	var (
		err    error
		result int64
	)

	if client.isTrace {
		client.traceEntry(35)

		entryTime := time.Now()

		defer func() { client.traceExit(36, result, err, time.Since(entryTime)) }()
	}

	result = client.GetActiveConfigIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8017, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method GetEntityByEntityID returns information about a resolved identity.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output

  - A JSON document.
*/
func (client *Szengine) GetEntityByEntityID(ctx context.Context, entityID int64, flags int64) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(37, entityID, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(38, entityID, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.GetEntityByEntityIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": formatEntityID(entityID),
				"flags":    strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8018, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method GetEntityByRecordID returns information about a resolved entity identified by a record
which is a member of the entity.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) GetEntityByRecordID(
	ctx context.Context,
	dataSourceCode string,
	recordID string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(39, dataSourceCode, recordID, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(40, dataSourceCode, recordID, flags, result, err, time.Since(entryTime))
		}()
	}

	result = client.GetEntityByRecordIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"flags":          strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8019, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method GetRecord returns a JSON document containing a single record from the Senzing repository.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) GetRecord(
	ctx context.Context,
	dataSourceCode string,
	recordID string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(45, dataSourceCode, recordID, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(46, dataSourceCode, recordID, flags, result, err, time.Since(entryTime))
		}()
	}

	result = client.GetRecordResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"flags":          strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8020, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method GetRedoRecord returns the next maintenance record from the Senzing repository.
Usually, [Szengine.ProcessRedoRecord] is called to process the maintenance record retrieved by GetRedoRecord.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document. If no redo records exist, an empty string is returned.
*/
func (client *Szengine) GetRedoRecord(ctx context.Context) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(47)

		entryTime := time.Now()

		defer func() { client.traceExit(48, result, err, time.Since(entryTime)) }()
	}

	result = client.GetRedoRecordResult

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8021, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method GetStats retrieves workload statistics for the current process.
These statistics are automatically reset after each call.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
*/
func (client *Szengine) GetStats(ctx context.Context) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(49)

		entryTime := time.Now()

		defer func() { client.traceExit(50, result, err, time.Since(entryTime)) }()
	}

	result = client.GetStatsResult

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8022, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method GetVirtualEntityByRecordID describes a hypothetical entity based on a list of records.

Input
  - ctx: A context to control lifecycle.
  - recordKeys: A JSON document listing records to include in the hypothetical entity.
    Example: `{"RECORDS": [{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1001"}]}`
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) GetVirtualEntityByRecordID(
	ctx context.Context,
	recordKeys string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(51, recordKeys, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(52, recordKeys, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.GetVirtualEntityByRecordIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"recordKeys": recordKeys,
				"flags":      strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8023, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method HowEntityByEntityID explains how an entity was constructed from its constituent records.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) HowEntityByEntityID(ctx context.Context, entityID int64, flags int64) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(53, entityID, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(54, entityID, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.HowEntityByEntityIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": formatEntityID(entityID),
				"flags":    strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8024, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method PreprocessRecord tests adding a record into the Senzing repository.

Input
  - ctx: A context to control lifecycle.
  - recordDefinition: A JSON document containing the record to be tested against the Senzing repository.
  - flags: Flags used to control information returned.

Output
  - A JSON document containing metadata as specified by the flags.
*/
func (client *Szengine) PreprocessRecord(ctx context.Context, recordDefinition string, flags int64) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(77, recordDefinition, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(78, recordDefinition, flags, result, err, time.Since(entryTime))
		}()
	}

	result = client.PreprocessRecordResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"flags": strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8035, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method PrimeEngine pre-initializes some of the heavier weight internal resources of the Senzing engine.

Input
  - ctx: A context to control lifecycle.
*/
func (client *Szengine) PrimeEngine(ctx context.Context) error {
	var err error

	if client.isTrace {
		client.traceEntry(57)

		entryTime := time.Now()

		defer func() { client.traceExit(58, err, time.Since(entryTime)) }()
	}

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8026, err, details)
		}()
	}

	return wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method ProcessRedoRecord processes a redo record retrieved by [Szengine.GetRedoRecord].
Calling ProcessRedoRecord has the potential to create more redo records in certain situations.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document.
*/
func (client *Szengine) ProcessRedoRecord(ctx context.Context, redoRecord string, flags int64) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(59, redoRecord, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(60, redoRecord, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.ProcessRedoRecordResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"flags": strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8027, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method ReevaluateEntity verifies that the entity is consistent with its records.
If inconsistent, ReevaluateEntity() adjusts the entity definition, splits entities, and/or merges entities.
Usually, the ReevaluateEntity method is called after a Senzing configuration change to impact
entities immediately.

Input
  - ctx: A context to control lifecycle.
  - entityID: The unique identifier of an entity.
  - flags: Flags used to control information returned.
*/
func (client *Szengine) ReevaluateEntity(ctx context.Context, entityID int64, flags int64) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(61, entityID, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(62, entityID, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.ReevaluateEntityResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID": formatEntityID(entityID),
				"flags":    strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8028, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method ReevaluateRecord verifies that a record is consistent with the entity to which it belongs.
If inconsistent, ReevaluateRecord() adjusts the entity definition, splits entities, and/or merges entities.
Usually, the ReevaluateRecord method is called after a Senzing configuration change to impact
the record immediately.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.
*/
func (client *Szengine) ReevaluateRecord(
	ctx context.Context,
	dataSourceCode string,
	recordID string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(63, dataSourceCode, recordID, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(64, dataSourceCode, recordID, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.ReevaluateRecordResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"flags":          strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8029, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method Reinitialize re-initializes the Senzing engine with a specific Senzing configuration JSON document identifier.

Input
  - ctx: A context to control lifecycle.
  - configID: The Senzing configuration JSON document identifier used for the initialization.
*/
func (client *Szengine) Reinitialize(ctx context.Context, configID int64) error {
	var err error

	if client.isTrace {
		entryTime := time.Now()

		client.traceEntry(65, configID)

		defer func() { client.traceExit(66, configID, err, time.Since(entryTime)) }()
	}

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"configID": strconv.FormatInt(configID, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8030, err, details)
		}()
	}

	return err
}

/*
Method SearchByAttributes retrieves entity data based on entity attributes
and an optional search profile.

Input
  - ctx: A context to control lifecycle.
  - attributes: A JSON document containing the attributes desired in the result set.
    Example: `{"NAME_FULL": "BOB SMITH", "EMAIL_ADDRESS": "bsmith@work.com"}`
  - searchProfile: The name of the search profile to use in the search.
    An empty string will use the default search profile.
    Example: "SEARCH"
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) SearchByAttributes(
	ctx context.Context,
	attributes string,
	searchProfile string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(69, attributes, searchProfile, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(70, attributes, searchProfile, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.SearchByAttributesResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"attributes":    attributes,
				"searchProfile": searchProfile,
				"flags":         strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8031, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method WhyEntities explains the ways in which two entities are related to each other.

Input
  - ctx: A context to control lifecycle.
  - entityID1: The first of two entity IDs.
  - entityID2: The second of two entity IDs.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) WhyEntities(
	ctx context.Context,
	entityID1 int64,
	entityID2 int64,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(71, entityID1, entityID2, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(72, entityID1, entityID2, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.WhyEntitiesResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"entityID1": formatEntityID(entityID1),
				"entityID2": formatEntityID(entityID2),
				"flags":     strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8032, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method WhyRecordInEntity explains why a record belongs to its resolved entitiy.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode: Identifies the provenance of the data.
  - recordID: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) WhyRecordInEntity(
	ctx context.Context,
	dataSourceCode string,
	recordID string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(73, dataSourceCode, recordID, flags)

		entryTime := time.Now()

		defer func() { client.traceExit(74, dataSourceCode, recordID, flags, result, err, time.Since(entryTime)) }()
	}

	result = client.WhyRecordInEntityResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"recordID":       recordID,
				"flags":          strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8033, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method WhyRecords describes ways in which two records are related to each other.

Input
  - ctx: A context to control lifecycle.
  - dataSourceCode1: Identifies the provenance of the data.
  - recordID1: The unique identifier within the records of the same data source.
  - dataSourceCode2: Identifies the provenance of the data.
  - recordID2: The unique identifier within the records of the same data source.
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) WhyRecords(
	ctx context.Context,
	dataSourceCode1 string,
	recordID1 string,
	dataSourceCode2 string,
	recordID2 string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(75, dataSourceCode1, recordID1, dataSourceCode2, recordID2, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(
				76,
				dataSourceCode1,
				recordID1,
				dataSourceCode2,
				recordID2,
				flags,
				result,
				err,
				time.Since(entryTime),
			)
		}()
	}

	result = client.WhyRecordsResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode1": dataSourceCode1,
				"recordID1":       recordID1,
				"dataSourceCode2": dataSourceCode2,
				"recordID2":       recordID2,
				"flags":           strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8034, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method WhySearch ...

Input
  - ctx: A context to control lifecycle.
  - attributes: A JSON document containing the attributes desired in the result set.
    Example: `{"NAME_FULL": "BOB SMITH", "EMAIL_ADDRESS": "bsmith@work.com"}`
  - entityID:
  - searchProfile: The name of the search profile to use in the search.
    An empty string will use the default search profile.
    Example: "SEARCH"
  - flags: Flags used to control information returned.

Output
  - A JSON document.
*/
func (client *Szengine) WhySearch(
	ctx context.Context,
	attributes string,
	entityID int64,
	searchProfile string,
	flags int64,
) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(69, attributes, entityID, searchProfile, flags)

		entryTime := time.Now()

		defer func() {
			client.traceExit(70, attributes, entityID, searchProfile, flags, result, err, time.Since(entryTime))
		}()
	}

	result = client.WhySearchResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"attributes":    attributes,
				"entityID":      formatEntityID(entityID),
				"searchProfile": searchProfile,
				"flags":         strconv.FormatInt(flags, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8031, err, details)
		}()
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

// ----------------------------------------------------------------------------
// Public non-interface methods
// ----------------------------------------------------------------------------

/*
Method GetObserverOrigin returns the "origin" value of past Observer messages.

Input
  - ctx: A context to control lifecycle.

Output
  - The value sent in the Observer's "origin" key/value pair.
*/
func (client *Szengine) GetObserverOrigin(ctx context.Context) string {
	_ = ctx

	return client.observerOrigin
}

/*
Method RegisterObserver adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *Szengine) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error

	if client.isTrace {
		client.traceEntry(703, observer.GetObserverID(ctx))

		entryTime := time.Now()

		defer func() { client.traceExit(704, observer.GetObserverID(ctx), err, time.Since(entryTime)) }()
	}

	if client.observers == nil {
		client.observers = &subject.SimpleSubject{}
	}

	err = client.observers.RegisterObserver(ctx, observer)

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"observerID": observer.GetObserverID(ctx),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8702, err, details)
		}()
	}

	return wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method SetLogLevel sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevelName: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (client *Szengine) SetLogLevel(ctx context.Context, logLevelName string) error {
	var err error

	if client.isTrace {
		client.traceEntry(705, logLevelName)

		entryTime := time.Now()

		defer func() { client.traceExit(706, logLevelName, err, time.Since(entryTime)) }()
	}

	if !logging.IsValidLogLevelName(logLevelName) {
		return wraperror.Errorf(szerror.ErrSzSdk, "invalid error level: %s", logLevelName)
	}

	err = client.getLogger().SetLogLevel(logLevelName)
	client.isTrace = (logLevelName == logging.LevelTraceName)

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"logLevelName": logLevelName,
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8703, err, details)
		}()
	}

	return wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method SetObserverOrigin sets the "origin" value in future Observer messages.

Input
  - ctx: A context to control lifecycle.
  - origin: The value sent in the Observer's "origin" key/value pair.
*/
func (client *Szengine) SetObserverOrigin(ctx context.Context, origin string) {
	_ = ctx
	client.observerOrigin = origin
}

/*
Method UnregisterObserver removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *Szengine) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error

	if client.isTrace {
		client.traceEntry(707, observer.GetObserverID(ctx))

		entryTime := time.Now()

		defer func() { client.traceExit(708, observer.GetObserverID(ctx), err, time.Since(entryTime)) }()
	}

	if client.observers != nil {
		// Tricky code:
		// client.notify is called synchronously before client.observers is set to nil.
		// In client.notify, each observer will get notified in a goroutine.
		// Then client.observers may be set to nil, but observer goroutines will be OK.
		details := map[string]string{
			"observerID": observer.GetObserverID(ctx),
		}
		notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8704, err, details)
		err = client.observers.UnregisterObserver(ctx, observer)

		if !client.observers.HasObservers(ctx) {
			client.observers = nil
		}
	}

	return wraperror.Errorf(err, wraperror.NoMessage)
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// --- Logging ----------------------------------------------------------------

// Get the Logger singleton.
func (client *Szengine) getLogger() logging.Logging {
	if client.logger == nil {
		client.logger = helper.GetLogger(ComponentID, szengine.IDMessages, baseCallerSkip)
	}

	return client.logger
}

// Trace method entry.
func (client *Szengine) traceEntry(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (client *Szengine) traceExit(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

func formatEntityID(entityID int64) string {
	return strconv.FormatInt(entityID, baseTen)
}
