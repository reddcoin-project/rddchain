// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcchain_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/conformal/btcchain"
)

// TestMedianTime tests the medianTime implementation.
func TestMedianTime(t *testing.T) {
	tests := []struct {
		in         []int64
		wantOffset int64
		useDupID   bool
	}{
		// Not enough samples must result in an offset of 0.
		{in: []int64{1}, wantOffset: 0},
		{in: []int64{1, 2}, wantOffset: 0},
		{in: []int64{1, 2, 3}, wantOffset: 0},
		{in: []int64{1, 2, 3, 4}, wantOffset: 0},

		// Various number of entries.  The expected offset is only
		// updated on odd number of elements.
		{in: []int64{-13, 57, -4, -23, -12}, wantOffset: 12},
		{in: []int64{55, -13, 61, -52, 39, 55}, wantOffset: -39},
		{in: []int64{-62, -58, -30, -62, 51, -30, 15}, wantOffset: 30},
		{in: []int64{29, -47, 39, 54, 42, 41, 8, -33}, wantOffset: -39},
		{in: []int64{37, 54, 9, -21, -56, -36, 5, -11, -39}, wantOffset: 11},
		{in: []int64{57, -28, 25, -39, 9, 63, -16, 19, -60, 25}, wantOffset: -9},
		{in: []int64{-5, -4, -3, -2, -1}, wantOffset: 3, useDupID: true},

		// The offset stops being updated once the max number of entries
		// has been reached.  This is actually a bug from Bitcoin Core,
		// but since the time is ultimately used as a part of the
		// consensus rules, it must be mirrored.
		{in: []int64{-67, 67, -50, 24, 63, 17, 58, -14, 5, -32, -52}, wantOffset: -17},
		{in: []int64{-67, 67, -50, 24, 63, 17, 58, -14, 5, -32, -52, 45}, wantOffset: -17},
		{in: []int64{-67, 67, -50, 24, 63, 17, 58, -14, 5, -32, -52, 45, 4}, wantOffset: -17},

		// Offsets that are too far away from the local time should
		// be ignored.
		{in: []int64{-4201, 4202, -4203, 4204, -4205}, wantOffset: 0},

		// Excerise the condition where the median offset is greater
		// than the max allowed adjustment, but there is at least one
		// sample that is close enough to the current time to avoid
		// triggering a warning about an invalid local clock.
		{in: []int64{4201, 4202, 4203, 4204, -299}, wantOffset: 0},
	}

	// Modify the max number of allowed median time entries for these tests.
	btcchain.TstSetMaxMedianTimeEntries(10)
	defer btcchain.TstSetMaxMedianTimeEntries(200)

	for i, test := range tests {
		filter := btcchain.NewMedianTime()
		for j, offset := range test.in {
			id := strconv.Itoa(j)
			tOffset := time.Now().Add(time.Duration(offset) *
				time.Second)
			filter.AddTimeSample(id, tOffset)

			// Ensure the duplicate IDs are ignored.
			if test.useDupID {
				// Modify the offsets to ensure the final median
				// would be different if the duplicate is added.
				tOffset = tOffset.Add(time.Duration(offset) *
					time.Second)
				filter.AddTimeSample(id, tOffset)
			}
		}

		gotOffset := filter.Offset()
		wantOffset := time.Duration(test.wantOffset) * time.Second
		if gotOffset != wantOffset {
			t.Errorf("Offset #%d: unexpected offset -- got %v, "+
				"want %v", i, gotOffset, wantOffset)
			continue
		}

		adjustedTime := time.Unix(filter.AdjustedTime().Unix(), 0)
		wantTime := time.Now().Add(filter.Offset())
		wantTime = time.Unix(wantTime.Unix(), 0)
		if !adjustedTime.Equal(wantTime) {
			t.Errorf("AdjustedTime #%d: unexpected result -- got %v, "+
				"want %v", i, adjustedTime, wantTime)
			continue
		}
	}
}