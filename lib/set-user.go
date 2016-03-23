// Copyright 2015 The appc Authors
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

package lib

import (
	"fmt"

	"github.com/appc/acbuild/util"

	"github.com/appc/spec/schema"
)

// SetUser sets the user the pod will run as in the untarred ACI stored at
// a.CurrentACIPath.
func (a *ACBuild) SetUser(user string) (err error) {
	if err = a.lock(); err != nil {
		return err
	}
	defer func() {
		if err1 := a.unlock(); err == nil {
			err = err1
		}
	}()

	if user == "" {
		return fmt.Errorf("user cannot be empty")
	}
	fn := func(s *schema.ImageManifest) error {
		if s.App == nil {
			s.App = newManifestApp()
		}
		s.App.User = user
		return nil
	}
	return util.ModifyManifest(fn, a.CurrentACIPath)
}
