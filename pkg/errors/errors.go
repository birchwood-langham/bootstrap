package errors

import (
	ge "errors"
)

var NotImplementedError = ge.New("NotImplemented")
var NoServiceInitializersError = ge.New("No service initializers have been defined")
var ServicePropertiesNotDefinedError = ge.New("Service properties have not been defined")
