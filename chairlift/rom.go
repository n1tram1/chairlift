package chairlift

import (
    "io/ioutil"
)

type Rom struct {
    bytes []byte
    cfg *BasicBlock
}

func OpenRom(filename string) (*Rom, error) {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    rom := new(Rom)
    rom.bytes = bytes

    cfg := AnalyzeFlow(bytes)
    Dump(cfg)

    panic("")

    return rom, nil
}

func (r *Rom) GetInstruction(addr int) Instruction {
    return &Sys{}
    // return r.instructions[addr - 0x200]
}

func (r *Rom) Iterate(f func(addr int, inst Instruction) error) error {
    return nil
    // for addr := 0x200; addr < r.maxCode; addr += 2 {
    //     inst := r.GetInstruction(addr)
    //     if inst == nil {
    //         continue
    //     }
    //
    //     err := f(addr, inst)
    //     if err != nil {
    //         return err
    //     }
    // }
    //
    // return nil
}

