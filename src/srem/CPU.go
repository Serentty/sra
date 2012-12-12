package main


type CPU struct {
	// Registers 
	registerStack [16]uint64 // Registers
	Register_RSP int // Top of the register stack
	Register_PC uint64
	// Simple general purpose registers
	Registers [16]uint64
	// RAM
	RAM [1024]byte
}

func (c *CPU) add(firstRegister int8, secondRegister int8, destination int8) {
	(*c).Registers[destination] = (*c).Registers[firstRegister] + (*c).Registers[secondRegister]
}

func (c *CPU) subtract(firstRegister int8, secondRegister int8, destination int8) {
	(*c).Registers[destination] = (*c).Registers[firstRegister] - (*c).Registers[secondRegister]
}

func (c *CPU) multiply(firstRegister int8, secondRegister int8, destination int8) {
	(*c).Registers[destination] = (*c).Registers[firstRegister] * (*c).Registers[secondRegister]
}

func (c *CPU) divide(firstRegister int8, secondRegister int8, destination int8) {
	(*c).Registers[destination] = (*c).Registers[firstRegister] / (*c).Registers[secondRegister]
}

func (c *CPU) modulo(firstRegister int8, secondRegister int8, destination int8) {
	(*c).Registers[destination] = (*c).Registers[firstRegister] % (*c).Registers[secondRegister]
}

func (c *CPU) registerStackPush(value uint64) {
	(*c).Register_RSP = ((*c).Register_RSP + 1) % 16 // Move the top up
	(*c).registerStack[(*c).Register_RSP] = value
}

func (c *CPU) registerStackPop() uint64 {
	value := (*c).registerStack[(*c).Register_RSP]
	(*c).Register_RSP -= 1
	if((*c).Register_RSP > 15) {
		(*c).Register_RSP = 15 
	}
	return value
}

func (c *CPU) halt() {
}

func (c *CPU) storeByte(address uint64, value byte) {
	(*c).RAM[address] = value
}

func (c *CPU) loadByte(address uint64) byte {
	return (*c).RAM[address]
}

func (c *CPU) store(address uint64, value uint64) {
	byteArray := int64ToBytes(value)
	(*c).RAM[address] = byteArray[0]
	(*c).RAM[address + 1] = byteArray[1]
	(*c).RAM[address + 2] = byteArray[2]
	(*c).RAM[address + 3] = byteArray[3]
	(*c).RAM[address + 4] = byteArray[4]
	(*c).RAM[address + 5] = byteArray[5]
	(*c).RAM[address + 6] = byteArray[6]
	(*c).RAM[address + 7] = byteArray[7]
}

func (c *CPU) load(address uint64) uint64 {
	byteArray := []byte{(*c).RAM[address], (*c).RAM[address + 1], (*c).RAM[address + 2], (*c).RAM[address + 3], (*c).RAM[address + 4], (*c).RAM[address + 5], (*c).RAM[address + 6], (*c).RAM[address + 7]}
	return bytesToInt64(byteArray)
}

func (c *CPU) run() {
	var instructionLength uint64

	while(true) {
		instructionLength = getInstructionLength(RAM[Register_PC])
		if(instructionLength == 10) {
			instruction := []byte{RAM[Register_PC], RAM[Register_PC + 1], RAM[Register_PC + 2], RAM[Register_PC = 3], RAM[Register_PC + 4], RAM[Register_PC + 5], RAM[Register_PC + 6], RAM[Register_PC + 7], RAM[Register_PC + 8], RAM[Register_PC + 9]}
			executeInstruction(instruction)
			Register_PC += 10
		}

		else { // If length is 4
			instruction := []byte{RAM[Register_PC], RAM[Register_PC + 1], RAM[Register_PC + 2], RAM[Register_PC + 3]}
			executeInstruction(instruction)
			Register_PC += 4
		}
	}
}

func (c *CPU) executeInstruction(instruction []byte) {
	switch instruction[0] {
	case 0x00:
		(*c).halt()
	case 0x01:
		(*c).Registers[instruction[1]] = bytesToInt64(getDataFromInstruction(instruction)) // Set Register to Value
	case 0x02:
		(*c).Registers[instruction[1]] = (*c).Registers[instruction[2]]	// Set Register to Register
	case 0x20:
		(*c).add(instruction[1], instruction[2], instruction[3])
	case 0x21:
		(*c).subtract(instruction[1], instruction[2], instruction[3])
	case 0x22:
		(*c).multiply(instruction[1], instruction[2], instruction[3])
	case 0x23:
		(*c).divide(instruction[1], instruction[2], instruction[3])
	case 0x24:
		(*c).modulo(instruction[1], instruction[2], instruction[3])
	}
}

func getInstructionLength(opcode byte) uint64 {
	switch opcode {
	case 0x01:
		return 10
	default:
		return 4
	}
}