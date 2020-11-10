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

    _, rom.cfg = AnalyzeFlow(bytes)
    Dump(rom.cfg)

    return rom, nil
}
