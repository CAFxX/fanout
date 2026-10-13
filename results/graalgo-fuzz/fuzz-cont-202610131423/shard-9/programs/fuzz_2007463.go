package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16, p1 int8, p2 int32) int16 {
    v4 := GenS1{F0: 1, F1: 765, F2: p1}
    p0 += -577
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int32 = p2
        _ = fi5
        _ = v6
    }
    v7 := []int{0, 32767}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 1000
}

func main() {
    var v1 uint = 561
    v2 := []uint32{}
    var v3 uint8 = 255
    var v8 int16 = -1
    v9 := GenS2{F0: v3}
    var v10 int16 = (v8 << ((v1) & 63))
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v11 uint64 = 4294967295
        _ = v11
    }
    fmt.Println(v8, v8, v3)
    v10 -= (v10 / (((^(v8))) | 1))
    switch 9223372036854775807 {
    case 127:
        v3 = v3
    default:
        switch 2147483647 {
        case 2147483647:
            fmt.Println(v3, v10, v1)
        }
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int64 = -1
        _ = fi12
        _ = v13
    }
    var v14 float32 = (1.5 - (0.0 + (3.14159 / 0.1)))
    v15 := GenS1{F0: 4294967295, F1: 877, F2: 58}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
