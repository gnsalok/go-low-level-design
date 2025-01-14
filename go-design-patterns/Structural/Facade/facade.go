package main

/*
The Facade design pattern provides a simplified (and often unified) interface to a complex subsystem.
It hides the complexity of the subsystem from the client by exposing only high-level methods.
This allows clients to interact with the system in a more straightforward way, without needing to understand every detail of its internal components.

*/
import "fmt"

// ----- Subsystem: CPU -----
type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("CPU: Freezing processor")
}

func (c *CPU) Jump(position int) {
	fmt.Printf("CPU: Jumping to address %d\n", position)
}

func (c *CPU) Execute() {
	fmt.Println("CPU: Executing instructions")
}

// ----- Subsystem: Memory -----
type Memory struct{}

func (m *Memory) Load(position int, data string) {
	fmt.Printf("Memory: Loading \"%s\" into address %d\n", data, position)
}

// ----- Subsystem: HardDrive -----
type HardDrive struct{}

func (hd *HardDrive) Read(lba int, size int) string {
	fmt.Printf("HardDrive: Reading %d bytes from LBA %d\n", size, lba)
	// For simplicity, let's say it returns a string
	return "BOOT_DATA"
}

// ----- Facade: Computer -----
type Computer struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

// NewComputer creates a new instance of the Facade
func NewComputer() *Computer {
	return &Computer{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// StartComputer abstracts the complexity of starting a computer
func (c *Computer) StartComputer() {
	fmt.Println("Facade: Starting computer...")

	c.cpu.Freeze()
	bootData := c.hardDrive.Read(0, 1024)
	c.memory.Load(0, bootData)
	c.cpu.Jump(0)
	c.cpu.Execute()
}

// ShutdownComputer abstracts the complexity of shutting down a computer
func (c *Computer) ShutdownComputer() {
	fmt.Println("Facade: Shutting down computer...")
	// Potentially do a series of steps:
	// flush caches, send shutdown interrupts, stop CPU, etc.
	fmt.Println("Facade: Computer is turned off.")
}

func main() {
	// Client code only interacts with the facade (Computer),
	// without worrying about CPU, Memory, or HardDrive details.
	computer := NewComputer()
	computer.StartComputer()
	fmt.Println("---- Computer is running ----")
	computer.ShutdownComputer()
}
