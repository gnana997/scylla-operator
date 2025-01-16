// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// RemoteKubernetesClusterLister helps list RemoteKubernetesClusters.
// All objects returned here must be treated as read-only.
type RemoteKubernetesClusterLister interface {
	// List lists all RemoteKubernetesClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RemoteKubernetesCluster, err error)
	// Get retrieves the RemoteKubernetesCluster from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RemoteKubernetesCluster, error)
	RemoteKubernetesClusterListerExpansion
}

// remoteKubernetesClusterLister implements the RemoteKubernetesClusterLister interface.
type remoteKubernetesClusterLister struct {
	listers.ResourceIndexer[*v1alpha1.RemoteKubernetesCluster]
}

// NewRemoteKubernetesClusterLister returns a new RemoteKubernetesClusterLister.
func NewRemoteKubernetesClusterLister(indexer cache.Indexer) RemoteKubernetesClusterLister {
	return &remoteKubernetesClusterLister{listers.New[*v1alpha1.RemoteKubernetesCluster](indexer, v1alpha1.Resource("remotekubernetescluster"))}
}
