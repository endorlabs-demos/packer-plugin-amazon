// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build windows
// +build windows

package chroot

import (
	"errors"
	"os"
)

func lockFile(*os.File) error {
	return errors.New("not supported on Windows")
}

func unlockFile(f *os.File) error {
	return nil
}
