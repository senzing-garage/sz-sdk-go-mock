//go:build linux

package szabstractfactory

import (
	"context"
	"fmt"

	"github.com/senzing-garage/sz-sdk-go/senzing"
)

// ----------------------------------------------------------------------------
// Interface methods - Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSzabstractfactory_CreateSzConfig() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szabstractfactory/szabstractfactory_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactoryExample(ctx)
	szConfig, err := szAbstractFactory.CreateSzConfig(ctx)
	if err != nil {
		fmt.Println(err)
	}
	_ = szConfig // szConfig can now be used.
	// Output:
}

func ExampleSzabstractfactory_CreateSzConfigManager() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szabstractfactory/szabstractfactory_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactoryExample(ctx)
	szConfigManager, err := szAbstractFactory.CreateSzConfigManager(ctx)
	if err != nil {
		fmt.Println(err)
	}
	_ = szConfigManager // szConfigManager can now be used.
	// Output:
}

func ExampleSzabstractfactory_CreateSzDiagnostic() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szabstractfactory/szabstractfactory_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactoryExample(ctx)
	szDiagnostic, err := szAbstractFactory.CreateSzDiagnostic(ctx)
	if err != nil {
		fmt.Println(err)
	}
	_ = szDiagnostic // szDiagnostic can now be used.
	// Output:
}

func ExampleSzabstractfactory_CreateSzEngine() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szabstractfactory/szabstractfactory_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactoryExample(ctx)
	szEngine, err := szAbstractFactory.CreateSzEngine(ctx)
	if err != nil {
		fmt.Println(err)
	}
	_ = szEngine // szEngine can now be used.
	// Output:
}

func ExampleSzabstractfactory_CreateSzProduct() {
	// For more information, visit https://github.com/senzing-garage/sz-sdk-go-mock/blob/main/szabstractfactory/szabstractfactory_examples_test.go
	ctx := context.TODO()
	szAbstractFactory := getSzAbstractFactoryExample(ctx)
	szProduct, err := szAbstractFactory.CreateSzProduct(ctx)
	if err != nil {
		fmt.Println(err)
	}
	_ = szProduct // szProduct can now be used.
	// Output:
}

// ----------------------------------------------------------------------------
// Helper functions
// ----------------------------------------------------------------------------

func getSzAbstractFactoryExample(ctx context.Context) senzing.SzAbstractFactory {
	result, err := getSzAbstractFactory(ctx)
	if err != nil {
		panic(err)
	}
	return result
}
