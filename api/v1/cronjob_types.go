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
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:MinLength=0
	// cron 형식의 값 (예: "* * * * *")
	Schedule string `Json:"schedule"`

	// +kubebuilder:validation:minimum=0
	// +optional
	// 예정된 시간에 작업을 시작하지 못한 경우, 이 시간내에 작업하지 못하면 실패로 간주하고 다음 예약 시간에 작업을 시작함
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`

	// +optional
	// 작업의 동시 실행 정책
	// -"Allow" : 작업이 동시에 실행될 수 있음
	// -"Forbid" : 작업이 동시에 실행될 수 없음
	// -"Replace" : 새 작업이 실행되면 이전 작업이 종료됨
	ConcurrencyPolicy batchv1.ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`

	// +optional
	// 작업이 중지되는 경우 다음 작업을 계속할지 여부
	Suspend *bool `json:"suspend,omitempty"`

	// +optional
	// job 템플릿
	JobTemplate batchv1.JobTemplateSpec `json:"jobTemplate"`

	// kubebuilder:validation:Minimum=0
	// +optional
	// 유지할 성공된 작업의 수
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// kubeubuilder:validation:Minimum=0
	// +optional
	// 유지할 실패한 작업의 수
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}

// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	// 현재 실행 중인 작업에 대한 포인터 리스트
	Active []v1.ObjectReference `json:"active,omitempty"`

	// +optional
	// 작업이 마지막으로 성공한 예약된 시간
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
