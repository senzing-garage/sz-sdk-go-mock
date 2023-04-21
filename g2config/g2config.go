/*
 *
 */

// Package g2config implements a client for the service.
package g2config

import (
	"context"
	"strconv"
	"time"

	g2configapi "github.com/senzing/g2-sdk-go/g2config"
	"github.com/senzing/go-logging/logging"
	"github.com/senzing/go-observing/notifier"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2config struct {
	isTrace               bool
	logger                logging.LoggingInterface
	observers             subject.Subject
	AddDataSourceResult   string
	CreateResult          uintptr
	ListDataSourcesResult string
	SaveResult            string
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// --- Logging ----------------------------------------------------------------

// Get the Logger singleton.
func (client *G2config) getLogger() logging.LoggingInterface {
	var err error = nil
	if client.logger == nil {
		options := []interface{}{
			&logging.OptionCallerSkip{Value: 4},
		}
		client.logger, err = logging.NewSenzingSdkLogger(ProductId, g2configapi.IdMessages, options...)
		if err != nil {
			panic(err)
		}
	}
	return client.logger
}

// Trace method entry.
func (client *G2config) traceEntry(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// Trace method exit.
func (client *G2config) traceExit(errorNumber int, details ...interface{}) {
	client.getLogger().Log(errorNumber, details...)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The AddDataSource method adds a data source to an existing in-memory configuration.
The configHandle is created by the Create() method.

Input
  - ctx: A context to control lifecycle.
  - configHandle: An identifier of an in-memory configuration.
  - inputJson: A JSON document in the format `{"DSRC_CODE": "NAME_OF_DATASOURCE"}`.

Output
  - A string containing a JSON document listing the newly created data source.
    See the example output.
*/
func (client *G2config) AddDataSource(ctx context.Context, configHandle uintptr, inputJson string) (string, error) {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(1, configHandle, inputJson)
		defer client.traceExit(2, configHandle, inputJson, client.AddDataSourceResult, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"inputJson": inputJson,
				"return":    client.AddDataSourceResult,
			}
			notifier.Notify(ctx, client.observers, ProductId, 8001, err, details)
		}()
	}
	return client.AddDataSourceResult, err
}

/*
The Close method cleans up the Senzing G2Config object pointed to by the handle.
The handle was created by the Create() method.

Input
  - ctx: A context to control lifecycle.
  - configHandle: An identifier of an in-memory configuration.
*/
func (client *G2config) Close(ctx context.Context, configHandle uintptr) error {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(5, configHandle)
		defer client.traceExit(6, configHandle, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, ProductId, 8002, err, details)
		}()
	}
	return err
}

/*
The Create method creates an in-memory Senzing configuration from the g2config.json
template configuration file located in the PIPELINE.RESOURCEPATH path.
A handle is returned to identify the in-memory configuration.
The handle is used by the AddDataSource(), ListDataSources(), DeleteDataSource(), Load(), and Save() methods.
The handle is terminated by the Close() method.

Input
  - ctx: A context to control lifecycle.

Output
  - A Pointer to an in-memory Senzing configuration.
*/
func (client *G2config) Create(ctx context.Context) (uintptr, error) {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(7)
		defer client.traceExit(8, client.CreateResult, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, ProductId, 8003, err, details)
		}()
	}
	return client.CreateResult, err
}

/*
The DeleteDataSource method removes a data source from an existing configuration.
The configHandle is created by the Create() method.

Input
  - ctx: A context to control lifecycle.
  - configHandle: An identifier of an in-memory configuration.
  - inputJson: A JSON document in the format `{"DSRC_CODE": "NAME_OF_DATASOURCE"}`.
*/
func (client *G2config) DeleteDataSource(ctx context.Context, configHandle uintptr, inputJson string) error {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(9, configHandle, inputJson)
		defer client.traceExit(10, configHandle, inputJson, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"inputJson": inputJson,
			}
			notifier.Notify(ctx, client.observers, ProductId, 8004, err, details)
		}()
	}
	return err
}

/*
The Destroy method will destroy and perform cleanup for the Senzing G2Config object.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2config) Destroy(ctx context.Context) error {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(11)
		defer client.traceExit(12, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, ProductId, 8005, err, details)
		}()
	}
	return err
}

/*
The GetSdkId method returns the identifier of this particular Software Development Kit (SDK).
It is handy when working with multiple implementations of the same G2configInterface.
For this implementation, "mock" is returned.

Input
  - ctx: A context to control lifecycle.
*/
func (client *G2config) GetSdkId(ctx context.Context) string {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(31)
		defer client.traceExit(32, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, ProductId, 8010, err, details)
		}()
	}
	return "mock"
}

