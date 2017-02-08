// Copyright 2016 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

// +build !darwin

package sysutils

// NotImplementedError is returned on platforms where GetExecPathFromPID is not
// implemented.
type NotImplementedError struct{}

// GetExecPathFromPID returns the process's executable path for given PID.
func GetExecPathFromPID(pid int) (string, error) {
	return "", NotImplementedError{}
}
