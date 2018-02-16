/*
Copyright 2016 The Kubernetes Authors.

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
	"testing"

	"k8s.io/test-infra/prow/plugins"
)

// Make sure that our plugins are valid.
func TestPlugins(t *testing.T) {
	pa := &plugins.PluginAgent{}
	if err := pa.Load("../../plugins.yaml"); err != nil {
		t.Fatalf("Could not load plugins: %v.", err)
	}
}

func TestOptions_Validate(t *testing.T) {
	var testCases = []struct {
		name        string
		opt         options
		expectedErr bool
	}{
		{
			name: "all ok without dry-run",
			opt: options{
				dryRun: false,
			},
			expectedErr: false,
		},
		{
			name: "all ok with dry-run",
			opt: options{
				dryRun:  true,
				deckURL: "internet",
			},
			expectedErr: false,
		},
		{
			name: "missing deck endpoint with dry-run",
			opt: options{
				dryRun: true,
			},
			expectedErr: true,
		},
	}

	for _, testCase := range testCases {
		err := testCase.opt.Validate()
		if testCase.expectedErr && err == nil {
			t.Errorf("%s: expected an error but got none", testCase.name)
		}
		if !testCase.expectedErr && err != nil {
			t.Errorf("%s: expected no error but got one: %v", testCase.name, err)
		}
	}
}
