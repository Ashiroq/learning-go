package learninggo

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
)

// Transaction represents data carried in Blocks.
type Transaction struct {
	TXID uint64
}

// Block groups transactions
type Block struct {
	Index    uint32        // Block index in chain
	PrevHash [32]byte      // Previous block's hash
	Nonce    uint64        // Arbitrary variable used for mining
	Target   [32]byte      // Threshold block's hash must be lower than in order to be valid (mining)
	Txs      []Transaction // Stored transactions
}

func (b *Block) computeHash() [32]byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, b)
	return sha256.Sum256(buf.Bytes())
}

func verifyBlock(b *Block, prev *Block) error {
	hash := b.computeHash()

	if b.Index != prev.Index+1 {
		return errors.New("invalid index")
	} else if b.PrevHash != prev.computeHash() {
		return errors.New("invalid hash of previous block")
	} else if bytes.Compare(hash[:], b.Target[:]) == 1 {
		return errors.New("invalid hash")
	}
	return nil
}

// Chain stores a array of blocks that form blockchain.
type Chain struct {
	blockchain   []Block
	currentIndex uint64
	currentTx    uint64
}

// GenerateGenesis generates genesis block for empty chain
func (c *Chain) GenerateGenesis() error {
	if len(c.blockchain) != 0 {
		return errors.New("genesis can be only first block in a chain")
	}
	c.blockchain = append(c.blockchain, Block{0, [32]byte{0}, 0, [32]byte{0}, []Transaction{}})
	c.currentIndex++
	return nil
}

// AddBlock adds new block to chain with genesis already in
func (c *Chain) AddBlock(b Block) error {
	block, err := c.getLastBlock()
	if err != nil {
		return errors.New("cannot add without genesis block")
	}
	if err := verifyBlock(&b, block); err != nil {
		return err
	}
	c.blockchain = append(c.blockchain, b)
	c.currentIndex++
	return nil
}

func (c *Chain) getLastHash() ([32]byte, error) {
	if len(c.blockchain) == 0 {
		return [32]byte{}, errors.New("blockchain empty")
	}
	return c.blockchain[len(c.blockchain)-1].computeHash(), nil
}

func (c *Chain) getLastIndex() uint64 {
	return c.currentIndex
}

func (c *Chain) getLastBlock() (*Block, error) {
	if len(c.blockchain) == 0 {
		return nil, errors.New("blockchain empty")
	}
	return &c.blockchain[len(c.blockchain)-1], nil
}
