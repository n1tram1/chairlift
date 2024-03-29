package chairlift

import (
    "fmt"
)

type BasicBlock struct {
    addr int

    instructions []Instruction
    willNeedTermination bool

    fallthrough_successor *BasicBlock
    jump_successor *BasicBlock
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

    if bb.jump_successor != nil {
        fmt.Printf("\"%p\" -> \"%p\"\n", bb, bb.jump_successor)
    }

    if bb.fallthrough_successor != nil {
        fmt.Printf("\"%p\" -> \"%p\" [color=green]\n", bb, bb.fallthrough_successor)
    }

    visited[bb] = true

    node_label := ""

    for _, inst := range bb.instructions {
        node_label += fmt.Sprintf("[0x%X] %#v\\n", addr, inst)
        addr += 2
    }

    if bb.willNeedTermination {
        fmt.Printf("\"%p\" [label=\"%v\" color=red]\n", bb, node_label)
    } else {
        fmt.Printf("\"%p\" [label=\"%v\"]\n", bb, node_label)
    }

    if bb.fallthrough_successor != nil {
        bb.fallthrough_successor.dump(visited)
    }

    if bb.jump_successor != nil {
        bb.jump_successor.dump(visited)
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
        new_bb.willNeedTermination = true

        if addr == 0x200 || bb.addr == 0x200 {
            fmt.Println("")
        }

        new_bb_inst_start := (addr - bb.addr) / 2

        new_bb.instructions = bb.instructions[:new_bb_inst_start]
        bb.instructions = bb.instructions[new_bb_inst_start:]

        new_bb.jump_successor = bb
        bb.addr = addr

        return new_bb
}

func (an *FlowAnalyzer) redirectSuccessors(from, to *BasicBlock) {
    for _, block := range an.addrToBlock {
        if block.jump_successor == from {
            block.jump_successor = to
        }

        if block.fallthrough_successor == from {
            block.fallthrough_successor = to
        }
    }
}

func (an *FlowAnalyzer) correspondingBlock(addr int) *BasicBlock {
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
    } else if found != nil && found.addr != addr {
        split := found.split(addr)

        an.redirectSuccessors(found, split)

        an.addrToBlock[split.addr] = split
    }

    an.addrToBlock[found.addr] = found

    return found 
}

type InstructionPair struct {
    addr int
    instruction *Instruction
}

type FlowAnalyzer struct {
    instructions map[int]Instruction

    addrToBlock map[int]*BasicBlock
}

func AnalyzeFlow(bytes []byte) (*FlowAnalyzer, *BasicBlock) {
    analyzer := FlowAnalyzer{}
    analyzer.instructions = map[int]Instruction{}
    analyzer.addrToBlock = map[int]*BasicBlock{}

    analyzer.discoverCode(bytes, 0)

    visited := map[int]bool{}
    analyzer.analyze(bytes, 0, visited)

    return &analyzer, analyzer.addrToBlock[0x200]
}

func (an *FlowAnalyzer) discoverCode(bytes []byte, index int) error {
    an.addrToBlock[index + 0x200] = &BasicBlock{addr:index + 0x200, instructions: []Instruction{}}

    for ; index < len(bytes) - 1 && an.instructions[index] == nil; index += INSTRUCTION_SIZE {
        opcode := uint16(bytes[index]) << 8 | uint16(bytes[index + 1])

        inst, err := disassemble(opcode)
        if err != nil {
            return nil
        }

        an.instructions[index] = inst

        if isJump(inst) {
            // Nothing to discover after this because it's a jump.
            destination := int(getJumpDestination(inst))
            destinationAsIndex := destination - 0x200

            err := an.discoverCode(bytes, destinationAsIndex)
            if err != nil {
                return err
            }

            break
        }

        if isBranch(inst) {
            possibleDestinationIndex := getBranchPossibleDestination(index)

            err := an.discoverCode(bytes, possibleDestinationIndex)
            if err != nil {
                return err
            }

            err = an.discoverCode(bytes, index + 2)
            if err != nil {
                return err
            }

            break
        }
    }

    return nil
}

func (an *FlowAnalyzer) analyze(bytes []byte, index int, visited map[int]bool) *BasicBlock {
    // bb :=  an.correspondingBlock(index + 0x200)
    bb := an.addrToBlock[index + 0x200]

    for ; index < len(bytes) - 1 && !visited[index + 0x200]; index += 2 {
        visited[index + 0x200] = true
        // opcode := uint16(bytes[index]) << 8 | uint16(bytes[index + 1])
        //
        // inst, err := disassemble(opcode)
        // if err != nil {
        //     return nil
        // }

        inst := an.instructions[index]
        bb.instructions = append(bb.instructions, inst)

        if isJump(inst) {
            // Nothing to discover after this because it's a jump.
            destination := int(getJumpDestination(inst))
            destinationAsIndex := destination - 0x200

            discovered := an.analyze(bytes, destinationAsIndex, visited)

            if discovered != nil {
                bb.jump_successor = discovered
            }

            break;
        }

        if isBranch(inst) {
            // destinationAsIndex := getBranchPossibleDestination(index)


            discovered_branch := an.analyze(bytes, index + 4, visited)
            discovered_fallthrough := an.analyze(bytes, index + 2, visited)

            if discovered_branch != nil {
                bb.jump_successor = discovered_branch
            }

            if discovered_fallthrough != nil {
                bb.fallthrough_successor = discovered_fallthrough

                if discovered_fallthrough.jump_successor == nil {
                    discovered_fallthrough.jump_successor = discovered_branch
                    discovered_fallthrough.willNeedTermination = true
                }
            }

            break
        }

        if an.addrToBlock[index + 0x200 + INSTRUCTION_SIZE] != nil {
            bb.jump_successor = an.analyze(bytes, index + INSTRUCTION_SIZE, visited)

            break
        }
    }

    if len(bb.instructions) > 0{
        last_inst := bb.instructions[len(bb.instructions) - 1]

        if !isTerminator(last_inst) {
            bb.willNeedTermination = true
        }
    }

    return bb
}

func isTerminator(inst Instruction) bool {
    return isJump(inst) || isBranch(inst)
}

func isJump(inst Instruction) bool {
    switch inst.(type) {
    case *Sys, *JpAddr, *CallAddr:
        return true
    }

    return false
}

func isBranch(inst Instruction) bool {
    switch inst.(type) {
    case *SeVxByte, *SneVxByte, *SeVxVy, *SneVxVy:
        return true
    }

    return false
}

func getJumpDestination(inst Instruction) uint16 {
    switch inst.(type) {
    case *Sys:
        sys := inst.(*Sys)
        return sys.addr
    case *JpAddr:
        jp := inst.(*JpAddr)
        return jp.addr
    case *CallAddr:
        call := inst.(*CallAddr)
        return call.addr
    default:
        panic("unreachable")
    }
}

func getBranchPossibleDestination(addr_or_offset int) int {
    return addr_or_offset + 4
}
