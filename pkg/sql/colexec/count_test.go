// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import "testing"

func TestCount(t *testing.T) {
	tcs := []struct {
		tuples   tuples
		expected tuples
	}{
		{
			tuples:   tuples{{1}, {1}},
			expected: tuples{{2}},
		},
		{
			tuples:   tuples{{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}},
			expected: tuples{{10}},
		},
	}
	for _, tc := range tcs {
		// The tuples consisting of all nulls still count as separate rows, so if
		// we replace all values with nulls, we should get the same output.
		runTestsWithoutAllNullsInjection(t, []tuples{tc.tuples}, nil /* typs */, tc.expected, orderedVerifier, []int{0}, func(input []Operator) (Operator, error) {
			return NewCountOp(input[0]), nil
		})
	}
}
