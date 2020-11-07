package main

import (
    "chairlift/chairlift"
    "fmt"
    "os"
    "log"
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

        _, err = chairlift.CompileRomToFile(rom, "a.out")
        if err != nil {
            log.Fatal(err)
            continue
        }

        fmt.Println()
    }
}
