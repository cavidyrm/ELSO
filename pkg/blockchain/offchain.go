package blockchain

type OffChain struct {
	Database map[string]models.Patient
}

func (o *OffChain) AddPatient(p models.Patient) {
	o.Database[p.ID] = p // No integrity guarantees
}
