package tag

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSingle(t *testing.T) {
	single := Single{"foo", "bar", "baz"}.Add("qux").Add("foo")
	expectedSingle := Single{"foo", "bar", "baz", "qux"}

	if len(single) != 4 {
		t.Errorf("expected single length to be %d but was %d", 4, len(single))
	}

	if !reflect.DeepEqual(single, expectedSingle) {
		t.Errorf("expected single to be equal to %v but was %v", expectedSingle, single)
	}

}

func TestAppend(t *testing.T) {
	single := Single{"foo", "bar"}.Append(Single{"baz", "qux"})
	expectedSingle := Single{"foo", "bar", "baz", "qux"}

	if len(single) != 4 {
		t.Errorf("expected single length to be %d but was %d", 4, len(single))
	}

	if !reflect.DeepEqual(single, expectedSingle) {
		t.Errorf("expected single to be equal to %v but was %v", expectedSingle, single)
	}
}

func TestMatch(t *testing.T) {
	single := Single{"foo", "bar", "baz"}

	for _, val := range single {
		if !single.Match(val) {
			t.Errorf("expected single to match %s but didn't", val)
		}
	}

	if single.Match("qux") {
		t.Error("expected single to NOT match qux but did")
	}
}

func TestAffinity(t *testing.T) {
	single := Single{"foo", "bar", "baz"}

	affinity := single.affinity("foo", "bar", "baz")
	if affinity != 3 {
		t.Errorf("expected affinity to be %d but was %d", 3, affinity)
	}

	affinity = single.affinity("foo", "bar")
	if affinity != 2 {
		t.Errorf("expected affinity to be %d but was %d", 2, affinity)
	}

	affinity = single.affinity("foo", "bar", "qux")
	if affinity != 2 {
		t.Errorf("expected affinity to be %d but was %d", 2, affinity)
	}
}

func TestNewFromJSON(t *testing.T) {

	bytes := []byte(`["foo", "bar", "baz"]`)

	single := NewSingleFromJSON(bytes)
	expectedSingle := Single{"foo", "bar", "baz"}

	if !reflect.DeepEqual(single, expectedSingle) {
		t.Errorf("expected single to be equal to %v but was %v", expectedSingle, single)
	}

}

func TestJSONSingleIsLikeArray(t *testing.T) {
	bytes := []byte(`{"tags": ["foo", "bar", "baz"]}`)

	var obj map[string]Single

	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		t.Error("failed to parse tags from json")
	}

	single := obj["tags"]
	expectedSingle := Single{"foo", "bar", "baz"}

	if !reflect.DeepEqual(single, expectedSingle) {
		t.Errorf("expected single to be equal to %v but was %v", expectedSingle, single)
	}
}
