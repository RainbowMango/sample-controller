/*
Copyright 2017 The Kubernetes Authors.

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FooSpec   `json:"spec"`
	Status FooStatus `json:"status"`
}

// FooSpec is the spec for a Foo resource
type FooSpec struct {
	Resource       metav1.GroupVersionResource `json:"resource,omitempty"`
	Placement      GenericPlacementFields      `json:"placement,omitempty"`
	Override       []GenericOverrideItem       `json:"overrides,omitempty"`
	DeploymentName string                      `json:"deploymentName"`
	Replicas       *int32                      `json:"replicas"`
}

// GenericClusterReference represents a signal cluster name.
type GenericClusterReference struct {
	Name string `json:"name"`
}

// GenericPlacementFields tells which clusters will be propagate to.
type GenericPlacementFields struct {
	Clusters        []GenericClusterReference `json:"clusters,omitempty"`
	ClusterSelector *metav1.LabelSelector     `json:"clusterSelector,omitempty"`
}

type ClusterOverrideValue interface {
	DeepCopy() ClusterOverrideValue
}

type ClusterOverride struct {
	Op   string `json:"op,omitempty"`
	Path string `json:"path"`
	// TODO(RainbowMango): value should be `interface{}`, but deepcopy not support this as:
	// F0730 14:35:07.878733   29978 deepcopy.go:876] DeepCopy of "interface{}" is unsupported. Instead, use named interfaces with DeepCopy<named-interface> as one of the methods.
	// Just use a dummy interface 'ClusterOverrideValue'.
	Value ClusterOverrideValue `json:"value,omitempty"`
}

type GenericOverrideItem struct {
	ClusterName      string            `json:"clusterName"`
	ClusterOverrides []ClusterOverride `json:"clusterOverrides,omitempty"`
}

// FooStatus is the status for a Foo resource
type FooStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Foo `json:"items"`
}
