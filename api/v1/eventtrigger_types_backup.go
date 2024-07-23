package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type TestTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestTriggerSpec   `json:"spec,omitempty"`
	Status TestTriggerStatus `json:"status,omitempty"`
}

type TestTriggerSpec struct {
	ResourceEvents []ResourceEvent `json:"resourceEvents,omitempty"`
}

type ResourceEvent struct {
	Resource      EventResource    `json:"resource"`
	Conditions    []EventCondition `json:"conditions"`
	Message       EventMessage     `json:"message"`
	Notifications []Notification   `json:"notifications"`
}

type EventResource struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Namespace  string            `json:"namespace"`
	Name       string            `json:"name"`
	Labels     map[string]string `json:"labels,omitempty"`
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

type TestTriggerStatus struct {
	ResourceStatuses []ResourceStatus `json:"resourceStatuses"`
}

type ResourceStatus struct {
	Name            string      `json:"name"`
	Kind            string      `json:"kind"`
	Status          string      `json:"status"`
	Reported        bool        `json:"reported"`
	LastReportTime  metav1.Time `json:"lastReportTime"`
	FirstReportTime metav1.Time `json:"firstReportTime"`
}
