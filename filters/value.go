/*
   Copyright The containerd Authors.

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

package filters

// Generic value type
type Value interface{}

// TypeConstraint specifies how a fieldpath should map to a type and tests 
// whether its value is appropriate, returning a transformation on success. 
type TypeConstraint interface {
	Test(raw string) (value Value, matches bool)
}

type TypeConstraintFunc func(raw string) (Value, bool)

func (fn TypeConstraintFunc) Test(raw string) (Value, bool) {
	return fn(raw)
}

// Doesn't apply any constraints and consequentially doesn't perform
// any transformations
var Self = func(raw string) (Value, bool) {
	return raw, true
}