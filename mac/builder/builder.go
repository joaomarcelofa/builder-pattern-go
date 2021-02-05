package builder

import "github.com/joaomarcelofa/builder_factory/mac/hardware"

// Builder is a compurter Build Interface
type Builder interface {
	SetRamMemory(ramSize hardware.RAMSize) Builder
	SetStorageSize(storageSize hardware.StorageSize) Builder
	SetProcessor(processor hardware.Processor) Builder
	SetScreenSize() Builder
	Build() (basicInterface, error)
}

func NewMacBookAirBuilder() Builder {
	return &MacBookAirBuilder{}
}

func NewMacBookPro13Builder() Builder {
	return &MacBookPro13Builder{}
}

func NewMacBookPro16Builder() Builder {
	return &MacBookPro16Builder{}
}
