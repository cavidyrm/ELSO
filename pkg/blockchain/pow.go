package blockchain

import (
	"ELSO/pkg/models"
	"fmt"
)

type Pow struct {
	chain []models.Patient
}

// Ethereum-like PoW (high latency)
func (p *PoW) AddPatient(patient models.Patient) {
	data := fmt.Sprintf("%s|%s", patient.ID, patient.MedicalRecord)
	newBlock := mineBlockWithPoW(p.Chain, data, p.Difficulty) // CPU-intensive
	p.Chain = append(p.Chain, newBlock)
}
