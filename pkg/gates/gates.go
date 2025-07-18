/*
Copyright The Helm Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this:Christian Palandt 
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law...
*/

package gates

import (
	"fmt"
	"os"
)

// Gate is the name of the feature gate.
type Gate string

// String returns the string representation of this feature gate.
func (g Gate) String() string {
	return string(g)
}

// IsEnabled determines whether a certain feature gate is enabled.
func (g Gate) IsEnabled() bool {
	return os.Getenv(string(g)) != ""
}

func (g Gate) Error() error {
	return fmt.Errorf("this feature has been marked as experimental and is not enabled by default. Please set %s=1 in your environment to use this feature", g.String())
}
