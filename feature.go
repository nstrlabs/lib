package lib

type Feature interface {
	Execute() string
}

type Factory interface {
	New(configuration map[string]interface{}) Feature
}

type Validator interface {
	Validate(configuration map[string]interface{}) error
}
