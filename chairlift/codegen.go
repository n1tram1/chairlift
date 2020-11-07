package chairlift

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
    panic("unimplemented")
}

func (ret  *Ret) compile(c *Compiler) error {
    panic("unimplemented")
}

func (jp *JpAddr) compile(c *Compiler) error {
    panic("unimplemented")
}

func (call *CallAddr) compile(c *Compiler) error {
    panic("unimplemented")
}

func (se *SeVxByte) compile(c *Compiler) error {
    panic("unimplemented")
}

func (sne *SneVxByte) compile(c *Compiler) error {
    panic("unimplemented")
}

func (se *SeVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdVxByte) compile(c *Compiler) error {
    panic("unimplemented")
}

func (add *AddVxByte) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (or *OrVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (and *AndVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (xor *XorVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (add *AddVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (sub *SubVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (shr *ShrVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (subn *SubnVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (shl *ShlVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (sne *SneVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdIAddr) compile(c *Compiler) error {
    panic("unimplemented")
}

func (jp *JpV0Addr) compile(c *Compiler) error {
    panic("unimplemented")
}

func (rnd *Rnd) compile(c *Compiler) error {
    panic("unimplemented")
}

func (drw *DrwVxVy) compile(c *Compiler) error {
    panic("unimplemented")
}

func (skp *SkpVx) compile(c *Compiler) error {
    panic("unimplemented")
}

func (sknp *SknpVx) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdVxDt) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdVxK) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdDtVx) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdStVx) compile(c *Compiler) error {
    panic("unimplemented")
}

func (add *AddIVx) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdFVx) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdBVx) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdIVx) compile(c *Compiler) error {
    panic("unimplemented")
}

func (ld *LdVxI) compile(c *Compiler) error {
    panic("unimplemented")
}
