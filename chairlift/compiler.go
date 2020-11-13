package chairlift

import (
    // "llvm.org/llvm/bindings/go/llvm"
    llvm "github.com/tinygo-org/go-llvm"
    "os"
    "fmt"
)

const (
    PROG_START = 0x200
    RAM_SIZE = 0xFFF
)

type Compiler struct {
    builder llvm.Builder
    mod llvm.Module

    currentBlock *llvm.BasicBlock
    addrToBlock map[int]*llvm.BasicBlock

    bb *BasicBlock

    mainFn llvm.Value

    // C runtime functions
    init_runtime_fn llvm.Value
    random_uint8_fn llvm.Value
    draw_fn llvm.Value
    clear_display_fn llvm.Value

    ram llvm.Value

    // Registers
    reg_i llvm.Value

    reg_v0 llvm.Value
    reg_v1 llvm.Value
    reg_v2 llvm.Value
    reg_v3 llvm.Value
    reg_v4 llvm.Value
    reg_v5 llvm.Value
    reg_v6 llvm.Value
    reg_v7 llvm.Value
    reg_v8 llvm.Value
    reg_v9 llvm.Value
    reg_vA llvm.Value
    reg_vB llvm.Value
    reg_vC llvm.Value
    reg_vD llvm.Value
    reg_vE llvm.Value
    reg_vF llvm.Value
}

func (c *Compiler) VRegToLLVMValue(vreg VReg) llvm.Value {
    switch vreg {
        case V0:
            return c.reg_v0
        case V1:
            return c.reg_v1
        case V2:
            return c.reg_v2
        case V3:
            return c.reg_v3
        case V4:
            return c.reg_v4
        case V5:
            return c.reg_v5
        case V6:
            return c.reg_v6
        case V7:
            return c.reg_v7
        case V8:
            return c.reg_v8
        case V9:
            return c.reg_v9
        case VA:
            return c.reg_vA
        case VB:
            return c.reg_vB
        case VC:
            return c.reg_vC
        case VD:
            return c.reg_vD
        case VE:
            return c.reg_vE
        case VF:
            return c.reg_vF
        default:
            panic(fmt.Sprintf("unkown register %X", vreg))
    }
}

func (c *Compiler) ConstUint8(x uint8) llvm.Value {
    return llvm.ConstInt(llvm.Int8Type(), uint64(x), false)
}

func (c *Compiler) ConstUint16(x uint16) llvm.Value {
    return llvm.ConstInt(llvm.Int16Type(), uint64(x), false)
}

func (c *Compiler) LoadFromAddr(addr llvm.Value) llvm.Value {
    ram_ptr := c.builder.CreateAdd(c.ram, addr, "")

    return c.builder.CreateLoad(ram_ptr, "")
}

func (c *Compiler) selectBlock(bb llvm.BasicBlock) {
    c.builder.SetInsertPointAtEnd(bb)
    c.currentBlock = &bb
}

func newCompiler() (*Compiler) {
    c := new(Compiler)

    c.builder = llvm.NewBuilder()
    c.mod = llvm.NewModule("asm_module")
    c.addrToBlock = map[int]*llvm.BasicBlock{}

    return c
}

func (c *Compiler) createNamedGlobal(intType llvm.Type, name string) llvm.Value {
    val := llvm.ConstInt(intType, 0, false)
    glob := llvm.AddGlobal(c.mod, val.Type(), name)
    glob.SetLinkage(llvm.PrivateLinkage)
    glob.SetInitializer(val)
    return glob
}

