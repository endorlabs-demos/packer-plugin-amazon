// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chroot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// StepCheckRootDevice makes sure the root device on the AMI is EBS-backed.
type StepCheckRootDevice struct{}

func (s *StepCheckRootDevice) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	image := state.Get("source_image").(*ec2.Image)
	ui := state.Get("ui").(packersdk.Ui)

	ui.Say("Checking the root device on source AMI...")

	// It must be EBS-backed otherwise the build won't work
	if *image.RootDeviceType != "ebs" {
		err := fmt.Errorf("The root device of the source AMI must be EBS-backed.")
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepCheckRootDevice) Cleanup(multistep.StateBag) {}
