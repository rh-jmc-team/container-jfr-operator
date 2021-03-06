// Copyright (c) 2021 Red Hat, Inc.
//
// The Universal Permissive License (UPL), Version 1.0
//
// Subject to the condition set forth below, permission is hereby granted to any
// person obtaining a copy of this software, associated documentation and/or data
// (collectively the "Software"), free of charge and under any and all copyright
// rights in the Software, and any and all patent rights owned or freely
// licensable by each licensor hereunder covering either (i) the unmodified
// Software as contributed to or provided by such licensor, or (ii) the Larger
// Works (as defined below), to deal in both
//
// (a) the Software, and
// (b) any piece of software and/or hardware listed in the lrgrwrks.txt file if
// one is included with the Software (each a "Larger Work" to which the Software
// is contributed by such licensors),
//
// without restriction, including without limitation the rights to copy, create
// derivative works of, display, perform, and distribute the Software and make,
// use, sell, offer for sale, import, export, have made, and have sold the
// Software and the Larger Work(s), and to sublicense the foregoing rights on
// either these or other terms.
//
// This license is subject to the following condition:
// The above copyright notice and either this complete permission notice or at
// a minimum a reference to the UPL must be included in all copies or
// substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package client

import (
	"fmt"
)

// RecordingDescriptor contains various metadata for a particular
// flight recording retrieved from the JVM
type RecordingDescriptor struct {
	// An identifier used by the JVM to uniquely identify a recording
	ID int64 `json:"id"`
	// Name of the recording specified during creation
	Name string `json:"name"`
	// State of the recording within its lifecycle
	State string `json:"state"`
	// Time when the recording first started, in milliseconds since Unix epoch
	StartTime int64 `json:"startTime"`
	// How long the recording was configured to run for, in milliseconds
	Duration int64 `json:"duration"`
	// Whether the recording was configured to record indefinitely
	Continuous bool `json:"continuous"`
	// Whether this recording was dumped to disk in the host containing the JVM
	ToDisk bool `json:"toDisk"`
	// The maximum configured size of the recording file
	MaxSize int64 `json:"maxSize"`
	// The maximum configured age of recorded events
	MaxAge int64 `json:"maxAge"`
	// URL to download the raw flight recording file
	DownloadURL string `json:"downloadUrl"`
	// URL to the automated analysis report for this recording
	ReportURL string `json:"reportUrl"`
}

// SavedRecording represents a recording file that has been archived in
// persistent storage by Container JFR
type SavedRecording struct {
	Name        string `json:"name"`
	DownloadURL string `json:"downloadUrl"`
	ReportURL   string `json:"reportUrl"`
}

// TargetAddress contains an address that Container JFR can use to connect
// to a particular JVM
type TargetAddress struct {
	Host string
	Port int32
}

func (target TargetAddress) String() string {
	return fmt.Sprintf("%s:%d", target.Host, target.Port)
}
