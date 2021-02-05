package builder

import "github.com/joaomarcelofa/builder_factory/mac/hardware"

// basicInterface is a simple computer interface
type basicInterface interface {
	TurnOn() error
	TurnOff() error
	GetProcessor() hardware.Processor
	GetRAMSize() hardware.RAMSize
	GetStorageSize() hardware.StorageSize
	GetScreenSize() hardware.ScreenSize
}
