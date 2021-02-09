package builder

import "github.com/joaomarcelofa/builder_factory/mac/hardware"

//MacBookPro16Builder is a MacBook Pro 16 implementation
type MacBookPro16Builder struct {
	Builder
	processor   hardware.Processor
	ramSize     hardware.RAMSize
	storageSize hardware.StorageSize
	screenSize  hardware.ScreenSize
}

func (bi *MacBookPro16Builder) SetRamMemory(ramSize hardware.RAMSize) Builder {
	bi.ramSize = ramSize
	return bi
}

func (bi *MacBookPro16Builder) SetStorageSize(storageSize hardware.StorageSize) Builder {
	bi.storageSize = storageSize
	return bi
}

func (bi *MacBookPro16Builder) SetProcessor(processor hardware.Processor) Builder {
	bi.processor = processor
	return bi
}

func (bi *MacBookPro16Builder) SetScreenSize() Builder {
	bi.screenSize = hardware.Screen16
	return bi
}

// Build a new MacBook Pro 16"
func (bi *MacBookPro16Builder) Build() (basicInterface, error) {
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
