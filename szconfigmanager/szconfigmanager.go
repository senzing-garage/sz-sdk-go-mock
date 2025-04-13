/*
The [Szconfigmanager] implementation of the [senzing.SzConfigManager] interface
communicates with the Senzing native C binary, libSz.so.
*/
package szconfigmanager

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-observing/notifier"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
	"github.com/senzing-garage/sz-sdk-go-mock/helper"
	"github.com/senzing-garage/sz-sdk-go-mock/szconfig"
	"github.com/senzing-garage/sz-sdk-go-mock/testdata"
	"github.com/senzing-garage/sz-sdk-go/senzing"
	"github.com/senzing-garage/sz-sdk-go/szconfigmanager"
	"github.com/senzing-garage/sz-sdk-go/szerror"
)

type Szconfigmanager struct {
	RegisterConfigResult     int64
	GetConfigResult          string
	GetConfigsResult         string
	GetDefaultConfigIDResult int64
	isTrace                  bool
	logger                   logging.Logging
	observerOrigin           string
	observers                subject.Subject
}

const (
	baseCallerSkip       = 4
	baseTen              = 10
	initialByteArraySize = 65535
	noError              = 0
)

// ----------------------------------------------------------------------------
// sz-sdk-go.SzConfigManager interface methods
// ----------------------------------------------------------------------------

/*
Method CreateConfigFromConfigID retrieves a specific Senzing configuration JSON document from the Senzing datastore.

Input
  - ctx: A context to control lifecycle.
  - configID: The identifier of the desired Senzing configuration JSON document to retrieve.

Output
  - senzing.SzConfig:
*/
func (client *Szconfigmanager) CreateConfigFromConfigID(ctx context.Context, configID int64) (senzing.SzConfig, error) {
	var (
		err    error
		result senzing.SzConfig
	)

	if client.isTrace {
		client.traceEntry(7, configID)

		entryTime := time.Now()
		defer func() { client.traceExit(8, configID, result, err, time.Since(entryTime)) }()
	}

	result = getSzConfig(ctx)

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8003, err, details)
		}()
	}

	return result, wraperror.Errorf(err, "szconfigmanager.CreateConfigFromConfigID error: %w", err)
}

func (client *Szconfigmanager) CreateConfigFromString(
	ctx context.Context,
	configDefinition string,
) (senzing.SzConfig, error) {
	var (
		err    error
		result senzing.SzConfig
	)

	if client.isTrace {
		client.traceEntry(23, configDefinition)

		entryTime := time.Now()
		defer func() { client.traceExit(24, configDefinition, result, err, time.Since(entryTime)) }()
	}

	result = getSzConfig(ctx)

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8009, err, details)
		}()
	}

	return result, wraperror.Errorf(err, "szconfigmanager.CreateConfigFromString error: %w", err)
}

/*
Method CreateConfigFromTemplate creates an SzConfig from the template Senzing configuration JSON document.
This document is found in a file on the gRPC server at PIPELINE.RESOURCEPATH/templates/g2config.json

Input
  - ctx: A context to control lifecycle.

Output
  - senzing.SzConfig:
*/
func (client *Szconfigmanager) CreateConfigFromTemplate(ctx context.Context) (senzing.SzConfig, error) {
	var (
		err    error
		result senzing.SzConfig
	)

	if client.isTrace {
		client.traceEntry(25)

		entryTime := time.Now()
		defer func() { client.traceExit(26, result, err, time.Since(entryTime)) }()
	}

	result = getSzConfig(ctx)

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8010, err, details)
		}()
	}

	return result, wraperror.Errorf(err, "szconfigmanager.CreateConfigFromTemplate error: %w", err)
}

/*
Method GetConfigs retrieves a list of Senzing configuration JSON documents from the Senzing datastore.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document listing Senzing configuration JSON document metadata.
*/
func (client *Szconfigmanager) GetConfigs(ctx context.Context) (string, error) {
	var (
		err    error
		result string
	)

	if client.isTrace {
		client.traceEntry(9)

		entryTime := time.Now()
		defer func() { client.traceExit(10, result, err, time.Since(entryTime)) }()
	}

	result = client.GetConfigsResult

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8004, err, details)
		}()
	}

	return result, wraperror.Errorf(err, "szconfigmanager.GetConfigs error: %w", err)
}

