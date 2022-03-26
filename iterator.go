package iterator

import "errors"

const (
	// Version is the current version of the library.
	Version = "0.0.1"
)

var (
	// Done is returned when an iterator is read past its end.
	Done = errors.New("iterator: no more items in iterator")
)

type (
	// Iterator is the interface that allows to sequentially
	// access the elements of a collection or another entity
	// that can be represented as a sequence of elements.
	Iterator[T any] interface {
		// Next returns the next item in the iteration.
		// It returns a nil pointer and ErrIteratorOverread
		// if the iterator was read past its end.
		// Otherwise it returns a reference to the next item.
		Next() (*T, error)
		// HasNext returns true if the iteration has more elements.
		HasNext() bool
	}

	// sliceIterator is the struct type that implements the
	// Iterator interface for slices.
	sliceIterator[T any] struct {
		index int
		slice []T
	}
)

// Next returns the next item in the iteration.
func (s *sliceIterator[T]) Next() (*T, error) {
	if s.index >= len(s.slice) {
		return nil, Done
	}
	item := s.slice[s.index]
	s.index++
	return &item, nil
}

// HasNext returns true if the iteration has more elements.
func (s *sliceIterator[T]) HasNext() bool {
	return s.index < len(s.slice)
}

// FromSlice returns a new Iterator for the given slice.
func FromSlice[T any](slice []T) Iterator[T] {
	return &sliceIterator[T]{
		index: 0,
		slice: slice,
	}
}
