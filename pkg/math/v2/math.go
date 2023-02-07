package v2

import (
	"errors"
	"reflect"
	"sync"
)

func Sum(numbers []int64, ch chan<- int64) {
	var sum int64
	for _, n := range numbers {
		sum += n
	}

	ch <- sum
}

func Average(numbers []float64, ch chan<- float64) {
	var mtx sync.Mutex
	var wg sync.WaitGroup

	var sum int64 = 0

	for _, n := range numbers {
		wg.Add(1)

		go func(x float64) {
			mtx.Lock()
			sum += int64(x)
			mtx.Unlock()
			wg.Done()
		}(n)
	}

	wg.Wait()
	avg := float64(sum) / float64(len(numbers))

	ch <- avg
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
