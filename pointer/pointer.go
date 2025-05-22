// Package pointer provides utilities for referencing or dereferencing values.
package pointer

// Or returns the value of the given pointer, or the fallback value.
func Or[T any](v *T, fallback T) T {
	if v == nil {
		return fallback
	}

	return *v
}

// OrDefault returns the value of the given pointer, or the default value for its type.
func OrDefault[T any](v *T) T {
	if v == nil {
		var t T

		return t
	}

	return *v
}

// To returns a pointer to the given value.
func To[T any](v T) *T {
	return &v
}

// ToUnlessDefault returns nil if the given value is the default value for its type,
// otherwise a pointer to the given value.
func ToUnlessDefault[T comparable](v T) *T {
	var def T

	if v == def {
		return nil
	}

	return &v
}

// ValuesOrDefaults returns a slice of the dereferenced values of the given pointers,
// replacing nils with the default value for the type.
func ValuesOrDefaults[T any](values []*T) []T {
	result := make([]T, 0, len(values))

	for _, value := range values {
		deref := OrDefault(value)
		result = append(result, deref)
	}

	return result
}
