package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -882
    var v3 uint = 4294967295
    var v4 int32 = -942
    var v5 float64 = 0.0
    fmt.Println(v3, v3)
    v6 := []int{42, int(v4)}
    var v7 int = 653
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint64 = 2
        v4 += 2147483647
        v10 := GenS1{F0: v2, F1: 0}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 bool = (504 == 747)
        _ = fi11
        _ = v12
    }
    var v13 int16 = 90
    var v14 int16 = 510
    v4 -= 1
    v13 = int16(v4)
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 uint = 824
        fmt.Println(v2, v13, fi15)
        _ = fi15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
