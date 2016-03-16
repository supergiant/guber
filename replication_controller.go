package guber

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
