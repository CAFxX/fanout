package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int8 }
type GenS2 struct { F0 int64 }
func genFn0() string {
    var v5 bool = true
    var v6 int64 = -62
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]string)
    var v3 uint16 = 138
    v4 := []uint32{42, 42, 75, 65535}
    var v7 uint32 = 719
    v3 = 1
    if (true || ((0.1 + (1.5 - 0.1)) <= (float64(0.0) + 0.0))) {
        var v8 uint8 = uint8(v3)
        _ = v8
    } else {
        fmt.Println(v7, v3, v1)
    }
    var v9 string = "\n\t"
    var v10 float32 = 0.1
    v9 = v9
    v11 := GenS1{F0: 255, F1: false, F2: 15}
    v7 += (v7 ^ 281)
    if len(v4) > 0 {
        v4[0] = v7
    }
    var v12 uint8 = 0
    var v13 uint64 = 0
    if false {
        v10 = v10
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS2{F0: 517}
            _ = fi14
            _ = v15
        }
        fmt.Println(v12)
    }
    var v16 uint32 = (v7 % (((v7 | uint32(v3))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
