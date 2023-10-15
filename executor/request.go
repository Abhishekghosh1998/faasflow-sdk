package executor

import (
	"encoding/json"
	"fmt"
)

// Request defines the body of async forward request to faasflow
type Request struct {
	Sign        string `json: "sign"`         // request signature
	ID          string `json: "id"`           // request ID
	Query       string `json: "query"`        // query string
	CallbackUrl string `json: "callback-url"` // callback url

	ExecutionState string `json: "state"` // Execution State (execution position / execution vertex)

	Data []byte `json: "data"` // Partial execution data
	// (empty if intermediate_storage enabled

	ContextStore map[string][]byte `json: "store"` // Context State for default DataStore
	// (empty if external Store is used)
}

func buildRequest(id string,
	state string,
	query string,
	data []byte,
	contextState map[string][]byte,
	sign string) *Request {

	fmt.Println("sdk/executor/request.go: buildRequest start")

	request := &Request{
		Sign:           sign,
		ID:             id,
		ExecutionState: state,
		Query:          query,
		Data:           data,
		ContextStore:   contextState,
	}

	fmt.Println("sdk/executor/request.go: buildRequest end")
	return request
}

func decodeRequest(data []byte) (*Request, error) {

	fmt.Println("sdk/executor/request.go: decodeRequest start")
	request := &Request{}
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, err
	}
	fmt.Println("sdk/executor/request.go: decodeRequest end")
	return request, nil
}

func (req *Request) encode() ([]byte, error) {

	fmt.Println("sdk/executor/request.go: encode start")
	fmt.Println("sdk/executor/request.go: encode end")
	return json.Marshal(req)
}

func (req *Request) getData() []byte {

	fmt.Println("sdk/executor/request.go: getData start")
	fmt.Println("sdk/executor/request.go: getData end")
	return req.Data
}

func (req *Request) getID() string {

	fmt.Println("sdk/executor/request.go: getID start")
	fmt.Println("sdk/executor/request.go: getID end")
	return req.ID
}

func (req *Request) getExecutionState() string {

	fmt.Println("sdk/executor/request.go: getExecutionState start")
	fmt.Println("sdk/executor/request.go: getExecutionState end")
	return req.ExecutionState
}

func (req *Request) getContextStore() map[string][]byte {

	fmt.Println("sdk/executor/request.go: getContextStore start")
	fmt.Println("sdk/executor/request.go: getContextStore end")
	return req.ContextStore
}

func (req *Request) getQuery() string {

	fmt.Println("sdk/executor/request.go: getQuery start")
	fmt.Println("sdk/executor/request.go: getQuery end")
	return req.Query
}
