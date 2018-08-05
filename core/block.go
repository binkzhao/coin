package core

import (
	"time"
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

// 区块结构体定义
type Block struct {
	Timestamp     int64 // 区块创建时间戳
	Transactions  []*Transaction
	PrevBlockHash []byte // 前一个区块的哈希值
	Hash          []byte // 区块自身的哈希值，用于校验区块数据有效
	Nonce         int // 工作量证明值,用来校验数据的
}

// 创建创世区块
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// NewBlock create and return Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:time.Now().UnixNano(),
		Transactions: transactions,
		PrevBlockHash: prevBlockHash,
		Hash: []byte{},
	}

	// 采用工作量证明得出的新区块
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

// 把Block序列化为一个字节数组
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	HandleErr(err)

	return result.Bytes()
}

// 把字节数组反序列化为一个Block
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	if err := decoder.Decode(&block); err != nil {
		log.Panic(err)
	}

	return &block
}

// 把区块的所有交易ID做个hash处理
func (b *Block) HashTransaction() []byte {
	var txHashes [][]byte
	var txHash   [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// 设置区块自身的Hash值
//func (b *Block) SetHash()  {
//	timestamp := IntToHex(b.Timestamp)
//	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
//	hash := sha256.Sum256(headers)
//	b.Hash = hash[:]
//}

