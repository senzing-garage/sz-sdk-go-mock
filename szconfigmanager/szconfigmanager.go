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

	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/go-observing/notifier"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
	"github.com/senzing-garage/sz-sdk-go-mock/helper"
	"github.com/senzing-garage/sz-sdk-go/szconfigmanager"
)

type Szconfigmanager struct {
	AddConfigResult          int64
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
Method AddConfig adds a Senzing configuration JSON document to the Senzing datastore.

Input
  - ctx: A context to control lifecycle.
  - configDefinition: The Senzing configuration JSON document.
  - configComment: A free-form string describing the Senzing configuration JSON document.

Output
  - configID: A Senzing configuration JSON document identifier.
*/
func (client *Szconfigmanager) AddConfig(ctx context.Context, configDefinition string, configComment string) (int64, error) {
	var err error
	result := client.AddConfigResult
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(1, configDefinition, configComment)
		defer func() {
			client.traceExit(2, configDefinition, configComment, result, err, time.Since(entryTime))
		}()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"configComment": configComment,
			}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8001, err, details)
		}()
	}
	return result, err
}

/*
Method GetConfig retrieves a specific Senzing configuration JSON document from the Senzing datastore.

Input
  - ctx: A context to control lifecycle.
  - configID: The identifier of the desired Senzing configuration JSON document to retrieve.

Output
  - configDefinition: A Senzing configuration JSON document.
*/
func (client *Szconfigmanager) GetConfig(ctx context.Context, configID int64) (string, error) {
	var err error
	result := client.GetConfigResult
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(7, configID)
		defer func() { client.traceExit(8, configID, result, err, time.Since(entryTime)) }()
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
Method GetConfigs retrieves a list of Senzing configuration JSON documents from the Senzing datastore.

Input
  - ctx: A context to control lifecycle.

Output
  - A JSON document listing Senzing configuration JSON document metadata.
*/
func (client *Szconfigmanager) GetConfigs(ctx context.Context) (string, error) {
	var err error
	result := client.GetConfigsResult
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(9)
		defer func() { client.traceExit(10, result, err, time.Since(entryTime)) }()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8004, err, details)
		}()
	}
	return result, err
}

/*
Method GetDefaultConfigID retrieves the default Senzing configuration JSON document identifier from the Senzing datastore.
Note: this may not be the currently active in-memory configuration.
See [Szconfigmanager.SetDefaultConfigID] and [Szconfigmanager.ReplaceDefaultConfigID] for more details.

Input
  - ctx: A context to control lifecycle.

Output
  - configID: The default Senzing configuration JSON document identifier. If none exists, zero (0) is returned.
*/
func (client *Szconfigmanager) GetDefaultConfigID(ctx context.Context) (int64, error) {
	var err error
	result := client.GetDefaultConfigIDResult
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(11)
		defer func() { client.traceExit(12, result, err, time.Since(entryTime)) }()
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, client.observerOrigin, ComponentID, 8005, err, details)
		}()
	}
	return result, err
}

/*
Similar to the [Szconfigmanager.SetDefaultConfigID] method,
method ReplaceDefaultConfigID sets which Senzing configuration JSON document is used when initializing or reinitializing the system.
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
func (client *Szconfigmanager) ReplaceDefaultConfigID(ctx context.Context, currentDefaultConfigID int64, newDefaultConfigID int64) error {
	var err error
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(19, currentDefaultConfigID, newDefaultConfigID)
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
	return err
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
		entryTime := time.Now()
		client.traceEntry(21, configID)
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
	return err
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
func (client *Szconfigmanager) SetLogLevel(ctx context.Context, logLevelName string) error {
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
