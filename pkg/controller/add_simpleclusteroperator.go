package controller

import (
	"github.com/openshift/simple-clusteroperator-operator/pkg/controller/simpleclusteroperator"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, simpleclusteroperator.Add)
}
