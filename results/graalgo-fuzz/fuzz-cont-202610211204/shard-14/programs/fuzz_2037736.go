package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int16; F1 uint; F2 bool }
func genFn0(p0 uint16) int64 {
    switch -2147483648 {
    case -9223372036854775808:
        fmt.Println(p0)
    default:
        p0 = (p0 ^ p0)
    }
    if (!((p0 == (p0 >> ((1) & 63))) && true)) {
        var v6 int16 = 1
        _ = v6
    } else {
        var v7 uint8 = 1
        p0 = p0
        _ = v7
    }
    _ = p0
    return -457
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]uint32)
    var v3 uint8 = 128
    var v4 uint8 = (v3 << ((4294967295) & 63))
    v5 := []int{-675, 128, -335}
    v8 := GenS1{F0: (v1 / ((v1) | 1)), F1: genFn0(666)}
    v9 := GenS1{F0: v1, F1: -787}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v5 = append(v5, -255)
        for fi11 := 0; fi11 < 14; fi11++ {
            v4 = (v3 ^ 127)
            _ = fi11
        }
        _ = fi10
    }
    v3 += 127
    var v12 int16 = 32767
    fmt.Println(v1, v3)
    v13 := GenS1{F0: 4294967295, F1: -224}
    v14 := GenS1{F0: v1, F1: -9223372036854775808}
    v15 := GenS2{F0: -1, F1: 515, F2: (!(2 != 255))}
    v1 = (^(v1))
    for fi16 := 0; fi16 < 15; fi16++ {
        if true {
            var v17 uint16 = uint16(v3)
            _ = v17
        }
        for fi18 := 0; fi18 < 4; fi18++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi18
        }
        _ = fi16
    }
    var v19 float32 = 0.0
    v12 -= (v12 / ((v12) | 1))
    fmt.Println(v4, v4)
    var v20 int32 = 655
    var v21 int16 = v12
    var v22 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
