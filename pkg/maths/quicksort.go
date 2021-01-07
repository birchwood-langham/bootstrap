package maths

import (
	"errors"

	"github.com/shopspring/decimal"
)

var IndexOutOfBounds = errors.New("index out of bounds")
var EmptyArrayError = errors.New("array is empty")

// swapDecimal swaps two elements in an array of decimals
func swapDecimal(v []decimal.Decimal, i, j int) (err error) {
	lastIndex := len(v) - 1

	if i > lastIndex || j > lastIndex || i < 0 || j < 0 {
		return IndexOutOfBounds
	}

	v[j], v[i] = v[i], v[j]

	return nil
}

// partition orders the values so they are smaller to the left of our pivot point and
// larger to the right of our pivot point.
func partition(v []decimal.Decimal, low, high int) (int, error) {
	pivot := v[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if v[j].LessThan(pivot) {
			i++
			if err := swapDecimal(v, i, j); err != nil {
				return 0, err
			}
		}
	}

	if err := swapDecimal(v, i+1, high); err != nil {
		return 0, err
	}

	return i + 1, nil
}

// DecimalQuickSort performs the quick sort algorithm on an array of
// Decimal values using the last element as the pivot without destroying
// the original array.
func DecimalQuickSort(v []decimal.Decimal, low, high int) ([]decimal.Decimal, error) {
	if len(v) == 0 {
		return nil, EmptyArrayError
	}

	if len(v) == 1 {
		return v, nil
	}

	l := len(v)

	sorted := make([]decimal.Decimal, l)

	for i, val := range v {
		sorted[i] = val
	}

	if err := decimalQuickSort(sorted, low, high); err != nil {
		return nil, err
	}

	return sorted, nil
}

func decimalQuickSort(v []decimal.Decimal, low, high int) error {
	if low >= high {
		return nil
	}

	pi, err := partition(v, low, high)

	if err != nil {
		return err
	}

	// everything less than the pivot value
	if err = decimalQuickSort(v, low, pi-1); err != nil {
		return err
	}

	// everything greater than the pivot value
	if err = decimalQuickSort(v, pi+1, high); err != nil {
		return err
	}

	return nil
}
