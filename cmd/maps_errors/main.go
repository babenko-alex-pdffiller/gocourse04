package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	clinic := NewClinic()

	// Додавання пацієнтів
	clinic.AddPatient(Patient{ID: "p1", Name: "John Doe", Age: 30, BloodType: "A"})
	clinic.AddPatient(Patient{ID: "p2", Name: "Jane Doe", Age: 25, BloodType: "B"})

	// Серіалізація пацієнтів
	serializedData, err := clinic.SerializePatients()
	if err != nil {
		fmt.Println("Помилка серіалізації:", err)
		return
	}
	fmt.Println("Серіалізовані дані:", serializedData)

	// Десеріалізація пацієнтів
	if err := clinic.DeserializePatients(serializedData); err != nil {
		fmt.Println("Помилка десеріалізації:", err)
		return
	}

	// Перевірка десеріалізації
	if patient, exists := clinic.GetPatient("p1"); exists {
		fmt.Println("Десеріалізований пацієнт:", patient)
	}
}
