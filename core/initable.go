package core

// Interface to indicate a instance should be initialize before using.
type Initable interface {
	Init() // do init
}

// Initialize the object if it implements Initable interface
func InitializeIfNeeded(obj interface{}) interface{} {
	if ia, ok := obj.(Initable); ok {
		ia.Init()
	}
	return obj
}
