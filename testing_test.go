// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0
// Modified: packege name changed

package lrucache_test

import (
	"crypto/rand"
	"math"
	"math/big"
	"testing"
)

func getRand(tb testing.TB) int64 {
	out, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		tb.Fatal(err)
	}
	return out.Int64()
}
