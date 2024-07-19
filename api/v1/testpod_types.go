/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TestPodSpec defines the desired state of TestPod
type TestPodSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".spec.replicas"
	Replicas int32 `json:"replicas"`

	// 테스트를 위해 추가된 임의의 값, Spec에 의도한 필드는 아님
	// +kubebuilder:default=KIMHO
	// +kubebuilder:validation:MinLength=0
	// +kubebuilder:printcolumn:name="Creator",type="string",JSONPath=".spec.creator"
	Creator string `json:"creator"`
}

// TestPodStatus defines the observed state of TestPod
type TestPodStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Phase string `json:"phase"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".spec.replicas"
// +kubebuilder:printcolumn:name="Creator",type="string",JSONPath=".spec.creator"
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"

// TestPod is the Schema for the testpods API
type TestPod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestPodSpec   `json:"spec,omitempty"`
	Status TestPodStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TestPodList contains a list of TestPod
type TestPodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestPod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestPod{}, &TestPodList{})
}
