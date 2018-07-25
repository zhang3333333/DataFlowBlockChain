// Copyright 2018 The go-ethereum Authors
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

package rawdb

import (
	"github.com/ethereum/go-ethereum/common"
	"DataFlowBlockChain/core/types"
	"log"
	"github.com/ethereum/go-ethereum/rlp"
	"fmt"
)

// ReadTxLookupEntry retrieves the positional metadata associated with a transaction
// hash to allow retrieving the transaction or receipt by hash.
func ReadTxLookupEntry(db DatabaseReader, hash common.Hash) (common.Hash, uint64, uint64) {
	data, _ := db.Get(txLookupKey(hash))
	if len(data) == 0 {
		return common.Hash{}, 0, 0
	}
	var entry TxLookupEntry
	if err := rlp.DecodeBytes(data, &entry); err != nil {
		log.Fatal("Invalid transaction lookup entry RLP", "hash", hash, "err", err)
		return common.Hash{}, 0, 0
	}
	return entry.BlockHash, entry.BlockIndex, entry.Index
}

// WriteTxLookupEntries stores a positional metadata for every transaction from
// a block, enabling hash based transaction and receipt lookups.
func WriteTxLookupEntries(db DatabaseWriter, block *types.Block) {
	for i, tx := range block.Transactions() {
		entry := TxLookupEntry{
			BlockHash:  block.Hash(),
			BlockIndex: block.NumberU64(),
			Index:      uint64(i),
		}
		data, err := rlp.EncodeToBytes(entry)
		if err != nil {
			log.Fatal("Failed to encode transaction lookup entry", "err", err)
		}
		if err := db.Put(txLookupKey(tx.Hash()), data); err != nil {
			log.Fatal("Failed to store transaction lookup entry", "err", err)
		}
	}
}

// DeleteTxLookupEntry removes all transaction data associated with a hash.
func DeleteTxLookupEntry(db DatabaseDeleter, hash common.Hash) {
	db.Delete(txLookupKey(hash))
}

// ReadTransaction retrieves a specific transaction from the database, along with
// its added positional metadata.
func ReadTransaction(db DatabaseReader, hash common.Hash) (*types.Transaction, common.Hash, uint64, uint64) {
	blockHash, blockNumber, txIndex := ReadTxLookupEntry(db, hash)
	if blockHash == (common.Hash{}) {
		return nil, common.Hash{}, 0, 0
	}
	body := ReadBody(db, blockHash, blockNumber)
	if body == nil || len(body.Transactions) <= int(txIndex) {
		log.Fatal("Transaction referenced missing", "number", blockNumber, "hash", blockHash, "index", txIndex)
		return nil, common.Hash{}, 0, 0
	}
	return body.Transactions[txIndex], blockHash, blockNumber, txIndex
}

// WriteVtLookupEntries store a positional metadata for every vote from a block,
func WriteVtLookupEntries(db DatabaseWriter, block *types.Block) {
	txs := block.Transactions()
	for i, vt := range block.VoteCollection() {
		entry := VtLookupEntry{
			BlockHash:  block.Hash(),
			BlockIndex: block.NumberU64(),
			Index:      uint64(i),
		}
		data, err := rlp.EncodeToBytes(entry)
		if err != nil {
			log.Fatal("Failed to encode vt lookup entry", "err", err)
		}
		// @mode
		fmt.Println(vtLookupKey(txs[vt.TxIndex.Uint64()].Hash(), vt.NodeID))
		if err := db.Put(vtLookupKey(txs[vt.TxIndex.Uint64()].Hash(), vt.NodeID), data); err != nil {
			log.Fatal("Failed to store vt lookup entry", "err", err)
		}
	}
}

// ReadVtLookupEntry retrieves the positional vote metadata associated with a transaction
// hash to allow retrieving the vote by hash.
func ReadVtLookupEntry(db DatabaseReader, hash common.Hash, nodeID string) (common.Hash, uint64, uint64) {
	// @mode
	fmt.Println(vtLookupKey(hash, nodeID))
	data, _ := db.Get(vtLookupKey(hash, nodeID))
	if len(data) == 0 {
		return common.Hash{}, 0, 0
	}
	var entry VtLookupEntry
	if err := rlp.DecodeBytes(data, &entry); err != nil {
		log.Fatal("Invalid transaction lookup entry RLP", "hash", hash, "err", err)
		return common.Hash{}, 0, 0
	}
	return entry.BlockHash, entry.BlockIndex, entry.Index
}

// DeleteVtLookupEntry removes vote data associated with a hash and node.
func DeleteVtLookupEntry(db DatabaseDeleter, hash common.Hash, NodeID string) {
	db.Delete(vtLookupKey(hash, NodeID))
}



// ReadVote retrieves a specific transaction from the database, along with
// its added positional metadata.
func ReadVote(db DatabaseReader, hash common.Hash, nodeID string) (*types.Vote, common.Hash, uint64, uint64) {
	blockHash, blockNumber, voteIndex := ReadVtLookupEntry(db, hash, nodeID)
	if blockHash == (common.Hash{}) {
		return nil, common.Hash{}, 0, 0
	}
	body := ReadBody(db, blockHash, blockNumber)
	if body == nil || len(body.Votes) <= int(voteIndex) {
		log.Fatal("Vote referenced missing", "number", blockNumber, "hash", blockHash, "index", voteIndex)
		return nil, common.Hash{}, 0, 0
	}
	return body.Votes[voteIndex], blockHash, blockNumber, voteIndex
}



// ReadBloomBits retrieves the compressed bloom bit vector belonging to the given
// section and bit index from the.
func ReadBloomBits(db DatabaseReader, bit uint, section uint64, head common.Hash) ([]byte, error) {
	return db.Get(bloomBitsKey(bit, section, head))
}

// WriteBloomBits stores the compressed bloom bits vector belonging to the given
// section and bit index.
func WriteBloomBits(db DatabaseWriter, bit uint, section uint64, head common.Hash, bits []byte) {
	if err := db.Put(bloomBitsKey(bit, section, head), bits); err != nil {
		log.Fatal("Failed to store bloom bits", "err", err)
	}
}