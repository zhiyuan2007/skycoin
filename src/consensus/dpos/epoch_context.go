package dpos

import (
	"errors"

	"github.com/skycoin/skycoin/src/cipher"
)

type EpochContext struct {
	DposContext DposContext
	TimeStamp   int64
}

func NewEpochFromXX(dc DposContext, ts int64) *EpochContext {
	return &EpochContext{
		DposContext: dc,
		TimeStamp:   ts,
	}
}

func (ec *EpochContext) LookupValidator(now int64) (validator cipher.Address, err error) {
	validator = cipher.Address{}
	offset := now % epochInterval
	if offset%blockInterval != 0 {
		return cipher.Address{}, ErrInvalidMintBlockTime
	}
	offset /= blockInterval

	validators, err := ec.DposContext.GetValidators()
	if err != nil {
		return cipher.Address{}, err
	}
	validatorSize := len(validators)
	if validatorSize == 0 {
		return cipher.Address{}, errors.New("failed to lookup validator")
	}
	offset %= int64(validatorSize)
	return validators[offset], nil
}
