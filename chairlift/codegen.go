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

func (ld *LdVxByte) compile(c *Compiler) error {
    panic("unimplemented")
}
