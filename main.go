package main

import (
    "crypto/sha256"
    "strconv"
    "time"
)

type Block struct {
    Index     int
    Timestamp int64
    PrevHash  []byte
    Hash      []byte
    Data      []byte
    Nonce     int
}

type AIRequest struct {
    UserID    string
    Operation string
    Data      []byte
}

type AIResponse struct {
    Result []byte
    Err    error
}

func aiInterface(request AIRequest) AIResponse {
    // Simulate AI processing
    return AIResponse{Result: []byte("Processed: " + string(request.Data)), Err: nil}
}

func createBlock(prevBlock Block, data []byte) Block {
    block := Block{
        Index:     prevBlock.Index + 1,
        Timestamp: time.Now().Unix(),
        PrevHash:  prevBlock.Hash,
        Data:      data,
        Nonce:     0,
    }
    block.Hash = calculateHash(block)
    return block
}

func calculateHash(block Block) []byte {
    record := strconv.Itoa(block.Index) + strconv.FormatInt(block.Timestamp, 10) + string(block.PrevHash) + string(block.Data) + strconv.Itoa(block.Nonce)
    h := sha256.New()
    h.Write([]byte(record))
    return h.Sum(nil)
}

func main() {
    //block creation sample
    genesisBlock := Block{Index: 0, Timestamp: time.Now().Unix(), PrevHash: []byte{}, Data: []byte("Genesis Block"), Nonce: 0}
    genesisBlock.Hash = calculateHash(genesisBlock)

    newBlock := createBlock(genesisBlock, []byte("New Block Data"))
    println("New Block Created with Hash:", string(newBlock.Hash))

    //AI request sample
    request := AIRequest{UserID: "user123", Operation: "process", Data: []byte("Sample Data")}
    response := aiInterface(request)
    println("AI Response:", string(response.Result))
}
