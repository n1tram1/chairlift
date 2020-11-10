package chairlift

import (
    "errors"
    "fmt"
)

const INSTRUCTION_SIZE = 2

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
    V7
    V8
    V9
    VA
    VB
    VC
    VD
    VE
    VF
)

type Cls struct {}

type Ret struct {}

type Sys struct {
    addr uint16
}

type JpAddr struct {
    addr uint16
}

type CallAddr struct {
    addr uint16
}

type SeVxByte struct {
    vx VReg
    kk byte
}

type SneVxByte struct {
    vx VReg
    kk byte
}

type SeVxVy struct {
    vx VReg
    vy VReg
}

type LdVxByte struct {
    vx VReg
    kk byte
}

type AddVxByte struct {
    vx VReg
    kk byte
}

type LdVxVy struct {
    vx VReg
    vy VReg
}

type OrVxVy struct {
    vx VReg
    vy VReg
}

type AndVxVy struct {
    vx VReg
    vy VReg
}

type XorVxVy struct {
    vx VReg
    vy VReg
}

type AddVxVy struct {
    vx VReg
    vy VReg
}

type SubVxVy struct {
    vx VReg
    vy VReg
}

type ShrVxVy struct {
    vx VReg
    vy VReg
}

type SubnVxVy struct {
    vx VReg
    vy VReg
}

type ShlVxVy struct {
    vx VReg
    vy VReg
}

type SneVxVy struct {
    vx VReg
    vy VReg
}

type LdIAddr struct {
    addr uint16
}

type JpV0Addr struct {
    addr uint16
}

type Rnd struct {
    vx VReg
    kk byte
}

type DrwVxVy struct {
    vx VReg
    vy VReg
    n uint8
}

type SkpVx struct {
    vx VReg
}

type SknpVx struct {
    vx VReg
}

type LdVxDt struct {
    vx VReg
}

type LdVxK struct {
    vx VReg
}

type LdDtVx struct {
    vx VReg
}

type LdStVx struct {
    vx VReg
}

type AddIVx struct {
    vx VReg
}

type LdFVx struct {
    vx VReg
}

type LdBVx struct {
    vx VReg
}

type LdIVx struct {
    vx VReg
}

type LdVxI struct {
    vx VReg
}

func newCls() *Cls {
    return &Cls{}
}

func newRet() *Ret {
    return &Ret{}
}

func newSys(opcode uint16) *Sys {
    return &Sys{extract_nnn(opcode)}
}

func newJpAddr(opcode uint16) *JpAddr {
    return &JpAddr{extract_nnn(opcode)}
}

func newCallAddr(opcode uint16) *CallAddr {
    return &CallAddr{extract_nnn(opcode)}
}

func newSeVxByte(opcode uint16) *SeVxByte {
    sevxbyte := new(SeVxByte)

    sevxbyte.vx = extract_vx(opcode)
    sevxbyte.kk = extract_kk(opcode)

    return sevxbyte
}

func newSneVxByte(opcode uint16) *SneVxByte {
    snevxbyte := new(SneVxByte)

    snevxbyte.vx = extract_vx(opcode)
    snevxbyte.kk = extract_kk(opcode)

    return snevxbyte
}