/*
Method GetDefaultConfigID retrieves the default Senzing configuration JSON
document identifier from the Senzing datastore.
Note: this may not be the currently active in-memory configuration.
See [Szconfigmanager.SetDefaultConfigID] and [Szconfigmanager.ReplaceDefaultConfigID] for more details.

Input
  - ctx: A context to control lifecycle.

Output
  - configID: The default Senzing configuration JSON document identifier. If none exists, zero (0) is returned.
*/
func (client *Szconfigmanager) GetDefaultConfigID(ctx context.Context) (int64, error) {
	var (
		err    error
		result int64
	)

	if client.isTrace {
		client.traceEntry(11)

		entryTime := time.Now()
		defer func() { client.traceExit(12, result, err, time.Since(entryTime)) }()
	}

	result = client.GetDefaultConfigIDResult

	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8005, err, details)
		}()
	}

	return result, wraperror.Errorf(err, "szconfigmanager.GetDefaultConfigID error: %w", err)
}

/*
Method RegisterConfig adds a Senzing configuration JSON document to the Senzing datastore.

Input
  - ctx: A context to control lifecycle.
  - configDefinition: The Senzing configuration JSON document.
  - configComment: A free-form string describing the Senzing configuration JSON document.

Output
  - configID: A Senzing configuration JSON document identifier.
*/
func (client *Szconfigmanager) RegisterConfig(
	ctx context.Context,
	configDefinition string,
	configComment string) (int64, error) {
	var (
		err    error
		result int64
	)

	if client.isTrace {
		client.traceEntry(1, configDefinition, configComment)

		entryTime := time.Now()
		defer func() {
			client.traceExit(2, configDefinition, configComment, result, err, time.Since(entryTime))
		}()
	}

	result = client.RegisterConfigResult

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"configComment": configComment,
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8001, err, details)
		}()
	}

	return result, wraperror.Errorf(err, "szconfigmanager.RegisterConfig error: %w", err)
}

/*
Similar to the [Szconfigmanager.SetDefaultConfigID] method,
method ReplaceDefaultConfigID sets which Senzing configuration JSON document
is used when initializing or reinitializing the system.
The difference is that ReplaceDefaultConfigID only succeeds when the old Senzing configuration JSON document identifier
is the existing default when the new identifier is applied.
In other words, if currentDefaultConfigID is no longer the "old" identifier, the operation will fail.
It is similar to a "compare-and-swap" instruction to avoid a "race condition".
Note that calling the ReplaceDefaultConfigID method does not affect the currently running in-memory configuration.
To simply set the default Senzing configuration JSON document identifier, use [Szconfigmanager.SetDefaultConfigID].

Input
  - ctx: A context to control lifecycle.
  - currentDefaultConfigID: The Senzing configuration JSON document identifier to replace.
  - newDefaultConfigID: The Senzing configuration JSON document identifier to use as the default.
*/
func (client *Szconfigmanager) ReplaceDefaultConfigID(
	ctx context.Context,
	currentDefaultConfigID int64,
	newDefaultConfigID int64) error {
	var err error

	if client.isTrace {
		client.traceEntry(19, currentDefaultConfigID, newDefaultConfigID)

		entryTime := time.Now()
		defer func() { client.traceExit(20, currentDefaultConfigID, newDefaultConfigID, err, time.Since(entryTime)) }()
	}

	if client.observers != nil {
		go func() {
			details := map[string]string{
				"newDefaultConfigID": strconv.FormatInt(newDefaultConfigID, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8007, err, details)
		}()
	}

	return wraperror.Errorf(err, "szconfigmanager.ReplaceDefaultConfigID error: %w", err)
}

/*
Method SetDefaultConfig sets which Senzing configuration JSON document
is used when initializing or reinitializing the system.
Note that calling the SetDefaultConfig method does not affect the currently
running in-memory configuration.
SetDefaultConfig is susceptible to "race conditions".
To avoid race conditions, see  [Szconfigmanager.ReplaceDefaultConfigID].

Input
  - ctx: A context to control lifecycle.
  - configDefinition: The Senzing configuration JSON document.
  - configComment: A free-form string describing the Senzing configuration JSON document.
*/
func (client *Szconfigmanager) SetDefaultConfig(
	ctx context.Context,
	configDefinition string,
	configComment string) (int64, error) {
	_ = ctx
	_ = configComment
	_ = configDefinition
	return 0, nil
}

/*
Method SetDefaultConfigID sets which Senzing configuration JSON document identifier
is used when initializing or reinitializing the system.
Note that calling the SetDefaultConfigID method does not affect the currently
running in-memory configuration.
SetDefaultConfigID is susceptible to "race conditions".
To avoid race conditions, see  [Szconfigmanager.ReplaceDefaultConfigID].

Input
  - ctx: A context to control lifecycle.
  - configID: The Senzing configuration JSON document identifier to use as the default.
*/
func (client *Szconfigmanager) SetDefaultConfigID(ctx context.Context, configID int64) error {
	var err error

	if client.isTrace {
		client.traceEntry(21, configID)

		entryTime := time.Now()
		defer func() { client.traceExit(22, configID, err, time.Since(entryTime)) }()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"configID": strconv.FormatInt(configID, baseTen),
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8008, err, details)
		}()
	}

	return wraperror.Errorf(err, "szconfigmanager.SetDefaultConfigID error: %w", err)
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
func (client *Szconfigmanager) GetObserverOrigin(ctx context.Context) string {
	_ = ctx

	return client.observerOrigin
}

