package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
func genFn0(p0 int, p1 uint64) int8 {
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(p0)
        p0 = (fi5 >> ((349) & 63))
        var v6 uint64 = (p1 / ((p1) | 1))
        _ = fi5
        _ = v6
    }
    fmt.Println(p1, p0)
    v7 := []int{(p0 ^ (p0 / ((42) | 1)))}
    _ = p0
    _ = p1
    _ = v7
    return 90
}

func main() {
    v1 := []int{127, 1, 2147483647}
    var v2 int8 = 116
    v3 := []int{32768}
    var v4 uint8 = 149
    var v8 int32 = -666
    v1 = append(v1, 32767)
    if true {
        var v9 uint8 = (v4 + (v4 << ((4294967295) & 63)))
        _ = v9
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            v8 += 0
            v8 -= int32(v4)
            _ = fi10
        }
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v2 -= int8(v8)
        _ = fi11
    }
    v12 := GenS1{F0: (v4 >> ((2) & 63))}
    v13 := GenS1{F0: 92}
    v14 := GenS2{F0: (^(v4)), F1: -87, F2: 255}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
