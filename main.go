package main

import (
    "bytes"
    "crypto/sha256"
    "encoding/gob"
    "fmt"
    "time"
)

// AIRequest represents a request to the AI system
type AIRequest struct {
    UserID    string
    Operation string
    Data      []byte
}

// AIResponse represents a response from the AI system
type AIResponse struct {
    Result []byte
    Err    error
}

// aiInterface simulates interaction with the AI system
func aiInterface(request AIRequest) AIResponse {
    // Simulate AI processing
    return AIResponse{Result: []byte("Processed: " + string(request.Data)), Err: nil}
}

// Block represents a single block in the blockchain
type Block struct {
    Index     int
    Timestamp int64
    PrevHash  []byte
    Hash      []byte
    Data      []byte
    Nonce     int
}

// Blockchain represents the entire chain of blocks
type Blockchain struct {
    Blocks []Block
}

// Transaction represents a single transaction in the blockchain
type Transaction struct {
    Sender    string
    Recipient string
    Amount    int
}

// createBlock creates a new block with the given data and previous block hash.
func createBlock(prevBlock Block, data []byte) Block {
    block := Block{
        Index:     prevBlock.Index + 1,
        Timestamp: time.Now().Unix(),
        PrevHash:  prevBlock.Hash,
        Data:      data,
        Nonce:     0, // Initialize Nonce for PoW
    }
    block = proofOfWork(block)
    return block
}

// calculateHash calculates the SHA-256 hash of a block using gob encoding.
func calculateHash(block Block) []byte {
    var buf bytes.Buffer
    encoder := gob.NewEncoder(&buf)
    err := encoder.Encode(block)
    if err != nil {
        // Handle the error gracefully
        panic(err)
    }
    hash := sha256.Sum256(buf.Bytes())
    return hash[:]
}

// proofOfWork performs the proof of work algorithm to find a valid hash
func proofOfWork(block Block) Block {
    target := bytes.Repeat([]byte{0}, 2) // Example difficulty: hash must start with two 0 bytes
    for {
        hash := calculateHash(block)
        if bytes.HasPrefix(hash, target) {
            block.Hash = hash
            break
        } else {
            block.Nonce++
        }
    }
    return block
}

// isValidChain validates the blockchain
func isValidChain(chain Blockchain) bool {
    for i := 1; i < len(chain.Blocks); i++ {
        currentBlock := chain.Blocks[i]
        prevBlock := chain.Blocks[i-1]
        if !bytes.Equal(currentBlock.PrevHash, prevBlock.Hash) {
            return false
        }
        if !bytes.Equal(calculateHash(currentBlock), currentBlock.Hash) {
            return false
        }
    }
    return true
}

func main() {
    // Initialize blockchain with genesis block
    genesisBlock := Block{Index: 0, Timestamp: time.Now().Unix(), PrevHash: []byte{}, Data: []byte("Genesis Block"), Nonce: 0}
    genesisBlock.Hash = calculateHash(genesisBlock)

    blockchain := Blockchain{Blocks: []Block{genesisBlock}}

    // Create a new block and add it to the blockchain
    newBlock := createBlock(genesisBlock, []byte("New Block Data"))
    blockchain.Blocks = append(blockchain.Blocks, newBlock)
    fmt.Println("New Block Created with Hash:", newBlock.Hash)

    // Validate the blockchain
    if isValidChain(blockchain) {
        fmt.Println("Blockchain is valid.")
    } else {
        fmt.Println("Blockchain is invalid.")
    }

    // AI request sample
    request := AIRequest{UserID: "user123", Operation: "process", Data: []byte("Sample Data")}
    response := aiInterface(request)
    fmt.Println("AI Response:", string(response.Result))
}
