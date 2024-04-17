//go:build linux

package szconfig

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-logging/logging"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSzconfig_AddDataSource() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	configHandle, err := szConfig.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	inputJson := `{"DSRC_CODE": "GO_TEST"}`
	result, err := szConfig.AddDataSource(ctx, configHandle, inputJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"DSRC_ID":1001}
}

func ExampleSzconfig_Close() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	configHandle, err := szConfig.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	err = szConfig.Close(ctx, configHandle)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSzconfig_Create() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	configHandle, err := szConfig.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configHandle > 0) // Dummy output.
	// Output: true
}

func ExampleSzconfig_DeleteDataSource() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	configHandle, err := szConfig.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	inputJson := `{"DSRC_CODE": "TEST"}`
	err = szConfig.DeleteDataSource(ctx, configHandle, inputJson)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSzconfig_ListDataSources() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	configHandle, err := szConfig.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	result, err := szConfig.ListDataSources(ctx, configHandle)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"DATA_SOURCES":[{"DSRC_ID":1,"DSRC_CODE":"TEST"},{"DSRC_ID":2,"DSRC_CODE":"SEARCH"}]}
}

func ExampleSzconfig_Load() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	mockConfigHandle, err := szConfig.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	jsonConfig, err := szConfig.Save(ctx, mockConfigHandle)
	if err != nil {
		fmt.Println(err)
	}
	configHandle, err := szConfig.Load(ctx, jsonConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configHandle == 0) // Dummy output.
	// Output: true
}

func ExampleSzconfig_Save() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	configHandle, err := szConfig.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	jsonConfig, err := szConfig.Save(ctx, configHandle)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(truncate(jsonConfig, 207))
	// Output: {"G2_CONFIG":{"CFG_ATTR":[{"ATTR_ID":1001,"ATTR_CODE":"DATA_SOURCE","ATTR_CLASS":"OBSERVATION","FTYPE_CODE":null,"FELEM_CODE":null,"FELEM_REQ":"Yes","DEFAULT_VALUE":null,"ADVANCED":"Yes","INTERNAL":"No"},...
}

func ExampleSzconfig_SetLogLevel() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	err := szConfig.SetLogLevel(ctx, logging.LevelInfoName)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSzconfig_Init() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	moduleName := "Test module name"
	iniParams := "{}"
	verboseLogging := int64(0)
	err := szConfig.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		// This should produce a "senzing-60114002" error.
	}
	// Output:
}

func ExampleSzconfig_Destroy() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	err := szConfig.Destroy(ctx)
	if err != nil {
		// This should produce a "senzing-60114001" error.
	}
	// Output:
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func ExampleSzconfig_SetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szConfig.SetObserverOrigin(ctx, origin)
	// Output:
}

func ExampleSzconfig_GetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2config_examples_test.go
	ctx := context.TODO()
	szConfig := getSzConfig(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szConfig.SetObserverOrigin(ctx, origin)
	result := szConfig.GetObserverOrigin(ctx)
	fmt.Println(result)
	// Output: Machine: nn; Task: UnitTest
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------
