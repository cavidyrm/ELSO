package main

import (
	"encoding/csv"
	"os"
	"pkg/models"
)

func LoadPatients(path string) ([]models.Patient, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var patients []models.Patient
	for _, record := range records[1:] { // Skip header
		patients = append(patients, models.Patient{
			ID:            record[0],
			Name:          record[1],
			Age:           atoi(record[2]),
			MedicalRecord: record[3],
			Visits:        atoi(record[4]),
		})
	}
	return patients, nil
}
