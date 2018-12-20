package storage

//Interface  wrap storage driver
type Interface interface {
	//storage type
	Type() string
}
