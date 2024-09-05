package main

import (
	"testing"
)

func TestMoveAnimal(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}

	ungulatesAnimals := z.Areas[AnimalTypeUngulates].Sectors[SectorTypeAnimals]
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

	if len(newSector.Animals) > 0 && newSector.Animals[0].Name != "Horse" {
		t.Errorf("expected 'Horse' in new sector, got %s", newSector.Animals[0].Name)
	}
}

func TestFindAnimalByNameSuccessful(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}
	// Act
	animal, err := z.FindAnimalByName("Eagle")
	// Assert
	if err != nil {
		t.Errorf("expected to find Eagle, but got error %s", err)
	}

	if animal.Name != "Eagle" {
		t.Errorf("expected animal name 'Eagle', but got %s", animal.Name)
	}
}

func TestFindAnimalByNameFailed(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}
	// Act
	_, err := z.FindAnimalByName("Duck")
	// Assert
	if err == nil {
		t.Error("expected get error, but got <nil>")
	}
}

func TestFindAnimalByIDSuccessful(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}
	// Act
	animal, err := z.FindAnimalByID(8)
	// Assert
	if err != nil {
		t.Errorf("expected to find Gorilla, but got error %s", err)
	}
	if animal.ID != 8 {
		t.Errorf("expected animal ID = 8, but got %d", animal.ID)
	}
}

func TestFindAnimalByIDFailed(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}
	// Act
	_, err := z.FindAnimalByID(18)
	// Assert
	if err == nil {
		t.Error("expected get error, but got <nil>")
	}
}
