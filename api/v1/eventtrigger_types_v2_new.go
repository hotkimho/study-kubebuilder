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

// EventTriggerSpec defines the desired state of EventTrigger
type EventTriggerSpec struct {
	// +kubeubuilder:default=false
	Suspend bool `json:"suspend"`

	// +kubeubuilder:validation:Minimum=0
	// +kubeubuilder:default=0
	MinDelaySeconds int `json:"minDelaySeconds"`

	// +kubeubuilder:validation:minLength=0
	// +kubeubuilder:default=Always
	RetryPolicy string `json:"retryPolicy"`

	// +kubeubuilder:validation:Minimum=0
	// +kubeubuilder:default=10
	TimeoutSeconds int `json:"timeoutSeconds"`

	Resource Resource `json:"resource"`

	Rule []Rule `json:"rule"`

	Message EventMessage `json:"message"`

	ReceiverConfig ReceiverConfig `json:"receiverConfig"`
}

type Resource struct {
	// +kubeubuilder:validation:MinLength=0
	Scope string `json:"scope"`

	// +kubeubuilder:validation:MinLength=0
	// +kubeubuilder:default=default
	Namespace *string `json:"namespace,omitempty"`

	// +kubeubuilder:validation:MinLength=0
	ApiVersion string `json:"apiVersion"`

	// +kubeubuilder:validation:MinLength=0
	Kind string `json:"kind"`

	// +kubeubuilder:validation:MinLength=0
	Name *string `json:"name,omitempty"`

	// +optional
	LabelSelector []LabelSelector `json:"labelSelector,omitempty"`

	// +optional
	AnnotationSelector []AnnotationSelector `json:"annotationSelector,omitempty"`
}

type LabelSelector struct {
	// +kubeubuilder:validation:MinLength=0
	Key string `json:"key"`

	// +kubeubuilder:validation:MinLength=0
	Operator string `json:"operator"`

	Values []string `json:"values"`
}

type AnnotationSelector struct {
	// +kubeubuilder:validation:MinLength=0
	Key string `json:"key"`

	// +kubeubuilder:validation:MinLength=0
	Operator string `json:"operator"`

	Values []string `json:"values"`
}

type Rule struct {
	// +kubeubuilder:validation:MinLength=0
	JSONPath string `json:"jsonPath"`

	// +kubeubuilder:validation:MinLength=0
	MatchExpression string `json:"matchExpression"`

	// +kubeubuilder:validation:MinLength=0
	Value string `json:"value"`
}

type EventMessage struct {
	// +kubeubuilder:validation:MinLength=0
	// +kubeubuilder:default=info
	Level string `json:"level"`

	// +kubeubuilder:validation:MinLength=0
	// +optional
	Message *string `json:"message"`
}

type ReceiverConfig struct {
	// +optional
	Gmail GmailConfig `json:"gmail,omitempty"`

	// +optional
	Mattermost MattermostConfig `json:"mattermost,omitempty"`
}

type GmailConfig struct {
	// +kubeubuilder:validation:MinLength=0
	SMTPServer string `json:"smtpServer"`

	// +kubeubuilder:validation:Minimum=0
	SMTPPort int `json:"smtpPort"`

	// +kubeubuilder:validation:MinLength=0
	Username string `json:"username"`

	// +kubeubuilder:validation:MinLength=0
	Password string `json:"password"`

	// +kubeubuilder:validation:MinLength=0
	FromAddress string `json:"fromAddress"`

	// +kubeubuilder:validation:MinLength=0
	ToAddress string `json:"toAddress"`

	// +kubeubuilder:validation:MinLength=0
	Subject string `json:"subject"`

	// +kubeubuilder:validation:MinLength=0
	Content string `json:"content"`
}

type MattermostConfig struct {
	// +kubeubuilder:validation:MinLength=0
	URL string `json:"url"`
}

// EventTriggerStatus defines the observed state of EventTrigger
type EventTriggerStatus struct {
	// +kubeubuilder:validation:MinLength=0
	// +kubeubuilder:default=0
	NotificationCount int `json:"notificationCount"`

	Conditions []Condition `json:"conditions"`

	ResourceStatus []ResourceStatus `json:"resourceStatus"`
}

type Condition struct {
	// +kubeubuilder:validation:MinLength=0
	Type string `json:"type"`

	// +kubeubuilder:validation:MinLength=0
	Status string `json:"status"`

	// +kubeubuilder:validation:MinLength=0
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// +kubeubuilder:validation:MinLength=0
	// +optional
	Message *string `json:"message,omitempty"`

	// +kubeubuilder:validation:MinLength=0
	// +optional
	Reason *string `json:"reason,omitempty"`
}

type ResourceStatus struct {
	// +kubeubuilder:validation:MinLength=0
	Name string `json:"name"`

	// +kubeubuilder:validation:MinLength=0
	ResourceVersion string `json:"resourceVersion"`

	// +kubeubuilder:default=false
	NotificationSent bool `json:"notificationSent"`

	// +kubeubuilder:validation:MinLength=0
	FirstNotificationTime metav1.Time `json:"firstNotificationTime,omitempty"`

	// +kubeubuilder:validation:MinLength=0
	LastNotificationTime metav1.Time `json:"lastNotificationTime,omitempty"`

	// +kubeubuilder:default=0
	NotificationCount int `json:"notificationCount"`

	ResponseStatus ResponseStatus `json:"responseStatus,omitempty"`
}

type ResponseStatus struct {
	// +kubeubuilder:vaidation:Minimum=0
	StatusCode int `json:"statusCode"`

	// +kubeubuilder:validation:MinLength=0
	// +optional
	ErrorMessage *string `json:"errorMessage"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

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
