package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 bool }
func genFn0(p0 int32) int32 {
    var v6 int8 = 102
    var v7 float32 = 1.5
    p0 = (^(p0))
    _ = p0
    _ = v6
    _ = v7
    return (p0 % ((p0) | 1))
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = ((0.0 + (((0.0 + 0.1) / -2.25) + 0.0)) / 1.0)
    var v3 int8 = 0
    var v4 int32 = 0
    var v5 uint8 = 204
    v2 += v2
    v4 += v4
    v5 = v5
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 0, F1: 42, F2: (1000 < 259)}
        var v10 int8 = v3
        v4 += v4
        _ = fi8
        _ = v10
        _ = v9
    }
    v4 += 1
    var v11 uint16 = uint16(v5)
    v12 := []int{0, -651}
    fmt.Println(v3)
    var v13 int64 = -321
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if false {
        fmt.Println(v11)
        if (false && ("a" == "")) {
            var v14 int8 = v3
            _ = v14
        } else {
            v11 += (v11 % (((v11 << ((270) & 63))) | 1))
        }
    } else {
        var v15 uint64 = 1
        var v16 uint16 = (v11 ^ v11)
        _ = v15
        _ = v16
    }
    var v17 int16 = 777
    var v18 uint64 = 4294967295
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
