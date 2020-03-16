package simpleclusteroperator

import (
	"testing"
)

func TestNewResource(t *testing.T) {
	apiVersion := "a"
	kind := "b"
	name := "c"
	namespace := "d"

	resource := newResource(apiVersion, kind, name, namespace)

	if resource.GetAPIVersion() != apiVersion {
		t.Errorf("GetAPIVersion: expected '%s', found '%s'", apiVersion, resource.GetAPIVersion())
	}
	if resource.GetKind() != kind {
		t.Errorf("GetAPIVersion: expected '%s', found '%s'", kind, resource.GetKind())
	}
	if resource.GetName() != name {
		t.Errorf("GetAPIVersion: expected '%s', found '%s'", name, resource.GetName())
	}
	if resource.GetNamespace() != namespace {
		t.Errorf("GetAPIVersion: expected '%s', found '%s'", namespace, resource.GetNamespace())
	}
}
