package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]uint32)
    var v3 bool = (!(42 != 1))
    v1 += v1
    var v4 bool = (0.1 != -2.25)
    v5 := []int{128}
    v2[154] = 42
    switch 127 {
    case -2147483648:
        v4 = (!true)
    case 42:
        var v6 int32 = v1
        _ = v6
    case -142:
        var v7 uint64 = 1
        _ = v7
    default:
        v1 = (v1 >> ((18446744073709551615) & 63))
    }
    var v8 uint32 = 4294967295
    v9 := GenS1{F0: (v8 & (v8 * (v8 & v8))), F1: 907}
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS1{F0: (v8 ^ 2147483647), F1: 231}
        var v12 uint64 = 649
        if ((v4 && true) && (!(1 >= 654))) {
            var v13 uint16 = 612
            v4 = ((v12 > (^(v12))) && v3)
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v4)
    v2[(v8 + 4294967295)] = (v8 >> ((2) & 63))
    v1 -= v1
    v3 = ((v1 + (v1 / (((v1 >> ((2) & 63))) | 1))) == (-(v1)))
    v14 := GenS1{F0: (v8 % ((v8) | 1)), F1: 65535}
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 int = fi15
        var v17 float32 = 1.5
        v18 := v2[v8]
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint64 = 4294967295
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS2{F0: "", F1: 42}
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: (v8 + (v8 << ((1) & 63))), F1: 997}
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(v4, v8)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
