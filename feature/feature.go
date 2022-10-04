package feature

type Feature interface {
	Execute()
}

type Factory interface {
	New() Feature
}
