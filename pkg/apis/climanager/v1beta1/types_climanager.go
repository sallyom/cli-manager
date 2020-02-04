package v1beta1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CLIManager is the Schema for the cluster cli-manager API
// +k8s:openapi-gen=true
// +genclient
type CLIManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CLIManagerSpec   `json:"spec,omitempty"`
	Status CLIManagerStatus `json:"status,omitempty"`
}

// CLIManagerSpec defines the desired state of CLIManager
type CLIManagerSpec struct {
	operatorv1.OperatorSpec `json:",inline"`

	// Description of CLIManager
	Description string `json:"description"`
	// Image for ClIManager
	Image string `json:"image"`
}

// CLIManagerStatus defines the observed state of CLIManager
type CLIManagerStatus struct {
	operatorv1.OperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CLIManagerList contains a list of CLIManagers
type CLIManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CLIManager `json:"items"`
}
