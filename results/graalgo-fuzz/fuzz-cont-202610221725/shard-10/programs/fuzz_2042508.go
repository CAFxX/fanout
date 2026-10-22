package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
type GenS2 struct { F0 string; F1 uint16; F2 uint32 }
func genFn0(p0 bool) int16 {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: 65535, F1: 128, F2: false}
        v6 := GenS1{F0: 2147483647, F1: 127, F2: ((-2147483648 < fi4) || p0)}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{127}
    _ = p0
    _ = v7
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 255
    v3 := make(map[uint32]bool)
    v8 := GenS2{F0: "\n\t", F1: 32767, F2: 44}
    v2 = (v2 << ((4294967295) & 63))
    v2 = (v2 % (((v2 % (((v2 | (^(v2)))) | 1))) | 1))
    var v9 string = ""
    fmt.Println(v2)
    if (195 == 2) {
        var v10 uint64 = 18446744073709551615
        _ = v10
    } else {
        for fi11 := 0; fi11 < 16; fi11++ {
            v2 -= (^(v2))
            _ = fi11
        }
        var v12 int64 = -1
        _ = v12
    }
    var v13 float64 = 1.5
    fmt.Println(v13, v9, v13)
    for fi14 := 0; fi14 < 9; fi14++ {
        fmt.Println(fi14, fi14, fi14)
        _ = fi14
    }
    var v15 int32 = 992
    var v16 int16 = 103
    var v17 string = ("foo bar" + v9)
    v13 -= ((0.0 - (v13 / v13)) - v13)
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v13 -= v13
    v17 = (v9 + "\n\t")
    v16 += -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
