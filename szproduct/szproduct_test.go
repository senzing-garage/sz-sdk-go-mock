package szproduct

import (
	"context"
	"fmt"
	"os"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing-garage/sz-sdk-go/sz"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
	printResults      = false
)

var (
	szProductSingleton *Szproduct
)

// ----------------------------------------------------------------------------
// Interface functions - test
// ----------------------------------------------------------------------------

func TestSzproduct_GetLicense(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(ctx, test)
	actual, err := szProduct.GetLicense(ctx)
	testError(test, err)
	printActual(test, actual)
}

func TestSzproduct_GetVersion(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(ctx, test)
	actual, err := szProduct.GetVersion(ctx)
	testError(test, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzproduct_SetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	szProduct.SetObserverOrigin(ctx, origin)
}

func TestSzproduct_GetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	szProduct.SetObserverOrigin(ctx, origin)
	actual := szProduct.GetObserverOrigin(ctx)
	assert.Equal(test, origin, actual)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzproduct_AsInterface(test *testing.T) {
	ctx := context.TODO()
	szProduct := getSzProductAsInterface(ctx)
	actual, err := szProduct.GetLicense(ctx)
	testError(test, err)
	printActual(test, actual)
}

func TestSzproduct_Initialize(test *testing.T) {
	ctx := context.TODO()
	szProduct := &Szproduct{}
	instanceName := "Test name"
	settings, err := getSettings()
	testError(test, err)
	verboseLogging := sz.SZ_NO_LOGGING
	err = szProduct.Initialize(ctx, instanceName, settings, verboseLogging)
	testError(test, err)
}

func TestSzproduct_Destroy(test *testing.T) {
	ctx := context.TODO()
	szProduct := getTestObject(ctx, test)
	err := szProduct.Destroy(ctx)
	testError(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getSettings() (string, error) {
	return "{}", nil
}

func getSzProduct(ctx context.Context) *Szproduct {
	_ = ctx
	if szProductSingleton == nil {
		szProductSingleton = &Szproduct{
			LicenseResult: `{"customer":"Senzing Public Test License","contract":"EVALUATION - support@senzing.com","issueDate":"2022-11-29","licenseType":"EVAL (Solely for non-productive use)","licenseLevel":"STANDARD","billing":"MONTHLY","expireDate":"2023-11-29","recordLimit":50000}`,
			VersionResult: `{"PRODUCT_NAME":"Senzing API","VERSION":"3.5.0","BUILD_VERSION":"3.5.0.23041","BUILD_DATE":"2023-02-09","BUILD_NUMBER":"2023_02_09__23_01","COMPATIBILITY_VERSION":{"CONFIG_VERSION":"10"},"SCHEMA_VERSION":{"ENGINE_SCHEMA_VERSION":"3.5","MINIMUM_REQUIRED_SCHEMA_VERSION":"3.0","MAXIMUM_REQUIRED_SCHEMA_VERSION":"3.99"}}`,
		}
	}
	return szProductSingleton
}

func getSzProductAsInterface(ctx context.Context) sz.SzProduct {
	return getSzProduct(ctx)
}

func getTestObject(ctx context.Context, test *testing.T) *Szproduct {
	_ = test
	return getSzProduct(ctx)
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

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
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
