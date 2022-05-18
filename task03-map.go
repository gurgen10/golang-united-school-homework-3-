package homework

func sortMapValues(input map[int]string) (result []string) {
	var idxs []int
	for i := range input {
		idxs = append(idxs, i)
	}
	sort.Ints(idxs)
	for _, val := range idxs {
		result = append(result,input[val])
	}
	return result
}
