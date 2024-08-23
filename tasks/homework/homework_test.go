package main

import (
	"testing"
)

func TestMoveAnimal(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}

	ungulatesAnimals := z.Areas["ungulates"].Sectors["animals"]
	newSector := Sector{
		Subtype: "animals",
		Animals: []Animal{},
	}

	// Act
	ungulatesAnimals.MoveAnimal(2, &newSector)

	// Assert
	if len(ungulatesAnimals.Animals) != 2 {
		t.Errorf("expected 2 animals, got %d", len(ungulatesAnimals.Animals))
	}

	if len(newSector.Animals) != 1 {
		t.Errorf("expected 1 animal in new sector, got %d", len(newSector.Animals))
	}

	if newSector.Animals[0].Name != "Horse" {
		t.Errorf("expected 'Horse' in new sector, got %s", newSector.Animals[0].Name)
	}
}

func TestFindAnimalByName(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}

	found := false

	//Ac
	for _, area := range z.Areas {
		animalsSector := area.Sectors["animals"]
		for _, animal := range animalsSector.Animals {
			if animal.Name == "Eagle" {
				found = true
				break
			}
		}
	}
	// Assert
	if !found {
		t.Error("expected to find Eagle, but it was not found")
	}
}

func TestFindAnimalByID(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}

	found := false
	// Act
	for _, area := range z.Areas {
		animalsSector := area.Sectors["animals"]
		for _, animal := range animalsSector.Animals {
			if animal.ID == 8 {
				found = true
				break
			}
		}
	}
	// Assert
	if !found {
		t.Error("expected to find animal with ID 8, but it was not found")
	}
}
