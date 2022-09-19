/*
Copyright 2022 The Kubernetes Authors.

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

package validation

import (
	"fmt"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

// errorsAggregate converts all arg validation errors to a single error interface.
// if no errors, it will return nil.
func errorsAggregate(errors ...error) error {
	return utilerrors.NewAggregate(errors)
}

func validatePodRestartThreshold(podRestartThreshold int32) error {
	if podRestartThreshold < 1 {
		return fmt.Errorf("PodsHavingTooManyRestarts threshold not set")
	}
	return nil
}
