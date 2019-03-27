package aws

import (
	"github.com/lyraproj/aws-example/cmd/goplugin-aws/resource"
	"github.com/lyraproj/pcore/pcore"
	"github.com/lyraproj/pcore/px"
	"github.com/lyraproj/servicesdk/grpc"
)

// Start this provider
func Start() {
	pcore.Do(func(c px.Context) {
		s := resource.Server(c)
		grpc.Serve(c, s)
	})
}
