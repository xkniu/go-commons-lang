package core

// Interface to indicate a instance should be destroy before dropping.
type Destroyable interface {
	Destroy() // do destroy
}

// Destroy the object if it implements Destroyable interface
func DestroyIfNeeded(obj interface{}) interface{} {
	if ia, ok := obj.(Destroyable); ok {
		ia.Destroy()
	}
	return obj
}
