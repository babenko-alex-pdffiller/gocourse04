package main

import (
	"reflect"
	"testing"
)

func TestAddPatient(t *testing.T) {
	clinic := NewClinic()
	patient := Patient{ID: "1", Name: "John Doe", Age: 30, BloodType: "A+"}
	clinic.AddPatient(patient)

	if p, exists := clinic.GetPatient("1"); !exists || !reflect.DeepEqual(p, patient) {
		t.Errorf("AddPatient failed: expected %v, got %v", patient, p)
	}
}

func TestFindPatientsByBloodType(t *testing.T) {
	clinic := NewClinic()
	clinic.AddPatient(Patient{ID: "1", Name: "John Doe", Age: 30, BloodType: "A+"})
	clinic.AddPatient(Patient{ID: "2", Name: "Jane Doe", Age: 28, BloodType: "O-"})

	found := clinic.FindPatientsByBloodType("O-")
	if len(found) != 1 || found[0].ID != "2" {
		t.Errorf("FindPatientsByBloodType failed: expected 1 patient with ID 2, got %v", found)
	}
}

func TestSerialization(t *testing.T) {
	clinic := NewClinic()
	clinic.AddPatient(Patient{ID: "1", Name: "John Doe", Age: 30, BloodType: "A+"})

	serialized, err := clinic.SerializePatients()
	if err != nil {
		t.Errorf("SerializePatients failed: %v", err)
	}

	newClinic := NewClinic()
	if err := newClinic.DeserializePatients(serialized); err != nil {
		t.Errorf("DeserializePatients failed: %v", err)
	}

	if _, exists := newClinic.GetPatient("1"); !exists {
		t.Errorf("DeserializePatients failed: patient ID 1 not found")
	}
}
