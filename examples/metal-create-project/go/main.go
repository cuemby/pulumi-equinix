package main

import (
	"github.com/cuemby/pulumi-equinix/sdk/go/equinix"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) (err error) {
		conf := config.New(ctx, "")

		p, err := equinix.NewMetalProject(ctx, ctx.Project(), &equinix.MetalProjectArgs{
			BackendTransfer: pulumi.Bool(true),
			BgpConfig: &equinix.MetalProjectBgpConfigArgs{
				Asn:            pulumi.Int(6500),
				DeploymentType: pulumi.String("local"),
				Md5:            pulumi.String("md5Password"),
			},
			Name: pulumi.String(conf.Require("projectName")),
		})

		if err != nil {
			return err
		}

		pulumi.Sprintf("Project name: %v", p.Name)

		return nil
	})
}
