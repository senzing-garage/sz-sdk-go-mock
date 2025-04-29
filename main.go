package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/senzing-garage/go-helpers/truthset"
	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/sz-sdk-go-mock/szabstractfactory"
	"github.com/senzing-garage/sz-sdk-go/senzing"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIDTemplate = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var Messages = map[int]string{
	1:    "%s",
	2:    "WithInfo: %s",
	2001: "Testing %s.",
	2002: "Physical cores: %d.",
	2003: "withInfo",
	2004: "License",
	2999: "Cannot retrieve last error message.",
}

// Values updated via "go install -ldflags" parameters.

var (
	programName    = "unknown"
	buildVersion   = "0.0.0"
	buildIteration = "0"
	logger         logging.Logging
)

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	var err error

	ctx := context.TODO()

	// Configure the "log" standard library.

	log.SetFlags(0)

	logger = getLogger(ctx)

	// Test logger.

	programmMetadataMap := map[string]interface{}{
		"ProgramName":    programName,
		"BuildVersion":   buildVersion,
		"BuildIteration": buildIteration,
	}

	// Create a SzAbstractFactory.

	szAbstractFactory := &szabstractfactory.Szabstractfactory{}

	outputf("\n-------------------------------------------------------------------------------\n\n")
	logger.Log(2001, "Just a test of logging", programmMetadataMap)

	// Demonstrate persisting a Senzing configuration to the Senzing repository.

	demonstrateConfigFunctions(ctx, szAbstractFactory)

	// Demonstrate tests.

	demonstrateAdditionalFunctions(ctx, szAbstractFactory)

	err = szAbstractFactory.Destroy(ctx)
	failOnError(5008, err)

	outputf("\n-------------------------------------------------------------------------------\n\n")
}

// ----------------------------------------------------------------------------
// Demonstrations
// ----------------------------------------------------------------------------

func demonstrateAdditionalFunctions(ctx context.Context, szAbstractFactory senzing.SzAbstractFactory) {
	// Using SzEngine: Add records with information returned.
	szEngine, err := szAbstractFactory.CreateEngine(ctx)
	failOnError(5100, err)

	withInfo, err := demonstrateAddRecord(ctx, szEngine)
	failOnError(5101, err)
	logger.Log(2101, withInfo)

	// Using SzProduct: Show license metadata.

	szProduct, err := szAbstractFactory.CreateProduct(ctx)
	failOnError(5102, err)

	license, err := szProduct.GetLicense(ctx)
	failOnError(5103, err)
	logger.Log(2102, license)
}

func demonstrateAddRecord(ctx context.Context, szEngine senzing.SzEngine) (string, error) {
	dataSourceCode := "TEST"
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000000000))
	failOnError(5201, err)

	recordID := randomNumber.String()
	recordDefinition := fmt.Sprintf(
		"%s%s%s",
		`{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "`,
		recordID,
		`", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
	)

	flags := senzing.SzNoFlags

	// Using SzEngine: Add record and return "withInfo".

	result, err := szEngine.AddRecord(ctx, dataSourceCode, recordID, recordDefinition, flags)

	return result, wraperror.Errorf(err, "demonstrateAddRecord.AddRecord error: %w", err)
}

func demonstrateConfigFunctions(ctx context.Context, szAbstractFactory senzing.SzAbstractFactory) {
	now := time.Now()

	// Create Senzing objects.

	szConfigManager, err := szAbstractFactory.CreateConfigManager(ctx)
	failOnError(5101, err)

	szConfig, err := szConfigManager.CreateConfigFromTemplate(ctx)
	failOnError(5102, err)

	// Using SzConfig: Add data source to in-memory configuration.

	for testDataSourceCode := range truthset.TruthsetDataSources {
		_, err := szConfig.AddDataSource(ctx, testDataSourceCode)
		failOnError(5104, err)
	}

	// Using SzConfig: Persist configuration to a string.

	configStr, err := szConfig.Export(ctx)
	failOnError(5105, err)

	// Using SzConfigManager: Persist configuration string to database.

	configComment := fmt.Sprintf("Created by main.go at %s", now.UTC())
	_, err = szConfigManager.SetDefaultConfig(ctx, configStr, configComment)
	failOnError(5106, err)
}

func failOnError(msgID int, err error) {
	if err != nil {
		logger.Log(msgID, err)
		panic(err.Error())
	}
}

func getLogger(ctx context.Context) logging.Logging {
	_ = ctx
	logger, err := logging.NewSenzingLogger(9999, Messages)
	failOnError(5401, err)

	return logger
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func outputf(format string, message ...any) {
	fmt.Printf(format, message...) //nolint
}
