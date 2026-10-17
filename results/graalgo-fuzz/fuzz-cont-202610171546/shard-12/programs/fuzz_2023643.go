package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 868
    var v3 uint16 = 2
    var v4 float32 = (3.14159 - float32(0.0))
    v5 := make(map[int]int)
    var v6 uint64 = v1
    var v7 uint8 = 1
    if true {
        v7 += v7
        var v8 int8 = -128
        _ = v8
    } else {
        var v9 int8 = 100
        v9 = 127
        _ = v9
    }
    var v10 float32 = (v4 - v4)
    var v11 float64 = 0.0
    if (-12 < 1) {
        v2 -= v2
        var v12 string = "a"
        _ = v12
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint64 = 0
            _ = fi13
            _ = v14
        }
        if (42 >= -32768) {
            v15 := GenS1{F0: -373, F1: (v2 & 757)}
            _ = v15
        }
    }
    if true {
        fmt.Println(v11, v2)
        v2 -= v2
        v16 := []int{-2147483648, -9223372036854775808, -339}
        _ = v16
    } else {
        delete(v5, 32767)
    }
    var v17 uint32 = 775
    for fi18 := 0; fi18 < 15; fi18++ {
        v10 = float32(((float64(((3.14159 * float32(v11)) + v10)) * (v11 + v11)) - 1e10))
        v11 += 0.0
        delete(v5, fi18)
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
