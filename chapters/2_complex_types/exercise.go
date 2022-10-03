package exercise

type Fish struct {
	Weight int
	Size   int
	Region string
}

// Return the names of the species of fish caught today,
// grouped by the average weight of a fish of each species
func GetFishNamesByWeight(
	fishesCaught []string, // species of fish caught today
	fishypedia map[string]*Fish, // reference of all fishes attributes associated to a given species
) map[int][]string
