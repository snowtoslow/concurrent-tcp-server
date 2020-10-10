package server

import "errors"

var (
	ErrFieldNotFound = errors.New("FIELD NOT FOUND")
	NotValidInput    = errors.New("INPUT IS NOT VALID")
)
