package sample

import (
	"math/rand"

	"github.com/google/uuid"
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

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2023",
			"RTX 2001",
			"GTX 0021",
			"GTX 7712",
		)
	}
	return randomStringFromSet(
		"RX 500",
		"RX 600",
		"RX 430",
	)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4300)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Width: uint32(width),
		Height: uint32(height),
	}
	return resolution
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("MacBook Air", "MacBook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Voster", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad K1", "Thinkpad P1", "Thinkpad P53")
	}
}