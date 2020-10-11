package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	CanaryFilterKind = "CanaryFilter"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CanaryFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CanaryFilterSpec `json:"spec"`
	Status CanaryStatus     `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CanaryFilterList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []CanaryFilter `json:"items"`
}

type CanaryFilterSpec struct {
	Service *CanaryFilterService `json:"service"`

	// Analysis defines the validation process of a release
	Analysis *CanaryAnalysis `json:"analysis,omitempty"`

	// ProgressDeadlineSeconds represents the maximum time in seconds for a
	// canary deployment to make progress before it is considered to be failed
	// +optional
	ProgressDeadlineSeconds *int32 `json:"progressDeadlineSeconds,omitempty"`

	// SkipAnalysis promotes the canary without analysing it
	// +optional
	SkipAnalysis bool `json:"skipAnalysis,omitempty"`
}

type CanaryFilterService struct {
	Labels map[string]string `json:"labels,omitempty"`

	// Defaults to http
	// +optional
	AppProtocol string `json:"appProtocol"`

	Port int32 `json:"port"`
}
