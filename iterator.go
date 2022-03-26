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

	// IndexedValue represents a value from a collection or sequence,
	// along with its associated index in that collection or sequence.
	IndexedValue[T any] struct {
		index int
		value T
	}
)

// ForEach iterates over the elements of the iterator, calling the
// provided function for each element.
//
// This is equivalent to calling Next() until HasNext() returns false.
// 	var it Iterator[T]
// 	for it.HasNext() {
// 		v, _ := it.Next()
// 	}
func ForEach[T any](iterator Iterator[T], visitor func(item *T)) {
	for iterator.HasNext() {
		item, _ := iterator.Next()
		visitor(item)
	}
}

// sliceIterator implements the Iterator interface for slices.
type sliceIterator[T any] struct {
	index int
	slice []T
}

func (s *sliceIterator[T]) Next() (*T, error) {
	if s.index >= len(s.slice) {
		return nil, Done
	}
	item := s.slice[s.index]
	s.index++
	return &item, nil
}

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

// indexedIterator implements the Iterator interface for
// indexed collections.
// It wraps an existing Iterator and each element produced
// by the original iterator into an IndexedValue containing
// the index of that element and the element itself.
type indexedIterator[T any] struct {
	index int
	iter  Iterator[T]
}

func (i *indexedIterator[T]) Next() (*IndexedValue[T], error) {
	item, err := i.iter.Next()
	if err != nil {
		return nil, err
	}
	i.index++
	return &IndexedValue[T]{
		index: i.index,
		value: *item,
	}, nil
}

func (i *indexedIterator[T]) HasNext() bool {
	return i.iter.HasNext()
}

// WithIndex returns a new Iterator that wraps each element
// produced by the original iterator into an IndexedValue
// containing the index of that element and the element itself.
func WithIndex[T any](iter Iterator[T]) Iterator[IndexedValue[T]] {
	return &indexedIterator[T]{
		index: -1,
		iter:  iter,
	}
}
