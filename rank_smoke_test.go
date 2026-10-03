package bitset

import "testing"

func TestRankMaxSmoke(t *testing.T) {
	for _, length := range []uint{0, 1, 64, 129} {
		b := New(length)
		_ = b.Rank(^uint(0) - 1)
		_ = b.Rank(^uint(0))
	}
}
