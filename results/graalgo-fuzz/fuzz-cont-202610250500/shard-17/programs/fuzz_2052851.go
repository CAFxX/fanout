package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
func genFn0(p0 int32, p1 bool, p2 int) uint {
    v5 := []int{9223372036854775807, -9223372036854775808, p2}
    switch "hello" {
    case "hello":
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: p0, F1: -651}
            _ = fi6
            _ = v7
        }
    case "":
        if len(v5) > 0 {
            v5[0] = -448
        }
    case "a":
        var v8 int16 = -552
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 int32) int {
    p0 = p0
    if ((2147483648 >= 1) || false) {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint8 = 8
            _ = fi9
            _ = v10
        }
    } else {
        fmt.Println(p0)
    }
    _ = p0
    return 127
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]uint32)
    var v3 int16 = -32768
    var v4 int16 = (-(v3))
    var v11 bool = true
    fmt.Println(v1, v1, v11)
    v12 := GenS1{F0: -1, F1: genFn1(-1)}
    var v13 float64 = 1.5
    var v14 int64 = -603
    var v15 float32 = 1.5
    v4 = v3
    var v16 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
