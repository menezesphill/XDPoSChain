// Copyright 2014 The go-ethereum Authors
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

package trie

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// MissingNodeError is returned by the trie functions (TryGet, TryUpdate, TryDelete)
// in the case where a trie node is not present in the local database. Contains
// information necessary for retrieving the missing node through an ODR service.
//
// NodeHash is the hash of the missing node
// RootHash is the original root of the trie that contains the node
// KeyPrefix is the prefix that leads from the root to the missing node (hex encoded)
// KeySuffix (optional) contains the rest of the key we were looking for, gives a
//  hint on which further nodes should also be retrieved (hex encoded)
type MissingNodeError struct {
	RootHash, NodeHash   common.Hash
	KeyPrefix, KeySuffix []byte
}

func (err *MissingNodeError) Error() string {
	return fmt.Sprintf("Missing trie node %064x", err.NodeHash)
}