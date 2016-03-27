package guber

// Common
//==============================================================================
type ResourceDefinition struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
}

type Metadata struct {
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels,omitempty"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
}

// Namespace
//==============================================================================
type Namespace struct {
	*ResourceDefinition
	Metadata *Metadata `json:"metadata"`
}

type NamespaceList struct {
	Items []*Namespace `json:"items"`
}

// Node
//==============================================================================
type NodeSpec struct {
	ExternalID string `json:"externalID"`
}

type NodeStatusCapacity struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

type NodeStatusCondition struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}

type NodeStatus struct {
	Capacity   *NodeStatusCapacity    `json:"capacity"`
	Conditions []*NodeStatusCondition `json:"conditions"`
}

type Node struct {
	*ResourceDefinition
	Metadata *Metadata   `json:"metadata"`
	Spec     *NodeSpec   `json:"spec"`
	Status   *NodeStatus `json:"status"`
}

// ReplicationController
//==============================================================================
type PodTemplate struct {
	Metadata *Metadata `json:"metadata"`
	Spec     *PodSpec  `json:"spec"`
}

type ReplicationControllerSpec struct {
	Selector map[string]string `json:"selector"`
	Replicas int               `json:"replicas"`
	Template *PodTemplate      `json:"template"`
}

type ReplicationControllerStatus struct {
	Replicas int `json:"replicas"`
}

type ReplicationController struct {
	*ResourceDefinition
	Metadata *Metadata                    `json:"metadata"`
	Spec     *ReplicationControllerSpec   `json:"spec"`
	Status   *ReplicationControllerStatus `json:"status,omitempty"`
}

type ReplicationControllerList struct {
	Items []*ReplicationController `json:"items"`
}

// Pod
//==============================================================================
type AwsElasticBlockStore struct {
	VolumeID string `json:"volumeID"`
	FSType   string `json:"fsType"`
}

type Volume struct {
	Name                 string                `json:"name"`
	AwsElasticBlockStore *AwsElasticBlockStore `json:"awsElasticBlockStore"`
}

type VolumeMount struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
}

type ResourceValues struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}

type Resources struct {
	Limits   *ResourceValues `json:"limits"`
	Requests *ResourceValues `json:"requests"`
}

type ContainerPort struct {
	Name          string `json:"name,omitempty"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol,omitempty"`
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Container struct {
	Name         string           `json:"name"`
	Image        string           `json:"image"`
	Resources    *Resources       `json:"resources"`
	Ports        []*ContainerPort `json:"ports"`
	VolumeMounts []*VolumeMount   `json:"volumeMounts"`
	Env          []*EnvVar        `json:"env"`
}

type ImagePullSecret struct {
	Name string `json:"name"`
}

type PodSpec struct {
	Volumes                       []*Volume          `json:"volumes"`
	Containers                    []*Container       `json:"containers"`
	ImagePullSecrets              []*ImagePullSecret `json:"imagePullSecrets"`
	TerminationGracePeriodSeconds int                `json:"terminationGracePeriodSeconds"`
}

type PodStatus struct {
	Phase string `json:"phase"`
}

type Pod struct {
	*ResourceDefinition
	Metadata *Metadata  `json:"metadata"`
	Spec     *PodSpec   `json:"spec"`
	Status   *PodStatus `json:"status"`
}

type PodList struct {
	Items []*Pod
}

// Service
//==============================================================================
type ServicePort struct {
	Name       string `json:"name"`
	Port       int    `json:"port"`
	Protocol   string `json:"protocol,omitempty"`
	NodePort   int    `json:"nodePort,omitempty"`
	TargetPort int    `json:"targetPort,omitempty"`
}

type ServiceSpec struct {
	Type     string            `json:"type,omitempty"`
	Selector map[string]string `json:"selector"`
	Ports    []*ServicePort    `json:"ports"`

	ClusterIP string `json:"clusterIP,omitempty"`
}

type Service struct {
	*ResourceDefinition
	Metadata *Metadata    `json:"metadata"`
	Spec     *ServiceSpec `json:"spec"`
	// Status   *ServiceStatus `json:"status"`
}

type ServiceList struct {
	Items []*Service `json:"items"`
}

// Secret
//==============================================================================
type Secret struct {
	*ResourceDefinition
	Metadata *Metadata         `json:"metadata"`
	Type     string            `json:"type"`
	Data     map[string]string `json:"data"`
}

type SecretList struct {
	Items []*Secret `json:"items"`
}

// Event
//==============================================================================
type Source struct {
	Host string `json:"host"`
}

type Event struct {
	*ResourceDefinition
	Message string  `json:"message"`
	Count   int     `json:"count"`
	Source  *Source `json:"source"`
}

type EventList struct {
	Items []*Event `json:"items"`
}

// TODO not sure if this should be in the types file.. related to queries, but is a Kube-specific thing
type QueryParams struct {
	LabelSelector string `url:"labelSelector"`
}
