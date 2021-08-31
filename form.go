package gmif

type Form struct {
	Action string  `toml:",omitempty"`
	Group  []Group `toml:",omitempty"`
	Error  string  `toml:",omitempty"`
}

type Group struct {
	Title string  `toml:",omitempty"`
	Field []Field `toml:",omitempty"`
	Error string  `toml:",omitempty"`
}

type DataType string

const (
	String   DataType = ""
	Int      DataType = "int"
	Float    DataType = "float"
	Double   DataType = "double"
	UnixSecs DataType = "unixsec"
)

type Field struct {
	Name        string
	Value       interface{}
	Description string   `toml:",omitempty"`
	ValueType   DataType `toml:",omitempty"`
	Error       string   `toml:",omitempty"`
}
