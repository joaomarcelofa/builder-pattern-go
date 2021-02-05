package builder

import (
	"fmt"

	"github.com/joaomarcelofa/builder_factory/mac/hardware"
)

func validateMacConfig(macbook *MacBook) error {
	if !isRAMMemorySetted(macbook.ramSize) {
		return fmt.Errorf("storage size not setted")
	}

	if !isStorageSizeSetted(macbook.storageSize) {
		return fmt.Errorf("storage size not setted")
	}

	if !isProcessorSetted(macbook.processor) {
		return fmt.Errorf("processor not setted")
	}

	if !isScreenSizeSetted(macbook.screenSize) {
		return fmt.Errorf("screen size not setted")
	}

	return nil
}

func isRAMMemorySetted(macRAMSize hardware.RAMSize) bool {
	for _, ramSize := range hardware.RAMSizes {
		if macRAMSize == ramSize {
			return true
		}
	}
	return false
}

func isStorageSizeSetted(macStorageSize hardware.StorageSize) bool {
	for _, storageSize := range hardware.StorageSizes {
		if macStorageSize == storageSize {
			return true
		}
	}
	return false
}

func isProcessorSetted(macProcessor hardware.Processor) bool {
	for _, processor := range hardware.Processors {
		if macProcessor == processor {
			return true
		}
	}
	return false
}

func isScreenSizeSetted(macScreenSize hardware.ScreenSize) bool {
	for _, screenSize := range hardware.ScreenSizes {
		if macScreenSize == screenSize {
			return true
		}
	}
	return false
}
