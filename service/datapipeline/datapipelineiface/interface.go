// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package datapipelineiface provides an interface to enable mocking the AWS Data Pipeline service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package datapipelineiface

import (
	"github.com/weAutomateEverything/aws-sdk-go/aws"
	"github.com/weAutomateEverything/aws-sdk-go/aws/request"
	"github.com/weAutomateEverything/aws-sdk-go/service/datapipeline"
)

// DataPipelineAPI provides an interface to enable mocking the
// datapipeline.DataPipeline service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Data Pipeline.
//    func myFunc(svc datapipelineiface.DataPipelineAPI) bool {
//        // Make svc.ActivatePipeline request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := datapipeline.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDataPipelineClient struct {
//        datapipelineiface.DataPipelineAPI
//    }
//    func (m *mockDataPipelineClient) ActivatePipeline(input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDataPipelineClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DataPipelineAPI interface {
	ActivatePipeline(*datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error)
	ActivatePipelineWithContext(aws.Context, *datapipeline.ActivatePipelineInput, ...request.Option) (*datapipeline.ActivatePipelineOutput, error)
	ActivatePipelineRequest(*datapipeline.ActivatePipelineInput) (*request.Request, *datapipeline.ActivatePipelineOutput)

	AddTags(*datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error)
	AddTagsWithContext(aws.Context, *datapipeline.AddTagsInput, ...request.Option) (*datapipeline.AddTagsOutput, error)
	AddTagsRequest(*datapipeline.AddTagsInput) (*request.Request, *datapipeline.AddTagsOutput)

	CreatePipeline(*datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error)
	CreatePipelineWithContext(aws.Context, *datapipeline.CreatePipelineInput, ...request.Option) (*datapipeline.CreatePipelineOutput, error)
	CreatePipelineRequest(*datapipeline.CreatePipelineInput) (*request.Request, *datapipeline.CreatePipelineOutput)

	DeactivatePipeline(*datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error)
	DeactivatePipelineWithContext(aws.Context, *datapipeline.DeactivatePipelineInput, ...request.Option) (*datapipeline.DeactivatePipelineOutput, error)
	DeactivatePipelineRequest(*datapipeline.DeactivatePipelineInput) (*request.Request, *datapipeline.DeactivatePipelineOutput)

	DeletePipeline(*datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error)
	DeletePipelineWithContext(aws.Context, *datapipeline.DeletePipelineInput, ...request.Option) (*datapipeline.DeletePipelineOutput, error)
	DeletePipelineRequest(*datapipeline.DeletePipelineInput) (*request.Request, *datapipeline.DeletePipelineOutput)

	DescribeObjects(*datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error)
	DescribeObjectsWithContext(aws.Context, *datapipeline.DescribeObjectsInput, ...request.Option) (*datapipeline.DescribeObjectsOutput, error)
	DescribeObjectsRequest(*datapipeline.DescribeObjectsInput) (*request.Request, *datapipeline.DescribeObjectsOutput)

	DescribeObjectsPages(*datapipeline.DescribeObjectsInput, func(*datapipeline.DescribeObjectsOutput, bool) bool) error
	DescribeObjectsPagesWithContext(aws.Context, *datapipeline.DescribeObjectsInput, func(*datapipeline.DescribeObjectsOutput, bool) bool, ...request.Option) error

	DescribePipelines(*datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error)
	DescribePipelinesWithContext(aws.Context, *datapipeline.DescribePipelinesInput, ...request.Option) (*datapipeline.DescribePipelinesOutput, error)
	DescribePipelinesRequest(*datapipeline.DescribePipelinesInput) (*request.Request, *datapipeline.DescribePipelinesOutput)

	EvaluateExpression(*datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error)
	EvaluateExpressionWithContext(aws.Context, *datapipeline.EvaluateExpressionInput, ...request.Option) (*datapipeline.EvaluateExpressionOutput, error)
	EvaluateExpressionRequest(*datapipeline.EvaluateExpressionInput) (*request.Request, *datapipeline.EvaluateExpressionOutput)

	GetPipelineDefinition(*datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error)
	GetPipelineDefinitionWithContext(aws.Context, *datapipeline.GetPipelineDefinitionInput, ...request.Option) (*datapipeline.GetPipelineDefinitionOutput, error)
	GetPipelineDefinitionRequest(*datapipeline.GetPipelineDefinitionInput) (*request.Request, *datapipeline.GetPipelineDefinitionOutput)

	ListPipelines(*datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error)
	ListPipelinesWithContext(aws.Context, *datapipeline.ListPipelinesInput, ...request.Option) (*datapipeline.ListPipelinesOutput, error)
	ListPipelinesRequest(*datapipeline.ListPipelinesInput) (*request.Request, *datapipeline.ListPipelinesOutput)

	ListPipelinesPages(*datapipeline.ListPipelinesInput, func(*datapipeline.ListPipelinesOutput, bool) bool) error
	ListPipelinesPagesWithContext(aws.Context, *datapipeline.ListPipelinesInput, func(*datapipeline.ListPipelinesOutput, bool) bool, ...request.Option) error

	PollForTask(*datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error)
	PollForTaskWithContext(aws.Context, *datapipeline.PollForTaskInput, ...request.Option) (*datapipeline.PollForTaskOutput, error)
	PollForTaskRequest(*datapipeline.PollForTaskInput) (*request.Request, *datapipeline.PollForTaskOutput)

	PutPipelineDefinition(*datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error)
	PutPipelineDefinitionWithContext(aws.Context, *datapipeline.PutPipelineDefinitionInput, ...request.Option) (*datapipeline.PutPipelineDefinitionOutput, error)
	PutPipelineDefinitionRequest(*datapipeline.PutPipelineDefinitionInput) (*request.Request, *datapipeline.PutPipelineDefinitionOutput)

	QueryObjects(*datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error)
	QueryObjectsWithContext(aws.Context, *datapipeline.QueryObjectsInput, ...request.Option) (*datapipeline.QueryObjectsOutput, error)
	QueryObjectsRequest(*datapipeline.QueryObjectsInput) (*request.Request, *datapipeline.QueryObjectsOutput)

	QueryObjectsPages(*datapipeline.QueryObjectsInput, func(*datapipeline.QueryObjectsOutput, bool) bool) error
	QueryObjectsPagesWithContext(aws.Context, *datapipeline.QueryObjectsInput, func(*datapipeline.QueryObjectsOutput, bool) bool, ...request.Option) error

	RemoveTags(*datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error)
	RemoveTagsWithContext(aws.Context, *datapipeline.RemoveTagsInput, ...request.Option) (*datapipeline.RemoveTagsOutput, error)
	RemoveTagsRequest(*datapipeline.RemoveTagsInput) (*request.Request, *datapipeline.RemoveTagsOutput)

	ReportTaskProgress(*datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error)
	ReportTaskProgressWithContext(aws.Context, *datapipeline.ReportTaskProgressInput, ...request.Option) (*datapipeline.ReportTaskProgressOutput, error)
	ReportTaskProgressRequest(*datapipeline.ReportTaskProgressInput) (*request.Request, *datapipeline.ReportTaskProgressOutput)

	ReportTaskRunnerHeartbeat(*datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error)
	ReportTaskRunnerHeartbeatWithContext(aws.Context, *datapipeline.ReportTaskRunnerHeartbeatInput, ...request.Option) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error)
	ReportTaskRunnerHeartbeatRequest(*datapipeline.ReportTaskRunnerHeartbeatInput) (*request.Request, *datapipeline.ReportTaskRunnerHeartbeatOutput)

	SetStatus(*datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error)
	SetStatusWithContext(aws.Context, *datapipeline.SetStatusInput, ...request.Option) (*datapipeline.SetStatusOutput, error)
	SetStatusRequest(*datapipeline.SetStatusInput) (*request.Request, *datapipeline.SetStatusOutput)

	SetTaskStatus(*datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error)
	SetTaskStatusWithContext(aws.Context, *datapipeline.SetTaskStatusInput, ...request.Option) (*datapipeline.SetTaskStatusOutput, error)
	SetTaskStatusRequest(*datapipeline.SetTaskStatusInput) (*request.Request, *datapipeline.SetTaskStatusOutput)

	ValidatePipelineDefinition(*datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error)
	ValidatePipelineDefinitionWithContext(aws.Context, *datapipeline.ValidatePipelineDefinitionInput, ...request.Option) (*datapipeline.ValidatePipelineDefinitionOutput, error)
	ValidatePipelineDefinitionRequest(*datapipeline.ValidatePipelineDefinitionInput) (*request.Request, *datapipeline.ValidatePipelineDefinitionOutput)
}

var _ DataPipelineAPI = (*datapipeline.DataPipeline)(nil)
