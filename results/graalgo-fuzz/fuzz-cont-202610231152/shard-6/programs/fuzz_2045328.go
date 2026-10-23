package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0() int64 {
    var v6 uint = 255
    fmt.Println(v6)
    _ = v6
    return 242
}

func main() {
    var v1 int8 = 1
    var v2 string = "a"
    var v3 int8 = (v1 | v1)
    v4 := make(map[uint32]uint32)
    var v5 int = 4294967295
    v3 = v1
    var v7 float32 = (1e10 * float32(1e10))
    v3 = -18
    fmt.Println(v3, v1, v1)
    _ = len(v4)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 bool = false
        var v10 int = (v5 % (((v5 | (^(fi8)))) | 1))
        v4[2147483648] = 65535
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int8 = -34
    v2 = ""
    v7 = float32(((float64((((1e10 * v7) / 1.5) - float32(float64(v7)))) * -2.25) + 0.0))
    v12 := GenS1{F0: 42, F1: v1}
    if (!((0 != 42) || ((2 < 65535) || (32767 >= 1000)))) {
        v13 := GenS1{F0: -1, F1: v1}
        _ = v13
    } else {
        v5 = 454
        var v14 uint32 = 2
        _ = v14
    }
    delete(v4, 65535)
    v15 := []int{(v5 + (-(v5))), (v5 ^ (v5 - v5)), (v5 % (((v5 - (-(v5)))) | 1)), (v5 ^ (v5 - -453))}
    var v16 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
