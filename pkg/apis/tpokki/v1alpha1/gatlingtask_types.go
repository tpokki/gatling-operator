package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GatlingTaskSpec defines the desired state of GatlingTask
type GatlingTaskSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Replicas                    int32                `json:"replicas,omitempty"`
	RestartPolicy               corev1.RestartPolicy `json:"restartPolicy,omitempty"`
	corev1.ResourceRequirements `json:"resources,omitempty"`
	ScenarioSpec                `json:"scenario,omitempty"`
}

// ScenarioSpec defines the loaded gatling scenario
type ScenarioSpec struct {
	Name       string            `json:"name"`
	DataSource map[string]string `json:"data,omitempty"`
	GitSource  GitSourceSpec     `json:"git,omitempty"`
}

// GitSourceSpec defines the git repository and the location in it for the scenario
type GitSourceSpec struct {
	Repo string `json:"repo"`
	Dir  string `json:"directory"`
}

// GatlingTaskStatus defines the observed state of GatlingTask
type GatlingTaskStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatlingTask is the Schema for the gatlingtasks API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=gatlingtasks,scope=Namespaced
type GatlingTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatlingTaskSpec   `json:"spec,omitempty"`
	Status GatlingTaskStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatlingTaskList contains a list of GatlingTask
type GatlingTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatlingTask `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GatlingTask{}, &GatlingTaskList{})
}
