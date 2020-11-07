package chairlift

import (
    "io/ioutil"
)

type Rom struct {
    bytes []byte
    instructions []Instruction
}

func OpenRom(filename string) (*Rom, error) {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    instructions, err := disassemble_bytes(bytes)
    if err != nil {
        return nil, err
    }

    rom := new(Rom)
    rom.bytes = bytes
    rom.instructions = instructions

    return rom, nil
}
