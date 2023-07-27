package main

import (
	"fmt"
	"sync"
	"time"

	"honnef.co/go/tools/lintcmd/cache"
)

// cacheValue is a simple wrapper of the cache data to support internal TTL implementation.
type cacheValue struct {
	data []byte
	ttl  time.Time
}

// newCacheValue is the default constructor for the cacheValue struct.
func newCacheValue(data []byte, ttl time.Time) *cacheValue {
	return &cacheValue{data, ttl}
}

// Repository provides key value storage implementation on memory.
type Repository[E cache.Cacheable] struct {
	options cache.DriverOptions
	dic     map[string]*cacheValue
}

// NewRepository is the default constructor for the Repository[E] struct.
func NewRepository[E cache.Cacheable](options DriverOptions) Repository[E] {
	return Repository[E]{options, make(map[string]*cacheValue, 0), sync.RWMutex{}}
}

// Encode provides serialization and compression for any write cache operation.
func (r *Repository[E]) Encode(entity E,
	compression compression,
	serialization Serialization[E]) ([]byte, error) {
	s, err := serialization.Serialize(entity)
	if err != nil {
		return nil, err
	}
	return compression.Encode(s)
}

// Decode provides deserialization and decompression for any read cache operation.
func (r *Repository[E]) Decode(chunk []byte,
	compression Compression,
	serialization Serialization[E]) (*E, error) {
	b, err := compression.Decode(chunk)
	if err != nil {
		return nil, err
	}
	return serialization.Deserialize(b)
}

// StandardizeKey standardizes the key format.
func (r *Repository[E]) StandardizeKey(key string) string {
	switch r.options.GetEnvironment() {
	case repo.EnvSandbox:
		return strings.ToLower(fmt.Sprintf("%s-%s", repo.EnvSandbox, key))
	case repo.EnvProduction:
		return strings.ToLower(fmt.Sprintf("%s-%s", repo.EnvProduction, key))
	default:
		return strings.ToLower(fmt.Sprintf("%s-%s", repo.EnvDevelopment, key))
	}
}

// Set inserts a key/value into the storage.
func (r *Repository[E]) Set(ctx context.Context,
	key string,
	value E,
	compression Compression,
	serialization Serialization[E],
	ttl time.Duration) error {
	key = r.StandardizeKey(key)
	b, err := r.Encode(value, compression, serialization)
	if err != nil {
		return err
	}
	r.dic[key] = newCacheValue(b, time.Now().Add(ttl))
	return nil
}

// Get retrieves a value from the storage via key.
func (r *Repository[E]) Get(ctx context.Context, key string, compression Compression, serialization Serialization[E]) (*E, error) {
	key = r.StandardizeKey(key)
	value := r.dic[key]
	if value == nil {
		return nil, nil
	}
	if time.Now().After(value.ttl) {
		return nil, nil
	}
	entity, err := r.Decode(value.data, compression, serialization)
	if err != nil {
		if _, err := r.Delete(ctx, key); err != nil {
			return nil, err
		}
	}
	return entity, nil
}

// Delete removes a key from the storage.
func (r *Repository[E]) Delete(ctx context.Context, key string) (bool, error) {
	key = r.StandardizeKey(key)
	if r.dic[key] == nil {
		return false, nil
	}
	delete(r.dic, key)
	return true, nil
}
