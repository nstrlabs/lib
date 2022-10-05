package lib

type Feature interface {
	Execute() string
}

type Factory interface {
	New() Feature
}
