package blockchain

import "fmt"

type Basic struct {
	Chain []Block
}

// All data on-chain (inefficient)
func (b *Basic) AddPatient(p models.Patient) {
	data := fmt.Sprintf("%s|%s|%s", p.ID, p.Name, p.MedicalRecord) // Raw data
	newBlock := createBlockWithPoW(b.Chain, data)                  // Slow PoW mining
	b.Chain = append(b.Chain, newBlock)
}
