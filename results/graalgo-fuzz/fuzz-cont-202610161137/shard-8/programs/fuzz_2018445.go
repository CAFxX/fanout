package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int8) int16 {
    if true {
        var v4 int32 = 191
        var v5 int32 = (v4 & v4)
        _ = v4
        _ = v5
    }
    p0 -= (-(p0))
    p0 += p0
    p0 = (-(p0))
    if (255 != 255) {
        var v6 string = "hello"
        var v7 int8 = (p0 + (p0 + -128))
        v7 += v7
        _ = v6
        _ = v7
    }
    _ = p0
    return 0
}

func main() {
    var v1 int32 = -55
    v2 := make(map[uint32]bool)
    var v3 int32 = v1
    v1 -= -1
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint16 = 42
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint64 = 18446744073709551615
            switch int(v1) {
            case 128:
                v2[2] = (("hello" + "\n\t") > "\n\t")
            case 2147483647:
                v1 += (^(v1))
            case 1:
                v12 := GenS1{F0: 579, F1: 127, F2: (v11 + v11)}
                _ = v12
            }
            var v13 uint64 = (v11 - (v11 << ((1) & 63)))
            _ = fi10
            _ = v11
            _ = v13
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 bool = true
            v16 := GenS2{F0: 2147483648, F1: 585}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi8
        _ = v9
    }
    if (v1 < (v1 + (^(v1)))) {
        var v17 int32 = (v3 + 0)
        _ = v17
    } else {
        v1 -= 0
        fmt.Println(v3, v1)
        var v18 int16 = 331
        _ = v18
    }
    v19 := []int{}
    if ((v1 / ((-369) | 1)) > 0) {
        var v20 int32 = 319
        var v21 int16 = genFn0(-73)
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 5; fi22++ {
            delete(v2, 4294967295)
            _ = fi22
        }
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v1 = 2147483647
        v24 := GenS2{F0: 557, F1: 42}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        delete(v2, 2147483647)
        var v26 uint = 752
        _ = fi25
        _ = v26
    }
    v1 -= 2147483647
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
