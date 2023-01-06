package lib

type Msg interface {
	Add(string, []byte)
	Get(string) []byte
	GetRaw() []byte
}

type Feature interface {
	Execute(msg Msg) error
}

type Factory interface {
	New(configuration map[string]interface{}) Feature
}

type Validator interface {
	Validate(configuration map[string]interface{}) error
}
