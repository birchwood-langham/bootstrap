package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func ExampleState() {
	s := NewStateStore().
		WithString("httpAddr", ":8080").
		WithInt64("buffer", 1000)

	fmt.Printf("HttpAddress: %s, Buffer: %d", s.String("httpAddr", "localhost:80"), s.Int64("buffer", 100))

	// Output: HttpAddress: :8080, Buffer: 1000
}

func TestState_String(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := "test"
	key := "key"
	fallback := "fallback"
	dummy := "dummy"

	s = s.WithString(key, want)

	gotValue := s.String(key, fallback)

	if gotValue != want {
		t.Errorf("WithString want: %s, got: %s", want, gotValue)
	}

	gotFallback := s.String(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithString fallback want: %s, got: %s", fallback, gotFallback)
	}
}

func TestState_Int8(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := int8(100)
	key := "key"
	fallback := int8(101)
	dummy := "dummy"

	s = s.WithInt8(key, want)

	gotValue := s.Int8(key, fallback)

	if gotValue != want {
		t.Errorf("WithInt8 want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Int8(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithInt8 fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func TestState_Int16(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := int16(100)
	key := "key"
	fallback := int16(101)
	dummy := "dummy"

	s = s.WithInt16(key, want)

	gotValue := s.Int16(key, fallback)

	if gotValue != want {
		t.Errorf("WithInt16 want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Int16(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithInt16 fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func TestState_Int32(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := int32(100)
	key := "key"
	fallback := int32(101)
	dummy := "dummy"

	s = s.WithInt32(key, want)

	gotValue := s.Int32(key, fallback)

	if gotValue != want {
		t.Errorf("WithInt32 want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Int32(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithInt32 fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func TestState_Int64(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := int64(100)
	key := "key"
	fallback := int64(101)
	dummy := "dummy"

	s = s.WithInt64(key, want)

	gotValue := s.Int64(key, fallback)

	if gotValue != want {
		t.Errorf("WithInt8 want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Int64(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithInt64 fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func TestState_Uint8(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := uint8(100)
	key := "key"
	fallback := uint8(101)
	dummy := "dummy"

	s = s.WithUint8(key, want)

	gotValue := s.Uint8(key, fallback)

	if gotValue != want {
		t.Errorf("WithUint8 want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Uint8(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithUint8 fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func TestState_Uint16(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := uint16(100)
	key := "key"
	fallback := uint16(101)
	dummy := "dummy"

	s = s.WithUint16(key, want)

	gotValue := s.Uint16(key, fallback)

	if gotValue != want {
		t.Errorf("WithUint16 want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Uint16(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithUint16 fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func TestState_Uint32(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := uint32(100)
	key := "key"
	fallback := uint32(101)
	dummy := "dummy"

	s = s.WithUint32(key, want)

	gotValue := s.Uint32(key, fallback)

	if gotValue != want {
		t.Errorf("WithUint32 want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Uint32(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithUint32 fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func TestState_Uint64(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := uint64(100)
	key := "key"
	fallback := uint64(101)
	dummy := "dummy"

	s = s.WithUint64(key, want)

	gotValue := s.Uint64(key, fallback)

	if gotValue != want {
		t.Errorf("WithUint64 want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Uint64(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithUint64 fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func TestState_Float32(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := float32(100)
	key := "key"
	fallback := float32(101)
	dummy := "dummy"

	s = s.WithFloat32(key, want)

	gotValue := s.Float32(key, fallback)

	if gotValue != want {
		t.Errorf("WithFloat32 want: %f, got: %f", want, gotValue)
	}

	gotFallback := s.Float32(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithFloat32 fallback want: %f, got: %f", fallback, gotFallback)
	}
}

func TestState_Float64(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := float64(100)
	key := "key"
	fallback := float64(101)
	dummy := "dummy"

	s = s.WithFloat64(key, want)

	gotValue := s.Float64(key, fallback)

	if gotValue != want {
		t.Errorf("WithFloat64 want: %f, got: %f", want, gotValue)
	}

	gotFallback := s.Float64(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithFloat64 fallback want: %f, got: %f", fallback, gotFallback)
	}
}

func TestState_Rune(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := rune(100)
	key := "key"
	fallback := rune(101)
	dummy := "dummy"

	s = s.WithRune(key, want)

	gotValue := s.Rune(key, fallback)

	if gotValue != want {
		t.Errorf("WithRune want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Rune(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithRune fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func TestState_Time(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := time.Unix(1604484825, 0)
	key := "key"
	fallback := time.Unix(0, 0)
	dummy := "dummy"

	s = s.WithTime(key, want)

	gotValue := s.Time(key, fallback)

	if gotValue != want {
		t.Errorf("WithTime want: %v, got: %v", want, gotValue)
	}

	gotFallback := s.Time(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithTime fallback want: %v, got: %v", fallback, gotFallback)
	}
}

func TestState_Duration(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := time.Hour
	key := "key"
	fallback := time.Minute
	dummy := "dummy"

	s = s.WithDuration(key, want)

	gotValue := s.Duration(key, fallback)

	if gotValue != want {
		t.Errorf("WithDuration want: %d, got: %d", want, gotValue)
	}

	gotFallback := s.Duration(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithDuration fallback want: %d, got: %d", fallback, gotFallback)
	}
}

func compareByteArray(a []byte, b []byte) bool {
	if a == nil && b == nil {
		return true
	}

	if (a == nil && b != nil) || (b == nil && a != nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i, j := range a {
		if j != b[i] {
			return false
		}
	}

	return true
}

func TestState_Bytes(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := []byte("Test")
	key := "key"
	fallback := []byte("fallback")
	dummy := "dummy"

	s = s.WithBytes(key, want)

	gotValue := s.Bytes(key, fallback)

	if !compareByteArray(gotValue, want) {
		t.Errorf("WithBytes want: %s, got: %s", string(want), string(gotValue))
	}

	gotFallback := s.Bytes(dummy, fallback)

	if !compareByteArray(gotFallback, fallback) {
		t.Errorf("WithBytes fallback want: %s, got: %s", string(fallback), string(gotFallback))
	}
}

func TestState_Any(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	type Person struct {
		FirstName string
		LastName  string
		Age       uint8
	}

	want := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       36,
	}

	key := "key"
	fallback := Person{
		FirstName: "Anne",
		LastName:  "Other",
		Age:       28,
	}

	dummy := "dummy"

	s = s.WithValue(key, want)

	gotValue := s.Value(key, fallback)

	if gotValue != want {
		t.Errorf("WithValue want: %v, got: %v", want, gotValue.(Person))
	}

	gotFallback := s.Value(dummy, fallback)

	if gotFallback != fallback {
		t.Errorf("WithInt fallback want: %v, got: %v", fallback, gotFallback.(Person))
	}
}

func TestState_Decimal(t *testing.T) {
	var s StateStore
	s = NewStateStore()

	want := decimal.NewFromInt(100)
	key := "key"
	fallback := decimal.NewFromInt(101)
	dummy := "dummy"

	s = s.WithDecimal(key, want)

	gotValue := s.Decimal(key, fallback)

	if !gotValue.Equal(want) {
		t.Errorf("WithDecimal want: %s, got: %s", want.StringFixed(2), gotValue.StringFixed(2))
	}

	gotFallback := s.Decimal(dummy, fallback)

	if !gotFallback.Equal(fallback) {
		t.Errorf("WithDecimal fallback want: %s, got: %s", fallback.StringFixed(2), gotFallback.StringFixed(2))
	}
}
