package blockchain

import (
	"errors"
	"time"

	"github.com/nomadcoderkor/dinocoin/utils"
	"github.com/nomadcoderkor/dinocoin/wallet"
)

const (
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx
}

// Mempool Memeory Pool
var Mempool *mempool = &mempool{}

// Tx transaction
type Tx struct {
	ID        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

// TxIn struct
type TxIn struct {
	TxID      string `json:"txId"`
	Index     int    `json:"index"`
	Signature string `json:"signature"`
}

// TxOut struct
type TxOut struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

// UTxOut .
type UTxOut struct {
	TxID   string `json:"txId"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

func (t *Tx) getID() {
	t.ID = utils.Hash(t)
}

func (t *Tx) sign() {
	for _, txIn := range t.TxIns {
		txIn.Signature = wallet.Sign(t.ID, wallet.Wallet())
	}
}

func validate(tx *Tx) bool {
	valid := true
	for _, txIn := range tx.TxIns {
		prevTx := FindTx(Blockchain(), txIn.TxID)
		if prevTx == nil {
			valid = false
			break
		}
		address := prevTx.TxOuts[txIn.Index].Address
		valid = wallet.Verify(txIn.Signature, tx.ID, address)
		if !valid {
			break
		}
	}
	return valid
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false
Outer:
	for _, tx := range Mempool.Txs {
		for _, input := range tx.TxIns {
			if input.TxID == uTxOut.TxID && input.Index == uTxOut.Index {
				exists = true
				break Outer
			}
		}
	}
	return exists
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "COINBASE"},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getID()
	return &tx
}

// ErrorNoMoney .
var ErrorNoMoney = errors.New("잔액부족")

// ErrorNotValid .
var ErrorNotValid = errors.New("Transaction Invalid")

func makeTx(from, to string, amount int) (*Tx, error) {
	if BalanceByAddress(from, Blockchain()) < amount {
		return nil, ErrorNoMoney //errors.New("잔액이 부족 합니다.")
	}
	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	uTxOuts := UTxOutsByAddress(from, Blockchain())
	for _, uTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}
	if change := total - amount; change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getID()
	tx.sign()
	valid := validate(tx)
	if !valid {
		return nil, ErrorNotValid
	}
	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx(wallet.Wallet().Address, to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

// TxToConfirm 블록 생성시 멤풀의 트랜잭션을 블록에 추가해주고 멤플의 트랜잭션은 비워준다.
func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx(wallet.Wallet().Address)
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}
