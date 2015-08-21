package clusternetwork

import (
	"fmt"

	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/fielderrors"

	"github.com/openshift/origin/pkg/sdn/api"
	"github.com/openshift/origin/pkg/sdn/api/validation"
)

// sdnStrategy implements behavior for ClusterNetworks
type sdnStrategy struct {
	runtime.ObjectTyper
}

// Strategy is the default logic that applies when creating and updating ClusterNetwork
// objects via the REST API.
var Strategy = sdnStrategy{kapi.Scheme}

func (sdnStrategy) PrepareForUpdate(obj, old runtime.Object) {}

// NamespaceScoped is false for sdns
func (sdnStrategy) NamespaceScoped() bool {
	return false
}

func (sdnStrategy) GenerateName(base string) string {
	return base
}

func (sdnStrategy) PrepareForCreate(obj runtime.Object) {
}

// Validate validates a new sdn
func (sdnStrategy) Validate(ctx kapi.Context, obj runtime.Object) fielderrors.ValidationErrorList {
	return validation.ValidateClusterNetwork(obj.(*api.ClusterNetwork))
}

// AllowCreateOnUpdate is false for sdns
func (sdnStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (sdnStrategy) AllowUnconditionalUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for a ClusterNetwork
func (sdnStrategy) ValidateUpdate(ctx kapi.Context, obj, old runtime.Object) fielderrors.ValidationErrorList {
	return validation.ValidateClusterNetworkUpdate(obj.(*api.ClusterNetwork), old.(*api.ClusterNetwork))
}

// MatchClusterNetwork returns a generic matcher for a given label and field selector.
func MatchClusterNetwork(label labels.Selector, field fields.Selector) generic.Matcher {
	return generic.MatcherFunc(func(obj runtime.Object) (bool, error) {
		_, ok := obj.(*api.ClusterNetwork)
		if !ok {
			return false, fmt.Errorf("not a ClusterNetwork")
		}
		return true, nil
	})
}