func (c *Compiler) addInterpreterData(ram *[]llvm.Value) {
    // Add '0' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xF0))

    // Add '1' sprite
    *ram = append(*ram, c.ConstUint8(0x20))
    *ram = append(*ram, c.ConstUint8(0x60))
    *ram = append(*ram, c.ConstUint8(0x20))
    *ram = append(*ram, c.ConstUint8(0x20))
    *ram = append(*ram, c.ConstUint8(0x70))

    // Add '3' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x10))
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x10))
    *ram = append(*ram, c.ConstUint8(0xF0))

    // Add '4' sprite
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x10))
    *ram = append(*ram, c.ConstUint8(0x10))

    // Add '5' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x80))
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x10))
    *ram = append(*ram, c.ConstUint8(0xF0))

    // Add '6' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x80))
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xF0))

    // Add '7' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x10))
    *ram = append(*ram, c.ConstUint8(0x20))
    *ram = append(*ram, c.ConstUint8(0x40))
    *ram = append(*ram, c.ConstUint8(0x40))

    // Add '8' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xF0))

    // Add '9' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x10))
    *ram = append(*ram, c.ConstUint8(0xF0))

    // Add 'A' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0x90))

    // Add 'B' sprite
    *ram = append(*ram, c.ConstUint8(0xE0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xE0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xE0))

    // Add 'C' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x80))
    *ram = append(*ram, c.ConstUint8(0x80))
    *ram = append(*ram, c.ConstUint8(0x80))
    *ram = append(*ram, c.ConstUint8(0xF0))

    // Add 'D' sprite
    *ram = append(*ram, c.ConstUint8(0xE0))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0x90))
    *ram = append(*ram, c.ConstUint8(0xE0))

    // Add 'E' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x80))
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x80))
    *ram = append(*ram, c.ConstUint8(0xF0))

    // Add 'F' sprite
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x80))
    *ram = append(*ram, c.ConstUint8(0xF0))
    *ram = append(*ram, c.ConstUint8(0x80))
    *ram = append(*ram, c.ConstUint8(0x80))

    for len(*ram) < PROG_START {
        *ram = append(*ram, c.ConstUint8(0))
    }
}

func (c *Compiler) addRemainingZeros(ram *[]llvm.Value) {
    for len(*ram) < RAM_SIZE {
        *ram = append(*ram, c.ConstUint8(0))
    }
}

func (c *Compiler) createRamArray(bytes []byte) {
    ramValues := make([]llvm.Value, 0, RAM_SIZE)

    c.addInterpreterData(&ramValues)

    for _, b := range bytes {
        val := c.ConstUint8(b)
        ramValues = append(ramValues, val)
    }

    c.addRemainingZeros(&ramValues)

    ramConst := llvm.ConstArray(llvm.ArrayType(llvm.Int8Type(), len(ramValues)), ramValues)

    c.ram = llvm.AddGlobal(c.mod, ramConst.Type(), "ram")
    c.ram.SetLinkage(llvm.PrivateLinkage)
    c.ram.SetInitializer(ramConst)
    c.ram.SetGlobalConstant(true)
}

func (c *Compiler) createRegisters() {
    c.reg_i = c.createWordRegister("I")

    c.reg_v0 = c.createByteRegister("V0")
    c.reg_v1 = c.createByteRegister("V1")
    c.reg_v2 = c.createByteRegister("V2")
    c.reg_v3 = c.createByteRegister("V3")
    c.reg_v4 = c.createByteRegister("V4")
    c.reg_v5 = c.createByteRegister("V5")
    c.reg_v6 = c.createByteRegister("V6")
    c.reg_v7 = c.createByteRegister("V7")
    c.reg_v8 = c.createByteRegister("V8")
    c.reg_v9 = c.createByteRegister("V9")
    c.reg_vA = c.createByteRegister("VA")
    c.reg_vB = c.createByteRegister("VB")
    c.reg_vC = c.createByteRegister("VC")
    c.reg_vD = c.createByteRegister("VD")
    c.reg_vE = c.createByteRegister("VE")
    c.reg_vF = c.createByteRegister("VF")
}

func (c *Compiler) createWordRegister(name string) llvm.Value {
    return c.createNamedGlobal(llvm.Int16Type(), name)
}

func (c *Compiler) createByteRegister(name string) llvm.Value {
    return c.createNamedGlobal(llvm.Int8Type(), name)
}

func (c *Compiler) createMain() {
    mainType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0), llvm.PointerType(llvm.Int8Type(), 0)}, false)

    c.mainFn = llvm.AddFunction(c.mod, "main", mainType)
    c.mainFn.SetFunctionCallConv(llvm.CCallConv)

    entry := llvm.AddBasicBlock(c.mainFn, "entry")
    c.selectBlock(entry)

    c.builder.CreateCall(c.init_runtime_fn, []llvm.Value{}, "")
}

