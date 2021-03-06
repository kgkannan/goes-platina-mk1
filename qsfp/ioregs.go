// Copyright © 2016-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qsfp

const ()

type regio8 byte
type regio16 [2]byte
type regio16r [2]byte
type regioi16 reg16

// Memory map
type regsio struct {
	Input    [2]regio8
	Output   [2]regio8
	Polarity [2]regio8
	Config   [2]regio8
}
