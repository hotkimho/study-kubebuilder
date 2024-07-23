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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EventTriggerSpec defines the desired state of EventTrigger
type EventTriggerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of EventTrigger. Edit eventtrigger_types.go to remove/update
	Foo string `json:"foo,omitempty"`
	v1.Volume
	v1.Pod

	ResourceEvents []ResourceEvent `json:"resourceEvents,omitempty"`
}

// EventTriggerStatus defines the observed state of EventTrigger
type EventTriggerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

type ResourceEvent struct {
	EventSource `json:",inline"`
}

type EventSource struct {
	Pod         *PodEvent         `json:"pod,omitempty"`
	Deployment  *DeploymentEvent  `json:"deployment,omitempty"`
	ClusterInfo *ClusterInfoEvent `json:"clusterInfo,omitempty"`
}

type PodEvent struct {
	Resource      EventResource    `json:"resource,omitempty"`
	Conditions    []EventCondition `json:"conditions"`
	Message       EventMessage     `json:"message"`
	Notifications []Notification   `json:"notifications"`
}

type DeploymentEvent struct {
	Resource      EventResource    `json:"resource,omitempty"`
	Conditions    []EventCondition `json:"conditions"`
	Message       EventMessage     `json:"message"`
	Notifications []Notification   `json:"notifications"`
}

type ClusterInfoEvent struct {
	Resource      EventResource    `json:"resource,omitempty"`
	Conditions    []EventCondition `json:"conditions"`
	Message       EventMessage     `json:"message"`
	Notifications []Notification   `json:"notifications"`
}

type EventResource struct {
	Namespace string            `json:"namespace"`
	Name      string            `json:"name"`
	Labels    map[string]string `json:"labels,omitempty"`
}

type EventCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type EventMessage struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

type Notification struct {
	Type        string  `json:"type"`
	URL         *string `json:"url,omitempty"`
	SMTPServer  *string `json:"smtpServer,omitempty"`
	SMTPPort    *int    `json:"smtpPort,omitempty"`
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`
	FromAddress *string `json:"fromAddress,omitempty"`
	ToAddress   *string `json:"toAddress,omitempty"`
	Subject     *string `json:"subject,omitempty"`
	Message     *string `json:"message,omitempty"`
}

// EventTrigger is the Schema for the eventtriggers API
type EventTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventTriggerSpec   `json:"spec,omitempty"`
	Status EventTriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventTriggerList contains a list of EventTrigger
type EventTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventTrigger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventTrigger{}, &EventTriggerList{})
}
