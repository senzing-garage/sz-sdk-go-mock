/*
Package szconfig implements a client for the service.
*/
package szconfig

import (
	"context"
	"fmt"
	"time"

	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-observing/notifier"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
	"github.com/senzing-garage/sz-sdk-go-mock/helper"
	"github.com/senzing-garage/sz-sdk-go/szconfig"
)

type Szconfig struct {
	AddDataSourceResult  string
	CreateConfigResult   uintptr
	isTrace              bool
	GetDataSourcesResult string
	ImportConfigResult   uintptr
	logger               logging.Logging
	observerOrigin       string
	observers            subject.Subject
	ExportConfigResult   string
}

const (
	baseCallerSkip       = 4
	baseTen              = 10
	initialByteArraySize = 65535
	noError              = 0
)

// ----------------------------------------------------------------------------
// sz-sdk-go.SzConfig interface methods
// ----------------------------------------------------------------------------

/*
Method AddDataSource adds a new data source to an existing in-memory configuration.

Input
  - ctx: A context to control lifecycle.
  - configHandle: Identifier of an in-memory configuration. It was created by the
    [Szconfig.CreateConfig] or [Szconfig.ImportConfig] methods.
  - dataSourceCode: Unique identifier of the data source (e.g. "TEST_DATASOURCE").

Output
  - A JSON document listing the newly created data source.
*/
func (client *Szconfig) AddDataSource(ctx context.Context, configHandle uintptr, dataSourceCode string) (string, error) {
	var err error
	result := client.AddDataSourceResult
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(1, configHandle, dataSourceCode)
		defer func() {
			client.traceExit(2, configHandle, dataSourceCode, result, err, time.Since(entryTime))
		}()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
				"return":         result,
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8001, err, details)
		}()
	}
	return result, err
}

/*
Method CloseConfig terminates an in-memory configuration and cleans up system resources.
After calling CloseConfig, the configuration handle can no longer be used and is invalid.

Input
  - ctx: A context to control lifecycle.
  - configHandle: Identifier of the in-memory configuration. It was created by the
    [Szconfig.CreateConfig] or [Szconfig.ImportConfig] methods.
*/
func (client *Szconfig) CloseConfig(ctx context.Context, configHandle uintptr) error {
	var err error
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(5, configHandle)
		defer func() { client.traceExit(6, configHandle, err, time.Since(entryTime)) }()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8002, err, details)
		}()
	}
	return err
}

/*
Method CreateConfig creates an in-memory configuration using the default template.
The default template is the Senzing configuration JSON document file, g2config.json, located in the PIPELINE.RESOURCEPATH path.
The returned configHandle is used by the [Szconfig.AddDataSource], [Szconfig.DeleteDataSource],
[Szconfig.ExportConfig], and [Szconfig.GetDataSources] methods.
The configHandle is terminated by the [Szconfig.CloseConfig] method.

Input
  - ctx: A context to control lifecycle.

Output
  - configHandle: Identifier of an in-memory configuration.
*/
func (client *Szconfig) CreateConfig(ctx context.Context) (uintptr, error) {
	var err error
	result := client.CreateConfigResult
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(7)
		defer func() { client.traceExit(8, result, err, time.Since(entryTime)) }()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8003, err, details)
		}()
	}
	return result, err
}

/*
Method DeleteDataSource removes a data source from an in-memory configuration.

Input
  - ctx: A context to control lifecycle.
  - configHandle: Identifier of an in-memory configuration. It was created by the
    [Szconfig.CreateConfig] or [Szconfig.ImportConfig] methods.
  - dataSourceCode: Unique identifier of the data source (e.g. "TEST_DATASOURCE").
*/
func (client *Szconfig) DeleteDataSource(ctx context.Context, configHandle uintptr, dataSourceCode string) error {
	var err error
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(9, configHandle, dataSourceCode)
		defer func() { client.traceExit(10, configHandle, dataSourceCode, err, time.Since(entryTime)) }()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"dataSourceCode": dataSourceCode,
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8004, err, details)
		}()
	}
	return err
}

