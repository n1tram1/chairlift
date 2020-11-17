package chairlift

import (
    llvm "github.com/tinygo-org/go-llvm"
    "errors"
    "fmt"
)

func (cls *Cls) compile(c *Compiler) error {
    c.builder.CreateCall(c.clear_display_fn, []llvm.Value{}, "")

    return nil
}

func (ret  *Ret) compile(c *Compiler) error {
    return errors.New("unimplemented instruction Ret in codegen")
}

func (sys *Sys) compile(c* Compiler) error {
    return errors.New("unimplemented instruction Sys in codegen")
}

func (jp *JpAddr) compile(c *Compiler) error {
    dest := c.addrToBlock[int(jp.addr)]

    c.builder.CreateBr(*dest)

    return nil
}

func (call *CallAddr) compile(c *Compiler) error {
    return errors.New("unimplemented instruction CallAddr in codegen")
}

func (se *SeVxByte) compile(c *Compiler) error {
    reg := c.builder.CreateLoad(c.VRegToLLVMValue(se.vx), "")
    val := c.ConstUint8(se.kk)

    regEqualsVal := c.builder.CreateICmp(llvm.IntEQ, reg, val, "")

    err := c.CreateCondBranch(regEqualsVal)
    if err != nil {
        return err
    }

    return nil
}

func (sne *SneVxByte) compile(c *Compiler) error {
    vx_val := c.builder.CreateLoad(c.VRegToLLVMValue(sne.vx), "")
    val := c.ConstUint8(sne.kk)

    vxNotEqualsVal := c.builder.CreateICmp(llvm.IntNE, vx_val, val, "")

    c.CreateCondBranch(vxNotEqualsVal)

    return nil
}

func (se *SeVxVy) compile(c *Compiler) error {
    vx_val := c.builder.CreateLoad(c.VRegToLLVMValue(se.vx), "")
    vy_val := c.builder.CreateLoad(c.VRegToLLVMValue(se.vy), "")

    vxEqualsVy := c.builder.CreateICmp(llvm.IntEQ, vx_val, vy_val, "")

    c.CreateCondBranch(vxEqualsVy)

    return nil
}

func (ld *LdVxByte) compile(c *Compiler) error {
    dst := c.VRegToLLVMValue(ld.vx)
    imm := c.ConstUint8(ld.kk)

    c.builder.CreateStore(imm, dst)

    return nil
}

func (add *AddVxByte) compile(c *Compiler) error {
    dest := c.VRegToLLVMValue(add.vx)
    vx_val := c.builder.CreateLoad(dest, "")
    operation_res := c.builder.CreateAdd(vx_val, c.ConstUint8(add.kk), "")

    c.builder.CreateStore(operation_res, dest)

    return nil
}

func (ld *LdVxVy) compile(c *Compiler) error {
    vy := c.builder.CreateLoad(c.VRegToLLVMValue(ld.vy), "")

    c.builder.CreateStore(vy, c.VRegToLLVMValue(ld.vx))

    return nil
}

func (or *OrVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction OrVxVy in codegen")
}

func (and *AndVxVy) compile(c *Compiler) error {
    return errors.New("unimplemented instruction AndVxVy in codegen")
}

func (xor *XorVxVy) compile(c *Compiler) error {
    vx_val := c.builder.CreateLoad(c.VRegToLLVMValue(xor.vx), "")
    vy_val := c.builder.CreateLoad(c.VRegToLLVMValue(xor.vy), "")
    result := c.builder.CreateXor(vx_val, vy_val, "")

    c.builder.CreateStore(result, c.VRegToLLVMValue(xor.vx))

    return nil
}

func (add *AddVxVy) compile(c *Compiler) error {
    vx_val := c.builder.CreateLoad(c.VRegToLLVMValue(add.vx), "")
    vy_val := c.builder.CreateLoad(c.VRegToLLVMValue(add.vy), "")
    result := c.builder.CreateAdd(vx_val, vy_val, "")

    smallerThanVx := c.builder.CreateICmp(llvm.IntULT, result, vx_val, "smallerThanVx")
    smallerThanVy := c.builder.CreateICmp(llvm.IntULT, result, vy_val, "smallerThanVy")
    tmp := c.builder.CreateAdd(c.CastU8(smallerThanVx), c.CastU8(smallerThanVy), "tmp")
    check_overflow := c.CastU8(c.builder.CreateICmp(llvm.IntUGE, tmp, c.ConstUint8(1), "check_overflow"))

    c.builder.CreateStore(check_overflow, c.reg_vF)

    c.builder.CreateStore(result, c.VRegToLLVMValue(add.vx))

    return nil
}

