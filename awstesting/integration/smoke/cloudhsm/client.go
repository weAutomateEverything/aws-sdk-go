// +build integration

//Package cloudhsm provides gucumber integration tests support.
package cloudhsm

import (
	"github.com/weAutomateEverything/aws-sdk-go/awstesting/integration/smoke"
	"github.com/weAutomateEverything/aws-sdk-go/service/cloudhsm"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudhsm", func() {
		gucumber.World["client"] = cloudhsm.New(smoke.Session)
	})
}
