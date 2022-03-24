// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of Cilium

package v2alpha1

import (

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories={cilium},singular="ciliumenvoymix",path="ciliumenvoymixs",scope="Cluster",shortName={cem}
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",description="The age of the identity",name="Age",type=date
// +kubebuilder:storageversion

type CiliumEnvoyMix struct {
	// +k8s:openapi-gen=false
	// +deepequal-gen=false
	metav1.TypeMeta `json:",inline"`
	// +k8s:openapi-gen=false
	// +deepequal-gen=false
	metav1.ObjectMeta `json:"metadata"`

	// +k8s:openapi-gen=false
	// +kubebuilder:validation:Type=object
	Spec CiliumEnvoyMixSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +deepequal-gen=false

// CiliumEnvoyConfigList is a list of CiliumEnvoyMix objects.
type CiliumEnvoyMixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is a list of CiliumEnvoyMix.
	Items []CiliumEnvoyMix `json:"items"`
}

type CiliumEnvoyMixSpec struct {
	
	// +kubebuilder:validation:Required
	Host string `json:"host,omitempty"`

	// +kubebuilder:validation:Required
	Destinations []Destination `json:"destination,omitempty"`

	// +kubebuilder:validation:Required
	Subsets []Subset `json:"subsets,omitempty"`

}

type Destination struct {

	// +kubebuilder:validation:Required
	Host string `json:"host,omitempty"`

	// +kubebuilder:validation:Required
	Subset string `json:"subset,omitempty"`
}

type Subset struct {
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

}
