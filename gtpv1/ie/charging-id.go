// Copyright 2019-2021 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie

import (
	"encoding/binary"
	"io"
)

// NewChargingID creates a new ChargingID IE.
func NewChargingID(id uint32) *IE {
	return newUint32ValIE(ChargingID, id)
}

// ChargingID returns the ChargingID value in uint32 if the type of IE matches.
func (i *IE) ChargingID() (uint32, error) {
	if i.Type != ChargingID {
		return 0, &InvalidTypeError{Type: i.Type}
	}
	if len(i.Payload) < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	return binary.BigEndian.Uint32(i.Payload), nil
}

// MustChargingID returns ChargingID in uint32, ignoring errors.
// This should only be used if it is assured to have the value.
func (i *IE) MustChargingID() uint32 {
	v, _ := i.ChargingID()
	return v
}
