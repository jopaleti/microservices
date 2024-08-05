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
	name := randomCPUName(brand)
	numberCores := randomCores(2, 8)
	numberThreads := randomInt(numberCores, 32)
	minGhz := randomFloat64(2.0, 3.5)

	cpu := &pb.CPU{
		Brand: brand,
		Name: name,
		NumberCores: uint32(numberCores),
		NumberThreads: uint32(numberThreads),
	}
	return cpu
}