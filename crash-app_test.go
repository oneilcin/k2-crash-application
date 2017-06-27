/*
Copyright 2017 K2 authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"testing"
)

func TestGetFailedTask(t *testing.T) {

	cases := []struct {
		name     string
		expected string
		arg      string
		method   func(string) string
	}{
		{
			name:     "Logfile contains no ansible tasks",
			expected: "Log does not contain any ansible tasks",
			arg:      "Test string",
			method:   getFailedTask,
		},
		{
			name:     "Parse failed task from logfile",
			expected: "kraken.config : Include configuration variables from defaults file",
			arg:      "TASK [kraken.config : Include configuration variables from defaults file]",
			method:   getFailedTask,
		},
		{
			name:     "Logfile contains unexpected characters",
			expected: "Log does not contain any ansible tasks",
			arg:      "TASK /.,;'l[\\&%%%%%%%%%[kraken.config : Include configuration variables from defaults file]",
			method:   getFailedTask,
		},
		{
			name:     "Logfile contains unexpected strings",
			expected: "kraken.config : Include configuration variables from defaults file",
			arg:      "TASK TASKTASK TASK TASK [i] TASK [kraken.config : Include configuration variables from defaults file] TASK",
			method:   getFailedTask,
		},
	}

	for _, tc := range cases {
		actual := tc.method(tc.arg)
		if actual != tc.expected {
			t.Errorf("Test '%s' failed, expected: '%s', got:  '%s'\n", tc.name, tc.expected, actual)
		}
	}
}

func TestSendLogs(t *testing.T) {
	file := "bogusfile"

	cases := []struct {
		name     string
		expected error
		arg      string
		method   func(string) error
	}{
		{
			name:     "Missing or incorrect logfile name",
			expected: fmt.Errorf("Error reading file in sendLogs function: open %v: no such file or directory", file),
			arg:      file,
			method:   sendLogs,
		},
	}

	for _, tc := range cases {
		actual := tc.method(tc.arg)
		if actual.Error() != tc.expected.Error() {
			t.Errorf("Test '%s' failed, expected: '%s', got:  '%s'\n", tc.name, tc.expected, actual)
		}
	}
}
