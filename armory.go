package armory

//go:generate go-bindata -ignore=\\*_test.go -pkg main -o cmd/armory/files.go ll set stack queue

// Generic represents a generic type. It is a alias of interface{}.
type Generic interface{}
