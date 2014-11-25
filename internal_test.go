// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
This test file is part of the rddchain package rather than than the
rddchain_test package so it can bridge access to the internals to properly test
cases which are either not possible or can't reliably be tested via the public
interface.  The functions are only exported while the tests are being run.
*/

package rddchain

import (
	"sort"
	"time"

	"github.com/reddcoin-project/rddutil"
)

// TstSetCoinbaseMaturity makes the ability to set the coinbase maturity
// available to the test package.
func TstSetCoinbaseMaturity(maturity int64) {
	coinbaseMaturity = maturity
}

// TstTimeSorter makes the internal timeSorter type available to the test
// package.
func TstTimeSorter(times []time.Time) sort.Interface {
	return timeSorter(times)
}

// TstCheckSerializedHeight makes the internal checkSerializedHeight function
// available to the test package.
func TstCheckSerializedHeight(coinbaseTx *rddutil.Tx, wantHeight int64) error {
	return checkSerializedHeight(coinbaseTx, wantHeight)
}

// TstSetMaxMedianTimeEntries makes the ability to set the maximum number of
// median tiem entries available to the test package.
func TstSetMaxMedianTimeEntries(val int) {
	maxMedianTimeEntries = val
}
