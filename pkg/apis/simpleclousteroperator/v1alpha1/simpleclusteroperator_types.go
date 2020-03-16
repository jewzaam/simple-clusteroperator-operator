package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// DefaultSimpleClusterOperatorDaemonSetDegradedSeconds is the default time in seconds when desired count is not equal to current count before declaring a DaemonSet degraded.
	DefaultSimpleClusterOperatorDaemonSetDegradedSeconds = 600
	// DefaultSimpleClusterOperatorDaemonSetUnavailableCount is the default maximum number of containers deplyoed while declaring a DaemonSet unavailable.
	DefaultSimpleClusterOperatorDaemonSetUnavailableCount = 0
	// DefaultSimpleClusterOperatorDeploymentUnavailableCount is the default maximum number of containers deplyoed while declaring a Deployment unavailable.
	DefaultSimpleClusterOperatorDeploymentUnavailableCount = 0
)

// SimpleClusterOperatorCsv defines details of any CSV being watched.
// +k8s:openapi-gen=true
type SimpleClusterOperatorCsv struct {
	Name              string   `json:"name"`
	Namespace         string   `json:"namespace"`
	DegradedPhases    []string `json:"degradedPhases,omitempty"`
	ProgressingPhases []string `json:"progressingPhases,omitempty"`
	UnavailablePhases []string `json:"unavailablePhases,omitempty"`
	UpgradeablePhases []string `json:"upgradeablePhases,omitempty"`
}

// SimpleClusterOperatorDaemonSet defines details of any DaemonSet being watched.  If Degraded or Unavailable conditions are not met the DaemonSet is assumed to be Progressing unless current and desired counts are equal.
// +k8s:openapi-gen=true
type SimpleClusterOperatorDaemonSet struct {
	Name              string   `json:"name"`
	Namespace         string   `json:"namespace"`
	// DegradedSeconds is the time in seconds before declaring a degraded state when desired count is not equal to current count.  Default is SimpleClusterOperatorDaemonSetDegradedSeconds.
	DegradedSeconds int `json:"degradedSeconds,,omitempty"`
	// UnavailableCount is the maximum number of containers deployed while declaring an unavailable state.  Default is DefaultSimpleClusterOperatorDaemonSetUnavailableCount.
	UnavailableCount int `json:"unaviailableCound,omitempty"`
}

// SimpleClusterOperatorDeployment defines details of any Deployment being watched.
// +k8s:openapi-gen=true
type SimpleClusterOperatorDeployment struct {
	Name              string   `json:"name"`
	Namespace         string   `json:"namespace"`
	// DegradedSeconds is the time in seconds before declaring a degraded state when desired count is not equal to current count.  Default is SimpleClusterOperatorDeploymentDegradedSeconds.
	DegradedSeconds int `json:"degradedSeconds,omitempty"`
	// UnavailableCount is the maximum number of containers deployed before declaring an unavailable state.  Default is DefaultSimpleClusterOperatorDeploymentUnavailableCount.
	UnavailableCount int `json:"unaviailableCound,omitempty"`
}

// SimpleClusterOperatorResources defines details of arbitrary Resources checked for existance.
type SimpleClusterOperatorResources struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Kind	  string `json:"kind"`
}

// SimpleClusterOperatorWatch defines the details of an operator being watched
// +k8s:openapi-gen=true
type SimpleClusterOperatorWatch struct {
	Csvs []SimpleClusterOperatorCsv `json:"csvs,omitempty"`
	DaemonSets []SimpleClusterOperatorDaemonSet `json:"daemonSets,omitempty"`
	Deployments []SimpleClusterOperatorDeployment `json:"deployments,omitempty"`
	Resources []SimpleClusterOperatorResources `json:"resources,omitempty"`
}

// SimpleClusterOperatorSpec defines the desired state of SimpleClusterOperator
// +k8s:openapi-gen=true
type SimpleClusterOperatorSpec struct {
	Watches []SimpleClusterOperatorWatch `json:"watches"`
}

// SimpleClusterOperatorStatus defines the observed state of SimpleClusterOperator
// +k8s:openapi-gen=true
type SimpleClusterOperatorStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SimpleClusterOperator is the Schema for the simpleclusteroperators API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type SimpleClusterOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SimpleClusterOperatorSpec   `json:"spec,omitempty"`
	Status SimpleClusterOperatorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SimpleClusterOperatorList contains a list of SimpleClusterOperator
type SimpleClusterOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SimpleClusterOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SimpleClusterOperator{}, &SimpleClusterOperatorList{})
}
