package main

import (
	"flag"
	"log"
	"net/http"
	"pkg/blockchain"
	"scripts"
)

func main() {
	algorithm := flag.String("algorithm", "elso", "Algorithm to benchmark")
	flag.Parse()

	// Load patient data from CSV
	patients, err := scripts.LoadPatients("data/patients.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize selected blockchain
	var bc blockchain.Blockchain
	switch *algorithm {
	case "elso":
		bc = blockchain.NewELSO()
	case "basic":
		bc = blockchain.NewBasic()
	case "offchain":
		bc = blockchain.NewOffChain()
	}

	// HTTP API
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		for _, p := range patients {
			bc.AddPatient(p) // Process all patients
		}
		w.Write([]byte("Data processed"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
