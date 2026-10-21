package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
func genFn0() uint8 {
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int = fi5
        _ = fi5
        _ = v6
    }
    v4 -= v4
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v4, fi7)
        v8 := GenS1{F0: -1, F1: 1}
        _ = fi7
        _ = v8
    }
    fmt.Println(v4)
    _ = v4
    return 128
}
func genFn1(p0 uint64, p1 int8) int8 {
    switch "hello" {
    case "\n\t":
        var v9 uint64 = 480
        _ = v9
    }
    p0 = 42
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 uint8 = 14
    v2 := make(map[uint32]int)
    v3 := []uint32{0, 2, 0, 4294967295}
    v1 = genFn0()
    v10 := GenS1{F0: 0, F1: 0}
    fmt.Println(v1)
    for fi11 := 0; fi11 < 7; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi11
    }
    switch 4294967295 {
    case 2147483647:
        for fi12 := 0; fi12 < 18; fi12++ {
            v1 = (v1 ^ v1)
            _ = fi12
        }
    case 127:
        v13 := GenS1{F0: -9, F1: -1}
        _ = v13
    case 269:
        var v14 uint32 = 2147483647
        _ = v14
    default:
        var v15 uint32 = 42
        _ = v15
    }
    v3 = append(v3, 2147483647)
    var v16 uint32 = 65535
    switch 798 {
    case 32768:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -2147483648:
        v3 = append(v3, (v16 << ((547) & 63)))
    default:
        var v17 uint8 = (v1 | genFn0())
        _ = v17
    }
    var v18 int = 42
    v19 := GenS1{F0: 0, F1: -1}
    var v20 string = "\n\t"
    v3 = append(v3, (^(v16)))
    v21 := GenS1{F0: 42, F1: -502}
    if (!false) {
        v1 = genFn0()
        var v22 bool = (false && (1 >= 1))
        for fi23 := 0; fi23 < 17; fi23++ {
            v22 = v22
            v16 += v16
            _ = fi23
        }
        _ = v22
    } else {
        v24 := v2[684]
        _ = v24
    }
    switch 857 {
    case -523:
        fmt.Println(v1, v18, v1)
    case 0:
        v18 = v18
    case 32767:
        v1 = 255
    }
    var v25 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
}
