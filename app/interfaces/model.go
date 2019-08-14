package interfaces

type Model interface {
	Collection() string
	Instance() interface{}
}
