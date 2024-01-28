package dry

// MapEqual returns true if the maps a and b are equal.
func MapEqual[K comparable, V comparable](a, b map[K]V) bool {
	// Lengths are different, maps are not equal
	if len(a) != len(b) {
		return false
	}

	// Compare each key/value pair
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

// MapCopy returns a copy of the map m.
func MapCopy[K comparable, V any](m map[K]V) map[K]V {
	c := make(map[K]V, len(m))
	for k, v := range m {
		c[k] = v
	}
	return c
}

// MapContains returns true if the map m contains the key k.
func MapContains[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

// MapFilter returns a new map containing only the key/value pairs of m for which f returns true.
func MapFilter[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if f(k, v) {
			r[k] = v
		}
	}
	return r
}

// MapMap returns a new map containing the results of applying f to each key/value pair of m.
func MapMap[K comparable, V1 any, V2 any](m map[K]V1, f func(K, V1) V2) map[K]V2 {
	r := make(map[K]V2, len(m))
	for k, v := range m {
		r[k] = f(k, v)
	}
	return r
}

// MapReduce returns a single value by applying f to each key/value pair of m and the previous result.
func MapReduce[K comparable, V any, U any](m map[K]V, f func(U, K, V) U, initial U) U {
	r := initial
	for k, v := range m {
		r = f(r, k, v)
	}
	return r
}

// MapKeys returns a slice containing the keys of m.
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// MapValues returns a slice containing the values of m.
func MapValues[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// MapMerge returns a new map containing the key/value pairs of a and b.
// If a and b contain the same key, the value from b is used.
func MapMerge[K comparable, V any](a, b map[K]V) map[K]V {
	r := MapCopy(a)
	for k, v := range b {
		r[k] = v
	}
	return r
}

// MapMergeWith returns a new map containing the key/value pairs of a and b.
// If a and b contain the same key, the value from f is used.
func MapMergeWith[K comparable, V any](a, b map[K]V, f func(K, V, V) V) map[K]V {
	r := MapCopy(a)
	for k, v := range b {
		if _, ok := r[k]; ok {
			r[k] = f(k, r[k], v)
		} else {
			r[k] = v
		}
	}
	return r
}

// MapDelete returns a new map containing the key/value pairs of m without the key k.
func MapDelete[K comparable, V any](m map[K]V, k K) map[K]V {
	r := MapCopy(m)
	delete(r, k)
	return r
}

// MapTraversal is a function type for traversing a map.
type MapTraversal[K comparable, V any, R any] func(K, V) R

// MapEach calls f for each key/value pair of m.
func MapEach[K comparable, V any, R any](m map[K]V, f MapTraversal[K, V, R]) {
	for k, v := range m {
		f(k, v)
	}
}

// MapKeyFunc applies f to the value of m at key k.
func MapKeyFunc[K comparable, V any](m map[K]V, k K, f func(V) V) map[K]V {
	if val, ok := m[k]; ok {
		m[k] = f(val)
	}
	return m
}
