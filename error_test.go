// Copyright (c) 2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rddchain_test

import (
	"testing"

	"github.com/reddcoin-project/rddchain"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	tests := []struct {
		in   rddchain.ErrorCode
		want string
	}{
		{rddchain.ErrDuplicateBlock, "ErrDuplicateBlock"},
		{rddchain.ErrBlockTooBig, "ErrBlockTooBig"},
		{rddchain.ErrBlockVersionTooOld, "ErrBlockVersionTooOld"},
		{rddchain.ErrInvalidTime, "ErrInvalidTime"},
		{rddchain.ErrTimeTooOld, "ErrTimeTooOld"},
		{rddchain.ErrTimeTooNew, "ErrTimeTooNew"},
		{rddchain.ErrDifficultyTooLow, "ErrDifficultyTooLow"},
		{rddchain.ErrUnexpectedDifficulty, "ErrUnexpectedDifficulty"},
		{rddchain.ErrHighHash, "ErrHighHash"},
		{rddchain.ErrBadMerkleRoot, "ErrBadMerkleRoot"},
		{rddchain.ErrBadCheckpoint, "ErrBadCheckpoint"},
		{rddchain.ErrForkTooOld, "ErrForkTooOld"},
		{rddchain.ErrCheckpointTimeTooOld, "ErrCheckpointTimeTooOld"},
		{rddchain.ErrNoTransactions, "ErrNoTransactions"},
		{rddchain.ErrTooManyTransactions, "ErrTooManyTransactions"},
		{rddchain.ErrNoTxInputs, "ErrNoTxInputs"},
		{rddchain.ErrNoTxOutputs, "ErrNoTxOutputs"},
		{rddchain.ErrTxTooBig, "ErrTxTooBig"},
		{rddchain.ErrBadTxOutValue, "ErrBadTxOutValue"},
		{rddchain.ErrDuplicateTxInputs, "ErrDuplicateTxInputs"},
		{rddchain.ErrBadTxInput, "ErrBadTxInput"},
		{rddchain.ErrBadCheckpoint, "ErrBadCheckpoint"},
		{rddchain.ErrMissingTx, "ErrMissingTx"},
		{rddchain.ErrUnfinalizedTx, "ErrUnfinalizedTx"},
		{rddchain.ErrDuplicateTx, "ErrDuplicateTx"},
		{rddchain.ErrOverwriteTx, "ErrOverwriteTx"},
		{rddchain.ErrImmatureSpend, "ErrImmatureSpend"},
		{rddchain.ErrDoubleSpend, "ErrDoubleSpend"},
		{rddchain.ErrSpendTooHigh, "ErrSpendTooHigh"},
		{rddchain.ErrBadFees, "ErrBadFees"},
		{rddchain.ErrTooManySigOps, "ErrTooManySigOps"},
		{rddchain.ErrFirstTxNotCoinbase, "ErrFirstTxNotCoinbase"},
		{rddchain.ErrMultipleCoinbases, "ErrMultipleCoinbases"},
		{rddchain.ErrBadCoinbaseScriptLen, "ErrBadCoinbaseScriptLen"},
		{rddchain.ErrBadCoinbaseValue, "ErrBadCoinbaseValue"},
		{rddchain.ErrMissingCoinbaseHeight, "ErrMissingCoinbaseHeight"},
		{rddchain.ErrBadCoinbaseHeight, "ErrBadCoinbaseHeight"},
		{rddchain.ErrScriptMalformed, "ErrScriptMalformed"},
		{rddchain.ErrScriptValidation, "ErrScriptValidation"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestRuleError tests the error output for the RuleError type.
func TestRuleError(t *testing.T) {
	tests := []struct {
		in   rddchain.RuleError
		want string
	}{
		{
			rddchain.RuleError{Description: "duplicate block"},
			"duplicate block",
		},
		{
			rddchain.RuleError{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
