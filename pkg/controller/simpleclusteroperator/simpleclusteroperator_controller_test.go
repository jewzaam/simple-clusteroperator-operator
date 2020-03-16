package simpleclusteroperator

import (
	"context"
	"reflect"
	"testing"

	"github.com/kubernetes/client-go/dynamic/fake"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
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

func newMockClient(t *testing.T, localObjects []runtime.Object) fake.FakeDynamicClient {
	return fake.NewSimpleDynamicClient(runtime.NewScheme(), localObjects...)
}

func TestGetUnstructured(t *testing.T) {
	// prepare the resources to feed to the mock client
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cm-name",
			Namespace: "cm-namespace",
		},
		Data: map[string]string{
			"KEY": "VALUE",
		},
	}

	// create mock client w/ resources
	client := newMockClient(t, []runtime.Object{
		configMap,
	})

	resourceGkv := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "configmaps"}

	// try to find the config map
	resource := newResource(
		configMap.APIVersion,
		configMap.Kind,
		configMap.GetName(),
		configMap.GetNamespace(),
	)

	err := client.Resource(resourceGkv).Get(context.TODO(), types.NamespacedName{Name: resource.GetName(), Namespace: resource.GetNamespace()}, resource)

	if err != nil {
		t.Errorf("Error getting resource: %s", err)
	}

	// the fleshed out `resource` should have Data now
	if resource.Object["data"] == nil {
		t.Errorf("Did not get back Data on ConfigMap.  type=%s", reflect.TypeOf(resource.Object))
		for k := range resource.Object {
			t.Errorf("key == %s", k)
		}
	}
}
