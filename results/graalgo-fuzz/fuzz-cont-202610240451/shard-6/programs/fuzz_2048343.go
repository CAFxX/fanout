package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
func genFn0() int32 {
    var v5 uint = 631
    var v6 float64 = (3.14159 / 1.0)
    _ = v5
    _ = v6
    return 972
}

func main() {
    var v1 uint16 = 32767
    v2 := []int{}
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    var v7 uint32 = 1
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int16 = 1000
        var v10 uint64 = 4294967295
        fmt.Println(v7, fi8)
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = v1
    v11 := GenS1{F0: int64(v7), F1: "\n\t"}
    v7 -= v7
    var v12 int64 = -9223372036854775808
    v1 = (v1 << ((931) & 63))
    v13 := GenS1{F0: (v12 | v12), F1: "hello"}
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
