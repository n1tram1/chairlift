package chairlift

import (
    // "llvm.org/llvm/bindings/go/llvm"
    llvm "github.com/tinygo-org/go-llvm"
    "os"
)

type Compiler struct {
    builder llvm.Builder
    mod llvm.Module

    currentBlock *llvm.BasicBlock

    mainFn llvm.Value

    v0 llvm.Value
    v1 llvm.Value
}

func (c *Compiler) selectBlock(bb llvm.BasicBlock) {
    c.builder.SetInsertPointAtEnd(bb)
    c.currentBlock = &bb
}

func newCompiler() (*Compiler) {
    c := new(Compiler)

    c.builder = llvm.NewBuilder()
    c.mod = llvm.NewModule("asm_module")

    return c
}

func (c *Compiler) createNamedGlobal(intType llvm.Type, name string) llvm.Value {
    val := llvm.ConstInt(intType, 0, false)
    glob := llvm.AddGlobal(c.mod, val.Type(), name)
    glob.SetLinkage(llvm.PrivateLinkage)
    glob.SetInitializer(val)
    return glob
}

func (c *Compiler) createByteRegister(name string) llvm.Value {
    return c.createNamedGlobal(llvm.Int8Type(), name)
}

func (c *Compiler) createRegisters() {
    c.v0 = c.createByteRegister("V0")
    c.v1 = c.createByteRegister("V1")
}

func (c *Compiler) compile(rom *Rom) error {
    c.createRegisters()

    mainType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0), llvm.PointerType(llvm.Int8Type(), 0)}, false)

    c.mainFn = llvm.AddFunction(c.mod, "main", mainType)
    c.mainFn.SetFunctionCallConv(llvm.CCallConv)

    entry := llvm.AddBasicBlock(c.mainFn, "entry")
    c.builder.SetInsertPointAtEnd(entry)

    err := compile_instructions(rom.instructions, c)
    if err != nil {
        return err
    }

    result := llvm.ConstInt(llvm.Int32Type(), 42, false)
    c.builder.CreateRet(result)

    err = llvm.VerifyModule(c.mod, llvm.ReturnStatusAction)
    if err != nil {
        return err
    }

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
