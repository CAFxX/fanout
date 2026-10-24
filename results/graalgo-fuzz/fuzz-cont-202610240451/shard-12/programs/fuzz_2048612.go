package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint64; F1 int; F2 uint32 }
func genFn0() int64 {
    var v4 uint16 = 32767
    v4 = (v4 + (v4 << ((1) & 63)))
    _ = v4
    return 0
}

func main() {
    v1 := []uint32{448, 2147483647}
    v2 := make(map[uint32]int)
    var v3 int16 = -847
    var v5 uint = 18446744073709551615
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(v3, v5, v5)
            _ = fi7
        }
        _ = len(v2)
        var v8 uint16 = 32767
        _ = fi6
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v5 = 2
        v5 = (v5 * 532)
        _ = fi9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: 42, F1: v3}
        _ = fi10
        _ = v11
    }
    v3 = (v3 ^ v3)
    v3 += (-(v3))
    v5 += (v5 & (^(v5)))
    v5 = v5
    v5 -= 4294967295
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int8 = -1
        switch "foo bar" {
        case "":
            var v14 uint16 = 65535
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v5, v3)
    v15 := v2[637]
    v16 := GenS2{F0: 228, F1: (v15 & v15), F2: 487}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
}
