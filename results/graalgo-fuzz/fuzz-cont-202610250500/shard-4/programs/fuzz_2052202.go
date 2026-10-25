package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0() uint32 {
    var v4 string = ("" + "")
    v4 = (v4 + v4)
    _ = v4
    return 42
}
func genFn1(p0 uint8) string {
    var v5 int32 = 101
    p0 -= p0
    v6 := GenS1{F0: -84, F1: 2, F2: 42}
    _ = p0
    _ = v5
    _ = v6
    return "a"
}

func main() {
    var v1 uint8 = 0
    var v2 uint8 = v1
    v3 := []uint32{2}
    v7 := GenS1{F0: 100, F1: 2, F2: 1}
    for fi8 := 0; fi8 < 8; fi8++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi8
    }
    var v9 uint8 = v2
    v9 = (v2 * (v9 >> ((201) & 63)))
    v10 := GenS1{F0: 0, F1: 1, F2: 42}
    v3 = append(v3, 2147483647)
    if (true && (-102 < 0)) {
        if len(v3) > 0 {
            v3[0] = 740
        }
        v11 := GenS1{F0: 1, F1: genFn0(), F2: 2}
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v2 = v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi12
    }
    var v13 uint64 = 2
    v14 := GenS2{F0: 260, F1: 0}
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(v2, v1, v1)
            _ = fi16
        }
        v17 := GenS1{F0: 42, F1: 2, F2: (v13 | v13)}
        var v18 string = "\n\t"
        _ = fi15
        _ = v17
        _ = v18
    }
    var v19 uint64 = (^(v13))
    fmt.Println(v9, v2, v2)
    switch 9223372036854775807 {
    case -172:
        v3 = append(v3, 2147483648)
    }
    v1 += (v2 >> ((uint(v9)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
