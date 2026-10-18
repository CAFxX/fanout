package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }

func main() {
    var v1 uint = 1
    var v2 uint64 = 654
    var v3 uint16 = 255
    var v4 uint16 = uint16(v2)
    v5 := make(map[uint32]int)
    v6 := []string{"foo bar", "", ("\n\t" + ("hello" + "a"))}
    delete(v5, 268)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint32 = 1
        v4 = (v4 & v4)
        fmt.Println(fi7, v2, fi7)
        _ = fi7
        _ = v8
    }
    var v9 int8 = -128
    _ = len(v6)
    v9 += 100
    v10 := GenS1{F0: v9, F1: (42 < 42), F2: -819}
    _ = len(v6)
    var v11 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
