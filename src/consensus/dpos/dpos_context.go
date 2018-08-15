package dpos

import (
	"errors"

	"github.com/skycoin/skycoin/src/cipher"
)

type DposContext struct {
	candidate []cipher.Address
}

func NewDposContext() *DposContext {
	return &DposContext{
		candidate: []cipher.Address{},
	}
}

func (dc *DposContext) GetValidators() ([]cipher.Address, error) {
	if len(dc.candidate) == 0 {
		return dc.candidate, errors.New("zero validator")
	}
	return dc.candidate, nil
}

func (dc *DposContext) SetValidators(validators []cipher.Address) error {
	dc.candidate = validators
	return nil
}
