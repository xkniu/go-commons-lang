package iterator

var IteratorUtils *iteratorUtils

type iteratorUtils struct {
}



// Applies the action to each element of the iterator.
func ForEach(iterator Iterator, action func(obj interface{})) {
	for iterator.HasNext() {
		obj := iterator.Next()
		action(obj)
	}
}

// Return the size of the iterator.
func Size(iterator Iterator) int {
	var size int
	for ; iterator.HasNext(); iterator.Next() {
		size++
	}
	return size
}

// Get a slice from the iterator.
func ToSlice(iterator Iterator) []interface{} {
	r := make([]interface{}, 0)
	for iterator.HasNext() {
		obj := iterator.Next()
		r = append(r, obj)
	}
	return r
}
