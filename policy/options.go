// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package policy

import (
	"github.com/google/cel-policy-templates-go/policy/model"
	"github.com/google/cel-policy-templates-go/policy/runtime"

	"github.com/google/cel-go/interpreter"
)

// EngineOption is a functional option for configuring the policy engine.
type EngineOption func(*Engine) (*Engine, error)

// Selector functions take a compiled representation of a policy instance 'selector' and the input
// argument set to determine whether the policy instance is applicable to the current evaluation
// context.
type Selector func(model.Selector, interpreter.Activation) bool

// Selectors is a functional option which may be configured to select a subset of policy instances
// which are applicable to the current evaluation context.
func Selectors(selectors ...Selector) EngineOption {
	return func(e *Engine) (*Engine, error) {
		e.selectors = selectors
		return e, nil
	}
}

// RangeLimit sets the range limit supported by the compilation and runtime components.
func RangeLimit(limit int) EngineOption {
	return func(e *Engine) (*Engine, error) {
		e.limits.RangeLimit = limit
		return e, nil
	}
}

// RuntimeTemplateOptions collects a set of runtime specific options to be configured on runtime
// templates.
func RuntimeTemplateOptions(rtOpts ...runtime.TemplateOption) EngineOption {
	return func(e *Engine) (*Engine, error) {
		e.rtOpts = append(e.rtOpts, rtOpts...)
		return e, nil
	}
}
