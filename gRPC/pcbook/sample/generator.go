package sample

import (
	"github.com/golang/protobuf/ptypes"
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
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand: brand,
		Name: name,
		NumberCores: uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz: minGhz,
		MaxGhz: maxGhz,
	}
	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit: pb.Memory_GIGABYTE,
	}
	
	gpu := &pb.GPU{
		Brand: brand,
		Name: name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

func NewRAM() *pb.Memory {
	ram := &pb.Memory {
		Value: uint64(randomInt(4, 64)),
		Unit: pb.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *pb.Storage {
	ssd := &pb.Storage {
		Driver: pb.Storage_SDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit: pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

func NewHDD() *pb.Storage {
	hdd := &pb.Storage {
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit: pb.Memory_TERABYTE,
		},
	}
	return hdd
}

// New screen return a new map
func NewScreen() *pb.Screen {
	screen := &pb.Screen {
		SizeInch: randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel: randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// Generate a new simple laptop
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id: randomID(),
		Brand: brand,
		Name: name,
		Cpu: NewCPU(),
		Ram: NewRAM(),
		Gpus: []*pb.CPU{NewCPU()},
		Storage: []*pb.Storage{NewSSD(), NewHDD()},
		Screen: NewScreen(),
		Keyboard: NewKeyBoard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd: randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2018, 2024)),
		UpdatedAt: ptypes.TimestampNow(),
	}

	return laptop
}

// RandomLaptopScore returns a random laptop score
func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}