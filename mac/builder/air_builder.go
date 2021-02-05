package builder

import (
	"github.com/joaomarcelofa/builder_factory/mac/hardware"
)

//MacBookAirBuilder is a MacBook Air implementation
type MacBookAirBuilder struct {
	processor   hardware.Processor
	ramSize     hardware.RAMSize
	storageSize hardware.StorageSize
	screenSize  hardware.ScreenSize
}

func (bi *MacBookAirBuilder) SetRamMemory(ramSize hardware.RAMSize) Builder {
	bi.ramSize = ramSize
	return bi
}

func (bi *MacBookAirBuilder) SetStorageSize(storageSize hardware.StorageSize) Builder {
	bi.storageSize = storageSize
	return bi
}

func (bi *MacBookAirBuilder) SetProcessor(processor hardware.Processor) Builder {
	bi.processor = processor
	return bi
}

func (bi *MacBookAirBuilder) SetScreenSize() Builder {
	bi.screenSize = hardware.Screen13
	return bi
}

func (bi *MacBookAirBuilder) Build() (basicInterface, error) {
	macbook := &MacBook{
		processor:   bi.processor,
		ramSize:     bi.ramSize,
		storageSize: bi.storageSize,
		screenSize:  bi.screenSize,
	}

	err := validateMacConfig(macbook)
	if err != nil {
		return nil, err
	}

	return macbook, nil
}
