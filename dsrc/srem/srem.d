module main;

import std.stdio;

import CPU;
import Registers;

void main()
{
	CPU cpu = new CPU();
	cpu.setRegister(REGISTER_A0, 142);
	writeln(cpu.getRegister(REGISTER_A0));
}