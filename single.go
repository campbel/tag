package tag

import (
	"encoding/json"
	"log"
)

// Single represents a collection of values
// denotes single as it contains only a single value for each entry
// compared to doubles key-value and triples namespace-key-value
type Single []string

// NewSingleFromJSON creates a new Single from the JSON bytes
func NewSingleFromJSON(bytes []byte) Single {
	var single Single

	err := json.Unmarshal(bytes, &single)
	if err != nil {
		log.Println("error creating single from JSON")
		return Single{}
	}

	return single
}

// Add a value to the single, returning the new copy
func (single Single) Add(val string) Single {
	for _, s := range single {
		if s == val {
			return single
		}
	}

	return append(single, val)
}

// Append a single to the end of this single
func (single Single) Append(other Single) Single {
	return append(single, other...)
}

// Match determines if the given value matches the single
func (single Single) Match(val string) bool {
	for _, str := range single {
		if str == val {
			return true
		}
	}
	return false
}

func (single Single) arrayAffinity(vals []string) int {
	count := 0

	for _, val := range vals {
		if single.Match(val) {
			count++
		}
	}

	return count
}

func (single Single) affinity(vals ...string) int {
	return single.arrayAffinity(vals)
}

// Affinity determines the affinity score of the two singles
func (single Single) Affinity(val Single) int {
	return single.arrayAffinity(val)
}
