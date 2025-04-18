#!/bin/bash

# Test ELSO
echo "Testing ELSO..."
ab -n 10000 -c 100 -p data/patients.csv http://localhost:8080/add > results/elso.txt

# Test Basic
echo "Testing Basic Blockchain..."
ab -n 10000 -c 100 -p data/patients.csv http://localhost:8081/add > results/basic.txt

# Test Off-Chain
echo "Testing Off-Chain..."
ab -n 10000 -c 100 -p data/patients.csv http://localhost:8082/add > results/offchain.txt

# Parse results
grep "Requests per second" results/*.txt