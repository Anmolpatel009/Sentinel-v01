package state

import "errors"

var (

	// ErrRepositoryNil indicates that the repository
	// has not been initialized.
	ErrRepositoryNil = errors.New("state: repository is nil")

	// ErrComponentNil indicates that the state component
	// has not been initialized.
	ErrComponentNil = errors.New("state: component is nil")

	// ErrInvalidState indicates that an invalid state
	// object was provided to the repository.
	ErrInvalidState = errors.New("state: invalid state")
)
