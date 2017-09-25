package testing

import (
	"testing"
	"reflect"
)

func AssertArrayDeepEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("actual not equal expected", actual, ",", expected)
	}
}
