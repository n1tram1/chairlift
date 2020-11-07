package chairlift

import (
    llvm "github.com/tinygo-org/go-llvm"
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
    return errors.New("unimplemented instruction Cls in codegen")
}

func (ret  *Ret) compile(c *Compiler) error {
    return errors.New("unimplemented instruction Ret in codegen")
}

func (jp *JpAddr) compile(c *Compiler) error {
    return errors.New("unimplemented instruction JpAddr in codegen")
}

func (call *CallAddr) compile(c *Compiler) error {
    return errors.New("unimplemented instruction CallAddr in codegen")
}

func (se *SeVxByte) compile(c *Compiler) error {
    // reg := c.VRegToLLVMValue(se.vx)
    // val := c.ConstUint8(se.kk)

    // regEqualsVal := c.builder.CreateICmp(llvm.IntEQ, reg, val, "")


    return errors.New("unimplemented instruction SeVxByte in codegen")
}

func (sne *SneVxByte) compile(c *Compiler) error {
    return errors.New("unimplemented instruction SneVxByte in codegen")
}

func (se *SeVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction SeVxVy in codegen")
}

func (ld *LdVxByte) compile(c *Compiler) error {
    dst := c.VRegToLLVMValue(ld.vx)
    imm := c.ConstUint8(ld.kk)

    c.builder.CreateStore(imm, dst)

    return nil
}

func (add *AddVxByte) compile(c *Compiler) error {
    return errors.New("unimplemented instruction AddVxVy in codegen")
}

func (ld *LdVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdVxVy in codegen")
}

func (or *OrVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction OrVxVy in codegen")
}

func (and *AndVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction AndVxVy in codegen")
}

func (xor *XorVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction XorVxVy in codegen")
}

func (add *AddVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction AddVxVy in codegen")
}

func (sub *SubVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction SubVxVy in codegen")
}

func (shr *ShrVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction ShrVxVy in codegen")
}

func (subn *SubnVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction SubnVxVy in codegen")
}

func (shl *ShlVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction ShlVxVy in codegen")
}

func (sne *SneVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction SneVxVy in codegen")
}

func (ld *LdIAddr) compile(c *Compiler) error {
    c.builder.CreateStore(c.ConstUint16(ld.addr), c.reg_i)

    return nil
}

func (jp *JpV0Addr) compile(c *Compiler) error {
    return errors.New("unimplemented instruction JpV0Addr in codegen")
}

func (rnd *Rnd) compile(c *Compiler) error {
    random_val := c.builder.CreateCall(c.random_uint8_fn, []llvm.Value{}, "")
    res := c.builder.CreateAnd(random_val, c.ConstUint8(rnd.kk), "")
    dst := c.VRegToLLVMValue(rnd.vx)

    c.builder.CreateStore(res, dst)

    return nil
}

func (drw *DrwVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction DrwVxVy in codegen")
}

func (skp *SkpVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction SkpVx in codegen")
}

func (sknp *SknpVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction SknpVx in codegen")
}

func (ld *LdVxDt) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdVxDt in codegen")
}

func (ld *LdVxK) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdVxK in codegen")
}

func (ld *LdDtVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdDtVx in codegen")
}

func (ld *LdStVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdStVx in codegen")
}

func (add *AddIVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction AddIVx in codegen")
}

func (ld *LdFVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdFVx in codegen")
}

func (ld *LdBVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdBVx in codegen")
}

func (ld *LdIVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdIVx in codegen")
}

func (ld *LdVxI) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdVxI in codegen")
}
