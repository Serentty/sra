typedef int Register // For Register IDs

import Registers

class CPU
{
	uint64 registers[32];
	byte RAM[1024];

	void add(Register a, Register b, Register destination)
	{
		Registers[destination] = Registers[a] + Registers[b]		
	}

	void subtract(Register a, Register b, Register destination)
	{
		Registers[destination] = Registers[a] - Registers[b]		
	}

	void multiply(Register a, Register b, Register destination)
	{
		Registers[destination] = Registers[a] * Registers[b]		
	}

	void divide(Register a, Register b, Register destination)
	{
		Registers[destination] = Registers[a] / Registers[b]		
	}

	void modulo(Register a, Register b, Register destination)
	{
		Registers[destination] = Registers[a] % Registers[b]		
	}

	void halt()
	{

	}

}

uint