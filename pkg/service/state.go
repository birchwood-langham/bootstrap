package service

import (
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

// StateStore provides an interface for implementing a state store for your application
// The default state store is an in-memory state store accessible by the service
// By implementing this interface users can implement state stores that can utilise
// different services such as Redis, Consul, Etcd etc. for storing their state
type StateStore interface {
	WithString(name string, value string) StateStore
	WithInt8(name string, value int8) StateStore
	WithInt16(name string, value int16) StateStore
	WithInt32(name string, value int32) StateStore
	WithInt64(name string, value int64) StateStore
	WithUint8(name string, value uint8) StateStore
	WithUint16(name string, value uint16) StateStore
	WithUint32(name string, value uint32) StateStore
	WithUint64(name string, value uint64) StateStore
	WithFloat32(name string, value float32) StateStore
	WithFloat64(name string, value float64) StateStore
	WithRune(name string, value rune) StateStore
	WithTime(name string, value time.Time) StateStore
	WithDuration(name string, value time.Duration) StateStore
	WithBytes(name string, value []byte) StateStore
	WithValue(name string, value interface{}) StateStore
	WithDecimal(name string, value decimal.Decimal) StateStore
	String(name string, fallback string) string
	Int8(name string, fallback int8) int8
	Int16(name string, fallback int16) int16
	Int32(name string, fallback int32) int32
	Int64(name string, fallback int64) int64
	Uint8(name string, fallback uint8) uint8
	Uint16(name string, fallback uint16) uint16
	Uint32(name string, fallback uint32) uint32
	Uint64(name string, fallback uint64) uint64
	Float32(name string, fallback float32) float32
	Float64(name string, fallback float64) float64
	Rune(name string, fallback rune) rune
	Time(name string, fallback time.Time) time.Time
	Duration(name string, fallback time.Duration) time.Duration
	Bytes(name string, fallback []byte) []byte
	Value(name string, fallback interface{}) interface{}
	Decimal(name string, fallback decimal.Decimal) decimal.Decimal
}

type state struct {
	mu         sync.RWMutex
	strings    map[string]string
	i8         map[string]int8
	i16        map[string]int16
	i32        map[string]int32
	i64        map[string]int64
	u8         map[string]uint8
	u16        map[string]uint16
	u32        map[string]uint32
	u64        map[string]uint64
	f32        map[string]float32
	f64        map[string]float64
	runes      map[string]rune
	timestamps map[string]time.Time
	durations  map[string]time.Duration
	bs         map[string][]byte
	any        map[string]interface{}
	dec        map[string]decimal.Decimal
}

// NewStateStore creates an in-memory state store for you to save state information for your application
func NewStateStore() *state {
	return &state{
		strings:    make(map[string]string),
		i8:         make(map[string]int8),
		i16:        make(map[string]int16),
		i32:        make(map[string]int32),
		i64:        make(map[string]int64),
		u8:         make(map[string]uint8),
		u16:        make(map[string]uint16),
		u32:        make(map[string]uint32),
		u64:        make(map[string]uint64),
		f32:        make(map[string]float32),
		f64:        make(map[string]float64),
		runes:      make(map[string]rune),
		timestamps: make(map[string]time.Time),
		durations:  make(map[string]time.Duration),
		bs:         make(map[string][]byte),
		any:        make(map[string]interface{}),
		dec:        make(map[string]decimal.Decimal),
	}
}

// WithString sets a string value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithString(name string, value string) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.strings[name] = value
	return s
}

// WithInt8 sets a int8 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithInt8(name string, value int8) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.i8[name] = value
	return s
}

// WithInt16 sets a int16 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithInt16(name string, value int16) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.i16[name] = value
	return s
}

// WithInt32 sets a int32 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithInt32(name string, value int32) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.i32[name] = value
	return s
}

// WithInt64 sets a int64 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithInt64(name string, value int64) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.i64[name] = value
	return s
}

// WithUint8 sets a uint8 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithUint8(name string, value uint8) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.u8[name] = value
	return s
}

// WithUint16 sets a uint16 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithUint16(name string, value uint16) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.u16[name] = value
	return s
}

// WithUint32 sets a uint32 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithUint32(name string, value uint32) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.u32[name] = value
	return s
}

// WithUint64 sets a uint64 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithUint64(name string, value uint64) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.u64[name] = value
	return s
}

