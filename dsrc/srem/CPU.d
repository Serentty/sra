module CPU;

alias int Register; // For Register IDs

import Registers;

class CPU
{
	private ulong registers[32];
	public RAM memory;

	ulong getRegister(Register r)
	{
		return registers[r];
	}

	void setRegister(Register r, ulong value)
	{
		registers[r] = value;
	}

	void add(Register a, Register b, Register destination)
	{
		registers[destination] = registers[a] + registers[b];		
	}

	void subtract(Register a, Register b, Register destination)
	{
		registers[destination] = registers[a] - registers[b];		
	}

	void multiply(Register a, Register b, Register destination)
	{
		registers[destination] = registers[a] * registers[b];		
	}

	void divide(Register a, Register b, Register destination)
	{
		registers[destination] = registers[a] / registers[b];		
	}

	void modulo(Register a, Register b, Register destination)
	{
		registers[destination] = registers[a] % registers[b];		
	}

	void halt()
	{

	}

	this()
	{
		memory = new RAM(1024);
	}	

}

class RAM // Note: Due to this class's design, it can only hold 4 GiB of data - Fix later
{
	private ulong size;
	private ubyte data[];

	this(ulong ramSize)
	{
		size = ramSize;
		data = new ubyte[cast(uint) ramSize];
	}

	public ulong getSize() { return size; }
	public ubyte getByte(ulong address) { return data[cast(uint)address % size]; }
	public void setByte(ulong address, ubyte value) { data[cast(uint)address % size] = value; }
}