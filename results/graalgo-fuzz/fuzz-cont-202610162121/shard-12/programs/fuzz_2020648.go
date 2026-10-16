package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int16 }

func main() {
    v1 := []string{"\n\t", "a", ("\n\t" + "hello")}
    var v2 bool = (((!((32767 > -240) && false)) || true) || false)
    var v3 uint16 = 32767
    var v4 uint8 = 128
    var v5 int64 = 0
    var v6 uint16 = (v3 % ((768) | 1))
    var v7 uint = 322
    if ((v2 && (v5 <= (-(v5)))) || false) {
        var v8 int = -907
        v3 += (^(v6))
        _ = v8
    }
    var v9 uint8 = 128
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 float32 = (0.1 - 3.14159)
        _ = len(v1)
        v12 := GenS1{F0: (v9 / (((^(v4))) | 1))}
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 = 790
    var v13 uint = (v7 % ((v7) | 1))
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(v3)
        var v15 uint32 = 2
        var v16 uint8 = uint8(v15)
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: -1, F1: 0}
    v18 := GenS1{F0: uint8(v6)}
    fmt.Println(v2, v6)
    _ = len(v1)
    v4 += uint8(v6)
    v19 := GenS1{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
