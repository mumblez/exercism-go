package allergies

const (
	EGGS = 1 << iota
	PEANUTS
	SHELLFISH
	STRAWBERRIES
	TOMATOES
	CHOCOLATE
	POLLEN
	CATS
	testVersion = 1
)

// AllergiesTo does a bitwise calculation to see what allergies apply
func Allergies(score uint) []string {
	var allergies []string
	if score&EGGS > 0 {
		allergies = append(allergies, "eggs")
	}
	if score&PEANUTS > 0 {
		allergies = append(allergies, "peanuts")
	}
	if score&SHELLFISH > 0 {
		allergies = append(allergies, "shellfish")
	}
	if score&STRAWBERRIES > 0 {
		allergies = append(allergies, "strawberries")
	}
	if score&TOMATOES > 0 {
		allergies = append(allergies, "tomatoes")
	}
	if score&CHOCOLATE > 0 {
		allergies = append(allergies, "chocolate")
	}
	if score&POLLEN > 0 {
		allergies = append(allergies, "pollen")
	}
	if score&CATS > 0 {
		allergies = append(allergies, "cats")
	}
	return allergies
}

// AllergicTo checks if given number and allergy match
func AllergicTo(score uint, allergy string) bool {
	results := Allergies(score)
	if len(results) == 0 {
		return false
	}
	for _, a := range results {
		if a == allergy {
			return true
		}
	}
	return false
}
