package builder

import "github.com/joaomarcelofa/builder_factory/mac/hardware"

//MacBookPro13Builder is a MacBook Air implementation
type MacBookPro13Builder struct {
	Builder
	processor   hardware.Processor
	ramSize     hardware.RAMSize
	storageSize hardware.StorageSize
	screenSize  hardware.ScreenSize
}

func (bi *MacBookPro13Builder) SetRamMemory(ramSize hardware.RAMSize) Builder {
	bi.ramSize = ramSize
	return bi
}

func (bi *MacBookPro13Builder) SetStorageSize(storageSize hardware.StorageSize) Builder {
	bi.storageSize = storageSize
	return bi
}

func (bi *MacBookPro13Builder) SetProcessor(processor hardware.Processor) Builder {
	bi.processor = processor
	return bi
}

func (bi *MacBookPro13Builder) SetScreenSize() Builder {
	bi.screenSize = hardware.Screen13
	return bi
}

func (bi *MacBookPro13Builder) Build() (basicInterface, error) {
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
