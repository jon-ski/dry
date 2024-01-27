package dry

// Equal returns true if the slices a and b are equal.
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Copy returns a copy of the slice s.
func Copy[T any](s []T) []T {
	c := make([]T, len(s))
	copy(c, s)
	return c
}

// Contains returns true if the slice s contains the value v.
func Contains[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

// Filter returns a new slice containing only the elements of s for which f returns true.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

// Map returns a new slice containing the results of applying f to each element of s.
func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

// Reduce returns a single value by applying f to each element of s and the previous result.
func Reduce[T, U any](s []T, f func(U, T) U, initial U) U {
	r := initial
	for _, e := range s {
		r = f(r, e)
	}
	return r
}

// Reverse returns a new slice with the elements of s in reverse order.
func Reverse[T any](s []T) []T {
	r := make([]T, len(s))
	for i, e := range s {
		r[len(s)-1-i] = e
	}
	return r
}

// Unique returns a new slice with duplicate elements removed.
func Unique[T comparable](s []T) []T {
	var r []T
	for _, e := range s {
		if !Contains(r, e) {
			r = append(r, e)
		}
	}
	return r
}

// Flatten returns a new slice with all elements of s flattened into a single slice.
func Flatten[T any](s [][]T) []T {
	var r []T
	for _, e := range s {
		r = append(r, e...)
	}
	return r
}

// Zip returns a new slice with the elements of s and t interleaved.
func Zip[T any](s, t []T) []T {
	var r []T
	for i := 0; i < len(s) || i < len(t); i++ {
		if i < len(s) {
			r = append(r, s[i])
		}
		if i < len(t) {
			r = append(r, t[i])
		}
	}
	return r
}