// WithFloat32 sets a float32 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithFloat32(name string, value float32) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.f32[name] = value
	return s
}

// WithFloat64 sets a float64 value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithFloat64(name string, value float64) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.f64[name] = value
	return s
}

// WithRune sets a rune value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithRune(name string, value rune) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.runes[name] = value
	return s
}

// WithTime sets a time value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithTime(name string, value time.Time) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.timestamps[name] = value
	return s
}

// WithDuration sets a duration value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithDuration(name string, value time.Duration) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.durations[name] = value
	return s
}

// WithValue sets a value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithValue(name string, value interface{}) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.any[name] = value
	return s
}

// WithBytes sets a bytes value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithBytes(name string, value []byte) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.bs[name] = value
	return s
}

// WithDecimal sets a decimal value for the given property.
// If the property already exists, the new property will be
// replaced with the new value
func (s *state) WithDecimal(name string, value decimal.Decimal) StateStore {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.dec[name] = value
	return s
}

// String returns the string value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) String(name string, fallback string) (value string) {
	var ok bool
	if value, ok = s.strings[name]; !ok {
		value = fallback
	}

	return
}

// Int8 returns the int8 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Int8(name string, fallback int8) (value int8) {
	var ok bool
	if value, ok = s.i8[name]; !ok {
		value = fallback
	}

	return
}

// Int16 returns the int16 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Int16(name string, fallback int16) (value int16) {
	var ok bool
	if value, ok = s.i16[name]; !ok {
		value = fallback
	}

	return
}

// Int32 returns the int32 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Int32(name string, fallback int32) (value int32) {
	var ok bool
	if value, ok = s.i32[name]; !ok {
		value = fallback
	}

	return
}

// Int64 returns the int64 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Int64(name string, fallback int64) (value int64) {
	var ok bool
	if value, ok = s.i64[name]; !ok {
		value = fallback
	}

	return
}

// Uint8 returns the uint8 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Uint8(name string, fallback uint8) (value uint8) {
	var ok bool
	if value, ok = s.u8[name]; !ok {
		value = fallback
	}

	return
}

// Uint16 returns the uint16 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Uint16(name string, fallback uint16) (value uint16) {
	var ok bool
	if value, ok = s.u16[name]; !ok {
		value = fallback
	}

	return
}

// Uint32 returns the uint32 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Uint32(name string, fallback uint32) (value uint32) {
	var ok bool
	if value, ok = s.u32[name]; !ok {
		value = fallback
	}

	return
}

// Uint64 returns the uint64 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Uint64(name string, fallback uint64) (value uint64) {
	var ok bool
	if value, ok = s.u64[name]; !ok {
		value = fallback
	}

	return
}

// Float32 returns the float32 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Float32(name string, fallback float32) (value float32) {
	var ok bool
	if value, ok = s.f32[name]; !ok {
		value = fallback
	}

	return
}

// Float64 returns the float64 value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Float64(name string, fallback float64) (value float64) {
	var ok bool
	if value, ok = s.f64[name]; !ok {
		value = fallback
	}

	return
}

// Time returns the time value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Time(name string, fallback time.Time) (value time.Time) {
	var ok bool
	if value, ok = s.timestamps[name]; !ok {
		value = fallback
	}

	return
}

// Duration returns the duration value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Duration(name string, fallback time.Duration) (value time.Duration) {
	var ok bool
	if value, ok = s.durations[name]; !ok {
		value = fallback
	}

	return
}

// Value returns the value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Value(name string, fallback interface{}) (value interface{}) {
	var ok bool
	if value, ok = s.any[name]; !ok {
		value = fallback
	}

	return
}

// Rune returns the rune value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Rune(name string, fallback rune) (value rune) {
	var ok bool
	if value, ok = s.runes[name]; !ok {
		value = fallback
	}

	return
}

// Bytes returns the byte array value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Bytes(name string, fallback []byte) (value []byte) {
	var ok bool
	if value, ok = s.bs[name]; !ok {
		value = fallback
	}

	return
}

// Decimal returns the decimal value of the property with the given name
// if the name is not found, the fallback value is returned instead
func (s *state) Decimal(name string, fallback decimal.Decimal) (value decimal.Decimal) {
	var ok bool
	if value, ok = s.dec[name]; !ok {
		value = fallback
	}

	return
}
