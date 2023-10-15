package sdk

import (
	"encoding/json"
	"fmt"
)

const (
	DEPTH_INCREMENT = 1
	DEPTH_DECREMENT = -1
	DEPTH_SAME      = 0
)

// PipelineErrorHandler the error handler OnFailure() registration on pipeline
type PipelineErrorHandler func(error) ([]byte, error)

// PipelineHandler definition for the Finally() registration on pipeline
type PipelineHandler func(string)

type Pipeline struct {
	Dag *Dag `json:"-"` // Dag that will be executed

	ExecutionPosition map[string]string `json:"pipeline-execution-position"` // Denotes the node that is executing now
	ExecutionDepth    int               `json:"pipeline-execution-depth"`    // Denotes the depth of subgraph its executing

	CurrentDynamicOption map[string]string `json:"pipeline-dynamic-option"` // Denotes the current dynamic option mapped against the dynamic Node UQ id

	FailureHandler PipelineErrorHandler `json:"-"`
	Finally        PipelineHandler      `json:"-"`
}

// CreatePipeline creates a faasflow pipeline
func CreatePipeline() *Pipeline {

	fmt.Println("sdk/pipeline.go: CreatePipeline start")
	pipeline := &Pipeline{}
	pipeline.Dag = NewDag()

	pipeline.ExecutionPosition = make(map[string]string, 0)
	pipeline.ExecutionDepth = 0
	pipeline.CurrentDynamicOption = make(map[string]string, 0)

	fmt.Println("sdk/pipeline.go: CreatePipeline end")
	return pipeline
}

// CountNodes counts the no of node added in the Pipeline Dag.
// It doesn't count subdags node
func (pipeline *Pipeline) CountNodes() int {

	fmt.Println("sdk/pipeline.go: CountNodes start")
	fmt.Println("sdk/pipeline.go: CountNodes end")
	return len(pipeline.Dag.nodes)
}

// GetAllNodesId returns a recursive list of all nodes that belongs to the pipeline
func (pipeline *Pipeline) GetAllNodesUniqueId() []string {

	fmt.Println("sdk/pipeline.go: GetAllNodesUniqueId start")
	nodes := pipeline.Dag.GetNodes("")
	fmt.Println("sdk/pipeline.go: GetAllNodesUniqueId end")
	return nodes
}

// GetInitialNodeId Get the very first node of the pipeline
func (pipeline *Pipeline) GetInitialNodeId() string {

	fmt.Println("sdk/pipeline.go: GetInitialNodeId start")
	node := pipeline.Dag.GetInitialNode()
	if node != nil {
		return node.Id
	}
	fmt.Println("sdk/pipeline.go: GetInitialNodeId end")
	return "0"
}

// GetNodeExecutionUniqueId provide a ID that is unique in an execution
func (pipeline *Pipeline) GetNodeExecutionUniqueId(node *Node) string {

	fmt.Println("sdk/pipeline.go: GetNodeExecutionUniqueId start")
	depth := 0
	dag := pipeline.Dag
	depthStr := ""
	optionStr := ""
	for depth < pipeline.ExecutionDepth {
		depthStr = fmt.Sprintf("%d", depth)
		node := dag.GetNode(pipeline.ExecutionPosition[depthStr])
		option := pipeline.CurrentDynamicOption[node.GetUniqueId()]
		if node.subDag != nil {
			dag = node.subDag
		} else {
			dag = node.conditionalDags[option]
		}
		if optionStr == "" {
			optionStr = option
		} else {
			optionStr = option + "--" + optionStr
		}

		depth++
	}
	if optionStr == "" {
		return node.GetUniqueId()
	}
	fmt.Println("sdk/pipeline.go: GetNodeExecutionUniqueId end")
	return optionStr + "--" + node.GetUniqueId()
}

// GetCurrentNodeDag returns the current node and current dag based on execution position
func (pipeline *Pipeline) GetCurrentNodeDag() (*Node, *Dag) {

	fmt.Println("sdk/pipeline.go: GetCurrentNodeDag start")
	depth := 0
	dag := pipeline.Dag
	depthStr := ""
	for depth < pipeline.ExecutionDepth {
		depthStr = fmt.Sprintf("%d", depth)
		node := dag.GetNode(pipeline.ExecutionPosition[depthStr])
		option := pipeline.CurrentDynamicOption[node.GetUniqueId()]
		if node.subDag != nil {
			dag = node.subDag
		} else {
			dag = node.conditionalDags[option]
		}
		depth++
	}
	depthStr = fmt.Sprintf("%d", depth)
	node := dag.GetNode(pipeline.ExecutionPosition[depthStr])
	fmt.Println("sdk/pipeline.go: GetCurrentNodeDag end")
	return node, dag
}

// UpdatePipelineExecutionPosition updates pipeline execution position
// specified depthAdjustment and vertex denotes how the ExecutionPosition must be altered
func (pipeline *Pipeline) UpdatePipelineExecutionPosition(depthAdjustment int, vertex string) {

	fmt.Println("sdk/pipeline.go: UpdatePipelineExecutionPosition start")
	pipeline.ExecutionDepth = pipeline.ExecutionDepth + depthAdjustment
	depthStr := fmt.Sprintf("%d", pipeline.ExecutionDepth)
	pipeline.ExecutionPosition[depthStr] = vertex
	fmt.Println("sdk/pipeline.go: UpdatePipelineExecutionPosition end")
}

// SetDag overrides the default dag
func (pipeline *Pipeline) SetDag(dag *Dag) {

	fmt.Println("sdk/pipeline.go: SetDag start")
	pipeline.Dag = dag
	fmt.Println("sdk/pipeline.go: SetDag end")
}

// decodePipeline decodes a json marshaled pipeline
func decodePipeline(data []byte) (*Pipeline, error) {

	fmt.Println("sdk/pipeline.go: decodePipeline start")
	pipeline := &Pipeline{}
	err := json.Unmarshal(data, pipeline)
	if err != nil {
		return nil, err
	}
	fmt.Println("sdk/pipeline.go: decodePipeline end")
	return pipeline, nil
}

// GetState get a state of a pipeline by encoding in JSON
func (pipeline *Pipeline) GetState() string {

	fmt.Println("sdk/pipeline.go: GetState start")
	encode, _ := json.Marshal(pipeline)
	fmt.Println("sdk/pipeline.go: GetState end")
	return string(encode)
}

// ApplyState apply a state to a pipeline by from encoded JSON pipeline
func (pipeline *Pipeline) ApplyState(state string) {

	fmt.Println("sdk/pipeline.go: ApplyState start")
	temp, _ := decodePipeline([]byte(state))
	pipeline.ExecutionDepth = temp.ExecutionDepth
	pipeline.ExecutionPosition = temp.ExecutionPosition
	pipeline.CurrentDynamicOption = temp.CurrentDynamicOption
	fmt.Println("sdk/pipeline.go: ApplyState end")
}
