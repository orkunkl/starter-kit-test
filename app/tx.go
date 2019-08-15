package app

import (
	"github.com/iov-one/weave"
)

// TxDecoder creates a Tx and unmarshals bytes into it
func TxDecoder(bz []byte) (weave.Tx, error) {
	tx := new(Tx)
	err := tx.Unmarshal(bz)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// make sure tx fulfills all interfaces
var _ weave.Tx = (*Tx)(nil)

// GetMsg returns a single message instance that is represented by this transaction.
func (tx *Tx) GetMsg() (weave.Msg, error) {
	return weave.ExtractMsgFromSum(tx.GetSum())
}