/*
The Init method initializes the Senzing G2Config object.
It must be called prior to any other calls.

Input
  - ctx: A context to control lifecycle.
  - moduleName: A name for the auditing node, to help identify it within system logs.
  - iniParams: A JSON string containing configuration parameters.
  - verboseLogging: A flag to enable deeper logging of the G2 processing. 0 for no Senzing logging; 1 for logging.
*/
func (client *G2config) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(17, moduleName, iniParams, verboseLogging)
		defer client.traceExit(18, moduleName, iniParams, verboseLogging, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"iniParams":      iniParams,
				"moduleName":     moduleName,
				"verboseLogging": strconv.Itoa(verboseLogging),
			}
			notifier.Notify(ctx, client.observers, ProductId, 8006, err, details)
		}()
	}
	return err
}

/*
The ListDataSources method returns a JSON document of data sources.
The configHandle is created by the Create() method.

Input
  - ctx: A context to control lifecycle.
  - configHandle: An identifier of an in-memory configuration.

Output
  - A string containing a JSON document listing all of the data sources.
    See the example output.
*/
func (client *G2config) ListDataSources(ctx context.Context, configHandle uintptr) (string, error) {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(19, configHandle)
		defer client.traceExit(20, configHandle, client.ListDataSourcesResult, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, ProductId, 8007, err, details)
		}()
	}
	return client.ListDataSourcesResult, err
}

/*
The Load method initializes the Senzing G2Config object from a JSON string.
The configHandle is created by the Create() method.

Input
  - ctx: A context to control lifecycle.
  - configHandle: An identifier of an in-memory configuration.
  - jsonConfig: A JSON document containing the Senzing configuration.
*/
func (client *G2config) Load(ctx context.Context, configHandle uintptr, jsonConfig string) error {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(21, configHandle, jsonConfig)
		defer client.traceExit(22, configHandle, jsonConfig, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, ProductId, 8008, err, details)
		}()
	}
	return err
}

/*
The RegisterObserver method adds the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *G2config) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(27, observer.GetObserverId(ctx))
		defer client.traceExit(28, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	if client.observers == nil {
		client.observers = &subject.SubjectImpl{}
	}
	err = client.observers.RegisterObserver(ctx, observer)
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"observerID": observer.GetObserverId(ctx),
			}
			notifier.Notify(ctx, client.observers, ProductId, 8011, err, details)
		}()
	}
	return err
}

/*
The Save method creates a JSON string representation of the Senzing G2Config object.
The configHandle is created by the Create() method.

Input
  - ctx: A context to control lifecycle.
  - configHandle: An identifier of an in-memory configuration.

Output
  - A string containing a JSON Document representation of the Senzing G2Config object.
    See the example output.
*/
func (client *G2config) Save(ctx context.Context, configHandle uintptr) (string, error) {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(23, configHandle)
		defer client.traceExit(24, configHandle, client.SaveResult, err, time.Since(entryTime))
	}
	if client.observers != nil {
		go func() {
			details := map[string]string{}
			notifier.Notify(ctx, client.observers, ProductId, 8009, err, details)
		}()
	}
	return client.SaveResult, err
}

/*
The SetLogLevel method sets the level of logging.

Input
  - ctx: A context to control lifecycle.
  - logLevel: The desired log level. TRACE, DEBUG, INFO, WARN, ERROR, FATAL or PANIC.
*/
func (client *G2config) SetLogLevel(ctx context.Context, logLevelName string) error {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(25, logLevelName)
		defer client.traceExit(26, logLevelName, err, time.Since(entryTime))
	}
	client.getLogger().SetLogLevel(logLevelName)
	client.isTrace = (logLevelName == logging.LevelTraceName)
	if client.observers != nil {
		go func() {
			details := map[string]string{
				"logLevel": logLevelName,
			}
			notifier.Notify(ctx, client.observers, ProductId, 8012, err, details)
		}()
	}
	return err
}

/*
The UnregisterObserver method removes the observer to the list of observers notified.

Input
  - ctx: A context to control lifecycle.
  - observer: The observer to be added.
*/
func (client *G2config) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error = nil
	if client.isTrace {
		entryTime := time.Now()
		client.traceEntry(29, observer.GetObserverId(ctx))
		defer client.traceExit(30, observer.GetObserverId(ctx), err, time.Since(entryTime))
	}
	if client.observers != nil {
		// Tricky code:
		// client.notify is called synchronously before client.observers is set to nil.
		// In client.notify, each observer will get notified in a goroutine.
		// Then client.observers may be set to nil, but observer goroutines will be OK.
		details := map[string]string{
			"observerID": observer.GetObserverId(ctx),
		}
		notifier.Notify(ctx, client.observers, ProductId, 8013, err, details)
	}
	err = client.observers.UnregisterObserver(ctx, observer)
	if !client.observers.HasObservers(ctx) {
		client.observers = nil
	}
	return err
}
