package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func Equal(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		return
	}
	t.Fatal(fmt.Sprintf("got %v (%T), expected %v (%T)", a, a, b, b))
}

func NotEqual(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		return
	}
	t.Fatal(fmt.Sprintf("expected %v (%T) and %v (%T) to not be equal", a, a, b, b))
}
