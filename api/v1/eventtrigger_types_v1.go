package v1

//
//import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//
//// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
//// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
//
//type EventTrigger struct {
//	metav1.TypeMeta   `json:",inline"`
//	metav1.ObjectMeta `json:"metadata,omitempty"`
//
//	Spec   EventTriggerSpec   `json:"spec,omitempty"`
//	Status EventTriggerStatus `json:"status,omitempty"`
//}
//
//type EventTriggerSpec struct {
//	ResourceEvents []ResourceEvent `json:"resourceEvents,omitempty"`
//}
//
//type ResourceEvent struct {
//	// 리소스 정보
//	Resource EventResource `json:"resource"`
//
//	// 조건 목록(status, value)
//	Conditions []EventCondition `json:"conditions"`
//
//	// 메시지 정보
//	Message EventMessage `json:"message"`
//
//	// 전송할 매체(matterMost, Gmail) 정보
//	Notifications []Notification `json:"notifications"`
//}
//
//type EventResource struct {
//	// +kubeubuilder:validation:MinLength=0
//	// 리소스 버전
//	APIVersion string `json:"apiVersion"`
//
//	// +kubeubuilder:validation:MinLength=0
//	// 리로스 종류(kind)
//	Kind string `json:"kind"`
//
//	// +kubebuilder:default=default
//	// 리소스가 속한 네임스페이스
//	Namespace string `json:"namespace"`
//
//	// +kubebuiler:validation:MinLength=0
//	// 리소스 이름
//	Name string `json:"name"`
//
//	// 리소스 레이블
//	// +optional
//	Labels map[string]string `json:"labels,omitempty"`
//}
//
//type ValueCondition struct {
//	// +kubeubuilder:validation:MinLength=0
//	// 검사할 필드명
//	Field string `json:"field"`
//
//	// +kubeubuilder:validation:MinLength=0
//	// 비교 연산자 ex) equal, notEqual, lessThan, greaterThan
//	Operator string `json:"operator"`
//
//	// +kubeubuilder:validation:MinLength=0
//	// 비교할 값(수치)
//	Value string `json:"value"`
//}
//
//type StatusCondition struct {
//	// +kubeubuilder:validation:MinLength=0
//	// 검사할 필드명
//	Field string `json:"field"`
//
//	// +kubeubuilder:validation:MinLength=0
//	// 비교 연산자 ex) equal, notEqual, Pending, Running, Succeeded, Failed, Unknown
//	Operator string `json:"operator"`
//
//	// +kubeubuilder:validation:MinLength=0
//	// 비교할 상태
//	Status string `json:"status"`
//}
//
//type EventCondition struct {
//	// +optional
//	// 값에 대한 조건
//	ValueCondition []ValueCondition `json:"valueCondition,omitempty"`
//
//	// +optional
//	// 상태에 대한 조건
//	StatusCondition []StatusCondition `json:"statusCondition,omitempty"`
//}
//
//type EventMessage struct {
//	// +kubeubuilder:validation:MinLength=0
//	// 로그 레벨 ex) info, warning, error
//	Level string `json:"level"`
//
//	// +optional
//	// 알림 매체에 보낼 메시지 내용
//	Message *string `json:"message"`
//}
//
//type Notification struct {
//	// +kubeubuilder:validation:MinLength=0
//	// 매체 종류(mattermost, gmail)
//	Type string `json:"type"`
//
//	// +optional
//	// 매체 URL(webhook URL)
//	URL *string `json:"url,omitempty"`
//
//	// +optional
//	// 매체 채널
//	Channel *string `json:"channel,omitempty"`
//
//	// +optional
//	// SMTP 서버 주소
//	SMTPServer *string `json:"smtpServer,omitempty"`
//
//	// +optional
//	// SMTP 포트
//	SMTPPort *int `json:"smtpPort,omitempty"`
//
//	// +optional
//	// mail 계정 사용자 이름
//	Username *string `json:"username,omitempty"`
//
//	// +optional
//	// mail 계정 비밀번호
//	Password *string `json:"password,omitempty"`
//
//	// +optional
//	// 발신자 메일 주소
//	FromAddress *string `json:"fromAddress,omitempty"`
//
//	// +optional
//	// 수신자 메일 주소
//	ToAddress *string `json:"toAddress,omitempty"`
//
//	// +optional
//	// 메일 제목
//	Subject *string `json:"subject,omitempty"`
//
//	// +optional
//	// 메일 내용
//	Message *string `json:"message,omitempty"`
//}
//
//type EventTriggerStatus struct {
//	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
//	// Important: Run "make" to regenerate code after modifying this file
//
//	ResourceStatus []ResourceStatus `json:"resourceStatus,omitempty"`
//}
//
//type ResourceStatus struct {
//	// +kubeubuilder:validation:MinLength=0
//	// 리소스 이름
//	Name string `json:"name"`
//
//	// +kubeubuilder:validation:MinLength=0
//	// 리소스 종류
//	Kind string `json:"kind"`
//
//	// +optional
//	// 리소스에 대한 상태
//	Conditions []ConditionStatus `json:"conditions"`
//
//	// +optional
//	// 마지막으로 알림을 전송한 시간
//	LastNotificationTime metav1.Time `json:"lastNotificationTime"`
//
//	// +optional
//	// 알림을 처음 전송한 시간
//	FirstNotificationTime metav1.Time `json:"firstNotificationTime"`
//
//	// +kubeubuilder:defualt=0
//	NotificationCount int `json:"notificationCount"`
//}
//
//type ConditionStatus struct {
//	// +kubeubuilder:validation:MinLength=0
//	// 검사한 필드명
//	Field string `json:"field"`
//
//	// +kubeubuilder:validation:MinLength=0
//	// 비교 연산자 ex) equal, notEqual, lessThan, greaterThan
//	Operator string `json:"operator"`
//
//	// +optional
//	// 현재 리소스의 상태
//	Status *string `json:"status,omitempty"`
//
//	// +optional
//	// 현재 리소스의 값(수치)
//	Value *string `json:"value,omitempty"`
//
//	// 현재 리소스의 상태와 조건이 일치하는지 여부(true, false)
//	Result string `json:"result"`
//}
