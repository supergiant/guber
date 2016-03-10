package model

type AwsElasticsBlockStore struct {
	VolumeID string `json:"volumeID"`
}

type Volume struct {
	AwsElasticsBlockStore `json:"awsElasticBlockStore"`
}

type ResourceValues struct {
	Memory string `json:"memory"`
	CPU    string `json:"memory"`
}

type Resources struct {
	Limits   ResourceValues `json:"limits"`
	Requests ResourceValues `json:"requests"`
}

type Container struct {
	Name      string `json:"name"`
	Image     string `json:"image"`
	Resources `json:"resources"`
	// VolumeMounts []VolumeMount `json:"volumeMounts"`
}

type PodSpec struct {
	Volumes    []Volume    `json:"volumes"`
	Containers []Container `json:"containers"`
}

type Pod struct {
	Metadata `json:"metadata"`
	Spec     PodSpec `json:"spec"`
}
