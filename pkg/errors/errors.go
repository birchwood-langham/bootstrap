package errors

import (
	ge "errors"
)

var NotImplementedError = ge.New("NotImplemented")
var NoServiceInitializersError = ge.New("no service initializers have been defined")
var ServicePropertiesNotDefinedError = ge.New("service properties have not been defined")
