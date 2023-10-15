package executor

import (
	"fmt"
)

// json to encode
type requestEmbedDataStore struct {
	store map[string][]byte
}

// CreateDataStore creates a new requestEmbedDataStore
func createDataStore() *requestEmbedDataStore {

	fmt.Println("sdk/executor/default_datastore.go: createDataStore start")
	rstore := &requestEmbedDataStore{}
	rstore.store = make(map[string][]byte)
	fmt.Println("sdk/executor/default_datastore.go: createDataStore end")
	return rstore
}

// retrieveDataStore creates a store manager from a map
func retrieveDataStore(store map[string][]byte) *requestEmbedDataStore {

	fmt.Println("sdk/executor/default_datastore.go: retrieveDataStore start")
	rstore := &requestEmbedDataStore{}
	rstore.store = store
	fmt.Println("sdk/executor/default_datastore.go: retrieveDataStore end")
	return rstore
}

// Configure Configure with requestId and flowname
func (rstore *requestEmbedDataStore) Configure(flowName string, requestId string) {
	fmt.Println("sdk/executor/default_datastore.go: Configure start")
	fmt.Println("sdk/executor/default_datastore.go: Configure end")
}

// Init initialize the storemanager (called only once in a request span)
func (rstore *requestEmbedDataStore) Init() error {
	fmt.Println("sdk/executor/default_datastore.go: Init start")
	fmt.Println("sdk/executor/default_datastore.go: Init end")
	return nil
}

// Set sets a value (implement DataStore)
func (rstore *requestEmbedDataStore) Set(key string, value []byte) error {

	fmt.Println("sdk/executor/default_datastore.go: Set start")
	fmt.Println("sdk/executor/default_datastore.go: Set end")
	rstore.store[key] = value
	return nil
}

// Get gets a value (implement DataStore)
func (rstore *requestEmbedDataStore) Get(key string) ([]byte, error) {

	fmt.Println("sdk/executor/default_datastore.go: Get start")
	value, ok := rstore.store[key]
	if !ok {
		return nil, fmt.Errorf("no field name %s", key)
	}
	fmt.Println("sdk/executor/default_datastore.go: Get end")
	return value, nil
}

// Del delets a value (implement DataStore)
func (rstore *requestEmbedDataStore) Del(key string) error {

	fmt.Println("sdk/executor/default_datastore.go: Del start")
	if _, ok := rstore.store[key]; ok {
		delete(rstore.store, key)
	}
	fmt.Println("sdk/executor/default_datastore.go: Del end")
	return nil
}

// Cleanup
func (rstore *requestEmbedDataStore) Cleanup() error {

	fmt.Println("sdk/executor/default_datastore.go: Cleanup start")
	fmt.Println("sdk/executor/default_datastore.go: Cleanup end")
	return nil
}
