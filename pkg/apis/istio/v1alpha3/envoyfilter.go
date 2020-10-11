// proto: https://github.com/istio/api/blob/master/networking/v1alpha3/envoy_filter.proto
package v1alpha3

import (
	"github.com/gogo/protobuf/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
spec:
  configPatches:
  - applyTo: HTTP_ROUTE
    match:
      context: GATEWAY
      routeConfiguration:
        portNumber: 80
        vhost:
          name: "*:80"
          route:
            action: ROUTE
    patch:
      operation: MERGE
      value:
        route:
          request_mirror_policies:
          - cluster: outbound|80|grpc|bookinfo-canary
            runtime_fraction:
              default_value:
                numerator: 200000
                denominator: MILLION
*/

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EnvoyFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnvoyFilterSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EnvoyFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []EnvoyFilter `json:"items"`
}
type EnvoyFilterSpec struct {
	ConfigPatches []*EnvoyConfigObjectPatch `json:"configPatches"`
}

type EnvoyConfigObjectPatch struct {
	ApplyTo EnvoyFilterApplyTo      `json:"applyTo,omitempty"`
	Match   *EnvoyConfigObjectMatch `json:"match,omitempty"`
	Patch   *EnvoyFilterPatch       `json:"patch,omitempty"`
}

type EnvoyFilterApplyTo int32

const (
	EnvoyFilterInvalid            EnvoyFilterApplyTo = 0
	EnvoyFilterListener           EnvoyFilterApplyTo = 1
	EnvoyFilterFilterChain        EnvoyFilterApplyTo = 2
	EnvoyFilterNetworkFilter      EnvoyFilterApplyTo = 3
	EnvoyFilterHTTPFilter         EnvoyFilterApplyTo = 4
	EnvoyFilterRouteConfiguration EnvoyFilterApplyTo = 5
	EnvoyFilterVirtualHost        EnvoyFilterApplyTo = 6
	EnvoyFilterHTTPRoute          EnvoyFilterApplyTo = 7
	EnvoyFilterCluster            EnvoyFilterApplyTo = 8
	EnvoyFilterExtensionConfig    EnvoyFilterApplyTo = 9
)

type EnvoyConfigObjectMatch struct {
	Context            EnvoyFilterPatchContext  `json:"context,omitempty"`
	Proxy              *EnvoyFilterProxyMatch   `json:"proxy,omitempty"`
	Listener           *ListenerMatch           `json:"listener,omitempty"`
	RouteConfiguration *RouteConfigurationMatch `json:"routeConfiguration,omitempty"`
}

type EnvoyFilterPatchContext int32

const (
	EnvoyFilterAny             EnvoyFilterPatchContext = 0
	EnvoyFilterSidecarInbound  EnvoyFilterPatchContext = 1
	EnvoyFilterSidecarOutbound EnvoyFilterPatchContext = 2
	EnvoyFilterGateway         EnvoyFilterPatchContext = 3
)

type EnvoyFilterProxyMatch struct {
	ProxyVersion string            `json:"proxy_version,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type ListenerMatch struct {
	PortNumber uint32                         `json:"port_number,omitempty"`
	PortName   string                         `json:"port_name,omitempty"`
	FilterChan *ListenerMatchFilterChainMatch `json:"filter_chain,omitempty"`
}

type EnvoyFilterListenerMatchFilterChainMatch struct {
	Name   string                    `json:"name,omitempty"`
	Sni    string                    `json:"sni,omitempty"`
	Filter *ListenerMatchFilterMatch `json:"filter,omitempty"`
}

type ListenerMatchFilterMatch struct {
	// The filter name to match on.
	Name string `json:"name,omitempty"`

	// Typically used for HTTP Connection Manager filters and Thrift filters.
	SubFilter *ListenerMatchSubFilterMatch `json:"sub_filter,omitempty"`
}

type ListenerMatchSubFilterMatch struct {
	// The filter name to match on.
	Name string `json:"name,omitempty"`
}

type EnvoyFilterPatch struct {
	Operation   EnvoyFilterPatchOperation   `json:"operation,omitempty"`
	Value       *types.Struct               `json:"value,omitempty"`
	FilterClass EnvoyFilterPatchFilterClass `json:"filterClass,omitempty"`
}

type EnvoyFilterPatchOperation int32

const (
	EnvoyFilterPatchInvalid      EnvoyFilterPatchOperation = 0
	EnvoyFilterPatchMerge        EnvoyFilterPatchOperation = 1
	EnvoyFilterPatchAdd          EnvoyFilterPatchOperation = 2
	EnvoyFilterPatchRemove       EnvoyFilterPatchOperation = 3
	EnvoyFilterPatchInsertBefore EnvoyFilterPatchOperation = 4
	EnvoyFilterPatchInsertAfter  EnvoyFilterPatchOperation = 5
	EnvoyFilterPatchInsertFirst  EnvoyFilterPatchOperation = 6
	EnvoyFilterPatchReplace      EnvoyFilterPatchOperation = 7
)

type EnvoyFilterPatchFilterClass int32

const (
	EnvoyFilterPatchUnspecified EnvoyFilterPatchFilterClass = 0
	EnvoyFilterPatchAuthn       EnvoyFilterPatchFilterClass = 1
	EnvoyFilterPatchAuthz       EnvoyFilterPatchFilterClass = 2
	EnvoyFilterPatchStats       EnvoyFilterPatchFilterClass = 3
)

type RouteConfigurationMatch struct {
	PortNumber uint32            `json:"port_number,omitempty"`
	PortName   string            `json:"port_name,omitempty"`
	Gateway    string            `json:"gateway,omitempty"`
	Vhost      *VirtualHostMatch `json:"vhost,omitempty"`
	Name       string            `json:"name,omitempty"`
}

type VirtualHostMatch struct {
	Name  string      `json:"name,omitempty"`
	Route *RouteMatch `json:"route,omitempty"`
}

type RouteMatch struct {
	Name   string           `json:"name,omitempty"`
	Action RouteMatchAction `json:"action,omitempty"`
}

type RouteMatchAction int32

const (
	RouteMatchActionAny      RouteMatchAction = 0
	RouteMatchActionRoute    RouteMatchAction = 1
	RouteMatchActionRedirect RouteMatchAction = 2
	RouteMatchActionResponse RouteMatchAction = 3
)
