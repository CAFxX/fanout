package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]bool)
    var v3 float64 = v1
    v4 := make(map[string]uint32)
    var v5 uint16 = 32767
    if (-1 <= 0) {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 float64 = v1
            _ = fi6
            _ = v7
        }
        _ = len(v2)
    } else {
        fmt.Println(v1)
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(v5, fi8)
            var v10 float32 = (1.5 / 1e10)
            _ = fi9
            _ = v10
        }
        v11 := []int{-2147483648, fi8, 32768}
        fmt.Println(fi8, v5)
        _ = fi8
        _ = v11
    }
    v3 = 0.1
    v12 := GenS1{F0: 4294967295}
    switch 2147483647 {
    case 4294967295:
        v13 := GenS1{F0: -339}
        _ = v13
    case 2147483647:
        var v14 float32 = (3.14159 / ((-2.25 - 0.1) * (-2.25 - float32(1e10))))
        _ = v14
    }
    var v15 uint16 = v5
    fmt.Println(v1, v5, v15)
    var v16 int64 = -9223372036854775808
    v17 := GenS1{F0: 0}
    v3 += (v3 / v1)
    var v18 int64 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
