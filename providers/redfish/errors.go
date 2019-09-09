package redfish

import "errors"

// ErrNoSystemEntry will be returned when we aren't able to find systems entries for Redfish data
var ErrNoSystemEntry = errors.New("no system entries present")
