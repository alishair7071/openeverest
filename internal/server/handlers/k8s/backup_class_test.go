// Copyright (C) 2026 The OpenEverest Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	backupv1alpha1 "github.com/openeverest/openeverest/v2/api/backup/v1alpha1"
)

func TestFilterBackupClassesByProvider(t *testing.T) {
	t.Parallel()

	classes := []backupv1alpha1.BackupClass{
		{
			ObjectMeta: metav1.ObjectMeta{Name: "pxc-provider"},
			Spec: backupv1alpha1.BackupClassSpec{
				ExecutionMode:      backupv1alpha1.BackupExecutionModeProviderManaged,
				SupportedProviders: []string{"pxc"},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "pxc-job"},
			Spec: backupv1alpha1.BackupClassSpec{
				ExecutionMode:      backupv1alpha1.BackupExecutionModeJob,
				SupportedProviders: []string{"pxc"},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "psmdb-provider"},
			Spec: backupv1alpha1.BackupClassSpec{
				ExecutionMode:      backupv1alpha1.BackupExecutionModeProviderManaged,
				SupportedProviders: []string{"psmdb"},
			},
		},
	}

	tests := []struct {
		name     string
		provider string
		want     []string
	}{
		{name: "provider managed classes", provider: "pxc", want: []string{"pxc-provider"}},
		{name: "unsupported provider", provider: "mysql", want: []string{}},
		{name: "empty input", provider: "pxc", want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			input := classes
			if tt.name == "empty input" {
				input = nil
			}
			filtered := filterBackupClassesByProvider(input, tt.provider)
			if len(filtered) != len(tt.want) {
				t.Fatalf("got %d classes, want %d", len(filtered), len(tt.want))
			}
			for i, name := range tt.want {
				if filtered[i].Name != name {
					t.Errorf("class %d = %q, want %q", i, filtered[i].Name, name)
				}
			}
		})
	}
}
