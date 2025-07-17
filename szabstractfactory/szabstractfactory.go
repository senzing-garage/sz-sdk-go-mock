package szabstractfactory

import (
	"context"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/sz-sdk-go-mock/szconfigmanager"
	"github.com/senzing-garage/sz-sdk-go-mock/szdiagnostic"
	"github.com/senzing-garage/sz-sdk-go-mock/szengine"
	"github.com/senzing-garage/sz-sdk-go-mock/szproduct"
	"github.com/senzing-garage/sz-sdk-go/senzing"
)

/*
Szabstractfactory is an implementation of the [senzing.SzAbstractFactory] interface.

[senzing.SzAbstractFactory]: https://pkg.go.dev/github.com/senzing-garage/sz-sdk-go/senzing#SzAbstractFactory
*/
type Szabstractfactory struct {
	AddConfigResult                         int64
	AddRecordResult                         string
	CheckRepositoryPerformanceResult        string
	CountRedoRecordsResult                  int64
	CreateConfigResult                      uintptr
	DeleteRecordResult                      string
	ExportConfigResult                      string
	ExportCsvEntityReportResult             uintptr
	ExportJSONEntityReportResult            uintptr
	FetchNextResult                         string
	FindInterestingEntitiesByEntityIDResult string
	FindInterestingEntitiesByRecordIDResult string
	FindNetworkByEntityIDResult             string
	FindNetworkByRecordIDResult             string
	FindPathByEntityIDResult                string
	FindPathByRecordIDResult                string
	GetActiveConfigIDResult                 int64
	GetConfigRegistryResult                 string
	GetConfigResult                         string
	GetDataSourceRegistryResult             string
	GetDefaultConfigIDResult                int64
	GetEntityByEntityIDResult               string
	GetEntityByRecordIDResult               string
	GetFeatureResult                        string
	GetLicenseResult                        string
	GetRecordPreviewResult                  string
	GetRecordResult                         string
	GetRedoRecordResult                     string
	GetRepositoryInfoResult                 string
	GetStatsResult                          string
	GetVersionResult                        string
	GetVirtualEntityByRecordIDResult        string
	HowEntityByEntityIDResult               string
	ImportConfigResult                      uintptr
	ProcessRedoRecordResult                 string
	ReevaluateEntityResult                  string
	ReevaluateRecordResult                  string
	RegisterDataSourceResult                string
	SearchByAttributesResult                string
	UnregisterDataSourceResult              string
	WhyEntitiesResult                       string
	WhyRecordInEntityResult                 string
	WhyRecordsResult                        string
}

// ----------------------------------------------------------------------------
// senzing.SzAbstractFactory interface methods
// ----------------------------------------------------------------------------

