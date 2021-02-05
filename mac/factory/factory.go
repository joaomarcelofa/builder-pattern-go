package factory

import (
	"fmt"

	"github.com/joaomarcelofa/builder_factory/mac/builder"
	"github.com/joaomarcelofa/builder_factory/mac/hardware"
)

func CreateNewMac(macType MacType) (*builder.MacBook, error) {
	switch macType {
	case MacBookAir:
		return buildMacBookAir()
	case MacBookPro13:
		return buildMacBookPro13()
	case MacBookPro16:
		return buildMacBookPro16()
	}
	return nil, fmt.Errorf("mac type not supported")
}

func buildMacBookAir() (*builder.MacBook, error) {
	airBuilder := builder.NewMacBookAirBuilder()
	airBuilder.SetProcessor(hardware.ProcessorM1).
		SetRamMemory(hardware.RAMSize8GB).
		SetStorageSize(hardware.StorageSize256GB).
		SetScreenSize()

	macBookAirInterface, err := airBuilder.Build()
	if err != nil {
		return nil, err
	}

	return builder.TransformMacInterface(macBookAirInterface)
}

func buildMacBookPro13() (*builder.MacBook, error) {
	pro13Builder := builder.NewMacBookPro13Builder()
	pro13Builder.SetProcessor(hardware.ProcessorI7).
		SetRamMemory(hardware.RAMSize8GB).
		SetStorageSize(hardware.StorageSize256GB).
		SetScreenSize()

	macBookpro13Interface, err := pro13Builder.Build()
	if err != nil {
		return nil, err
	}

	return builder.TransformMacInterface(macBookpro13Interface)
}

func buildMacBookPro16() (*builder.MacBook, error) {
	pro16Builder := builder.NewMacBookPro16Builder()
	pro16Builder.SetProcessor(hardware.ProcessorM1).
		SetRamMemory(hardware.RAMSize8GB).
		SetStorageSize(hardware.StorageSize256GB).
		SetScreenSize()

	macBookpro16Interface, err := pro16Builder.Build()
	if err != nil {
		return nil, err
	}

	return builder.TransformMacInterface(macBookpro16Interface)
}
