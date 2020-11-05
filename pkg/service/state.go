package service

// StateStore provides an interface for implementing a state store for your application
// The default state store is an in-memory state store accessible by the service
// By implementing this interface users can implement state stores that can utilise
// different services such as Redis, Consul, Etcd etc. for storing their state
type StateStore interface{}
