package guber

// func (node *Node) IsReady() bool {
// 	var ready bool
// 	for _, cond := range node.Status.Conditions {
// 		if cond.Type == "Ready" && cond.Status == "True" {
// 			ready = true
// 			break
// 		}
// 	}
// 	return ready
// }
//
// func (node *Node) Cores() int {
// 	cores, err := strconv.Atoi(node.Status.Capacity.CPU)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return cores
// }
//
// func (node *Node) RamGB() float64 {
// 	mem := node.Status.Capacity.Memory
// 	memInt, err := strconv.ParseFloat(mem[:len(mem)-2], 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var ramGB float64
// 	switch mem[len(mem)-2:] {
// 	case "Ki":
// 		ramGB = memInt / 1048576
// 	default:
// 		panic(fmt.Sprintf("Could not parse node memory string \"%s\"", mem))
// 	}
// 	return ramGB
// }
//
// func (node *Node) RunningPods() []*Pod {
// 	query := fmt.Sprintf("fieldSelector=spec.nodeName=%s,status.phase=Running", node.Name())
// 	return node.Client.Pods(query)
// }
//
// func (node *Node) CreatedAt() time.Time {
// 	layout := "2006-01-02T15:04:05Z"
// 	createdAt, err := time.Parse(layout, node.Metadata.CreationTimestamp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return createdAt
// }
//
// func (node *Node) IsOlderThan(duration time.Duration) bool {
// 	return time.Since(node.CreatedAt()) > duration
// }
//
// func (node *Node) IsAlone() bool {
// 	return len(node.Client.Nodes("")) == 1
// }
//
// type byRAM []*Node
//
// func (r byRAM) Len() int           { return len(r) }
// func (r byRAM) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
// func (r byRAM) Less(i, j int) bool { return r[i].RamGB() < r[j].RamGB() }
//
// func (node *Node) IsSmallest() bool {
// 	nodes := node.Client.Nodes("")
// 	sort.Sort(byRAM(nodes))
// 	return node.RamGB() == nodes[0].RamGB()
// }
//
// func (node *Node) HasPodsWithVolumes() bool {
// 	for _, pod := range node.RunningPods() {
// 		if pod.NumExternalVolumes() > 0 {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func (node *Node) Destroy() {
// 	node.Client.Request("DELETE", fmt.Sprintf("nodes/%s", node.Name()), nil)
// }
