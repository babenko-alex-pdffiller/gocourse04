package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// Patient - структура, що представляє пацієнта
type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BloodType string `json:"blood_type"`
}

// Clinic - структура, що містить мапу пацієнтів
type Clinic struct {
	patients map[string]Patient
}

// NewClinic створює нову клініку з порожньою мапою пацієнтів
func NewClinic() *Clinic {
	return &Clinic{
		patients: make(map[string]Patient),
	}
}

// AddPatient додає нового пацієнта у мапу
func (c *Clinic) AddPatient(p Patient) {
	c.patients[p.ID] = p
}

// GetPatient повертає пацієнта за ID
func (c *Clinic) GetPatient(id string) (Patient, bool) {
	p, exists := c.patients[id]
	return p, exists
}

// UpdatePatient оновлює дані пацієнта
func (c *Clinic) UpdatePatient(id string, newPatient Patient) {
	if _, exists := c.patients[id]; exists {
		c.patients[id] = newPatient
	}
}

// DeletePatient видаляє пацієнта за ID
func (c *Clinic) DeletePatient(id string) {
	delete(c.patients, id)
}

// FindPatientsByBloodType знаходить пацієнтів за групою крові
func (c *Clinic) FindPatientsByBloodType(bloodType string) []Patient {
	var found []Patient
	for _, p := range c.patients {
		if p.BloodType == bloodType {
			found = append(found, p)
		}
	}
	return found
}

// SerializePatients серіалізує мапу пацієнтів в JSON
func (c *Clinic) SerializePatients() (string, error) {
	data, err := json.Marshal(c.patients)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// DeserializePatients десеріалізує JSON в мапу пацієнтів
func (c *Clinic) DeserializePatients(data string) error {
	var patients map[string]Patient
	if err := json.Unmarshal([]byte(data), &patients); err != nil {
		return err
	}
	c.patients = patients
	return nil
}

func generatePatients() []Patient {
	var patients []Patient
	for i := 0; i < 1000; i++ {
		patients = append(patients, Patient{
			ID:        fmt.Sprintf("%d", i),
			Name:      fmt.Sprintf("Patient %d", i),
			Age:       rand.Intn(100),
			BloodType: []string{"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"}[rand.Intn(8)],
		})
	}
	return patients
}

func findInArray(patients [1000]Patient, bloodType string) []Patient {
	var found []Patient
	for _, p := range patients {
		if p.BloodType == bloodType {
			found = append(found, p)
		}
	}
	return found
}

func findInSlice(patients []Patient, bloodType string) []Patient {
	var found []Patient
	for _, p := range patients {
		if p.BloodType == bloodType {
			found = append(found, p)
		}
	}
	return found
}
