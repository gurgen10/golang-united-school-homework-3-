package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0.0
	var len = 0
	for _, val := range input {
		if val != 0 {
			len += 1
			sum += val
		}
	}
	result = sum / float32(len)
	
	return result
}
