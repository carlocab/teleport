/*
Copyright 2022 Gravitational, Inc.

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

package aws

import "strings"

const (
	// TagKeyTeleportCreated defines a tag key that indicates the the cloud
	// resource is created by Teleport.
	TagKeyTeleportCreated = "teleport.dev/created"

	// TagKeyTeleportManaged defines a tag key that indicates the the cloud
	// resource is being managed by Teleport.
	TagKeyTeleportManaged = "teleport.dev/managed"

	// TagValueTrue is the tag value "true" in string format.
	TagValueTrue = "true"
)

// IsTagValueTrue checks whether a tag value is true.
func IsTagValueTrue(value string) bool {
	// Here doing a lenient negative check. Any other value is assumed to be
	// true.
	switch strings.ToLower(value) {
	case "false", "no", "disable", "disabled":
		return false
	default:
		return true
	}
}
