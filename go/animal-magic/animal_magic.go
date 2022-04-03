package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	dieNumber := rand.Intn(19)
	return dieNumber + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	wandEnergy := rand.Float64() * float64(rand.Intn(12))
	return wandEnergy
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	animalRoster := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animalRoster), func(i, j int) {
		animalRoster[i], animalRoster[j] = animalRoster[j], animalRoster[i]
	})
	return animalRoster
}
