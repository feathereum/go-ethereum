// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xe1\xe0\x94\x01\x9e{\xf4\x1c\xd9\xc1\xc8\xe0\xcb\u0479;>=\x00\xf0C\x00\xf1\x8a\xc0\x1bi\x19hf\x87\x00\x00\x00"
const testnetAllocData = "\xe3\xe2\x94\xf8\xc4\x18\xcc\x14\xb1\x8a\xfd\xedB&P\xb2'\xfd\x05\x1b\xdfm\x8b\x8c\xa7\x81\n\xff\x15\x19\xf7\u0680\x00\x00\x00"
const rinkebyAllocData = "\xe1\xe0\x94\x01\x9e{\xf4\x1c\xd9\xc1\xc8\xe0\xcb\u0479;>=\x00\xf0C\x00\xf1\x8am\u01853\x17\x16\x04\x00\x00\x00"
