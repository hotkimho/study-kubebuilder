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

//
//import (
//	v1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//)
//
//// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
//// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
//
//// EventTriggerSpec defines the desired state of EventTrigger
//type EventTriggerSpec struct {
//	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
//	// Important: Run "make" to regenerate code after modifying this file
//
//	// Foo is an example field of EventTrigger. Edit eventtrigger_types.go to remove/update
//	Foo string `json:"foo,omitempty"`
//	v1.Volume
//	v1.Pod
//
//	ResourceEvents []ResourceEvent `json:"resourceEvents,omitempty"`
//}
//
//// EventTriggerStatus defines the observed state of EventTrigger
//type EventTriggerStatus struct {
//	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
//	// Important: Run "make" to regenerate code after modifying this file
//}
//
//type ResourceEvent struct {
//	EventSource `json:",inline"`
//}
//
//type EventSource struct {
//	PodEvent *PodEvent `json:"podEvent,omitempty"`
//}
//
//type PodEvent struct {
//	Resource          EventResource      `json:"resource,omitempty"`
//	Phase             string             `json:"phase,omitempty"`
//	Conditions        v1.PodCondition    `json:"conditions,omitempty"`
//	ContainerStatuses v1.ContainerStatus `json:"containerStatuses,omitempty"`
//}
//
//type DeploymentEvent struct {
//	Resource EventResource `json:"resource,omitempty"`
//}
//
//type EventResource struct {
//	// ApiVersion string  `json:"apiVersion,omitempty"`
//	// Kind       string  `json:"kind,omitempty"`
//	Name *string `json:"name,omitempty"`
//	// default: "default"
//	Namespace string            `json:"namespace,omitempty"`
//	Labels    map[string]string `json:"labels,omitempty"`
//}
//
//// +kubebuilder:object:root=true
//// +kubebuilder:subresource:status
//
//// EventTrigger is the Schema for the eventtriggers API
//type EventTrigger struct {
//	metav1.TypeMeta   `json:",inline"`
//	metav1.ObjectMeta `json:"metadata,omitempty"`
//
//	Spec   EventTriggerSpec   `json:"spec,omitempty"`
//	Status EventTriggerStatus `json:"status,omitempty"`
//}
//
//// +kubebuilder:object:root=true
//
//// EventTriggerList contains a list of EventTrigger
//type EventTriggerList struct {
//	metav1.TypeMeta `json:",inline"`
//	metav1.ListMeta `json:"metadata,omitempty"`
//	Items           []EventTrigger `json:"items"`
//}
//
//func init() {
//	SchemeBuilder.Register(&EventTrigger{}, &EventTriggerList{})
//}
