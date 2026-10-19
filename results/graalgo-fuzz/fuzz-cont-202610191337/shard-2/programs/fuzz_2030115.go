package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32 }
func genFn0() uint16 {
    var v7 int32 = 236
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 2147483647, F1: 2, F2: -1}
        var v10 uint = 18446744073709551615
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v7
    return 0
}

func main() {
    var v1 int16 = -620
    v2 := make(map[uint32]uint32)
    var v3 uint = 2
    var v4 uint = (v3 / (((v3 | (v3 % ((v3) | 1)))) | 1))
    v5 := make(map[int]uint32)
    var v6 float64 = -2.25
    v6 = (3.14159 - v6)
    fmt.Println(v1, v6)
    if false {
        v11 := GenS1{F0: 747, F1: 1, F2: v1}
        _ = v11
    }
    v12 := GenS2{F0: 42}
    if (false || true) {
        v4 += (v3 | 4294967295)
        v13 := []int{-1, 9223372036854775807, 2147483647, 163}
        _ = v13
    } else {
        if true {
            v14 := []int{32767, -249, 32768, 65}
            _ = len(v14)
            _ = v14
        } else {
            var v15 int8 = 42
            v1 = (^(v1))
            _ = v15
        }
        var v16 uint64 = 18446744073709551615
        _ = v16
    }
    var v17 int64 = 9223372036854775807
    if (-1 < v1) {
        v4 += 255
    }
    _ = len(v2)
    v18 := []int{42, -9223372036854775808, 32767, -149}
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 float32 = (((3.14159 * ((-2.25 * 1.5) * 0.0)) * 1e10) - 1.5)
        v2[4294967295] = 65535
        _ = fi19
        _ = v20
    }
    var v21 int16 = 651
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
