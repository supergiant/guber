package kubernetes

type Pod struct {
	Metadata struct {
		Name   string
		Labels map[string]string
	}
	Spec struct {
		Volumes []struct {
			AwsElasticBlockStore struct {
				VolumeID string
			}
		}
		Containers []struct {
			Resources struct {
				Limits struct {
					Memory string
					CPU    string
				}
				Requests struct {
					Memory string
					CPU    string
				}
			}
		}
	}
}

func (pod *Pod) Name() string {
	return pod.Metadata.Name
}

// NOTE this will return either the limit, or the request if there is no limit
func (pod *Pod) Cores() float64 {
	var cores float64
	for _, container := range pod.Spec.Containers {
		val := ""
		if container.Resources.Limits.CPU != "" {
			val = container.Resources.Limits.CPU
		} else if container.Resources.Requests.CPU != "" {
			val = container.Resources.Requests.CPU
		} else {
			return 0.0
		}
		cores += ParseCores(val)
	}
	return cores
}

func (pod *Pod) RamGB() float64 {
	var ram float64
	for _, container := range pod.Spec.Containers {
		val := ""
		if container.Resources.Limits.Memory != "" {
			val = container.Resources.Limits.Memory
		} else if container.Resources.Requests.Memory != "" {
			val = container.Resources.Requests.Memory
		} else {
			return 0.0
		}
		ram += ParseRAM(val)
	}
	return ram
}

func (pod *Pod) NumExternalVolumes() int {
	vols := 0
	for _, volume := range pod.Spec.Volumes {
		if volume.AwsElasticBlockStore.VolumeID != "" {
			vols++
		}
	}
	return vols
}
