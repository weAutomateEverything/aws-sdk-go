// +build integration

//Package ecs provides gucumber integration tests support.
package ecs

import (
	"github.com/weAutomateEverything/aws-sdk-go/aws"
	"github.com/weAutomateEverything/aws-sdk-go/awstesting/integration/smoke"
	"github.com/weAutomateEverything/aws-sdk-go/service/ecs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@ecs", func() {
		// FIXME remove custom region
		gucumber.World["client"] = ecs.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
