package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int, p1 int64) uint64 {
    v6 := []int{860}
    var v7 uint16 = 1
    var v8 int8 = -91
    var v9 int16 = 32767
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]uint32)
    var v3 float32 = 1e10
    var v4 int64 = -9223372036854775808
    v5 := make(map[string]string)
    switch 2147483647 {
    case 0:
        fmt.Println(v4, v1, v3)
    }
    v10 := []int{}
    switch 128 {
    case 32768:
        var v11 int64 = (v4 % (((v4 * (v4 + (v4 << ((45) & 63))))) | 1))
        _ = v11
    case -1:
        _ = len(v10)
    default:
        fmt.Println(v1, v3)
    }
    if false {
        var v12 int32 = 2147483647
        v3 -= ((v3 * v3) * float32(((float64(v3) / 1e10) - (0.1 - 0.0))))
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 int64 = v4
        v15 := GenS1{F0: 4294967295, F1: v1}
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 18446744073709551615, F1: v1}
    fmt.Println(v3, v4)
    if true {
        if len(v10) > 0 {
            v10[0] = 698
        }
    } else {
        var v17 uint = 439
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
