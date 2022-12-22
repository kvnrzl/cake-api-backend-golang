package model

import "errors"

var (
	ErrRecordNotFound    = errors.New("Record not found.")
	ErrCannotCreate      = errors.New("Cannot save the data into the database due to a database error.")
	ErrCannotUpdate      = errors.New("Cannot update the data because there is something problem due to a database error.")
	ErrCannotDelete      = errors.New("Cannot delete the data from the database due to a database error.")
	ErrInputFieldInvalid = errors.New("Request could not be completed because the input fields provided was invalid.")
	ErrInternalServer    = errors.New("Internal server error.")
	// ErrDatabase          = errors.New("Unable to complete your request due to a database error.")
)
