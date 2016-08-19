package guber

// Common
//==============================================================================
type ResourceDefinition struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
}

type Metadata struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	Labels            map[string]string `json:"labels,omitempty"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
}

// Namespace
//==============================================================================
type Namespace struct {
	collection *Namespaces
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

type NodeAddress struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}

type NodeStatus struct {
	Capacity   *NodeStatusCapacity    `json:"capacity"`
	Conditions []*NodeStatusCondition `json:"conditions"`
	Addresses  []*NodeAddress         `json:"addresses"`
}

type Node struct {
	collection *Nodes
	*ResourceDefinition

	Metadata *Metadata   `json:"metadata"`
	Spec     *NodeSpec   `json:"spec"`
	Status   *NodeStatus `json:"status"`
}

type NodeList struct {
	Items []*Node `json:"items"`
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
	collection *ReplicationControllers
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
	Memory string `json:"memory,omitempty"`
	CPU    string `json:"cpu,omitempty"`
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

type SecurityContext struct {
	Privileged bool `json:"privileged"`
}

type Container struct {
	Name            string           `json:"name"`
	Image           string           `json:"image"`
	Command         []string         `json:"command"`
	Resources       *Resources       `json:"resources"`
	Ports           []*ContainerPort `json:"ports"`
	VolumeMounts    []*VolumeMount   `json:"volumeMounts"`
	Env             []*EnvVar        `json:"env"`
	SecurityContext *SecurityContext `json:"securityContext"`
	ImagePullPolicy string           `json:"imagePullPolicy"`
}

type ImagePullSecret struct {
	Name string `json:"name"`
}

type PodSpec struct {
	Volumes                       []*Volume          `json:"volumes"`
	Containers                    []*Container       `json:"containers"`
	ImagePullSecrets              []*ImagePullSecret `json:"imagePullSecrets"`
	TerminationGracePeriodSeconds int                `json:"terminationGracePeriodSeconds"`
	RestartPolicy                 string             `json:"restartPolicy"`
	NodeName                      string             `json:"nodeName"`
}

type ContainerStateRunning struct {
	StartedAt string `json:"startedAt"` // TODO should be time type
}

type ContainerStateTerminated struct {
	ExitCode   int    `json:"exitcode"`
	StartedAt  string `json:"startedAt"`  // TODO should be time type
	FinishedAt string `json:"finishedAt"` // TODO should be time type
	Reason     string `json:"reason"`
}

type ContainerState struct {
	Running    *ContainerStateRunning    `json:"running"`
	Terminated *ContainerStateTerminated `json:"terminated"`
}

type ContainerStatus struct {
	ContainerID  string          `json:"containerID"`
	Image        string          `json:"image"`
	ImageID      string          `json:"imageID"`
	Name         string          `json:"name"`
	Ready        bool            `json:"ready"`
	RestartCount int             `json:"restartCount"`
	State        *ContainerState `json:"state"`
	LastState    *ContainerState `json:"state"`
}

type PodStatusCondition struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}

type PodStatus struct {
	Phase             string                `json:"phase"`
	Conditions        []*PodStatusCondition `json:"conditions"`
	ContainerStatuses []*ContainerStatus    `json:"containerStatuses"`
}

type Pod struct {
	collection *Pods
	*ResourceDefinition

	Metadata *Metadata  `json:"metadata"`
	Spec     *PodSpec   `json:"spec"`
	Status   *PodStatus `json:"status"`
}

type PodList struct {
	Items []*Pod `json:"items"`
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
	collection *Services
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
	collection *Secrets
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
	collection *Events
	*ResourceDefinition

	Metadata *Metadata `json:"metadata"`
	Message  string    `json:"message"`
	Count    int       `json:"count"`
	Source   *Source   `json:"source"`
}

type EventList struct {
	Items []*Event `json:"items"`
}

// TODO not sure if this should be in the types file.. related to queries, but is a Kube-specific thing
type QueryParams struct {
	LabelSelector string
	FieldSelector string
}

type HeapsterStatMetric struct {
	Average    int64 `json:"average"`
	Percentile int64 `json:"percentile"`
	Max        int64 `json:"max"`
}

type HeapsterStatPeriods struct {
	Minute *HeapsterStatMetric `json:"minute"`
	Hour   *HeapsterStatMetric `json:"hour"`
	Day    *HeapsterStatMetric `json:"day"`
}

type HeapsterStats struct {
	Uptime int `json:"uptime"`
	Stats  map[string]*HeapsterStatPeriods
}
