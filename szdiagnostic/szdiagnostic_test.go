package szdiagnostic

import (
	"context"
	"fmt"
	"os"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/sz-sdk-go/sz"
	"github.com/senzing-garage/sz-sdk-go/szerror"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
	printResults      = false
)

var (
	szDiagnosticSingleton *Szdiagnostic
)

// ----------------------------------------------------------------------------
// Interface functions - test
// ----------------------------------------------------------------------------

func TestSzdiagnostic_CheckDatastorePerformance(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := getTestObject(ctx, test)
	secondsToRun := 1
	actual, err := szDiagnostic.CheckDatastorePerformance(ctx, secondsToRun)
	testError(test, err)
	printActual(test, actual)
}

func TestSzdiagnostic_GetDatastoreInfo(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := getTestObject(ctx, test)
	actual, err := szDiagnostic.GetDatastoreInfo(ctx)
	testError(test, err)
	printActual(test, actual)
}

func TestSzdiagnostic_GetFeature(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := getTestObject(ctx, test)
	featureId := int64(1)
	actual, err := szDiagnostic.GetFeature(ctx, featureId)
	testError(test, err)
	printActual(test, actual)
}

func TestSzdiagnostic_PurgeRepository(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := getTestObject(ctx, test)
	err := szDiagnostic.PurgeRepository(ctx)
	testError(test, err)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzdiagnostic_SetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	szDiagnostic.SetObserverOrigin(ctx, origin)
}

func TestSzdiagnostic_GetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	szDiagnostic.SetObserverOrigin(ctx, origin)
	actual := szDiagnostic.GetObserverOrigin(ctx)
	assert.Equal(test, origin, actual)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzdiagnostic_AsInterface(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := getSzDiagnosticAsInterface(ctx)
	secondsToRun := 1
	actual, err := szDiagnostic.CheckDatastorePerformance(ctx, secondsToRun)
	testError(test, err)
	printActual(test, actual)
}

func TestSzdiagnostic_Initialize(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := &Szdiagnostic{}
	instanceName := "Test name"
	settings, err := getSettings()
	testError(test, err)
	verboseLogging := sz.SZ_NO_LOGGING
	configId := sz.SZ_INITIALIZE_WITH_DEFAULT_CONFIGURATION
	err = szDiagnostic.Initialize(ctx, instanceName, settings, configId, verboseLogging)
	testError(test, err)
}

func TestSzdiagnostic_Initialize_withConfigId(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := &Szdiagnostic{}
	instanceName := "Test name"
	settings, err := getSettings()
	testError(test, err)
	verboseLogging := sz.SZ_NO_LOGGING
	configId := getDefaultConfigId()
	err = szDiagnostic.Initialize(ctx, instanceName, settings, configId, verboseLogging)
	testError(test, err)
}

func TestSzdiagnostic_Reinitialize(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := getTestObject(ctx, test)
	configId := getDefaultConfigId()
	err := szDiagnostic.Reinitialize(ctx, configId)
	testErrorNoFail(test, err)
}

func TestSzdiagnostic_Destroy(test *testing.T) {
	ctx := context.TODO()
	szDiagnostic := getTestObject(ctx, test)
	err := szDiagnostic.Destroy(ctx)
	testError(test, err)
	szDiagnosticSingleton = nil
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getDefaultConfigId() int64 {
	return int64(1)
}

func getSettings() (string, error) {
	return "{}", nil
}

func getSzDiagnostic(ctx context.Context) *Szdiagnostic {
	_ = ctx
	if szDiagnosticSingleton == nil {
		szDiagnosticSingleton = &Szdiagnostic{
			CheckDatastorePerformanceResult: `{"numRecordsInserted":76667,"insertTime":1000}`,
			GetFeatureResult:                `{}`,
			GetDatastoreInfoResult:          `{}`,
		}
	}
	return szDiagnosticSingleton
}

func getSzDiagnosticAsInterface(ctx context.Context) sz.SzDiagnostic {
	return getSzDiagnostic(ctx)
}

func getTestObject(ctx context.Context, test *testing.T) *Szdiagnostic {
	_ = test
	return getSzDiagnostic(ctx)
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
