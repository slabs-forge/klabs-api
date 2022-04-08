package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	slabs_forge "github.com/slabs-forge/klabs-api/apis/slabs.forge"
)

var SchemeGroupVersion = schema.GroupVersion{Group: slabs_forge.GroupName, Version: "v1"}

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	
	AddToScheme = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Module{},
		&ModuleList{},
		&ModuleLink{},
		&ModuleLinkList{},
	)
	
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	
	return nil
}
