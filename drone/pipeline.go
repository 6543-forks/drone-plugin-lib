// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

// Pipeline being executed.
//
// Represents the full Drone environment that the plugin is executing in.
type Pipeline struct {
	Build  Build
	Repo   Repo
	Commit Commit
	Stage  Stage
	Step   Step
	SemVer SemVer
	CalVer CalVer
	System System
}
