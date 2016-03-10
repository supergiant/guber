package guber

// // Returns GB
// func ParseRAM(mem string) float64 {
// 	memInt, err := strconv.ParseFloat(mem[:len(mem)-2], 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var ramGB float64
// 	switch mem[len(mem)-2:] {
// 	case "Ki":
// 		ramGB = memInt / 1048576
// 	case "Mi":
// 		ramGB = memInt / 1024
// 	case "Gi":
// 		ramGB = memInt
// 	default:
// 		panic(fmt.Sprintf("Could not parse memory string \"%s\"", mem))
// 	}
// 	return ramGB
// }
//
// func ParseCores(cpu string) float64 {
// 	var cores float64
// 	switch cpu[len(cpu)-1:] {
// 	case "m":
// 		coresStr := cpu[:len(cpu)-2]
// 		coresFloat, err := strconv.ParseFloat(coresStr, 64)
// 		if err != nil {
// 			panic(err)
// 		}
// 		cores = coresFloat / 1000
// 	default:
// 		coresFloat, err := strconv.ParseFloat(cpu, 64)
// 		if err != nil {
// 			panic(err)
// 		}
// 		cores = coresFloat
// 	}
// 	return cores
// }