/*
Method ExportConfig creates a Senzing configuration JSON document representation of an in-memory configuration.

Input
  - ctx: A context to control lifecycle.
  - configHandle: Identifier of an in-memory configuration. It was created by the
    [Szconfig.CreateConfig] or [Szconfig.ImportConfig] methods.

Output
  - configDefinition: A Senzing configuration JSON document representation of the in-memory configuration.
*/
func (client *Szconfig) ExportConfig(ctx context.Context, configHandle uintptr) (string, error) {
	var err error
	result := client.ExportConfigResult
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(13, configHandle)
		defer func() { client.traceExit(14, configHandle, result, err, time.Since(entryTime)) }()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8006, err, details)
		}()
	}
	return result, err
}

/*
Method GetDataSources returns a JSON document containing data sources defined in an in-memory configuration.

Input
  - ctx: A context to control lifecycle.
  - configHandle: Identifier of an in-memory configuration. It was created by the
    [Szconfig.CreateConfig] or [Szconfig.ImportConfig] methods.

Output
  - A JSON document listing data sources in the in-memory configuration.
*/
func (client *Szconfig) GetDataSources(ctx context.Context, configHandle uintptr) (string, error) {
	var err error
	result := client.GetDataSourcesResult
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(15, configHandle)
		defer func() { client.traceExit(16, configHandle, result, err, time.Since(entryTime)) }()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8008, err, details)
		}()
	}
	return result, err
}

/*
Method ImportConfig creates a new in-memory configuration from a JSON document.
The returned configHandle is used by the [Szconfig.AddDataSource], [Szconfig.DeleteDataSource],
[Szconfig.ExportConfig], and [Szconfig.GetDataSources] methods.
The configHandle is terminated by the [Szconfig.CloseConfig] method.

Input
  - ctx: A context to control lifecycle.
  - configDefinition: A Senzing configuration JSON document.

Output
  - configHandle: Identifier of the in-memory configuration.
*/
func (client *Szconfig) ImportConfig(ctx context.Context, configDefinition string) (uintptr, error) {
	var err error
	result := client.ImportConfigResult
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(21, configDefinition)
		defer func() { client.traceExit(22, configDefinition, result, err, time.Since(entryTime)) }()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8009, err, details)
		}()
	}
	return result, err
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
func (client *Szconfig) GetObserverOrigin(ctx context.Context) string {
	_ = ctx
	return client.observerOrigin
}

/*
Method RegisterObserver adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *Szconfig) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(703, observer.GetObserverID(ctx))
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
	return err
}

/*
Method SetLogLevel sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevelName: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (client *Szconfig) SetLogLevel(ctx context.Context, logLevelName string) error {
	var err error
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(705, logLevelName)
		defer func() { client.traceExit(706, logLevelName, err, time.Since(entryTime)) }()
	}
	if !logging.IsValidLogLevelName(logLevelName) {
		return fmt.Errorf("invalid error level: %s", logLevelName)
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
	return err
}

/*
Method SetObserverOrigin sets the "origin" value in future Observer messages.

Input
  - ctx: A context to control lifecycle.
  - origin: The value sent in the Observer's "origin" key/value pair.
*/
func (client *Szconfig) SetObserverOrigin(ctx context.Context, origin string) {
	_ = ctx
	client.observerOrigin = origin
}

/*
Method UnregisterObserver removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *Szconfig) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(707, observer.GetObserverID(ctx))
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
	return err
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// --- Logging ----------------------------------------------------------------

// Get the Logger singleton.
func (client *Szconfig) getLogger() logging.Logging {
	if client.logger == nil {
		client.logger = helper.GetLogger(ComponentID, szconfig.IDMessages, baseCallerSkip)
	}
	return client.logger
}

// Trace method entry.
func (client *Szconfig) traceEntry(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (client *Szconfig) traceExit(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}
