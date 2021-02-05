package hardware

type RAMSize int8
type StorageSize int
type Processor string
type ScreenSize string

const (
	RAMSize8GB  RAMSize = 8
	RAMSize16GB RAMSize = 16

	StorageSize256GB StorageSize = 256
	StorageSize512GB StorageSize = 512
	StorageSize1TB   StorageSize = 1024
	StorageSize2TB   StorageSize = 2048

	ProcessorI7 Processor = "i7"
	ProcessorI9 Processor = "i9"
	ProcessorM1 Processor = "M1"

	Screen13 ScreenSize = "13.3"
	Screen16 ScreenSize = "16"
)

var RAMSizes = []RAMSize{
	RAMSize8GB,
	RAMSize16GB,
}

var Processors = []Processor{
	ProcessorI7,
	ProcessorI9,
	ProcessorM1,
}

var StorageSizes = []StorageSize{
	StorageSize256GB,
	StorageSize512GB,
	StorageSize1TB,
	StorageSize2TB,
}

var ScreenSizes = []ScreenSize{
	Screen13,
	Screen16,
}
