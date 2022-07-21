package main

import (
  "github.com/cuemby/pulumi-equinix/sdk/go/equinix"
  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    equinix.NewMetalProject()
    return nil
  })
}
