// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pdata // import "go.opentelemetry.io/collector/model/pdata"

// This file contains aliases to data structures that are common for all
// signal types, such as timestamps, attributes, etc.

import "go.opentelemetry.io/collector/model/internal/pdata"

// ValueType is an alias for pdata.ValueType type.
type ValueType = pdata.ValueType

// AttributeValueType is an alias for pdata.ValueType type.
// Deprecated: [v0.48.0] Use ValueType instead.
type AttributeValueType = pdata.ValueType

const (
	ValueTypeEmpty  = pdata.ValueTypeEmpty
	ValueTypeString = pdata.ValueTypeString
	ValueTypeInt    = pdata.ValueTypeInt
	ValueTypeDouble = pdata.ValueTypeDouble
	ValueTypeBool   = pdata.ValueTypeBool
	ValueTypeMap    = pdata.ValueTypeMap
	ValueTypeArray  = pdata.ValueTypeArray
	ValueTypeBytes  = pdata.ValueTypeBytes

	// Deprecated: [v0.48.0] Use ValueTypeEmpty instead.
	AttributeValueTypeEmpty = pdata.ValueTypeEmpty

	// Deprecated: [v0.48.0] Use ValueTypeString instead.
	AttributeValueTypeString = pdata.ValueTypeString

	// Deprecated: [v0.48.0] Use ValueTypeInt instead.
	AttributeValueTypeInt = pdata.ValueTypeInt

	// Deprecated: [v0.48.0] Use ValueTypeDouble instead.
	AttributeValueTypeDouble = pdata.ValueTypeDouble

	// Deprecated: [v0.48.0] Use ValueTypeBool instead.
	AttributeValueTypeBool = pdata.ValueTypeBool

	// Deprecated: [v0.48.0] Use ValueTypeMap instead.
	AttributeValueTypeMap = pdata.ValueTypeMap

	// Deprecated: [v0.48.0] Use ValueTypeArray instead.
	AttributeValueTypeArray = pdata.ValueTypeArray

	// Deprecated: [v0.48.0] Use ValueTypeBytes instead.
	AttributeValueTypeBytes = pdata.ValueTypeBytes
)

// Value is an alias for pdata.Value struct.
type Value = pdata.Value

// Deprecated: [v0.48.0] Use Value instead.
type AttributeValue = pdata.Value

// Aliases for functions to create pdata.Value.
var (
	NewValueEmpty  = pdata.NewValueEmpty
	NewValueString = pdata.NewValueString
	NewValueInt    = pdata.NewValueInt
	NewValueDouble = pdata.NewValueDouble
	NewValueBool   = pdata.NewValueBool
	NewValueMap    = pdata.NewValueMap
	NewValueArray  = pdata.NewValueArray
	NewValueBytes  = pdata.NewValueBytes

	// Deprecated: [v0.48.0] Use NewValueEmpty instead.
	NewAttributeValueEmpty = pdata.NewValueEmpty

	// Deprecated: [v0.48.0] Use NewValueString instead.
	NewAttributeValueString = pdata.NewValueString

	// Deprecated: [v0.48.0] Use NewValueInt instead.
	NewAttributeValueInt = pdata.NewValueInt

	// Deprecated: [v0.48.0] Use NewValueDouble instead.
	NewAttributeValueDouble = pdata.NewValueDouble

	// Deprecated: [v0.48.0] Use NewValueBool instead.
	NewAttributeValueBool = pdata.NewValueBool

	// Deprecated: [v0.48.0] Use NewValueMap instead.
	NewAttributeValueMap = pdata.NewValueMap

	// Deprecated: [v0.48.0] Use NewValueArray instead.
	NewAttributeValueArray = pdata.NewValueArray

	// Deprecated: [v0.48.0] Use NewValueBytes instead.
	NewAttributeValueBytes = pdata.NewValueBytes
)

// AttributeMap is an alias for pdata.AttributeMap struct.
type AttributeMap = pdata.AttributeMap

// Aliases for functions to create pdata.AttributeMap.
var (
	NewAttributeMap        = pdata.NewAttributeMap
	NewAttributeMapFromMap = pdata.NewAttributeMapFromMap
)