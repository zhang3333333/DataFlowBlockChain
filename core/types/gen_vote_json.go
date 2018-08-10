// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"go-ethereum1/common/hexutil"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var _ = (*VoteMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (v Vote) MarshalJSON() ([]byte, error) {
	type Vote struct {
		DataHash common.Hash  `json:"txHash"	gencodec:"required"`
		IsExist  *hexutil.Big `json:"isExist" gencodec:"required"`
		NodeID   string       `json:"nodeID"	gencodec:"required"`
		Func     common.Hash  `json:"func"	gencodec:"required"`
		V        *hexutil.Big `json:"v" 		gencodec:"required"`
		R        *hexutil.Big `json:"r"		gencodec:"required"`
		S        *hexutil.Big `json:"s"		gencodec:"required"`
		PubKey   []byte       `json:"pubKey"	gencodec:"required"`
		Hash     common.Hash  `json:"hash"`
	}
	var enc Vote
	enc.DataHash = v.DataHash
	enc.IsExist = (*hexutil.Big)(v.IsExist)
	enc.NodeID = v.NodeID
	enc.Func = v.Func
	enc.V = (*hexutil.Big)(v.V)
	enc.R = (*hexutil.Big)(v.R)
	enc.S = (*hexutil.Big)(v.S)
	enc.PubKey = v.PubKey
	enc.Hash = v.Hash()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (v *Vote) UnmarshalJSON(input []byte) error {
	type Vote struct {
		DataHash *common.Hash `json:"txHash"	gencodec:"required"`
		IsExist  *hexutil.Big `json:"isExist" gencodec:"required"`
		NodeID   *string      `json:"nodeID"	gencodec:"required"`
		Func     *common.Hash `json:"func"	gencodec:"required"`
		V        *hexutil.Big `json:"v" 		gencodec:"required"`
		R        *hexutil.Big `json:"r"		gencodec:"required"`
		S        *hexutil.Big `json:"s"		gencodec:"required"`
		PubKey   []byte       `json:"pubKey"	gencodec:"required"`
	}
	var dec Vote
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.DataHash != nil {
		v.DataHash = *dec.DataHash
	}
	if dec.IsExist == nil {
		return errors.New("missing required field 'isExist' for Vote")
	}
	v.IsExist = (*big.Int)(dec.IsExist)
	if dec.NodeID != nil {
		v.NodeID = *dec.NodeID
	}
	if dec.Func != nil {
		v.Func = *dec.Func
	}
	if dec.V != nil {
		v.V = (*big.Int)(dec.V)
	}
	if dec.R != nil {
		v.R = (*big.Int)(dec.R)
	}
	if dec.S != nil {
		v.S = (*big.Int)(dec.S)
	}
	if dec.PubKey != nil {
		v.PubKey = dec.PubKey
	}
	return nil
}
