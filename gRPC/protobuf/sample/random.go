package sample

import (
	"math/rand"

	"github.com/jopaleti/pcbook/pb"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "MACOS")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"core i3",
			"core i4",
			"core i5",
			"core i6",
			"core i7",
			"core i9",
		)
	}

	return randomStringFromSet(
		"MAC core i5",
		"MAC core i7",
		"MAC core i9",
		"MAC M1",
		"MAC M1 Pro",
		"MAC M2",
		"MAC M2 Pro",
		"MAC M3",
		"MAC M3 Pro",
	)
}

// RandomStringFromSet
func randomStringFromSet(s ...string) string {
	n := len(s) 
	if n == 0 {
		return ""
	}
	return s[rand.Intn(n)]
}

func randomCores(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64() * max - min
}