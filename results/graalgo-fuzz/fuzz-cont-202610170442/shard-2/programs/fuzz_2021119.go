package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0(p0 bool) bool {
    var v7 bool = p0
    var v8 int16 = 711
    p0 = (42 != 4294967295)
    var v9 uint32 = 376
    switch 4294967295 {
    case -806:
        v9 += (v9 ^ v9)
    case 1:
        fmt.Println(v9, v9, v7)
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 int8 = 0
    v2 := []int{52, 42, 585, 127}
    v3 := make(map[uint32]uint32)
    var v4 int32 = -195
    v5 := make(map[string]int)
    v6 := make(map[string]string)
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            _ = len(v5)
            if (false && false) {
                v4 += (v4 | -1)
                var v12 uint32 = 2147483647
                v13 := GenS1{F0: 1, F1: 1, F2: 23}
                _ = v12
                _ = v13
            }
            _ = fi11
        }
        v4 = (v4 ^ v4)
        _ = fi10
    }
    v1 = v1
    var v14 bool = false
    delete(v6, ("" + ("" + "")))
    if ((-1 > 893) || v14) {
        if (719 > 1) {
            v15 := GenS1{F0: 0, F1: 941, F2: 65535}
            _ = v15
        }
        var v16 float64 = -2.25
        v14 = v14
        _ = v16
    } else {
        v2 = append(v2, 4294967295)
    }
    var v17 int8 = v1
    v2 = append(v2, 128)
    var v18 uint64 = 4294967295
    var v19 int8 = v1
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint16 = 32767
        var v22 uint64 = v18
        v17 = int8(fi20)
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(v18)
    var v23 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
