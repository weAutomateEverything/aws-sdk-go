// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"github.com/weAutomateEverything/aws-sdk-go/aws"
	"github.com/weAutomateEverything/aws-sdk-go/aws/client"
	"github.com/weAutomateEverything/aws-sdk-go/aws/client/metadata"
	"github.com/weAutomateEverything/aws-sdk-go/aws/request"
	"github.com/weAutomateEverything/aws-sdk-go/aws/signer/v4"
	"github.com/weAutomateEverything/aws-sdk-go/private/protocol/jsonrpc"
)

// DatabaseMigrationService provides the API operation methods for making requests to
// AWS Database Migration Service. See this package's package overview docs
// for details on the service.
//
// DatabaseMigrationService methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type DatabaseMigrationService struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "dms"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the DatabaseMigrationService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a DatabaseMigrationService client from just a session.
//     svc := databasemigrationservice.New(mySession)
//
//     // Create a DatabaseMigrationService client with additional configuration
//     svc := databasemigrationservice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *DatabaseMigrationService {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *DatabaseMigrationService {
	svc := &DatabaseMigrationService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-01-01",
				JSONVersion:   "1.1",
				TargetPrefix:  "AmazonDMSv20160101",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a DatabaseMigrationService operation and runs any
// custom request initialization.
func (c *DatabaseMigrationService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
