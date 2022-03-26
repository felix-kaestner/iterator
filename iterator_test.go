package iterator

import (
	"reflect"
	"testing"
)

func isNil(i any) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return v.IsNil()
	}

	return false
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if (isNil(expected) && isNil(actual)) || reflect.DeepEqual(expected, actual) {
		return
	}

	t.Errorf("Test %s: Expected `%v` (type %v), Received `%v` (type %v)", t.Name(), expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
}

// Test the slice iterator
func TestSliceIterator(t *testing.T) {
	{
		s := []int{1, 2, 3}
		i := FromSlice(s)

		c := 0
		for i.HasNext() {
			c++
			item, _ := i.Next()
			assertEqual(t, c, *item)
		}

		assertEqual(t, i.HasNext(), false)

		_, err := i.Next()
		assertEqual(t, Done, err)
	}

	{
		s := []int{1, 2, 3}
		i := FromSlice(s)

		assertEqual(t, i.HasNext(), true)
		item, _ := i.Next()
		assertEqual(t, 1, *item)

		assertEqual(t, i.HasNext(), true)
		item, _ = i.Next()
		assertEqual(t, 2, *item)

		assertEqual(t, i.HasNext(), true)
		item, _ = i.Next()
		assertEqual(t, 3, *item)

		assertEqual(t, i.HasNext(), false)
		_, err := i.Next()
		assertEqual(t, Done, err)
	}
}

func TestIndexedIterator(t *testing.T) {
	{
		s := []int{1, 2, 3}
		i := WithIndex(FromSlice(s))

		c := 0
		for i.HasNext() {
			c++
			item, _ := i.Next()
			assertEqual(t, c, item.value)
			assertEqual(t, c-1, item.index)
		}

		assertEqual(t, i.HasNext(), false)

		_, err := i.Next()
		assertEqual(t, Done, err)
	}

	{
		s := []int{1, 2, 3}
		i := WithIndex(FromSlice(s))

		assertEqual(t, i.HasNext(), true)
		item, _ := i.Next()
		assertEqual(t, 1, item.value)
		assertEqual(t, 0, item.index)

		assertEqual(t, i.HasNext(), true)
		item, _ = i.Next()
		assertEqual(t, 2, item.value)
		assertEqual(t, 1, item.index)

		assertEqual(t, i.HasNext(), true)
		item, _ = i.Next()
		assertEqual(t, 3, item.value)
		assertEqual(t, 2, item.index)

		assertEqual(t, i.HasNext(), false)
		_, err := i.Next()
		assertEqual(t, Done, err)
	}
}
