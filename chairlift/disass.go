package chairlift

import (
    "errors"
    "fmt"
)

type Instruction interface {
    compile(*Compiler) error
}

type VReg uint8

const (
    V0 VReg = iota
    V1
    V2
    V3
    V4
    V5
    V6
    V8
    V9
    VA
    VB
    VC
    VD
    VE
    VF
)

type LdVxByte struct {
    vx VReg
    b byte
}

func newLdVxByte(opcode uint16) *LdVxByte {
    ldvxbyte := new(LdVxByte)

    ldvxbyte.vx = extract_vx(opcode)
    ldvxbyte.b = extract_kk(opcode)

    return ldvxbyte
}

func disassemble_bytes(bytes []byte) ([]Instruction, error) {
    if len(bytes) < 2 {
        return nil, errors.New("expected at least 2 bytes")
    }

    instructions := make([]Instruction, 0, len(bytes) / 2)

    for i := 0; i < len(bytes); i += 2 {
        word := uint16(bytes[i]) << 8 | uint16(bytes[i + 1])

        inst, err := disassemble(word)
        if err != nil {
            return nil, err
        }

        instructions = append(instructions, inst)
    }

    return nil, nil
}

func disassemble(opcode uint16) (Instruction, error) {
    high_4_bits := opcode >> 12

    switch high_4_bits {
    case 0x6:
        return newLdVxByte(opcode), nil
    }

    return nil, errors.New(fmt.Sprintf("Unknown opcode %X", opcode))
}

func extract_nnn(opcode uint16) uint16 {
    return opcode & 0x0FFF
}

func extract_kk(opcode uint16) byte {
    return byte(opcode)
}

func extract_vx(opcode uint16) VReg {
    return VReg(opcode >> 8 & 0x0F)
}