/*
Method Close will destroy and perform cleanup for the Senzing objects created by the AbstractFactory.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
func (factory *Szabstractfactory) Close(ctx context.Context) error {
	var err error

	_ = ctx

	return err
}

/*
Method CreateConfigManager returns an SzConfigManager object
implemented to use the Senzing native C binary, libSz.so.

Input
  - ctx: A context to control lifecycle.

Output
  - An SzConfigManager object.
*/
func (factory *Szabstractfactory) CreateConfigManager(ctx context.Context) (senzing.SzConfigManager, error) {
	var err error

	_ = ctx
	result := &szconfigmanager.Szconfigmanager{
		RegisterConfigResult:     factory.AddConfigResult,
		GetConfigResult:          factory.GetConfigResult,
		GetConfigRegistryResult:  factory.GetConfigRegistryResult,
		GetDefaultConfigIDResult: factory.GetDefaultConfigIDResult,
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method CreateDiagnostic returns an SzDiagnostic object
implemented to use the Senzing native C binary, libSz.so.

Input
  - ctx: A context to control lifecycle.

Output
  - An SzDiagnostic object.
*/
func (factory *Szabstractfactory) CreateDiagnostic(ctx context.Context) (senzing.SzDiagnostic, error) {
	var err error

	_ = ctx
	result := &szdiagnostic.Szdiagnostic{
		CheckRepositoryPerformanceResult: factory.CheckRepositoryPerformanceResult,
		GetRepositoryInfoResult:          factory.GetRepositoryInfoResult,
		GetFeatureResult:                 factory.GetFeatureResult,
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method CreateEngine returns an SzEngine object
implemented to use the Senzing native C binary, libSz.so.

Input
  - ctx: A context to control lifecycle.

Output
  - An SzEngine object.
*/
func (factory *Szabstractfactory) CreateEngine(ctx context.Context) (senzing.SzEngine, error) {
	var err error

	_ = ctx
	result := &szengine.Szengine{
		AddRecordResult:                         factory.AddRecordResult,
		CountRedoRecordsResult:                  factory.CountRedoRecordsResult,
		DeleteRecordResult:                      factory.DeleteRecordResult,
		ExportConfigResult:                      factory.ExportConfigResult,
		ExportCsvEntityReportResult:             factory.ExportCsvEntityReportResult,
		ExportJSONEntityReportResult:            factory.ExportJSONEntityReportResult,
		FetchNextResult:                         factory.FetchNextResult,
		FindInterestingEntitiesByEntityIDResult: factory.FindInterestingEntitiesByEntityIDResult,
		FindInterestingEntitiesByRecordIDResult: factory.FindInterestingEntitiesByRecordIDResult,
		FindNetworkByEntityIDResult:             factory.FindNetworkByEntityIDResult,
		FindNetworkByRecordIDResult:             factory.FindNetworkByRecordIDResult,
		FindPathByEntityIDResult:                factory.FindPathByEntityIDResult,
		FindPathByRecordIDResult:                factory.FindPathByRecordIDResult,
		GetActiveConfigIDResult:                 factory.GetActiveConfigIDResult,
		GetEntityByEntityIDResult:               factory.GetEntityByEntityIDResult,
		GetEntityByRecordIDResult:               factory.GetEntityByRecordIDResult,
		GetRecordResult:                         factory.GetRecordResult,
		GetRedoRecordResult:                     factory.GetRedoRecordResult,
		GetStatsResult:                          factory.GetStatsResult,
		GetVirtualEntityByRecordIDResult:        factory.GetVirtualEntityByRecordIDResult,
		HowEntityByEntityIDResult:               factory.HowEntityByEntityIDResult,
		GetRecordPreviewResult:                  factory.GetRecordPreviewResult,
		ProcessRedoRecordResult:                 factory.ProcessRedoRecordResult,
		ReevaluateEntityResult:                  factory.ReevaluateEntityResult,
		ReevaluateRecordResult:                  factory.ReevaluateRecordResult,
		SearchByAttributesResult:                factory.SearchByAttributesResult,
		WhyEntitiesResult:                       factory.WhyEntitiesResult,
		WhyRecordInEntityResult:                 factory.WhyRecordInEntityResult,
		WhyRecordsResult:                        factory.WhyRecordsResult,
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method CreateProduct returns an SzProduct object
implemented to use the Senzing native C binary, libSz.so.

Input
  - ctx: A context to control lifecycle.

Output
  - An SzProduct object.
*/
func (factory *Szabstractfactory) CreateProduct(ctx context.Context) (senzing.SzProduct, error) {
	var err error

	_ = ctx
	result := &szproduct.Szproduct{
		GetLicenseResult: factory.GetLicenseResult,
		GetVersionResult: factory.GetVersionResult,
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

/*
Method Reinitialize re-initializes the Senzing objects created by the AbstractFactory
with a specific Senzing configuration JSON document identifier.

Input
  - ctx: A context to control lifecycle.
  - configID: The Senzing configuration JSON document identifier used for the initialization.
*/
func (factory *Szabstractfactory) Reinitialize(ctx context.Context, configID int64) error {
	var err error

	_ = ctx
	_ = configID

	return err
}
