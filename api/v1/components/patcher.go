/*
Copyright © 2021 Dell Inc. or its subsidiaries. All Rights Reserved.

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

package components

// Patcher represents scheduler patcher container, which tries to patch Kubernetes scheduler
type Patcher struct {
	Enable            bool   `json:"enable,omitempty"`
	Image             *Image `json:"image,omitempty"`
	Manifest          string `json:"manifest,omitempty"`
	SrcConfigPath     string `json:"srcConfigPath,omitempty"`
	SrcPolicyPath     string `json:"srcPolicyPath,omitempty"`
	TargetConfigPath  string `json:"targetConfigPath,omitempty"`
	TargetPolicyPath  string `json:"targetPolicyPath,omitempty"`
	Interval          int    `json:"interval,omitempty"`
	RestoreOnShutdown bool   `json:"restoreOnShutdown,omitempty"`
	ConfigMapName     string `json:"configMapName,omitempty"`
}