package controller

import (
	"github.com/tpokki/gatling-operator/pkg/controller/gatlingtask"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, gatlingtask.Add)
}
