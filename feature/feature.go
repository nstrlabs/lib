package feature

type Feature interface {
	Execute() string
}

type Factory interface {
	New() Feature
}
