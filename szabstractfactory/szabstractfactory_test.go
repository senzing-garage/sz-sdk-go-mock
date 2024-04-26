package szabstractfactory

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
	instanceName      = "SzAbstractFactory Test"
	printResults      = false
	verboseLogging    = sz.SZ_NO_LOGGING
)

// ----------------------------------------------------------------------------
// Interface functions - test
// ----------------------------------------------------------------------------

func TestSzAbstractFactory_CreateSzConfig(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	szConfig, err := szAbstractFactory.CreateSzConfig(ctx)
	testError(test, ctx, szAbstractFactory, err)
	defer szConfig.Destroy(ctx)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, ctx, szAbstractFactory, err)
	dataSources, err := szConfig.GetDataSources(ctx, configHandle)
	testError(test, ctx, szAbstractFactory, err)
	printActual(test, dataSources)
}

func TestSzAbstractFactory_CreateSzConfigManager(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	szConfigManager, err := szAbstractFactory.CreateSzConfigManager(ctx)
	testError(test, ctx, szAbstractFactory, err)
	defer szConfigManager.Destroy(ctx)
	configList, err := szConfigManager.GetConfigList(ctx)
	testError(test, ctx, szAbstractFactory, err)
	printActual(test, configList)
}

func TestSzAbstractFactory_CreateSzDiagnostic(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	szDiagnostic, err := szAbstractFactory.CreateSzDiagnostic(ctx)
	testError(test, ctx, szAbstractFactory, err)
	defer szDiagnostic.Destroy(ctx)
	result, err := szDiagnostic.CheckDatastorePerformance(ctx, 1)
	testError(test, ctx, szAbstractFactory, err)
	printActual(test, result)
}

func TestSzAbstractFactory_CreateSzEngine(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	testError(test, ctx, szAbstractFactory, err)
	defer szEngine.Destroy(ctx)
	stats, err := szEngine.GetStats(ctx)
	testError(test, ctx, szAbstractFactory, err)
	printActual(test, stats)
}

func TestSzAbstractFactory_CreateSzProduct(test *testing.T) {
	ctx := context.TODO()
	szAbstractFactory := getTestObject(ctx, test)
	szProduct, err := szAbstractFactory.CreateSzProduct(ctx)
	testError(test, ctx, szAbstractFactory, err)
	defer szProduct.Destroy(ctx)
	version, err := szProduct.GetVersion(ctx)
	testError(test, ctx, szAbstractFactory, err)
	printActual(test, version)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getSzAbstractFactory(ctx context.Context) sz.SzAbstractFactory {
	_ = ctx
	result := &Szabstractfactory{}
	return result
}

func getTestObject(ctx context.Context, test *testing.T) sz.SzAbstractFactory {
	_ = test
	return getSzAbstractFactory(ctx)
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func printResult(test *testing.T, title string, result interface{}) {
	if printResults {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result), defaultTruncation))
	}
}

func testError(test *testing.T, ctx context.Context, szAbstractFactory sz.SzAbstractFactory, err error) {
	_ = ctx
	_ = szAbstractFactory
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
	return nil
}
