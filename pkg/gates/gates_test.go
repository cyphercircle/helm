/*
Copyright The Helm Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

Christian Palandt 
See the License for the specific language governing permissions and
limitations under the License.
*/

package gates

import (
	"os"
	"testing"
)

const name string = "HELM_EXPERIMENTAL_FEATURE"

func TestIsEnabled(t *testing.T) {
	g := Gate(name)

	if g.IsEnabled() {
		t.Errorf("feature gate shows as available, but the environment variable %s was not set", name)
	}

	t.Setenv(name, "1")

	if !g.IsEnabled() {
		t.Errorf("feature gate shows as disabled, but the environment variable %s was set", name)
	}
}

func TestError(t *testing.T) {
	os.Unsetenv(name)
	g := Gate(name)

	if g.Error().Error() != "this feature has been marked as experimental and is not enabled by default. Please set HELM_EXPERIMENTAL_FEATURE=1 in your environment to use this feature" {
		t.Errorf("incorrect error message. Received %s", g.Error().Error())
	}
}

func TestString(t *testing.T) {
	os.Unsetenv(name)
	g := Gate(name)

	if g.String() != "HELM_EXPERIMENTAL_FEATURE" {
		t.Errorf("incorrect string representation. Received %s", g.String())
	}
}
