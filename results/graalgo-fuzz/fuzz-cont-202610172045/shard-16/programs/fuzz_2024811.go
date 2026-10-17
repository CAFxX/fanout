package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint64) uint {
    v5 := GenS1{F0: 2, F1: 912}
    var v6 uint8 = 127
    p0 += 324
    p0 -= p0
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]string)
    var v3 int = 459
    var v4 uint16 = 255
    v3 += -125
    fmt.Println(v4, v4)
    v3 += -9223372036854775808
    var v7 uint32 = 2147483648
    var v8 uint8 = uint8(v1)
    v9 := []int{v3, v3}
    var v10 int8 = -1
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(v10, v1)
        v10 -= (v10 >> ((512) & 63))
        _ = fi11
    }
    v8 -= (^(v8))
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int64 = -1
        _ = fi12
        _ = v13
    }
    v4 -= uint16(v1)
    var v14 int32 = 883
    v8 += v8
    v10 = 77
    var v15 bool = true
    var v16 int64 = 42
    v4 += (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
