package k8_extensions

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// SupportedThirdPartyResources contains all supported TPRs on bitesize
// cluster.
var SupportedThirdPartyResources = []string{
	"mongo", "mysql", "cassandra", "redis", "zookeeper", "kafka", "postgres", "neptune", "mks", "docdb", "cb",
}

// PrsnExternalResource represents ThirdpartyResources mapped from
// kubernetes to externally running services (e.g. RDS, cassandra, mongo etc.)
type PrsnExternalResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec PrsnExternalResourceSpec `json:"spec"`
}

// PrsnExternalResourceSpec represents format for these mappings - which is
// basically it's version and  options
type PrsnExternalResourceSpec struct {
	Version  string                 `json:"version"`
	Options  map[string]interface{} `json:"options"`
	Replicas int                    `json:"replicas"`
}

// PrsnExternalResourceList is a list of PrsnExternalResource
type PrsnExternalResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []PrsnExternalResource `json:"items"`
}

// DeepCopyObject required to satisfy Object interface
func (tpr PrsnExternalResource) DeepCopyObject() runtime.Object {
	return new(PrsnExternalResource)
}

// DeepCopyObject required to satisfy Object interface
func (tpr PrsnExternalResourceList) DeepCopyObject() runtime.Object {
	return new(PrsnExternalResource)
}