func newSeVxVy(opcode uint16) *SeVxVy {
    return &SeVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newLdVxByte(opcode uint16) *LdVxByte {
    ldvxbyte := new(LdVxByte)

    ldvxbyte.vx = extract_vx(opcode)
    ldvxbyte.kk = extract_kk(opcode)

    return ldvxbyte
}

func newAddVxByte(opcode uint16) *AddVxByte {
    addvxbyte := new(AddVxByte)

    addvxbyte.vx = extract_vx(opcode)
    addvxbyte.kk = extract_kk(opcode)

    return addvxbyte
}

func newLdVxVy(opcode uint16) *LdVxVy {
    return &LdVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newOrVxVy(opcode uint16) *OrVxVy {
    return &OrVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newAndVxVy(opcode uint16) *AndVxVy {
    return &AndVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newXorVxVy(opcode uint16) *XorVxVy {
    return &XorVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newAddVxVy(opcode uint16) *AddVxVy {
    return &AddVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newSubVxVy(opcode uint16) *SubVxVy {
    return &SubVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newShrVxVy(opcode uint16) *ShrVxVy {
    return &ShrVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newSubnVxVy(opcode uint16) *SubnVxVy {
    return &SubnVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newShlVxVy(opcode uint16) *ShlVxVy {
    return &ShlVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newSneVxVy(opcode uint16) *SneVxVy {
    return &SneVxVy{extract_vx(opcode), extract_vy(opcode)}
}

func newLdIAddr(opcode uint16) *LdIAddr {
    return &LdIAddr{extract_nnn(opcode)}
}

func newJpV0Addr(opcode uint16) *JpV0Addr {
    return &JpV0Addr{extract_nnn(opcode)}
}

func newRnd(opcode uint16) *Rnd {
    return &Rnd{vx: extract_vx(opcode), kk: extract_kk(opcode)}
}

func newDrwVxVy(opcode uint16) *DrwVxVy {
    return &DrwVxVy{extract_vx(opcode), extract_vy(opcode), extract_n(opcode)}
}

func newSkpVx(opcode uint16) *SkpVx {
    return &SkpVx{extract_vx(opcode)}
}

func newSknpVx(opcode uint16) *SknpVx {
    return &SknpVx{extract_vx(opcode)}
}

func newLdVxDt(opcode uint16) *LdVxDt {
    return &LdVxDt{extract_vx(opcode)}
}

func newLdVxK(opcode uint16) *LdVxK {
    return &LdVxK{extract_vx(opcode)}
}

func newLdDtVx(opcode uint16) *LdDtVx {
    return &LdDtVx{extract_vx(opcode)}
}

func newLdStVx(opcode uint16) *LdStVx {
    return &LdStVx{extract_vx(opcode)}
}

func newAddIVx(opcode uint16) *AddIVx {
    return &AddIVx{extract_vx(opcode)}
}

func newLdFVx(opcode uint16) *LdFVx {
    return &LdFVx{extract_vx(opcode)}
}

func newLdBVx(opcode uint16) *LdBVx {
    return &LdBVx{extract_vx(opcode)}
}

func newLdIVx(opcode uint16) *LdIVx {
    return &LdIVx{extract_vx(opcode)}
}

func newLdVxI(opcode uint16) *LdVxI {
    return &LdVxI{extract_vx(opcode)}
}

func disassemble_bytes(bytes []byte) ([]Instruction, error) {
    if len(bytes) < 2 {
        return nil, errors.New("expected at least 2 bytes")
    }

    instructions := make([]Instruction, 0, len(bytes) / INSTRUCTION_SIZE)

    for i := 0; i < len(bytes); i += INSTRUCTION_SIZE {
        word := uint16(bytes[i]) << 8 | uint16(bytes[i + 1])

        inst, err := disassemble(word)
        if err != nil {
            return nil, err
        }

        instructions = append(instructions, inst)
    }

    return instructions, nil
}

func disassemble(opcode uint16) (Instruction, error) {
    high_4_bits := opcode >> 12

    switch high_4_bits {
    case 0x0:
        return disass_0_prefix(opcode), nil
    case 0x1:
        return newJpAddr(opcode), nil
    case 0x2:
        return newCallAddr(opcode), nil
    case 0x3:
        return newSeVxByte(opcode), nil
    case 0x4:
        return newSneVxByte(opcode), nil
    case 0x5:
        if opcode & 0xF == 0 {
            return newSeVxVy(opcode), nil
        }
    case 0x6:
        return newLdVxByte(opcode), nil
    case 0x7:
        return newAddVxByte(opcode), nil
    case 0x8:
        inst, err := disass_8_prefix(opcode)
        if err != nil {
            return nil, err
        }
        return inst, nil
    case 0x9:
        if opcode & 0xF == 0 {
            return newSneVxVy(opcode), nil
        }
    case 0xA:
        return newLdIAddr(opcode), nil
    case 0xB:
        return newJpV0Addr(opcode), nil
    case 0xC:
        return newRnd(opcode), nil
    case 0xD:
        return newDrwVxVy(opcode), nil
    case 0xE:
        inst, err := disass_E_prefix(opcode)
        if err != nil {
            return nil, err
        }
        return inst, nil
    case 0xF:
        inst, err := disass_F_prefix(opcode)
        if err != nil {
            return nil, err
        }
        return inst, nil
    }

    return nil, errors.New(fmt.Sprintf("Unknown opcode %X", opcode))
}

func disass_0_prefix(opcode uint16) Instruction {
    switch opcode {
    case 0x00E0:
        return newCls()
    case 0x0EE:
        return newRet()
    default:
        return newSys(opcode)
    }
}

func disass_8_prefix(opcode uint16) (Instruction, error) {
    low_byte := opcode & 0xF
    switch low_byte {
    case 0x0:
        return newLdVxVy(opcode), nil
    case 0x1:
        return newOrVxVy(opcode), nil
    case 0x2:
        return newAndVxVy(opcode), nil
    case 0x3:
        return newXorVxVy(opcode), nil
    case 0x4:
        return newAddVxVy(opcode), nil
    case 0x5:
        return newSubVxVy(opcode), nil
    case 0x6:
        return newShrVxVy(opcode), nil
    case 0x7:
        return newSubnVxVy(opcode), nil
    case 0xE:
        return newShlVxVy(opcode), nil
    default:
        return nil, errors.New(fmt.Sprintf("invalid 8 prefix instruction %X", opcode))
    }
}

func disass_E_prefix(opcode uint16) (Instruction, error) {
    low_byte := opcode & 0xFF

    switch low_byte {
    case 0x9E:
        return newSkpVx(opcode), nil
    case 0xA1:
        return newSknpVx(opcode), nil
    default:
        return nil, errors.New(fmt.Sprintf("invalid E prefix instruciton %X", opcode))
    }
}

func disass_F_prefix(opcode uint16) (Instruction, error) {
    low_byte := opcode & 0xFF

    switch low_byte {
    case 0x07:
        return newLdVxDt(opcode), nil
    case 0x0A:
        return newLdVxK(opcode), nil
    case 0x15:
        return newLdDtVx(opcode), nil
    case 0x18:
        return newLdStVx(opcode), nil
    case 0x1E:
        return newAddIVx(opcode), nil
    case 0x29:
        return newLdFVx(opcode), nil
    case 0x33:
        return newLdBVx(opcode), nil
    case 0x55:
        return newLdIVx(opcode), nil
    case 0x65:
        return newLdVxI(opcode), nil
    default:
        return nil, errors.New(fmt.Sprintf("invalid F prefix instruction %X", opcode))
    }
}

func extract_nnn(opcode uint16) uint16 {
    return opcode & 0x0FFF
}

func extract_kk(opcode uint16) byte {
    return byte(opcode)
}

func extract_vx(opcode uint16) VReg {
    return VReg(opcode >> 8 & 0xF)
}

func extract_vy(opcode uint16) VReg {
    return VReg(opcode >> 4 & 0xF)
}

func extract_n(opcode uint16) uint8 {
    return uint8(opcode) & 0xF
}
