package model

type ServicePort struct {
	Name       string `json:"name"`
	Protocol   string `json:"protocol"`
	Port       int    `json:"port"`
	NodePort   int    `json:"nodePort"`
	TargetPort int    `json:"targetPort"`
}

type ServiceSpec struct {
	Type     string            `json:"type"`
	Selector map[string]string `json:"selector"`
	Ports    []*ServicePort    `json:"ports"`
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
