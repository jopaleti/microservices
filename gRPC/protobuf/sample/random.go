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
	return randomStringFromSet("Intel", "AMD", "MACOS")
}

func randomStringFromSet(s ...string) string {
	n := len(s) 
	if n == 0 {
		return ""
	}
	return s[rand.Intn(n)]
}