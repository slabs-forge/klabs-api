package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ModuleLinkSpec struct {
	LinkDefinition `json:",inline"`
	Context string `json:"context"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ModuleLink struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`	
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ModuleLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ModuleLink `json:"items"`
}