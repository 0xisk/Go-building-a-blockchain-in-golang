package blockchain

import (
	"fmt"
	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks) - 1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockchain() *Blockchain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handler(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handler(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handler(err)
			err = item.Value(func(val []byte) error {
				lastHash = val

				return err
			})
			return err
		}
	})
	Handler(err)

	blockchain := Blockchain{lastHash, db}
	return &blockchain
}