/*
Method RegisterObserver adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *Szconfigmanager) RegisterObserver(ctx context.Context, observer observer.Observer) error {
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

	return wraperror.Errorf(err, "szconfigmanager.RegisterObserver error: %w", err)
}

/*
Method SetLogLevel sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevelName: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (client *Szconfigmanager) SetLogLevel(ctx context.Context, logLevelName string) error {
	var err error

	if client.isTrace {
		client.traceEntry(705, logLevelName)

		entryTime := time.Now()
		defer func() { client.traceExit(706, logLevelName, err, time.Since(entryTime)) }()
	}

	if !logging.IsValidLogLevelName(logLevelName) {
		return fmt.Errorf("invalid error level: %s; %w", logLevelName, szerror.ErrSzSdk)
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

	return wraperror.Errorf(err, "szconfigmanager.SetLogLevel error: %w", err)
}

/*
Method SetObserverOrigin sets the "origin" value in future Observer messages.

Input
  - ctx: A context to control lifecycle.
  - origin: The value sent in the Observer's "origin" key/value pair.
*/
func (client *Szconfigmanager) SetObserverOrigin(ctx context.Context, origin string) {
	_ = ctx
	client.observerOrigin = origin
}

/*
Method UnregisterObserver removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *Szconfigmanager) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
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

	return wraperror.Errorf(err, "szconfigmanager.UnregisterObserver error: %w", err)
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getSzConfig(ctx context.Context) *szconfig.Szconfig {
	_ = ctx
	testValue := &testdata.TestData{
		Int64s:   testdata.Data1_int64s,
		Strings:  testdata.Data1_strings,
		Uintptrs: testdata.Data1_uintptrs,
	}
	result := &szconfig.Szconfig{
		AddDataSourceResult:  testValue.String("AddDataSourceResult"),
		CreateConfigResult:   testValue.Uintptr("CreateConfigResult"),
		GetDataSourcesResult: testValue.String("GetDataSourcesResult"),
		ImportConfigResult:   testValue.Uintptr("ImportConfigResult"),
		ExportResult:         testValue.String("ExportConfigResult"),
	}

	return result

}

// --- Logging ----------------------------------------------------------------

// Get the Logger singleton.
func (client *Szconfigmanager) getLogger() logging.Logging {
	if client.logger == nil {
		client.logger = helper.GetLogger(ComponentID, szconfigmanager.IDMessages, baseCallerSkip)
	}

	return client.logger
}

// Trace method entry.
func (client *Szconfigmanager) traceEntry(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (client *Szconfigmanager) traceExit(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}
