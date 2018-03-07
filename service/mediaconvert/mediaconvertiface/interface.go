// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediaconvertiface provides an interface to enable mocking the AWS Elemental MediaConvert service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediaconvertiface

import (
	"github.com/weAutomateEverything/aws-sdk-go/aws"
	"github.com/weAutomateEverything/aws-sdk-go/aws/request"
	"github.com/weAutomateEverything/aws-sdk-go/service/mediaconvert"
)

// MediaConvertAPI provides an interface to enable mocking the
// mediaconvert.MediaConvert service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Elemental MediaConvert.
//    func myFunc(svc mediaconvertiface.MediaConvertAPI) bool {
//        // Make svc.CancelJob request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := mediaconvert.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMediaConvertClient struct {
//        mediaconvertiface.MediaConvertAPI
//    }
//    func (m *mockMediaConvertClient) CancelJob(input *mediaconvert.CancelJobInput) (*mediaconvert.CancelJobOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMediaConvertClient{}
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
type MediaConvertAPI interface {
	CancelJob(*mediaconvert.CancelJobInput) (*mediaconvert.CancelJobOutput, error)
	CancelJobWithContext(aws.Context, *mediaconvert.CancelJobInput, ...request.Option) (*mediaconvert.CancelJobOutput, error)
	CancelJobRequest(*mediaconvert.CancelJobInput) (*request.Request, *mediaconvert.CancelJobOutput)

	CreateJob(*mediaconvert.CreateJobInput) (*mediaconvert.CreateJobOutput, error)
	CreateJobWithContext(aws.Context, *mediaconvert.CreateJobInput, ...request.Option) (*mediaconvert.CreateJobOutput, error)
	CreateJobRequest(*mediaconvert.CreateJobInput) (*request.Request, *mediaconvert.CreateJobOutput)

	CreateJobTemplate(*mediaconvert.CreateJobTemplateInput) (*mediaconvert.CreateJobTemplateOutput, error)
	CreateJobTemplateWithContext(aws.Context, *mediaconvert.CreateJobTemplateInput, ...request.Option) (*mediaconvert.CreateJobTemplateOutput, error)
	CreateJobTemplateRequest(*mediaconvert.CreateJobTemplateInput) (*request.Request, *mediaconvert.CreateJobTemplateOutput)

	CreatePreset(*mediaconvert.CreatePresetInput) (*mediaconvert.CreatePresetOutput, error)
	CreatePresetWithContext(aws.Context, *mediaconvert.CreatePresetInput, ...request.Option) (*mediaconvert.CreatePresetOutput, error)
	CreatePresetRequest(*mediaconvert.CreatePresetInput) (*request.Request, *mediaconvert.CreatePresetOutput)

	CreateQueue(*mediaconvert.CreateQueueInput) (*mediaconvert.CreateQueueOutput, error)
	CreateQueueWithContext(aws.Context, *mediaconvert.CreateQueueInput, ...request.Option) (*mediaconvert.CreateQueueOutput, error)
	CreateQueueRequest(*mediaconvert.CreateQueueInput) (*request.Request, *mediaconvert.CreateQueueOutput)

	DeleteJobTemplate(*mediaconvert.DeleteJobTemplateInput) (*mediaconvert.DeleteJobTemplateOutput, error)
	DeleteJobTemplateWithContext(aws.Context, *mediaconvert.DeleteJobTemplateInput, ...request.Option) (*mediaconvert.DeleteJobTemplateOutput, error)
	DeleteJobTemplateRequest(*mediaconvert.DeleteJobTemplateInput) (*request.Request, *mediaconvert.DeleteJobTemplateOutput)

	DeletePreset(*mediaconvert.DeletePresetInput) (*mediaconvert.DeletePresetOutput, error)
	DeletePresetWithContext(aws.Context, *mediaconvert.DeletePresetInput, ...request.Option) (*mediaconvert.DeletePresetOutput, error)
	DeletePresetRequest(*mediaconvert.DeletePresetInput) (*request.Request, *mediaconvert.DeletePresetOutput)

	DeleteQueue(*mediaconvert.DeleteQueueInput) (*mediaconvert.DeleteQueueOutput, error)
	DeleteQueueWithContext(aws.Context, *mediaconvert.DeleteQueueInput, ...request.Option) (*mediaconvert.DeleteQueueOutput, error)
	DeleteQueueRequest(*mediaconvert.DeleteQueueInput) (*request.Request, *mediaconvert.DeleteQueueOutput)

	DescribeEndpoints(*mediaconvert.DescribeEndpointsInput) (*mediaconvert.DescribeEndpointsOutput, error)
	DescribeEndpointsWithContext(aws.Context, *mediaconvert.DescribeEndpointsInput, ...request.Option) (*mediaconvert.DescribeEndpointsOutput, error)
	DescribeEndpointsRequest(*mediaconvert.DescribeEndpointsInput) (*request.Request, *mediaconvert.DescribeEndpointsOutput)

	GetJob(*mediaconvert.GetJobInput) (*mediaconvert.GetJobOutput, error)
	GetJobWithContext(aws.Context, *mediaconvert.GetJobInput, ...request.Option) (*mediaconvert.GetJobOutput, error)
	GetJobRequest(*mediaconvert.GetJobInput) (*request.Request, *mediaconvert.GetJobOutput)

	GetJobTemplate(*mediaconvert.GetJobTemplateInput) (*mediaconvert.GetJobTemplateOutput, error)
	GetJobTemplateWithContext(aws.Context, *mediaconvert.GetJobTemplateInput, ...request.Option) (*mediaconvert.GetJobTemplateOutput, error)
	GetJobTemplateRequest(*mediaconvert.GetJobTemplateInput) (*request.Request, *mediaconvert.GetJobTemplateOutput)

	GetPreset(*mediaconvert.GetPresetInput) (*mediaconvert.GetPresetOutput, error)
	GetPresetWithContext(aws.Context, *mediaconvert.GetPresetInput, ...request.Option) (*mediaconvert.GetPresetOutput, error)
	GetPresetRequest(*mediaconvert.GetPresetInput) (*request.Request, *mediaconvert.GetPresetOutput)

	GetQueue(*mediaconvert.GetQueueInput) (*mediaconvert.GetQueueOutput, error)
	GetQueueWithContext(aws.Context, *mediaconvert.GetQueueInput, ...request.Option) (*mediaconvert.GetQueueOutput, error)
	GetQueueRequest(*mediaconvert.GetQueueInput) (*request.Request, *mediaconvert.GetQueueOutput)

	ListJobTemplates(*mediaconvert.ListJobTemplatesInput) (*mediaconvert.ListJobTemplatesOutput, error)
	ListJobTemplatesWithContext(aws.Context, *mediaconvert.ListJobTemplatesInput, ...request.Option) (*mediaconvert.ListJobTemplatesOutput, error)
	ListJobTemplatesRequest(*mediaconvert.ListJobTemplatesInput) (*request.Request, *mediaconvert.ListJobTemplatesOutput)

	ListJobs(*mediaconvert.ListJobsInput) (*mediaconvert.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *mediaconvert.ListJobsInput, ...request.Option) (*mediaconvert.ListJobsOutput, error)
	ListJobsRequest(*mediaconvert.ListJobsInput) (*request.Request, *mediaconvert.ListJobsOutput)

	ListPresets(*mediaconvert.ListPresetsInput) (*mediaconvert.ListPresetsOutput, error)
	ListPresetsWithContext(aws.Context, *mediaconvert.ListPresetsInput, ...request.Option) (*mediaconvert.ListPresetsOutput, error)
	ListPresetsRequest(*mediaconvert.ListPresetsInput) (*request.Request, *mediaconvert.ListPresetsOutput)

	ListQueues(*mediaconvert.ListQueuesInput) (*mediaconvert.ListQueuesOutput, error)
	ListQueuesWithContext(aws.Context, *mediaconvert.ListQueuesInput, ...request.Option) (*mediaconvert.ListQueuesOutput, error)
	ListQueuesRequest(*mediaconvert.ListQueuesInput) (*request.Request, *mediaconvert.ListQueuesOutput)

	UpdateJobTemplate(*mediaconvert.UpdateJobTemplateInput) (*mediaconvert.UpdateJobTemplateOutput, error)
	UpdateJobTemplateWithContext(aws.Context, *mediaconvert.UpdateJobTemplateInput, ...request.Option) (*mediaconvert.UpdateJobTemplateOutput, error)
	UpdateJobTemplateRequest(*mediaconvert.UpdateJobTemplateInput) (*request.Request, *mediaconvert.UpdateJobTemplateOutput)

	UpdatePreset(*mediaconvert.UpdatePresetInput) (*mediaconvert.UpdatePresetOutput, error)
	UpdatePresetWithContext(aws.Context, *mediaconvert.UpdatePresetInput, ...request.Option) (*mediaconvert.UpdatePresetOutput, error)
	UpdatePresetRequest(*mediaconvert.UpdatePresetInput) (*request.Request, *mediaconvert.UpdatePresetOutput)

	UpdateQueue(*mediaconvert.UpdateQueueInput) (*mediaconvert.UpdateQueueOutput, error)
	UpdateQueueWithContext(aws.Context, *mediaconvert.UpdateQueueInput, ...request.Option) (*mediaconvert.UpdateQueueOutput, error)
	UpdateQueueRequest(*mediaconvert.UpdateQueueInput) (*request.Request, *mediaconvert.UpdateQueueOutput)
}

var _ MediaConvertAPI = (*mediaconvert.MediaConvert)(nil)