func (sub *SubVxVy) compile(c *Compiler) error {
    // TODO: I'm 100% sure `check_overflow` is wrong
    vx_val := c.builder.CreateLoad(c.VRegToLLVMValue(sub.vx), "")
    vy_val := c.builder.CreateLoad(c.VRegToLLVMValue(sub.vy), "")
    result := c.builder.CreateSub(vx_val, vy_val, "")

    greaterThanVx := c.builder.CreateICmp(llvm.IntUGT, result, vx_val, "greaterThanVx")
    greaterThanVy := c.builder.CreateICmp(llvm.IntUGT, result, vy_val, "greaterThanVy")
    tmp := c.builder.CreateAdd(c.CastU8(greaterThanVx), c.CastU8(greaterThanVy), "tmp")
    check_underflow := c.CastU8(c.builder.CreateICmp(llvm.IntUGE, tmp, c.ConstUint8(1), "check_overflow"))

    c.builder.CreateStore(check_underflow, c.reg_vF)

    c.builder.CreateStore(result, c.VRegToLLVMValue(sub.vx))

    return nil
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
    vx_val := c.builder.CreateLoad(c.VRegToLLVMValue(sne.vx), "")
    vy_val := c.builder.CreateLoad(c.VRegToLLVMValue(sne.vy), "")

    vxNotEqualsVy := c.builder.CreateICmp(llvm.IntNE, vx_val, vy_val, "")

    c.CreateCondBranch(vxNotEqualsVy)

    return nil
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
    vx := c.builder.CreateLoad(c.VRegToLLVMValue(drw.vx), "")
    vy := c.builder.CreateLoad(c.VRegToLLVMValue(drw.vy), "")
    I := c.builder.CreateLoad(c.reg_i, "")
    n := c.ConstUint8(drw.n)
    c.builder.CreateCall(c.draw_fn, []llvm.Value{vx, vy, I, c.ram, n, c.reg_vF}, "")

    return nil
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
    vx_val := c.builder.CreateLoad(c.VRegToLLVMValue(add.vx), "")
    vx_val_as_u16 := c.CastU16(vx_val)
    i_val := c.builder.CreateLoad(c.reg_i, "")
    res := c.builder.CreateAdd(vx_val_as_u16, i_val, fmt.Sprintf("AddIVx_%x", c.bb.addr))

    c.builder.CreateStore(res, c.reg_i)

    return nil
}

func (ld *LdFVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdFVx in codegen")
}

func (ld *LdBVx) compile(c *Compiler) error {
    return errors.New("unimplemented instruction LdBVx in codegen")
}

// TODO: GEP related stuff causes a segfault
func (ld *LdIVx) compile(c *Compiler) error {
    for i := VReg(0); i <= ld.vx; i++ {
        ptr := c.builder.CreateGEP(
            c.ram,
            []llvm.Value{
                c.builder.CreateLoad(c.reg_i, ""),
                c.ConstUint16(uint16(i)),
            },
            "",
        )

        mem_val := c.builder.CreateLoad(ptr, fmt.Sprintf("load_%v", i))
        c.builder.CreateStore(mem_val, c.VRegToLLVMValue(i))
    }

    return nil
}

func (ld *LdVxI) compile(c *Compiler) error {
    for i := VReg(0); i <= ld.vx; i++ {
        ptr := c.builder.CreateGEP(
            c.ram,
            []llvm.Value{
                c.builder.CreateLoad(c.reg_i, ""),
                c.ConstUint16(uint16(i)),
            },
            "",
        )

        c.builder.CreateStore(c.builder.CreateLoad(c.VRegToLLVMValue(i), ""), ptr)
    }

    return nil
}
