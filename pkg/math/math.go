package math

func Sum(numbers []int64) int64 {
	var sum int64
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func Average(numbers []float64) float64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}

	avg := float64(sum) / float64(len(numbers))

	return avg
}
