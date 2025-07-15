package gftools

import (
	"regexp"
)

// Equals returns a predicate (paramFunc) that checks if a value equals a given target.
func Equals[T comparable](target T) func(T) bool {
	return func(x T) bool {
		return x == target
	}
}

// GreaterThan returns a predicate (paramFunc) that checks if a value is greater than n.
func GreaterThan(n int) func(int) bool {
	return func(x int) bool {
		return x > n
	}
}

// GreaterOrEqual returns a predicate (paramFunc) that checks if a value is greater than or equal to n.
func GreaterOrEqual(n int) func(int) bool {
	return func(x int) bool {
		return x >= n
	}
}

// LessThan returns a predicate (paramFunc) that checks if a int value is less than n.
func LessThan(n int) func(int) bool {
	return func(x int) bool {
		return x < n
	}
}

// LessOrEqual returns a predicate (paramFunc) that checks if a int value is less than or equal to n.
func LessOrEqual(n int) func(int) bool {
	return func(x int) bool {
		return x <= n
	}
}

// Between returns a predicate (paramFunc) that checks if a int value is between min and max 
// (inclusive).
func Between(min, max int) func(int) bool {
	return func(x int) bool {
		return x >= min && x <= max
	}
}

// IsEven returns a predicate (paramFunc) that checks if a int value is even.
func IsEven() func(int) bool {
	return func(x int) bool {
		return x % 2 == 0
	}
}

// MatchRegex returns a predicate (paramFunc) that checks if a string matches the given regex pattern.
func MatchRegex(pattern string) (func(string) bool, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return func(s string) bool {
		return re.MatchString(s)
	}, nil
}

// HasKey returns a predicate (paramFunc) that checks if a given key exists in a map.
// When calling, include the type of the map keys. Example: HasKey[string]("myKey")
func HasKey[K comparable, V any](targetKey K) func(map[K]V) bool {
	return func(m map[K]V) bool {
		_, exists := m[targetKey]
		return exists
	}
}

// HasValue returns a predicate (paramFunc) that checks if a value exists in the given map.
// When calling, include the type of the keys in the map. Example: HasValue[string](targetValue)
func HasValue[K comparable, V comparable](targetValue V) func(map[K]V) bool {
	return func(m map[K]V) bool {
		for _, v := range m {
			if v == targetValue {
				return true
			}
		}
		return false
	}
}

// HasKeyValue returns a predicate (paramFunc) that checks if the map contains the given key with the exact value. Example: HasKeyValue[string, string]("foo", "bar")
func HasKeyValue[K comparable, V comparable](key K, value V) func(map[K]V) bool {
	return func(m map[K]V) bool {
		v, ok := m[key]
		return ok && v == value
	}
}

// Contains returns a predicate (paramFunc) that checks if a string contains a given substring.
func Contains(substr string) func(string) bool {
	return func(s string) bool {
		return len(substr) > 0 && len(s) > 0 && (len(substr) <= len(s)) && (stringContains(s, substr))
	}
}

// StartsWith returns a predicate (paramFunc) that checks if a string starts with a given prefix.
func StartsWith(prefix string) func(string) bool {
	return func(s string) bool {
		return len(s) >= len(prefix) && s[:len(prefix)] == prefix
	}
}

// EndsWith returns a predicate (paramFunc) that checks if a string ends with a given suffix.
func EndsWith(suffix string) func(string) bool {
	return func(s string) bool {
		return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
	}
}

// IsZero returns a predicate (paramFunc) that checks if a value is the zero value of its type.
func IsZero[T comparable]() func(T) bool {
	var zero T
	return func(x T) bool {
		return x == zero
	}
}

// Not returns a predicate (paramFunc) that negates the result of another predicate (paramFunc).
func Not[T any](paramFunc func(T) bool) func(T) bool {
	return func(x T) bool {
		return !paramFunc(x)
	}
}

// In returns a predicate (paramFunc) that checks if a value exists in a slice.
func In[T comparable](list []T) func(T) bool {
	return func(x T) bool {
		for _, v := range list {
			if v == x {
				return true
			}
		}
		return false
	}
}

// stringContains is a helper for substring checking.
func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