func (c *Compiler) createCBindings() {
    random_uint8_fn_type := llvm.FunctionType(llvm.Int8Type(), []llvm.Type{}, false)
    c.random_uint8_fn = llvm.AddFunction(c.mod, "random_uint8", random_uint8_fn_type)
    c.random_uint8_fn.SetLinkage(llvm.ExternalLinkage)

    // void draw(uint8_t vx, uint8_t vy, uint16_t I, uint8_t *ram, uint8_t n, uint8_t *collision)
    draw_fn_type := llvm.FunctionType(llvm.VoidType(), []llvm.Type{llvm.Int8Type(), llvm.Int8Type(), llvm.Int16Type(), c.ram.Type(), llvm.Int8Type(), llvm.PointerType(llvm.Int8Type(), 0)}, false)
    c.draw_fn = llvm.AddFunction(c.mod, "draw", draw_fn_type)
    c.draw_fn.SetLinkage(llvm.ExternalLinkage)

    init_runtime_type := llvm.FunctionType(llvm.VoidType(), []llvm.Type{}, false)
    c.init_runtime_fn = llvm.AddFunction(c.mod, "init_runtime", init_runtime_type)
    c.init_runtime_fn.SetLinkage(llvm.ExternalLinkage)

    clear_fn_type := llvm.FunctionType(llvm.VoidType(), []llvm.Type{}, false)
    c.clear_display_fn = llvm.AddFunction(c.mod, "clear_display", clear_fn_type)
    c.clear_display_fn.SetLinkage(llvm.ExternalLinkage)
}

func (c *Compiler) addBasicBlock(bb *BasicBlock) {
        block := llvm.AddBasicBlock(c.mainFn, fmt.Sprintf("block_%x", bb.addr))
        c.addrToBlock[bb.addr] = &block
}

func (c *Compiler) createBasicBlocks(cfg *BasicBlock) {
    visited := map[int]bool{}

    for bb := cfg; bb != nil && !visited[bb.addr]; bb = bb.jump_successor {
        c.addBasicBlock(bb)

        if bb.fallthrough_successor != nil {
            c.addBasicBlock(bb.fallthrough_successor)
        }

        visited[bb.addr] = true
    }
}

func (c *Compiler) compileBb(bb *BasicBlock) (*llvm.BasicBlock, error) {
    c.bb = bb

    block := c.addrToBlock[bb.addr]
    c.selectBlock(*block)


    for _, inst := range bb.instructions {
        err := inst.compile(c)

        if err != nil {
            return nil, err
        }
    }

    return block, nil
}

func (c *Compiler) fixUnterminatedBasicBlocks(cfg *BasicBlock) {
    visited := map[int]bool{}

    for bb := cfg; bb != nil && !visited[bb.addr]; bb = bb.jump_successor {
        visited[bb.addr] = true

        if !bb.willNeedTermination {
            continue
        }
        currBlock := c.addrToBlock[bb.addr]
        succBlock := c.addrToBlock[bb.jump_successor.addr]

        c.selectBlock(*currBlock)
        c.builder.CreateBr(*succBlock)
    }

}

func (c *Compiler) linkEntryToFirstBlock(cfg *BasicBlock) {
    firstBlock := c.addrToBlock[cfg.addr]

    c.selectBlock(c.mainFn.FirstBasicBlock())
    c.builder.CreateBr(*firstBlock)
}

func (c *Compiler) compile(rom *Rom) error {
    c.createRamArray(rom.bytes)
    c.createRegisters()

    c.createCBindings()

    c.createMain()


    _, cfg := AnalyzeFlow(rom.bytes)

    c.createBasicBlocks(cfg)

    visited := map[int]bool{}
    for bb := cfg; bb != nil && !visited[bb.addr]; bb = bb.jump_successor {
        _, err := c.compileBb(bb)
        if err != nil {
            c.mod.Dump()
            return err
        }

        visited[bb.addr] = true
    }

    c.fixUnterminatedBasicBlocks(cfg)
    c.linkEntryToFirstBlock(cfg)

    err := llvm.VerifyModule(c.mod, llvm.ReturnStatusAction)
    if err != nil {
        return err
    }
    c.mod.Dump()

    return nil
}

func CompileRomToFile(rom *Rom, filename string) (*Compiler, error) {
    file, err := os.Create(filename)
    if err != nil {
        return nil, err
    }

    c := newCompiler()

    err = c.compile(rom)
    if err != nil {
        return nil, err
    }

    err = llvm.WriteBitcodeToFile(c.mod, file)
    if err != nil {
        return nil, err
    }


    return c, nil
}
