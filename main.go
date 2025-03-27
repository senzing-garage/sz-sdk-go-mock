package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/senzing-garage/go-helpers/truthset"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/senzing-garage/sz-sdk-go-mock/szconfig"
	"github.com/senzing-garage/sz-sdk-go-mock/szconfigmanager"
	"github.com/senzing-garage/sz-sdk-go-mock/szengine"
	"github.com/senzing-garage/sz-sdk-go-mock/szproduct"
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

var programName = "unknown"
var buildVersion = "0.0.0"
var buildIteration = "0"
var logger logging.Logging

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	var err error
	ctx := context.TODO()

	// Configure the "log" standard library.

	log.SetFlags(0)
	logger, err = getLogger(ctx)
	failOnError(5001, err)

	// Test logger.

	programmMetadataMap := map[string]interface{}{
		"ProgramName":    programName,
		"BuildVersion":   buildVersion,
		"BuildIteration": buildIteration,
	}

	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
	logger.Log(2001, "Just a test of logging", programmMetadataMap)

	// Get Senzing objects for installing a Senzing Engine configuration.

	szConfig, err := getSzConfig(ctx)
	failOnError(5002, err)

	szConfigManager, err := getSzConfigManager(ctx)
	failOnError(5003, err)

	// Persist the Senzing configuration to the Senzing repository.

	err = demonstrateConfigFunctions(ctx, szConfig, szConfigManager)
	failOnError(5004, err)

	// Now that a Senzing configuration is installed, get the remainder of the Senzing objects.

	szEngine, err := getSzEngine(ctx)
	failOnError(5005, err)

	szProduct, err := getSzProduct(ctx)
	failOnError(5006, err)

	// Demonstrate tests.

	err = demonstrateAdditionalFunctions(ctx, szEngine, szProduct)
	failOnError(5007, err)

	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func demonstrateAdditionalFunctions(ctx context.Context, szEngine senzing.SzEngine, szProduct senzing.SzProduct) error {

	// Using SzEngine: Add records with information returned.

	withInfo, err := demonstrateAddRecord(ctx, szEngine)
	failOnError(5101, err)
	logger.Log(2101, withInfo)

	// Using SzProduct: Show license metadata.

	license, err := szProduct.GetLicense(ctx)
	failOnError(5102, err)
	logger.Log(2102, license)

	return err
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
		`", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "SEAMAN", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`)
	var flags = senzing.SzNoFlags

	// Using SzEngine: Add record and return "withInfo".

	return szEngine.AddRecord(ctx, dataSourceCode, recordID, recordDefinition, flags)
}

func demonstrateConfigFunctions(ctx context.Context, szConfig senzing.SzConfig, szConfigManager senzing.SzConfigManager) error {
	now := time.Now()

	// Using SzConfig: Create a default configuration in memory.

	configHandle, err := szConfig.CreateConfig(ctx)
	failOnError(5301, err)

	// Using SzConfig: Add data source to in-memory configuration.

	for _, testDataSource := range truthset.TruthsetDataSources {
		_, err := szConfig.AddDataSource(ctx, configHandle, testDataSource.JSON)
		failOnError(5302, err)
	}

	// Using SzConfig: Persist configuration to a string.

	configStr, err := szConfig.ExportConfig(ctx, configHandle)
	failOnError(5303, err)

	// Using SzConfigManager: Persist configuration string to database.

	configComments := fmt.Sprintf("Created by main at %s", now.UTC())
	configID, err := szConfigManager.AddConfig(ctx, configStr, configComments)
	failOnError(5304, err)

	// Using SzConfigManager: Set new configuration as the default.

	err = szConfigManager.SetDefaultConfigID(ctx, configID)
	failOnError(5305, err)

	return err
}

func failOnError(msgID int, err error) {
	if err != nil {
		logger.Log(msgID, err)
		panic(err.Error())
	}
}

func getLogger(ctx context.Context) (logging.Logging, error) {
	_ = ctx
	logger, err := logging.NewSenzingLogger(9999, Messages)
	failOnError(5401, err)
	return logger, err
}

func getSzConfig(ctx context.Context) (senzing.SzConfig, error) {
	var err error
	_ = ctx
	result := szconfig.Szconfig{}
	return &result, err
}

func getSzConfigManager(ctx context.Context) (senzing.SzConfigManager, error) {
	var err error
	_ = ctx
	result := szconfigmanager.Szconfigmanager{}
	return &result, err
}

func getSzEngine(ctx context.Context) (senzing.SzEngine, error) {
	var err error
	_ = ctx
	result := szengine.Szengine{}
	return &result, err
}

func getSzProduct(ctx context.Context) (senzing.SzProduct, error) {
	var err error
	_ = ctx
	result := szproduct.Szproduct{}
	return &result, err
}
