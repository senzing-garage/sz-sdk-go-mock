//go:build linux

package szproduct

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-logging/logging"
)

// ----------------------------------------------------------------------------
// Interface functions - Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSzproduct_GetLicense() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getSzProduct(ctx)
	result, err := g2product.GetLicense(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"customer":"Senzing Public Test License","contract":"EVALUATION - support@senzing.com","issueDate":"2022-11-29","licenseType":"EVAL (Solely for non-productive use)","licenseLevel":"STANDARD","billing":"MONTHLY","expireDate":"2023-11-29","recordLimit":50000}
}

func ExampleSzproduct_GetVersion() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getSzProduct(ctx)
	result, err := g2product.GetVersion(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(truncate(result, 43))
	// Output: {"PRODUCT_NAME":"Senzing API","VERSION":...
}

// ----------------------------------------------------------------------------
// Logging and observing
// ----------------------------------------------------------------------------

func ExampleSzproduct_SetLogLevel() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getSzProduct(ctx)
	err := g2product.SetLogLevel(ctx, logging.LevelInfoName)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleSzproduct_SetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getSzProduct(ctx)
	origin := "Machine: nn; Task: UnitTest"
	g2product.SetObserverOrigin(ctx, origin)
	// Output:
}

func ExampleSzproduct_GetObserverOrigin() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2config/g2product_test.go
	ctx := context.TODO()
	g2product := getSzProduct(ctx)
	origin := "Machine: nn; Task: UnitTest"
	g2product.SetObserverOrigin(ctx, origin)
	result := g2product.GetObserverOrigin(ctx)
	fmt.Println(result)
	// Output: Machine: nn; Task: UnitTest
}

// ----------------------------------------------------------------------------
// Object creation / destruction
// ----------------------------------------------------------------------------

func ExampleSzproduct_Init() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getSzProduct(ctx)
	moduleName := "Test module name"
	iniParams := "{}"
	verboseLogging := int64(0)
	err := g2product.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		// This should produce a "senzing-60164002" error.
	}
	// Output:
}

func ExampleSzproduct_Destroy() {
	// For more information, visit https://github.com/senzing-garage/g2-sdk-go-mock/blob/main/g2product/g2product_examples_test.go
	ctx := context.TODO()
	g2product := getSzProduct(ctx)
	err := g2product.Destroy(ctx)
	if err != nil {
		// This should produce a "senzing-60164001" error.
	}
	// Output:
}
