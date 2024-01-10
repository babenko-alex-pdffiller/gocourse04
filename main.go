package main

import (
	"fmt"
)

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
