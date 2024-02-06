package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster

// VNode is the struct created for Virtual Node to report and store node info
type VNode struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a NMNode.
	// +optional
	Spec VNodeSpec `json:"spec,omitempty"`

	// Status represents the current information about a NMNode. This data may not be up
	// to date.
	// +optional
	Status VNodeStatus `json:"status,omitempty"`
}

type VNodeSpec struct {
	// If specified, the nmnode's taints.
	// TODO:should we move this to NMNode Status ?
	// +optional
	Taints []v1.Taint `json:"taints,omitempty"`
}

type VNodeStatus struct {
	// store the resource info reported by RM(NM)
	// +optional
	ResourceCapacity *v1.ResourceList `json:"resourceCapacity,omitempty"`
	// +optional
	ResourceAllocatable *v1.ResourceList `json:"resourceAllocatable,omitempty"`
	// node status from Yarn perspective
	// +optional
	NodeStatus NodePhase `json:"nodeStatus,omitempty"`
	// +optional
	NodeCondition []*v1.NodeCondition `json:"nodeCondition,omitempty"`

	// machine status
	// +optional
	MachineStatus *MachineStatus `json:"machineStatus,omitempty"`
}

type NodePhase string

type MachineStatus struct {
	// +optional
	LoadAvgLastM *resource.Quantity `json:"loadAvgLastM,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VNodeList is a collection of VNode objects.
type VNodeList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// items is the list of ENode
	Items []VNode `json:"items"`
}
