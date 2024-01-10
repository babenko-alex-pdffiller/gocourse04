package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddPatient(t *testing.T) {
	clinic := NewClinic()
	patient := Patient{ID: "1", Name: "John Doe", Age: 30, BloodType: "A+"}
	clinic.AddPatient(patient)

	p, exists := clinic.GetPatient("1")
	assert.True(t, exists, "Patient should exist")
	assert.Equal(t, patient, p, "Patients should be equal")
}

func TestFindPatientsByBloodType(t *testing.T) {
	clinic := NewClinic()
	clinic.AddPatient(Patient{ID: "1", Name: "John Doe", Age: 30, BloodType: "A+"})
	clinic.AddPatient(Patient{ID: "2", Name: "Jane Doe", Age: 28, BloodType: "O-"})

	found := clinic.FindPatientsByBloodType("O-")
	assert.Len(t, found, 1, "Should find one patient")
	assert.Equal(t, "2", found[0].ID, "Patient ID should be 2")
}

func TestSerialization(t *testing.T) {
	clinic := NewClinic()
	clinic.AddPatient(Patient{ID: "1", Name: "John Doe", Age: 30, BloodType: "A+"})

	serialized, err := clinic.SerializePatients()
	assert.NoError(t, err, "SerializePatients should not error")

	newClinic := NewClinic()
	assert.NoError(t, newClinic.DeserializePatients(serialized), "DeserializePatients should not error")

	_, exists := newClinic.GetPatient("1")
	assert.True(t, exists, "Patient ID 1 should exist after deserialization")
}

func BenchmarkFindInArray(b *testing.B) {
	patients := generatePatients()
	var array [1000]Patient
	copy(array[:], patients[:])

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findInArray(array, "A+")
	}
}

func BenchmarkFindInSlice(b *testing.B) {
	patients := generatePatients()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findInSlice(patients, "A+")
	}
}

func BenchmarkFindInMap(b *testing.B) {
	clinic := NewClinic()
	patients := generatePatients()
	for _, p := range patients {
		clinic.AddPatient(p)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clinic.FindPatientsByBloodType("A+")
	}
}
