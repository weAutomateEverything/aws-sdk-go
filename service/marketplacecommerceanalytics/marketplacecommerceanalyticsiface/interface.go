// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package marketplacecommerceanalyticsiface provides an interface to enable mocking the AWS Marketplace Commerce Analytics service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package marketplacecommerceanalyticsiface

import (
	"github.com/weAutomateEverything/aws-sdk-go/aws"
	"github.com/weAutomateEverything/aws-sdk-go/aws/request"
	"github.com/weAutomateEverything/aws-sdk-go/service/marketplacecommerceanalytics"
)

// MarketplaceCommerceAnalyticsAPI provides an interface to enable mocking the
// marketplacecommerceanalytics.MarketplaceCommerceAnalytics service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Marketplace Commerce Analytics.
//    func myFunc(svc marketplacecommerceanalyticsiface.MarketplaceCommerceAnalyticsAPI) bool {
//        // Make svc.GenerateDataSet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := marketplacecommerceanalytics.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMarketplaceCommerceAnalyticsClient struct {
//        marketplacecommerceanalyticsiface.MarketplaceCommerceAnalyticsAPI
//    }
//    func (m *mockMarketplaceCommerceAnalyticsClient) GenerateDataSet(input *marketplacecommerceanalytics.GenerateDataSetInput) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMarketplaceCommerceAnalyticsClient{}
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
type MarketplaceCommerceAnalyticsAPI interface {
	GenerateDataSet(*marketplacecommerceanalytics.GenerateDataSetInput) (*marketplacecommerceanalytics.GenerateDataSetOutput, error)
	GenerateDataSetWithContext(aws.Context, *marketplacecommerceanalytics.GenerateDataSetInput, ...request.Option) (*marketplacecommerceanalytics.GenerateDataSetOutput, error)
	GenerateDataSetRequest(*marketplacecommerceanalytics.GenerateDataSetInput) (*request.Request, *marketplacecommerceanalytics.GenerateDataSetOutput)

	StartSupportDataExport(*marketplacecommerceanalytics.StartSupportDataExportInput) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error)
	StartSupportDataExportWithContext(aws.Context, *marketplacecommerceanalytics.StartSupportDataExportInput, ...request.Option) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error)
	StartSupportDataExportRequest(*marketplacecommerceanalytics.StartSupportDataExportInput) (*request.Request, *marketplacecommerceanalytics.StartSupportDataExportOutput)
}

var _ MarketplaceCommerceAnalyticsAPI = (*marketplacecommerceanalytics.MarketplaceCommerceAnalytics)(nil)
