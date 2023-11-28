package errors

import "fmt"

type ErrTerraformBackendFileIsNotFound struct {
	ErrWrapped      error
	BackendFilePath string
}

func (e *ErrTerraformBackendFileIsNotFound) Error() string {
	return fmt.Sprintf("The backend file %s is invalid, or it does not exist: %s",
		e.BackendFilePath, e.ErrWrapped)
}

type ErrTerraformInitFailedToStart struct {
	ErrWrapped error
	Details    string
}

func (e *ErrTerraformInitFailedToStart) Error() string {
	return fmt.Sprintf("Failed to start the terraform init command: %s: %s",
		e.Details, e.ErrWrapped)
}
