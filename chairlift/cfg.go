package chairlift

import (
    "fmt"
)

type BasicBlock struct {
    addr int

    instructions []Instruction
    successors []*BasicBlock
}

func (bb *BasicBlock) contains(addr int) bool {
    max_addr := bb.addr + len(bb.instructions) * 2
    return  bb.addr <= addr && addr < max_addr
}

func (bb *BasicBlock) dump(visited map[*BasicBlock]bool) {
    addr := bb.addr

    if visited[bb] == true {
        return
    }

    for _, succ := range bb.successors {
        fmt.Printf("\"%p\" -> \"%p\"\n", bb, succ)
    }

    visited[bb] = true

    node_label := ""

    for _, inst := range bb.instructions {
        node_label += fmt.Sprintf("[0x%X] %#v\\n", addr, inst)
        addr += 2
    }

    fmt.Printf("\"%p\" [label=\"%v\"]\n", bb, node_label)

    for _, succ := range bb.successors {
        succ.dump(visited)
    }
}

func Dump(bb *BasicBlock) {
    visited := map[*BasicBlock]bool{}

    fmt.Println("digraph {")

    bb.dump(visited)

    fmt.Println("}")
}

func (bb *BasicBlock) split(addr int) *BasicBlock {
        new_bb := &BasicBlock{}
        new_bb.addr = bb.addr
        new_bb.successors = []*BasicBlock{bb}

        new_bb_inst_start := (addr - bb.addr) / 2

        new_bb.instructions = bb.instructions[:new_bb_inst_start]
        bb.instructions = bb.instructions[new_bb_inst_start:]
        bb.addr = addr

        return new_bb
}

func (an *flowAnalyzer) redirectSuccessors(from, to *BasicBlock) {
    for _, block := range an.addrToBlock {
        for i := 0; i < len(block.successors); i += 1 {
            if block.successors[i] == from {
                block.successors[i] = to
            }
        }
    }
}

func (an *flowAnalyzer) correspondingBlock(addr int) *BasicBlock {
    var found *BasicBlock

    for _, bb := range an.addrToBlock {
        if bb.contains(addr) {
            found = bb
            break
        }
    }

    if found == nil {
        found = &BasicBlock{}
        found.addr = addr
        found.instructions = []Instruction{}
        found.successors = []*BasicBlock{}
    }

    if found != nil && found.addr != addr {
        split := found.split(addr)

        an.redirectSuccessors(found, split)

        an.addrToBlock[split.addr] = split
    }

    an.addrToBlock[found.addr] = found

    return found 
}

type flowAnalyzer struct {
    instructions map[int]Instruction

    addrToBlock map[int]*BasicBlock
}

func AnalyzeFlow(bytes []byte) *BasicBlock {
    analyzer := flowAnalyzer{}
    analyzer.instructions = map[int]Instruction{}
    analyzer.addrToBlock = map[int]*BasicBlock{}

    analyzer.analyze(bytes, 0)

    return analyzer.addrToBlock[0x200]
}

func (an *flowAnalyzer) analyze(bytes []byte, index int) *BasicBlock {
    bb :=  an.correspondingBlock(index + 0x200)

    for ; index < len(bytes) - 1 && an.instructions[index] == nil; index += 2 {
        opcode := uint16(bytes[index]) << 8 | uint16(bytes[index + 1])

        inst, err := disassemble(opcode)
        if err != nil {
            return nil
        }

        an.instructions[index] = inst
        bb.instructions = append(bb.instructions, inst)

        if isJump(inst) {
            // Nothing to discover after this because it's a jump.
            destination := int(getDestination(inst))
            destinationAsIndex := destination - 0x200

            discovered := an.analyze(bytes, destinationAsIndex)

            if discovered != nil /* && !bb.contains(destination) */ {
                bb.successors = append(bb.successors, discovered)
            }

            break;
        }
    }

    return bb
}