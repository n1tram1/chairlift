package main

import (
    "chairlift/chairlift"
    "fmt"
    "os"
    "log"
    "strings"
    "path"
)

func main() {
    for _, arg := range os.Args[1:] {
        fmt.Println("processing ", arg)

        rom, err := chairlift.OpenRom(arg)
        if err != nil {
            log.Fatal(err)
            continue
        }
        fmt.Println("rom: ", rom)

        name := strings.TrimSuffix(arg, path.Ext(arg))

        _, err = chairlift.CompileRomToFile(rom, name + ".bc")
        if err != nil {
            log.Fatal(err)
            continue
        }

        fmt.Println()
    }
}
