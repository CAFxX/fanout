package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8, p1 bool, p2 uint64) int32 {
    p2 -= (p2 >> ((2) & 63))
    p0 = (p0 - p0)
    v4 := []int{2147483647, 128, 128, 190}
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: p1, F1: 1, F2: 575}
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        var v7 uint = 2
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 582
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -2147483648
    var v3 bool = false
    v8 := GenS1{F0: (v2 == -2147483648), F1: -902, F2: 505}
    if (2 >= 65535) {
        v9 := GenS2{F0: 1}
        v10 := GenS1{F0: false, F1: 320, F2: 2}
        var v11 bool = (((float64((3.14159 * 0.0)) - 1e10) != 0.1) && (127 < 127))
        _ = v10
        _ = v11
        _ = v9
    } else {
        v3 = false
        v12 := []int{}
        v3 = (114 < 629)
        _ = v12
    }
    var v13 int8 = -128
    v14 := GenS2{F0: 65535}
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int = (fi15 | fi15)
        v17 := []int{v16, -177}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint16 = 193
    fmt.Println(v13)
    v2 += -641
    var v19 uint8 = 0
    var v20 int32 = -787
    fmt.Println(v3, v18, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
}
