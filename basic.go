package sliceNum

import (
	"reflect"
)

//SumSlice return the sum of all the elements in a slice
//Only accepts int, int32, int64, float32 and float64
//
func SumSlice(slice interface{}) (total float32, err error) {
	chkSlice := reflect.TypeOf(slice)
	if chkSlice.Kind() != reflect.TypeOf([]int{}).Kind() {
		err = GeneralError{2, "Input must be a slice."}
		return
	}

	switch chkSlice.Elem() {
	case reflect.TypeOf(int(0)):
		for _, it := range slice.([]int) {
			total += float32(it)
		}
	case reflect.TypeOf(int32(0)):
		for _, it := range slice.([]int32) {
			total += float32(it)
		}
	case reflect.TypeOf(int64(0)):
		for _, it := range slice.([]int64) {
			total += float32(it)
		}
	case reflect.TypeOf(float32(0)):
		for _, it := range slice.([]float32) {
			total += it
		}
	case reflect.TypeOf(float64(0)):
		for _, it := range slice.([]float64) {
			total += float32(it)
		}
	default:
		err = GeneralError{1, "Type" + chkSlice.Elem().String() + " is not supported"}
	}
	return
}
