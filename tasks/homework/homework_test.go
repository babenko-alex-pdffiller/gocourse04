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

func TestFindAnimalByNameSuccessful(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}
	// Act
	_, err := FindAnimalByName(z.Areas, "Eagle")
	// Assert
	if nil != err {
		t.Errorf("expected to find Eagle, but got error %s", err)
	}
}

func TestFindAnimalByNameFailed(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}
	// Act
	animal, err := FindAnimalByName(z.Areas, "Duck")
	// Assert
	if nil == err {
		t.Errorf("expected get error, but got animal %v", animal)
	}
}

func TestFindAnimalByIDSuccessful(t *testing.T) {
	// Arrange
	z := Zoo{
		Areas: buildAreas(),
	}
	// Act
	animal, err := FindAnimalByID(z.Areas, 8)
	// Assert
	if nil != err {
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
	animal, err := FindAnimalByID(z.Areas, 18)
	// Assert
	if nil == err {
		t.Errorf("expected get error, but got animal %v", animal)
	}
}
