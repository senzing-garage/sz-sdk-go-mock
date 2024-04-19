package szabstractfactory

import (
	"context"

	"github.com/senzing-garage/sz-sdk-go-mock/szconfig"
	"github.com/senzing-garage/sz-sdk-go-mock/szconfigmanager"
	"github.com/senzing-garage/sz-sdk-go-mock/szdiagnostic"
	"github.com/senzing-garage/sz-sdk-go-mock/szengine"
	"github.com/senzing-garage/sz-sdk-go-mock/szproduct"
	"github.com/senzing-garage/sz-sdk-go/sz"
)

// Szconfig is the default implementation of the Szconfig interface.
type Szabstractfactory struct {
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
TODO: Write description for CreateSzConfig
The CreateSzConfig method...

Input
  - ctx: A context to control lifecycle.

Output
  - An sz.SzConfig object.
    See the example output.
*/
func (factory *Szabstractfactory) CreateSzConfig(ctx context.Context) (sz.SzConfig, error) {
	result := &szconfig.Szconfig{}
	return result, nil
}

/*
TODO: Write description for CreateSzConfigManager
The CreateSzConfigManager method...

Input
  - ctx: A context to control lifecycle.

Output
  - An sz.CreateConfigManager object.
    See the example output.
*/
func (factory *Szabstractfactory) CreateSzConfigManager(ctx context.Context) (sz.SzConfigManager, error) {
	result := &szconfigmanager.Szconfigmanager{}
	return result, nil
}

/*
TODO: Write description for CreateSzDiagnostic
The CreateSzDiagnostic method...

Input
  - ctx: A context to control lifecycle.

Output
  - An sz.SzDiagnostic object.
    See the example output.
*/
func (factory *Szabstractfactory) CreateSzDiagnostic(ctx context.Context) (sz.SzDiagnostic, error) {
	result := &szdiagnostic.Szdiagnostic{}
	return result, nil
}

/*
TODO: Write description for CreateSzEngine
The CreateSzEngine method...

Input
  - ctx: A context to control lifecycle.

Output
  - An sz.SzEngine object.
    See the example output.
*/
func (factory *Szabstractfactory) CreateSzEngine(ctx context.Context) (sz.SzEngine, error) {
	result := &szengine.Szengine{}
	return result, nil
}

/*
TODO: Write description for CreateSzProduct
The CreateSzProduct method...

Input
  - ctx: A context to control lifecycle.

Output
  - An sz.SzProduct object.
    See the example output.
*/
func (factory *Szabstractfactory) CreateSzProduct(ctx context.Context) (sz.SzProduct, error) {
	result := &szproduct.Szproduct{}
	return result, nil
}
