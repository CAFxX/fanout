package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0(p0 uint64, p1 int64) uint64 {
    var v6 uint16 = 255
    var v7 float32 = 3.14159
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint, p1 bool) int64 {
    p1 = p1
    for fi8 := 0; fi8 < 14; fi8++ {
        p0 += (p0 & p0)
        fmt.Println(p0, p1, fi8)
        _ = fi8
    }
    var v9 uint32 = 0
    v10 := GenS1{F0: 42, F1: 2, F2: ""}
    var v11 bool = (!p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 313
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    var v3 uint8 = 44
    v4 := []uint32{}
    var v5 uint = 910
    var v12 uint = (v5 / ((375) | 1))
    fmt.Println(v3, v3)
    v5 = v5
    if false {
        v5 -= v5
        v5 -= (^(v5))
    } else {
        v3 = v3
        var v13 int8 = -21
        _ = v13
    }
    v5 = (v5 % ((v5) | 1))
    for fi14 := 0; fi14 < 8; fi14++ {
        v3 += (v3 >> (((v5 | v5)) & 63))
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 float32 = (1e10 + 1e10)
            _ = fi15
            _ = v16
        }
        _ = len(v1)
        _ = fi14
    }
    v17 := GenS1{F0: -128, F1: uint64(v3), F2: (("" + (("a" + "foo bar") + ((("foo bar" + "foo bar") + "") + "foo bar"))) + "")}
    v12 -= v5
    if false {
        fmt.Println(v3)
        var v18 bool = false
        _ = v18
    }
    v19 := GenS1{F0: 100, F1: genFn0(genFn0(uint64(v12), -1), 9223372036854775807), F2: ("foo bar" + ("hello" + "hello"))}
    fmt.Println(v5)
    v5 = v12
    v4 = append(v4, 1)
    v20 := GenS2{F0: true, F1: -117, F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
