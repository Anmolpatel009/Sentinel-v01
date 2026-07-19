package observation

import "errors"

var (

	// ErrComponentNil indicates that the Observation
	// component has not been initialized.
	ErrComponentNil = errors.New("observation: component is nil")

	// ErrManagerNil indicates that the Observation
	// manager has not been initialized.
	ErrManagerNil = errors.New("observation: manager is nil")
)
