package armory

//go:generate go-bindata -pkg main -o cmd/armory/set.go set.go

// Generic represents a generic type. It is a alias of interface{}.
type Generic interface{}
