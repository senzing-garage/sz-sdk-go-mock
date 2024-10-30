package szabstractfactory

import (
	"context"

	"github.com/senzing-garage/sz-sdk-go-mock/szconfig"
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
}

// ----------------------------------------------------------------------------
// senzing.SzAbstractFactory interface methods
// ----------------------------------------------------------------------------

/*
Method CreateSzConfig returns an SzConfig object
implemented to use the Senzing native C binary, libSz.so.

Input
  - ctx: A context to control lifecycle.

Output
  - An SzConfig object.
*/
func (factory *Szabstractfactory) CreateSzConfig(ctx context.Context) (senzing.SzConfig, error) {
	_ = ctx
	result := &szconfig.Szconfig{}
	return result, nil
}

/*
Method CreateSzConfigManager returns an SzConfigManager object
implemented to use the Senzing native C binary, libSz.so.

Input
  - ctx: A context to control lifecycle.

Output
  - An SzConfigManager object.
*/
func (factory *Szabstractfactory) CreateSzConfigManager(ctx context.Context) (senzing.SzConfigManager, error) {
	_ = ctx
	result := &szconfigmanager.Szconfigmanager{}
	return result, nil
}

/*
Method CreateSzDiagnostic returns an SzDiagnostic object
implemented to use the Senzing native C binary, libSz.so.

Input
  - ctx: A context to control lifecycle.

Output
  - An SzDiagnostic object.
*/
func (factory *Szabstractfactory) CreateSzDiagnostic(ctx context.Context) (senzing.SzDiagnostic, error) {
	_ = ctx
	result := &szdiagnostic.Szdiagnostic{}
	return result, nil
}

/*
Method CreateSzEngine returns an SzEngine object
implemented to use the Senzing native C binary, libSz.so.

Input
  - ctx: A context to control lifecycle.

Output
  - An SzEngine object.
*/
func (factory *Szabstractfactory) CreateSzEngine(ctx context.Context) (senzing.SzEngine, error) {
	_ = ctx
	result := &szengine.Szengine{}
	return result, nil
}

/*
Method CreateSzProduct returns an SzProduct object
implemented to use the Senzing native C binary, libSz.so.

Input
  - ctx: A context to control lifecycle.

Output
  - An SzProduct object.
*/
func (factory *Szabstractfactory) CreateSzProduct(ctx context.Context) (senzing.SzProduct, error) {
	_ = ctx
	result := &szproduct.Szproduct{}
	return result, nil
}

/*
Method Destroy will destroy and perform cleanup for the Senzing objects created by the AbstractFactory.
It should be called after all other calls are complete.

Input
  - ctx: A context to control lifecycle.
*/
func (factory *Szabstractfactory) Destroy(ctx context.Context) error {
	var err error
	_ = ctx
	return err
}

/*
Method Reinitialize re-initializes the Senzing objects created by the AbstractFactory with a specific Senzing configuration JSON document identifier.

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
