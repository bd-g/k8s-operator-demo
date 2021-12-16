/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DatabaseDependencyDemoSpec defines the desired state of DatabaseDependencyDemo
type DatabaseDependencyDemoSpec struct {
	//+kubebuilder:validation:Minimum=0
	// Size is the size of the deployment
	Size int32 `json:"size"`
}

// DatabaseDependencyDemoStatus defines the observed state of DatabaseDependencyDemo
type DatabaseDependencyDemoStatus struct {
	// Nodes are the names of the memcached pods
	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DatabaseDependencyDemo is the Schema for the databasedependencydemoes API
type DatabaseDependencyDemo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseDependencyDemoSpec   `json:"spec,omitempty"`
	Status DatabaseDependencyDemoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DatabaseDependencyDemoList contains a list of DatabaseDependencyDemo
type DatabaseDependencyDemoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseDependencyDemo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseDependencyDemo{}, &DatabaseDependencyDemoList{})
}
