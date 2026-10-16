package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int }

func main() {
    var v1 float32 = (3.14159 + 1.5)
    var v2 uint32 = 151
    var v3 uint64 = 831
    v4 := []uint32{65535}
    var v5 int8 = 61
    var v6 bool = true
    v3 = 2
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int32 = 0
            if (!((v9 << ((243) & 63)) == v9)) {
                v10 := GenS1{F0: 763}
                _ = v10
            }
            if v6 {
                _ = len(v4)
            }
            _ = fi8
            _ = v9
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            _ = len(v4)
            var v12 int32 = -1
            _ = fi11
            _ = v12
        }
        _ = fi7
    }
    var v13 uint64 = v3
    v14 := GenS2{F0: -32768, F1: 488}
    fmt.Println(v2)
    if v6 {
        if false {
            var v15 uint16 = 42
            fmt.Println(v2)
            var v16 int16 = 32767
            _ = v15
            _ = v16
        }
        if len(v4) > 0 {
            v4[0] = 898
        }
    } else {
        fmt.Println(v6, v2)
        if len(v4) > 0 {
            v4[0] = v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
