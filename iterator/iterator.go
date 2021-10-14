package iterator

// An iterator over a collection.
type Iterator interface {
	HasNext() bool
	Next() interface{}
}
