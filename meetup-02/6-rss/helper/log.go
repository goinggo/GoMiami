// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package helper : log.go provides log abstraction support.
package helper

import (
	"fmt"
	"time"
)

//** PUBLIC METHODS

// WriteStdout is used to write a system message directly to stdout
//  goRoutine: The Go routine making the call
//  namespace: The namespace the call is being made from
//  functionName: The function makeing the call
//	message: The message to be written
func WriteStdout(goRoutine string, namespace string, functionName string, message string) {
	fmt.Printf("%s : %s : %s : %s : %s\n", time.Now().Format("2006-01-02T15:04:05.000"), goRoutine, namespace, functionName, message)
}

// WriteStdoutf is used to write a formatted system message directly stdout
//  goRoutine: The Go routine making the call
//  namespace: The namespace the call is being made from
//  functionName: The function makeing the call
//  format: The message with formatting information
//  a: The set of parameters for the formatting
func WriteStdoutf(goRoutine string, namespace string, functionName string, format string, a ...interface{}) {
	WriteStdout(goRoutine, namespace, functionName, fmt.Sprintf(format, a...))
}
