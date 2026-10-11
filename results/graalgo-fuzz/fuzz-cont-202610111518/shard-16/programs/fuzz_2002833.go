package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 int; F1 uint; F2 int32 }

func main() {
    var v1 uint64 = 286
    var v2 float64 = -2.25
    var v3 uint32 = 0
    v4 := []uint32{uint32(v1)}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (!true) {
        v5 := GenS1{F0: (((1 != 42) || ("" == (("\n\t" + "a") + "\n\t"))) || (v3 < v3)), F1: -1, F2: (v1 / ((2) | 1))}
        for fi6 := 0; fi6 < 14; fi6++ {
            v1 += 2
            _ = len(v4)
            _ = fi6
        }
        v2 = v2
        _ = v5
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float64 = v2
        v2 = v2
        _ = fi7
        _ = v8
    }
    if len(v4) > 0 {
        v4[0] = (^(v3))
    }
    var v9 uint = 0
    fmt.Println(v1, v9, v2)
    var v10 uint8 = 10
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint64 = (v1 << ((255) & 63))
        if false {
            var v13 uint8 = (^(v10))
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    v14 := GenS2{F0: 42, F1: v9, F2: -2147483648}
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(v3)
        var v16 uint32 = v3
        for fi17 := 0; fi17 < 4; fi17++ {
            v4 = append(v4, (v3 << ((993) & 63)))
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    v1 = 4294967295
    v18 := GenS1{F0: (9223372036854775807 < -645), F1: 9223372036854775807, F2: (v1 ^ v1)}
    for fi19 := 0; fi19 < 9; fi19++ {
        if len(v4) > 0 {
            v4[0] = (v3 << ((v9) & 63))
        }
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 uint16 = 42
            fmt.Println(v9, v1)
            var v22 int = -486
            _ = fi20
            _ = v21
            _ = v22
        }
        v3 += (^(v3))
        _ = fi19
    }
    if (-41 == 100) {
        v1 = (^(v1))
        var v23 uint32 = uint32(v9)
        fmt.Println(v2)
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
