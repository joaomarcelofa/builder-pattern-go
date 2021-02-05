package builder

import (
	"fmt"
	"time"

	"github.com/joaomarcelofa/builder_factory/mac/hardware"
)

type MacBook struct {
	processor   hardware.Processor
	ramSize     hardware.RAMSize
	storageSize hardware.StorageSize
	screenSize  hardware.ScreenSize
}

func (m *MacBook) TurnOn() error {
	fmt.Println("Turning On!")
	time.Sleep(time.Second * 3)
	fmt.Println("Turned On!")
	return nil
}

func (m *MacBook) TurnOff() error {
	fmt.Println("Turning Off!")
	time.Sleep(time.Second * 3)
	fmt.Println("Turned Off!")
	return nil
}

func (m *MacBook) GetProcessor() hardware.Processor {
	return m.processor
}

func (m *MacBook) GetRAMSize() hardware.RAMSize {
	return m.ramSize
}

func (m *MacBook) GetStorageSize() hardware.StorageSize {
	return m.storageSize
}

func (m *MacBook) GetScreenSize() hardware.ScreenSize {
	return m.screenSize
}
