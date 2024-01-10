package main

import (
	"fmt"
	"time"
)

type Medicine struct {
	Name         string
	Manufactured time.Time
}

type Box struct {
	ToDiscard map[string]Medicine
	ForSale   map[string]Medicine
	ToUse     map[string]Medicine
}

func NewBox() *Box {
	return &Box{
		ToDiscard: make(map[string]Medicine),
		ForSale:   make(map[string]Medicine),
		ToUse:     make(map[string]Medicine),
	}
}

func (b *Box) SortMedicines(medicines []Medicine) {
	now := time.Now()
	halfAYearAgo := now.AddDate(0, -6, 0)

	for _, medicine := range medicines {
		switch {
		case medicine.Manufactured.Before(halfAYearAgo):
			b.ToDiscard[medicine.Name] = medicine
		case medicine.Manufactured.After(halfAYearAgo) && medicine.Manufactured.Before(now):
			b.ToUse[medicine.Name] = medicine
		case medicine.Manufactured.After(now):
			b.ForSale[medicine.Name] = medicine
		}
	}
}

// FindMedicine шукає ліки за назвою в усіх категоріях
func (b *Box) FindMedicine(name string) (*Medicine, string) {
	if med, exists := b.ToDiscard[name]; exists {
		return &med, "для викидання"
	}
	if med, exists := b.ForSale[name]; exists {
		return &med, "на продаж"
	}
	if med, exists := b.ToUse[name]; exists {
		return &med, "для використання"
	}
	return nil, ""
}

func main() {
	medicines := []Medicine{
		{"MedicineA", time.Date(2022, 1, 10, 0, 0, 0, 0, time.UTC)},
		{"MedicineB", time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC)},
		{"MedicineC", time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)},
	}

	box := NewBox()
	box.SortMedicines(medicines)

	// Пошук ліків
	if med, category := box.FindMedicine("MedicineA"); med != nil {
		fmt.Printf("Ліки '%s' знайдені у категорії '%s'\n", med.Name, category)
	} else {
		fmt.Println("Ліки 'MedicineA' не знайдені.")
	}
}
