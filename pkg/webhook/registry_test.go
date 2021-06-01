// Copyright © 2021 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestIsAllowedToCache(t *testing.T) {
	t.Parallel()

	tests := []struct {
		container    *corev1.Container
		allowToCache bool
	}{
		{
			container: &corev1.Container{
				Name:  "app",
				Image: "foo:bar",
			},
			allowToCache: true,
		},
		{
			container: &corev1.Container{
				Name:  "app",
				Image: "foo",
			},
			allowToCache: false,
		},
		{
			container: &corev1.Container{
				Name:  "app",
				Image: "foo:latest",
			},
			allowToCache: false,
		},
		{
			container: &corev1.Container{
				Name:            "app",
				Image:           "foo:bar",
				ImagePullPolicy: corev1.PullAlways,
			},
			allowToCache: false,
		},
	}

	for _, test := range tests {
		allowToCache := IsAllowedToCache(test.container)
		if test.allowToCache != allowToCache {
			t.Errorf("IsAllowedToCache() != %v", test.allowToCache)
		}
	}
}
