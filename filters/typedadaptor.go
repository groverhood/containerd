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

// Layer a type constraint on top of a result yielded by an adaptor.
type TypedAdaptor interface {
	TypedField(fieldpath []string, typecons TypeConstraint) (value Value, valid bool)
}

type TypedAdaptorFunc func(fieldpath []string, typecons TypeConstraint) (Value, bool)

func (fn TypedAdaptorFunc) Validate(fieldpath []string, typecons TypeConstraint) (Value, bool) {
	return fn(adaptor, typecons)
}
