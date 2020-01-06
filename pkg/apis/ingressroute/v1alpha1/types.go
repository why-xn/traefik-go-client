package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IngressRoute is a top-level type
type IngressRoute struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// your own custom spec
	Spec IngressRouteSpec `json:"spec,omitempty"`
}

// custom spec
type IngressRouteSpec struct {
	EntryPoints []string                `json:"entryPoints"`
	Routes      []IngressRouteSpecRoute `json:"routes"`
	Tls         IngressRouteSpecTls     `json:"tls"`
}

// custom object
type IngressRouteSpecRoute struct {
	Match    string                         `json:"match"`
	Kind     string                         `json:"kind"`
	Priority int64                          `json:"priority"`
	Services []IngressRouteSpecRouteService `json:"services"`
}

// custom object
type IngressRouteSpecRouteService struct {
	Name               string                         `json:"name"`
	Port               int64                          `json:"port"`
	Weight             int64                          `json:"weight"`
	PassHostHeader     bool                           `json:"passHostHeader"`
	ResponseForwarding RouteServiceResponseForwarding `json:"responseForwarding"`
}

// custom object
type RouteServiceResponseForwarding struct {
	FlushInterval int64 `json:"flushInterval"`
}

// custom object
type IngressRouteSpecTls struct {
	CertResolver string `json:"certResolver"`
}
