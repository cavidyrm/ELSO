package models

type Patient struct {
	ID            string `csv:"id"`
	Name          string `csv:"name"`
	Age           int    `csv:"age"`
	MedicalRecord string `csv:"medical_record"` // Sensitive data (off-chain)
	Visits        int    `csv:"visits"`
}
