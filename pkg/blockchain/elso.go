package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"pkg/models"
)

type ELSO struct {
	Chain      []Block
	OffChain   map[string]string // patientID -> encrypted medical record
	Validators []string          // PBFT validator nodes
}

func NewELSO() *ELSO {
	return &ELSO{
		Chain:    []Block{genesisBlock()},
		OffChain: make(map[string]string),
	}
}

// Hybrid on/off-chain storage with PBFT
func (e *ELSO) AddPatient(p models.Patient) {
	e.OffChain[p.ID] = encrypt(p.MedicalRecord) // Off-chain
	data := p.ID + string(p.Visits)             // On-chain (hashed)
	hash := sha256.Sum256([]byte(data))
	newBlock := createBlock(e.Chain, hex.EncodeToString(hash[:]))
	if validatePBFT(newBlock, e.Validators) { // Fast consensus
		e.Chain = append(e.Chain, newBlock)
	}
}
