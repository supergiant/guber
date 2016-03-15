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

type ReplicationController struct {
	*ResourceDefinition
	Metadata *Metadata                  `json:"metadata"`
	Spec     *ReplicationControllerSpec `json:"spec"`
	// Status   *ReplicationControllerStatus `json:"status"`
}

type ReplicationControllerList struct {
	Items []*ReplicationController `json:"items"`
}
