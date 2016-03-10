package model

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
	Capacity   NodeStatusCapacity    `json:"capacity"`
	Conditions []NodeStatusCondition `json:"conditions"`
}

type Node struct {
	Metadata `json:"metadata"`
	Spec     NodeSpec   `json:"spec"`
	Status   NodeStatus `json:"status"`
}
