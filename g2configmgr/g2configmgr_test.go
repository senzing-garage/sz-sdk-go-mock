package g2configmgr

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/g2-sdk-go/g2api"
	"github.com/senzing/go-logging/logger"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
	printResults      = false
)

var (
	g2configmgrSingleton *G2configmgr
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) *G2configmgr {
	return getG2Configmgr(ctx)
}

func getG2Configmgr(ctx context.Context) *G2configmgr {
	if g2configmgrSingleton == nil {
		g2configmgrSingleton = &G2configmgr{
			AddConfigResult:          1,
			GetConfigResult:          `{"G2_CONFIG":{"CFG_ATTR":[{"ATTR_ID":1001,"ATTR_CODE":"DATA_SOURCE","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"Yes","DEFAULT_VALUE":null,"ADVANCED":"Yes","INTERNAL":"No"},{"ATTR_ID":1002,"ATTR_CODE":"ROUTE_CODE","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"No","DEFAULT_VALUE":null,"ADVANCED":"Yes","INTERNAL":"No"},{"ATTR_ID":1003,"ATTR_CODE":"RECORD_ID","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"No","DEFAULT_VALUE":null,"ADVANCED":"No","INTERNAL":"No"},{"ATTR_ID":1004,"ATTR_CODE":"ENTITY_TYPE","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,`,
			GetConfigListResult:      `{"CONFIGS":[{"CONFIG_ID":41320074,"CONFIG_COMMENTS":"Example configuration","SYS_CREATE_DT":"2023-02-16 21:43:10.171"},{"CONFIG_ID":1111755672,"CONFIG_COMMENTS":"g2configmgr_test at 2023-02-16 21:43:10.154619801 +0000 UTC","SYS_CREATE_DT":"2023-02-16 21:43:10.159"},{"CONFIG_ID":3680541328,"CONFIG_COMMENTS":"Created by g2diagnostic_test at 2023-02-16 21:43:07.294747409 +0000 UTC","SYS_CREATE_DT":"2023-02-16 21:43:07.755"}]}`,
			GetDefaultConfigIDResult: 1,
		}
	}
	return g2configmgrSingleton
}

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}

func printResult(test *testing.T, title string, result interface{}) {
	if printResults {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result), defaultTruncation))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, ctx context.Context, g2configmgr g2api.G2configmgr, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
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

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestG2configmgr_AddConfig(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	now := time.Now()
	configStr := ``
	configComments := fmt.Sprintf("g2configmgr_test at %s", now.UTC())
	actual, err := g2configmgr.AddConfig(ctx, configStr, configComments)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgr_GetConfig(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	configID, err1 := g2configmgr.GetDefaultConfigID(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()")
	}
	actual, err := g2configmgr.GetConfig(ctx, configID)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgr_GetConfigList(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	actual, err := g2configmgr.GetConfigList(ctx)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgr_GetDefaultConfigID(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	actual, err := g2configmgr.GetDefaultConfigID(ctx)
	testError(test, ctx, g2configmgr, err)
	printActual(test, actual)
}

func TestG2configmgr_ReplaceDefaultConfigID(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	oldConfigID, err1 := g2configmgr.GetDefaultConfigID(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()")
	}

	// FIXME: This is kind of a cheeter.

	newConfigID, err2 := g2configmgr.GetDefaultConfigID(ctx)
	if err2 != nil {
		test.Log("Error:", err2.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()-2")
	}

	err := g2configmgr.ReplaceDefaultConfigID(ctx, oldConfigID, newConfigID)
	testError(test, ctx, g2configmgr, err)
}

func TestG2configmgr_SetDefaultConfigID(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	configID, err1 := g2configmgr.GetDefaultConfigID(ctx)
	if err1 != nil {
		test.Log("Error:", err1.Error())
		assert.FailNow(test, "g2configmgr.GetDefaultConfigID()")
	}
	err := g2configmgr.SetDefaultConfigID(ctx, configID)
	testError(test, ctx, g2configmgr, err)
}

func TestG2configmgr_Init(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams := "{}"
	err := g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2configmgr, err)
}

func TestG2configmgr_Destroy(test *testing.T) {
	ctx := context.TODO()
	g2configmgr := getTestObject(ctx, test)
	err := g2configmgr.Destroy(ctx)
	testError(test, ctx, g2configmgr, err)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleG2configmgr_AddConfig() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_test.go
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx)
	configStr := ``
	configComments := "Example configuration"
	configID, err := g2configmgr.AddConfig(ctx, configStr, configComments)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configID > 0) // Dummy output.
	// Output: true
}

func ExampleG2configmgr_GetConfig() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_test.go
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx)
	configID, err := g2configmgr.GetDefaultConfigID(ctx)
	if err != nil {
		fmt.Println(err)
	}
	configStr, err := g2configmgr.GetConfig(ctx, configID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(truncate(configStr, defaultTruncation))
	// Output: {"G2_CONFIG":{"CFG_ATTR":[{"ATTR_ID":1001,"ATTR_CODE":"DATA_SOURCE","ATTR...
}

func ExampleG2configmgr_GetConfigList() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_test.go
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx)
	jsonConfigList, err := g2configmgr.GetConfigList(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(truncate(jsonConfigList, 28))
	// Output: {"CONFIGS":[{"CONFIG_ID":...
}

func ExampleG2configmgr_GetDefaultConfigID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_test.go
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx)
	configID, err := g2configmgr.GetDefaultConfigID(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configID > 0) // Dummy output.
	// Output: true
}

func ExampleG2configmgr_ReplaceDefaultConfigID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_test.go
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx)
	oldConfigID := int64(1)
	newConfigID := int64(2)
	err := g2configmgr.ReplaceDefaultConfigID(ctx, oldConfigID, newConfigID)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2configmgr_SetDefaultConfigID() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_test.go
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx)
	configID, err := g2configmgr.GetDefaultConfigID(ctx) // For example purposes only. Normally would use output from GetConfigList()
	if err != nil {
		fmt.Println(err)
	}
	err = g2configmgr.SetDefaultConfigID(ctx, configID)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2configmgr_SetLogLevel() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_test.go
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx)
	err := g2configmgr.SetLogLevel(ctx, logger.LevelInfo)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2configmgr_Init() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_test.go
	ctx := context.TODO()
	g2configmgr := &G2configmgr{}
	moduleName := "Test module name"
	iniParams := "{}"
	verboseLogging := 0
	err := g2configmgr.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		// This should produce a "senzing-60124002" error.
	}
	// Output:
}

func ExampleG2configmgr_Destroy() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_test.go
	ctx := context.TODO()
	g2configmgr := getG2Configmgr(ctx)
	err := g2configmgr.Destroy(ctx)
	if err != nil {
		// This should produce a "senzing-60124001" error.
	}
	// Output:
}
