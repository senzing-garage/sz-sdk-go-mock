//go:build linux

package szconfigmanager

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-logging/logging"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSzconfigmanager_SetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szConfigManager.SetObserverOrigin(ctx, origin)
	// Output:
}

func ExampleSzconfigmanager_GetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2configmgr_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	origin := "Machine: nn; Task: UnitTest"
	szConfigManager.SetObserverOrigin(ctx, origin)
	result := szConfigManager.GetObserverOrigin(ctx)
	fmt.Println(result)
	// Output: Machine: nn; Task: UnitTest
}

func ExampleSzconfigmanager_AddConfig() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	configStr := ``
	configComments := "Example configuration"
	configID, err := szConfigManager.AddConfig(ctx, configStr, configComments)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configID > 0) // Dummy output.
	// Output: true
}

func ExampleSzconfigmanager_GetConfig() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	configID, err := szConfigManager.GetDefaultConfigID(ctx)
	if err != nil {
		fmt.Println(err)
	}
	configStr, err := szConfigManager.GetConfig(ctx, configID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(truncate(configStr, defaultTruncation))
	// Output: {"G2_CONFIG":{"CFG_ATTR":[{"ATTR_ID":1001,"ATTR_CODE":"DATA_SOURCE","ATTR...
}

func ExampleSzconfigmanager_GetConfigList() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	jsonConfigList, err := szConfigManager.GetConfigList(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(truncate(jsonConfigList, 28))
	// Output: {"CONFIGS":[{"CONFIG_ID":...
}

func ExampleSzconfigmanager_GetDefaultConfigID() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	configID, err := szConfigManager.GetDefaultConfigID(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configID > 0) // Dummy output.
	// Output: true
}

func ExampleSzconfigmanager_ReplaceDefaultConfigID() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	oldConfigID := int64(1)
	newConfigID := int64(2)
	err := szConfigManager.ReplaceDefaultConfigID(ctx, oldConfigID, newConfigID)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSzconfigmanager_SetDefaultConfigID() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	configID, err := szConfigManager.GetDefaultConfigID(ctx) // For example purposes only. Normally would use output from GetConfigList()
	if err != nil {
		fmt.Println(err)
	}
	err = szConfigManager.SetDefaultConfigID(ctx, configID)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSzconfigmanager_SetLogLevel() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	err := szConfigManager.SetLogLevel(ctx, logging.LevelInfoName)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSzconfigmanager_Init() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := &Szconfigmanager{}
	moduleName := "Test module name"
	iniParams := "{}"
	verboseLogging := int64(0)
	err := szConfigManager.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		// This should produce a "senzing-60124002" error.
	}
	// Output:
}

func ExampleSzconfigmanager_Destroy() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2configmgr/g2configmgr_examples_test.go
	ctx := context.TODO()
	szConfigManager := getSzConfigManager(ctx)
	err := szConfigManager.Destroy(ctx)
	if err != nil {
		// This should produce a "senzing-60124001" error.
	}
	// Output:
}
