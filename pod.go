package guber

type AwsElasticsBlockStore struct {
	VolumeID string `json:"volumeID"`
}

type Volume struct {
	AwsElasticsBlockStore *AwsElasticsBlockStore `json:"awsElasticBlockStore"`
}

type VolumeMount struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
}

type ResourceValues struct {
	Memory string `json:"memory"`
	CPU    string `json:"memory"`
}

type Resources struct {
	Limits   *ResourceValues `json:"limits"`
	Requests *ResourceValues `json:"requests"`
}

type ContainerPort struct {
	Name          string `json:"name"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
}

type Container struct {
	Name         string           `json:"name"`
	Image        string           `json:"image"`
	Resources    *Resources       `json:"resources"`
	Ports        []*ContainerPort `json:"ports"`
	VolumeMounts []*VolumeMount   `json:"volumeMounts"`
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

type Pod struct {
	*ResourceDefinition
	Metadata *Metadata `json:"metadata"`
	Spec     *PodSpec  `json:"spec"`
}
