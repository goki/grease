// Copyright (c) 2023, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

//go:generate goki generate

// TestEnum is an enum type for testing,
// and it should not be used by external code.
type TestEnum int32 //enums:enum

const (
	TestValue1 TestEnum = iota
	TestValue2
)
