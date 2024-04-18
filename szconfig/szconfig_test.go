package szconfig

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
	szConfigSingleton *Szconfig
)

// ----------------------------------------------------------------------------
// Interface functions - test
// ----------------------------------------------------------------------------

func TestSzconfig_AddDataSource(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	dataSourceCode := "GO_TEST"
	actual, err := szConfig.AddDataSource(ctx, configHandle, dataSourceCode)
	testError(test, err)
	printActual(test, actual)
	err = szConfig.CloseConfig(ctx, configHandle)
	testError(test, err)
}

func TestSzconfig_AddDataSource_withLoad(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	configDefinition, err := szConfig.ExportConfig(ctx, configHandle)
	testError(test, err)
	err = szConfig.CloseConfig(ctx, configHandle)
	testError(test, err)
	configHandle2, err := szConfig.ImportConfig(ctx, configDefinition)
	testError(test, err)
	dataSourceCode := "GO_TEST"
	actual, err := szConfig.AddDataSource(ctx, configHandle2, dataSourceCode)
	testError(test, err)
	printActual(test, actual)
	err = szConfig.CloseConfig(ctx, configHandle2)
	testError(test, err)
}

func TestSzconfig_CloseConfig(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	err = szConfig.CloseConfig(ctx, configHandle)
	testError(test, err)
}

func TestSzconfig_CreateConfig(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	actual, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	printActual(test, actual)
}

func TestSzconfig_DeleteDataSource(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	actual, err := szConfig.GetDataSources(ctx, configHandle)
	testError(test, err)
	printResult(test, "Original", actual)
	dataSourceCode := "GO_TEST"
	_, err = szConfig.AddDataSource(ctx, configHandle, dataSourceCode)
	testError(test, err)
	actual, err = szConfig.GetDataSources(ctx, configHandle)
	testError(test, err)
	printResult(test, "     Add", actual)
	err = szConfig.DeleteDataSource(ctx, configHandle, dataSourceCode)
	testError(test, err)
	actual, err = szConfig.GetDataSources(ctx, configHandle)
	testError(test, err)
	printResult(test, "  Delete", actual)
	err = szConfig.CloseConfig(ctx, configHandle)
	testError(test, err)
}

func TestSzconfig_DeleteDataSource_withLoad(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	actual, err := szConfig.GetDataSources(ctx, configHandle)
	testError(test, err)
	printResult(test, "Original", actual)
	dataSourceCode := "GO_TEST"
	_, err = szConfig.AddDataSource(ctx, configHandle, dataSourceCode)
	testError(test, err)
	actual, err = szConfig.GetDataSources(ctx, configHandle)
	testError(test, err)
	printResult(test, "     Add", actual)
	configDefinition, err := szConfig.ExportConfig(ctx, configHandle)
	testError(test, err)
	err = szConfig.CloseConfig(ctx, configHandle)
	testError(test, err)
	configHandle2, err := szConfig.ImportConfig(ctx, configDefinition)
	testError(test, err)
	err = szConfig.DeleteDataSource(ctx, configHandle2, dataSourceCode)
	testError(test, err)
	actual, err = szConfig.GetDataSources(ctx, configHandle2)
	testError(test, err)
	printResult(test, "  Delete", actual)
	err = szConfig.CloseConfig(ctx, configHandle2)
	testError(test, err)
}

func TestSzconfig_ExportConfig(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	actual, err := szConfig.ExportConfig(ctx, configHandle)
	testError(test, err)
	printActual(test, actual)
}

func TestSzconfig_GetDataSources(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	actual, err := szConfig.GetDataSources(ctx, configHandle)
	testError(test, err)
	printActual(test, actual)
	err = szConfig.CloseConfig(ctx, configHandle)
	testError(test, err)
}

func TestSzconfig_ImportConfig(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	configDefinition, err := szConfig.ExportConfig(ctx, configHandle)
	testError(test, err)
	actual, err := szConfig.ImportConfig(ctx, configDefinition)
	testError(test, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func TestSzconfig_SetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	szConfig.SetObserverOrigin(ctx, origin)
}

func TestSzconfig_GetObserverOrigin(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	origin := "Machine: nn; Task: UnitTest"
	szConfig.SetObserverOrigin(ctx, origin)
	actual := szConfig.GetObserverOrigin(ctx)
	assert.Equal(test, origin, actual)
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func TestSzconfig_AsInterface(test *testing.T) {
	ctx := context.TODO()
	szConfig := getSzConfigAsInterface(ctx)
	configHandle, err := szConfig.CreateConfig(ctx)
	testError(test, err)
	actual, err := szConfig.GetDataSources(ctx, configHandle)
	testError(test, err)
	printActual(test, actual)
	err = szConfig.CloseConfig(ctx, configHandle)
	testError(test, err)
}

func TestSzconfig_Initialize(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	instanceName := "Test name"
	verboseLogging := sz.SZ_NO_LOGGING
	settings, err := getSettings()
	testError(test, err)
	err = szConfig.Initialize(ctx, instanceName, settings, verboseLogging)
	testError(test, err)
}

func TestSzconfig_Destroy(test *testing.T) {
	ctx := context.TODO()
	szConfig := getTestObject(ctx, test)
	err := szConfig.Destroy(ctx)
	testError(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getSettings() (string, error) {
	return "{}", nil
}

func getSzConfig(ctx context.Context) *Szconfig {
	_ = ctx
	if szConfigSingleton == nil {
		szConfigSingleton = &Szconfig{
			AddDataSourceResult:  `{"DSRC_ID":1001}`,
			CreateResult:         1,
			GetDataSourcesResult: `{"DATA_SOURCES":[{"DSRC_ID":1,"DSRC_CODE":"TEST"},{"DSRC_ID":2,"DSRC_CODE":"SEARCH"}]}`,
			SaveResult:           `{"G2_CONFIG":{"CFG_ATTR":[{"ATTR_ID":1001,"ATTR_CODE":"DATA_SOURCE","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"Yes","DEFAULT_VALUE":null,"ADVANCED":"Yes","INTERNAL":"No"},{"ATTR_ID":1002,"ATTR_CODE":"ROUTE_CODE",`,
		}
	}
	return szConfigSingleton
}

func getSzConfigAsInterface(ctx context.Context) sz.SzConfig {
	return getSzConfig(ctx)
}

func getTestObject(ctx context.Context, test *testing.T) *Szconfig {
	_ = test
	return getSzConfig(ctx)
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
