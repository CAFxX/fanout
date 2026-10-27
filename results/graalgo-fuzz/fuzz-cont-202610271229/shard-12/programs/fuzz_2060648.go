package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
func genFn0(p0 int) int32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint32 = 2147483647
        var v9 uint8 = 197
        fmt.Println(v9)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: true, F1: 2}
    var v11 uint = 65535
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 int16 = -367
    var v5 int64 = 611
    var v6 int16 = (^(v4))
    v12 := []int{-9223372036854775808, 128, 128, 4294967295}
    v5 = (v5 - v5)
    var v13 uint8 = 103
    fmt.Println(v6)
    for fi14 := 0; fi14 < 7; fi14++ {
        v6 = (v4 ^ 0)
        _ = fi14
    }
    v15 := GenS1{F0: false, F1: 42}
    var v16 uint16 = 1
    v16 = (v16 ^ 583)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
