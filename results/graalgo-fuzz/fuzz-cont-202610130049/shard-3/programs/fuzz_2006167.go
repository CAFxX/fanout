package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 int) int32 {
    v5 := GenS1{F0: 807, F1: (^(p0))}
    fmt.Println(p0)
    v6 := []int{(^(p0)), (p0 - (p0 & 1))}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 float32 = -2.25
        v9 := GenS2{F0: -253}
        fmt.Println(v8)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]uint32)
    var v3 float64 = (1e10 / (-2.25 + 1.5))
    var v4 float32 = 1.5
    switch 127 {
    case -167:
        fmt.Println(v4)
    case 42:
        for fi10 := 0; fi10 < 15; fi10++ {
            v3 -= v3
            _ = fi10
        }
    }
    v4 -= v4
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := v2[2147483647]
        _ = fi11
        _ = v12
    }
    v3 += (v3 * float64((3.14159 + float32(3.14159))))
    if true {
        var v13 int32 = 42
        _ = v13
    } else {
        v14 := []int{32767, -9223372036854775808, 128}
        _ = v14
    }
    v1 -= 1
    var v15 uint64 = v1
    v15 -= 0
    fmt.Println(v3)
    var v16 int16 = 184
    if false {
        var v17 uint8 = 128
        if true {
            v18 := GenS1{F0: 2, F1: -1}
            _ = v18
        } else {
            var v19 int = 4294967295
            _ = v19
        }
        _ = v17
    }
    v3 -= 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
