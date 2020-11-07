package chairlift

import (
    "errors"
)

func compile_instructions(instructions []Instruction, c *Compiler) error {
    for _, inst := range instructions {
        err := inst.compile(c)
        if err != nil {
            return err
        }
    }

    return nil
}

func (cls *Cls) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ret  *Ret) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (jp *JpAddr) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (call *CallAddr) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (se *SeVxByte) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (sne *SneVxByte) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (se *SeVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdVxByte) compile(c *Compiler) error {
    dst := c.VRegToLLVMValue(ld.vx)
    imm := c.ConstUint8(ld.kk)

    c.builder.CreateStore(imm, dst)

    return nil
}

func (add *AddVxByte) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (or *OrVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (and *AndVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (xor *XorVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (add *AddVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (sub *SubVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (shr *ShrVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (subn *SubnVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (shl *ShlVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (sne *SneVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdIAddr) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (jp *JpV0Addr) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (rnd *Rnd) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (drw *DrwVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (skp *SkpVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (sknp *SknpVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdVxDt) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdVxK) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdDtVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdStVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (add *AddIVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdFVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdBVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdIVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}

func (ld *LdVxI) compile(c *Compiler) error {
    return errors.New("unimplemented instruction in codegen")
}
