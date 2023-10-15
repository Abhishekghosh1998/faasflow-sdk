package sdk

import "fmt"

type Operation interface {
	GetId() string
	Encode() []byte
	GetProperties() map[string][]string
	// Execute executes an operation, executor can pass configuration
	Execute([]byte, map[string]interface{}) ([]byte, error)
}

type BlankOperation struct {
}

func (ops *BlankOperation) GetId() string {

	fmt.Println("sdk/operation.go: GetId start")
	fmt.Println("sdk/operation.go: GetId end")
	return "end"
}

func (ops *BlankOperation) Encode() []byte {

	fmt.Println("sdk/operation.go: Encode start")
	fmt.Println("sdk/operation.go: Encode end")
	return []byte("")
}

func (ops *BlankOperation) GetProperties() map[string][]string {

	fmt.Println("sdk/operation.go: GetProperties start")
	fmt.Println("sdk/operation.go: GetProperties end")
	return make(map[string][]string)
}

func (ops *BlankOperation) Execute(data []byte, option map[string]interface{}) ([]byte, error) {

	fmt.Println("sdk/operation.go: Execute start")
	fmt.Println("sdk/operation.go: Execute end")
	return data, nil
}
