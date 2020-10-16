package router

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	flaggerv1 "github.com/weaveworks/flagger/pkg/apis/flagger/v1beta1"
	istiov1alpha3 "github.com/weaveworks/flagger/pkg/apis/istio/v1alpha3"
	clientset "github.com/weaveworks/flagger/pkg/client/clientset/versioned"
)

// IstioRouter is managing Istio envoy filter
type IstioRouter struct {
	kubeClient    kubernetes.Interface
	istioClient   clientset.Interface
	flaggerClient clientset.Interface
	logger        *zap.SugaredLogger
}

func (r IstioRouter) Reconcile(canary *flaggerv1.CanaryFilter) error {

	newSpec := istiov1alpha3.EnvoyFilterSpec{
		ConfigPatches: []*istiov1alpha3.EnvoyConfigObjectPatch{
			{
				ApplyTo: istiov1alpha3.EnvoyFilterHTTPRoute,
				Match: &istiov1alpha3.EnvoyConfigObjectMatch{
					RouteConfiguration: &istiov1alpha3.RouteConfigurationMatch{
						Vhost: &istiov1alpha3.VirtualHostMatch{
							Route: &istiov1alpha3.RouteMatch{
								Action: istiov1alpha3.RouteMatchActionRoute,
							},
						},
					},
				},
				Patch: &istiov1alpha3.EnvoyFilterPatch{
					Operation: istiov1alpha3.EnvoyFilterPatchMerge,
				},
			},
		},
	}

	envoyFilter, err := r.istioClient.NetworkingV1alpha3().EnvoyFilters(canary.Namespace).Get(context.TODO(), canary.Name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		envoyFilter = &istiov1alpha3.EnvoyFilter{
			ObjectMeta: metav1.ObjectMeta{
				Name:      canary.Name,
				Namespace: canary.Namespace,
				OwnerReferences: []metav1.OwnerReference{
					*metav1.NewControllerRef(canary, schema.GroupVersionKind{
						Group:   flaggerv1.SchemeGroupVersion.Group,
						Version: flaggerv1.SchemeGroupVersion.Version,
						Kind:    flaggerv1.CanaryFilterKind,
					}),
				},
			},
			Spec: newSpec,
		}
		if _, err := r.istioClient.NetworkingV1alpha3().EnvoyFilters(canary.Namespace).Create(context.TODO(), envoyFilter, metav1.CreateOptions{}); err != nil {
			return fmt.Errorf("EnvoyFilter %s.%s create error: %w", canary.Name, canary.Namespace, err)
		}

		r.logger.With("canaryfilter", fmt.Sprintf("%s.%s", canary.Name, canary.Namespace)).
			Infof("EnvoyFilter %s.%s created", envoyFilter.GetName(), canary.Namespace)

		return nil
	}

	if envoyFilter != nil {
		if diff := cmp.Diff(envoyFilter.Spec, newSpec); diff != "" {
			clone := envoyFilter.DeepCopy()
			clone.Spec = newSpec
			if _, err := r.istioClient.NetworkingV1alpha3().EnvoyFilters(canary.Namespace).Update(context.TODO(), clone, metav1.UpdateOptions{}); err != nil {
				return fmt.Errorf("EnvoyFilter %s.%s update error: %w", canary.Name, canary.Namespace, err)
			}

			r.logger.With("canaryfilter", fmt.Sprintf("%s.%s", canary.Name, canary.Namespace)).
				Infof("EnvoyFilter %s.%s updated", envoyFilter.GetName(), canary.Namespace)
		}
	}

	return nil
}

func (r IstioRouter) SetRoutes(canary *flaggerv1.CanaryFilter, primaryWeight int, canaryWeight int, mirrored bool) error {
	panic("implement me")
}

func (r IstioRouter) GetRoutes(canary *flaggerv1.CanaryFilter) (primaryWeight int, canaryWeight int, mirrored bool, err error) {
	panic("implement me")
}

func (r IstioRouter) Finalize(canary *flaggerv1.CanaryFilter) error {
	panic("implement me")
}
