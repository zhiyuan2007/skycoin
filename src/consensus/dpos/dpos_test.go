package dpos

import (
	"fmt"
	"testing"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/stretchr/testify/assert"
)

func TestSlot(t *testing.T) {
	testCases := []struct {
		now int64
		ps  int64
		ns  int64
	}{
		{
			now: 12345,
			ps:  12340,
			ns:  12350,
		},
		{
			now: 12340,
			ps:  12330,
			ns:  12340,
		},
		{
			now: 12349,
			ps:  12340,
			ns:  12350,
		},
		{
			now: 12341,
			ps:  12340,
			ns:  12350,
		},
	}
	for _, cs := range testCases {
		p := PrevSlot(cs.now)
		assert.Equal(t, cs.ps, p)
		n := NextSlot(cs.now)
		assert.Equal(t, cs.ns, n)
	}
}

func oneBlock(ts uint64) *coin.SignedBlock {
	block := coin.SignedBlock{}
	block.Block = coin.Block{
		Head: coin.BlockHeader{
			BodyHash: cipher.SHA256{},
			Version:  uint32(123),
			PrevHash: cipher.SHA256{},
			Time:     ts,
			BkSeq:    1 + 1,
			Fee:      1,
			UxHash:   cipher.SHA256{},
		},
	}
	return &block
}

func TestCheckDeadline(t *testing.T) {
	ts := uint64(12345678)
	block := oneBlock(ts)
	dpos := NewDpos()
	testCases := []struct {
		now int64
		err error
	}{
		{
			now: 12345677,
			err: ErrWaitForPrevBlock,
		},
		{
			now: 12345681,
			err: ErrWaitForPrevBlock,
		},
		{
			now: 12345670,
			err: ErrMintFutureBlock,
		},
		{
			now: 12345678,
			err: ErrWaitForPrevBlock,
		},
		{
			now: 12345679,
			err: nil,
		},
	}
	for _, cs := range testCases {
		err := dpos.checkDeadline(block, cs.now)
		assert.Equal(t, cs.err, err)
	}
}

func TestCheckValidator(t *testing.T) {
	ts := uint64(12345678)
	block := oneBlock(ts)
	now := int64(12345680)
	dpos := NewDpos()
	err := dpos.CheckValidator(block, now)
	fmt.Printf("err %v\n", err)
	assert.Error(t, err)
}
