package guber

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
