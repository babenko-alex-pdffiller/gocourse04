package main

import (
	"fmt"
)

type Animal struct {
	ID   int
	Name string
	Type string
}

type Areas map[string]Area

type Area struct {
	Name    string
	Type    string
	Sectors map[string]Sector
}

const (
	SectorTypeTechnical = "technical"
	SectorTypeAnimals   = "animals"
	AnimalUngulatesType = "ungulates"
	AnimalFeatheredType = "feathered"
	AnimalPrimatesType  = "primates"
)

type Sector struct {
	Subtype string
	Animals []Animal
}

func (s *Sector) AddAnimal(animal Animal) {
	s.Animals = append(s.Animals, animal)
}

func (s *Sector) MoveAnimal(animalID int, distSector *Sector) {
	for i, animal := range s.Animals {
		if animal.ID == animalID {
			distSector.AddAnimal(animal)
			s.Animals = append(s.Animals[:i], s.Animals[i+1:]...)
			return
		}
	}
}

func (s *Sector) CleanUp() {
	if s.Subtype == SectorTypeTechnical {
		fmt.Println("technical sector cleaned up")
	}
}

func (s *Sector) Feed(animal *Animal) {
	if s.Subtype == SectorTypeTechnical {
		fmt.Printf("feeding the %s\n", animal.Name)
	}
}

type Zoo struct {
	Areas Areas
}

func main() {
	z := Zoo{
		Areas: buildAreas(),
	}

	fmt.Println("Try to find by name Eagle")
	eagle, err := z.FindAnimalByName("Eagle")
	if err == nil {
		fmt.Printf("%s has ID: %d\n", eagle.Name, eagle.ID)
		sector := z.Areas[eagle.Type].Sectors[SectorTypeTechnical]
		sector.Feed(eagle)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Try to find by ID 8")
	gorilla, err := z.FindAnimalByID(8)
	if err == nil {
		fmt.Printf("Animal with ID = %d has Name %s\n", gorilla.ID, gorilla.Name)
		sector := z.Areas[gorilla.Type].Sectors[SectorTypeTechnical]
		sector.Feed(gorilla)
	} else {
		fmt.Println(err)
	}

	newAnimals := Sector{
		Subtype: SectorTypeAnimals,
		Animals: []Animal{
			{ID: 10, Name: "Cow"},
		},
	}

	ungulatesAnimals := z.Areas[AnimalUngulatesType].Sectors[SectorTypeAnimals]
	ungulatesAnimals.MoveAnimal(2, &newAnimals)

	z.Areas[AnimalUngulatesType].Sectors["newAnimals"] = newAnimals

	fmt.Println("Animals from new sector")
	for _, animal := range z.Areas[AnimalUngulatesType].Sectors["newAnimals"].Animals {
		fmt.Printf("%s found, animal ID %d\n", animal.Name, animal.ID)
	}
}

func (z *Zoo) FindAnimalByName(name string) (*Animal, error) {
	for _, area := range z.Areas {
		animalsSector := area.Sectors[SectorTypeAnimals]
		for _, animal := range animalsSector.Animals {
			if name == animal.Name {
				return &animal, nil
			}
		}
	}

	return nil, fmt.Errorf("%s not found", name)
}

func (z *Zoo) FindAnimalByID(id int) (*Animal, error) {
	for _, area := range z.Areas {
		animalsSector := area.Sectors[SectorTypeAnimals]
		for _, animal := range animalsSector.Animals {
			if id == animal.ID {
				return &animal, nil
			}
		}
	}

	return nil, fmt.Errorf("Animal with ID = %d not found", id)
}

func buildAreas() Areas {
	return Areas{
		AnimalUngulatesType: {
			Name: AnimalUngulatesType,
			Type: AnimalUngulatesType,
			Sectors: map[string]Sector{
				SectorTypeAnimals: {
					Subtype: SectorTypeAnimals,
					Animals: []Animal{
						{ID: 1, Name: "Deer"},
						{ID: 2, Name: "Horse"},
						{ID: 3, Name: "Bison"},
					},
				},
				SectorTypeTechnical: {
					Subtype: SectorTypeTechnical,
				},
			},
		},
		AnimalFeatheredType: {
			Name: AnimalFeatheredType,
			Type: AnimalFeatheredType,
			Sectors: map[string]Sector{
				SectorTypeAnimals: {
					Subtype: SectorTypeAnimals,
					Animals: []Animal{
						{ID: 4, Name: "Parrot"},
						{ID: 5, Name: "Eagle"},
						{ID: 6, Name: "Penguin"},
					},
				},
				SectorTypeTechnical: {
					Subtype: SectorTypeTechnical,
				},
			},
		},
		AnimalPrimatesType: {
			Name: AnimalPrimatesType,
			Type: AnimalPrimatesType,
			Sectors: map[string]Sector{
				SectorTypeAnimals: {
					Subtype: SectorTypeAnimals,
					Animals: []Animal{
						{ID: 7, Name: "Chimpanzee"},
						{ID: 8, Name: "Gorilla"},
						{ID: 9, Name: "Orangutan"},
					},
				},
				SectorTypeTechnical: {
					Subtype: SectorTypeTechnical,
				},
			},
		},
	}
}
