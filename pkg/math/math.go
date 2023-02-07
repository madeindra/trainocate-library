package math

import (
	"errors"
	"reflect"
)

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

func IsTooLong(list interface{}) (bool, error) {
	rt := reflect.TypeOf(list)
	if rt.Kind() != reflect.Slice {
		return false, errors.New("invalid data type")
	}

	var newList []interface{}
	newList = append(newList, list)

	if len(newList) >= 10 {
		return true, nil
	}

	return false, nil
}
