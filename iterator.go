package iterator

import "errors"

// Done is returned when an iterator is read past its end.
var Done = errors.New("iterator: no more items in iterator")

// Iterator is the interface that allows to sequentially
// access the elements of a collection or another entity
// that can be represented as a sequence of elements.
type Iterator[T any] interface {
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
type IndexedValue[T any] struct {
	index int
	value *T
}

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

func (it *sliceIterator[T]) Next() (*T, error) {
	if !it.HasNext() {
		return nil, Done
	}
	item := it.slice[it.index]
	it.index++
	return &item, nil
}

func (it *sliceIterator[T]) HasNext() bool {
	return it.index < len(it.slice)
}

// FromSlice returns a new Iterator for the given slice.
func FromSlice[T any](slice []T) Iterator[T] {
	return &sliceIterator[T]{
		index: 0,
		slice: slice,
	}
}

// channelIterator implements the Iterator interface for channels.
type channelIterator[T any] struct {
	next *T
	ch   <-chan T
}

func (it *channelIterator[T]) Next() (*T, error) {
	if !it.HasNext() {
		return nil, Done
	}
	item := it.next
	it.next = nil
	return item, nil
}

func (it *channelIterator[T]) HasNext() bool {
	if it.next == nil {
		if item, ok := <-it.ch; ok {
			it.next = &item
		} else {
			return false
		}
	}
	return true
}

// FromChannel returns a new Iterator for the given channel.
func FromChannel[T any](ch <-chan T) Iterator[T] {
	return &channelIterator[T]{
		next: nil,
		ch:   ch,
	}
}

// funcIterator implements the Iterator interface for functions.
type funcIterator[T any] struct {
	next *T
	err  error
	fn   func() (*T, error)
}

func (it *funcIterator[T]) Next() (*T, error) {
	if !it.HasNext() {
		return nil, it.err
	}
	item := it.next
	it.next = nil
	return item, nil
}

func (it *funcIterator[T]) HasNext() bool {
	if it.next == nil {
		if item, err := it.fn(); err == nil {
			it.next = item
		} else {
			it.err = err
			return false
		}
	}
	return true
}

// FromFunc returns a new Iterator for the given function.
func FromFunc[T any](fn func() (*T, error)) Iterator[T] {
	return &funcIterator[T]{
		fn: fn,
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

func (it *indexedIterator[T]) Next() (*IndexedValue[T], error) {
	next, err := it.iter.Next()
	if err != nil {
		return nil, err
	}
	item := &IndexedValue[T]{
		index: it.index,
		value: next,
	}
	it.index++
	return item, nil
}

func (it *indexedIterator[T]) HasNext() bool {
	return it.iter.HasNext()
}

// WithIndex returns a new Iterator that wraps each element
// produced by the original iterator into an IndexedValue
// containing the index of that element and the element itself.
func WithIndex[T any](iter Iterator[T]) Iterator[IndexedValue[T]] {
	return &indexedIterator[T]{
		index: 0,
		iter:  iter,
	}
}
