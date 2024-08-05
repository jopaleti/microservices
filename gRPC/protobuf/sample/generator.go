package sample

import (
	"github.com/jopaleti/pcbook/pb"
)

func NewKeyBoard() *pb.Keyboard {
	keyboard := &pb.Keyboard {
		Layout: randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	cpu := &pb.CPU{

	}
	return cpu
}