package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"


type ModuleSpec struct {
	Links []LinkDefinition `json:"interfaces"`
	Context string `json:"context"`
}


// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Module struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []Module `json:"items"`
}
