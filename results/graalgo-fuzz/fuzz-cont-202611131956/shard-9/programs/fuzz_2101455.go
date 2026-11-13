package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
func genFn0(p0 uint32, p1 uint32) int32 {
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: -1, F1: p0, F2: 461}
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(p0, v7, v7)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 4294967295
    var v3 uint64 = (v2 | (^(v2)))
    v4 := []int{32767, 953, 42, 2147483647}
    var v5 bool = true
    var v6 string = "a"
    if false {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 float32 = 0.0
            _ = fi10
            _ = v11
        }
        v6 = (v6 + (((v6 + v6) + "a") + v6))
        var v12 int16 = -427
        _ = v12
    }
    var v13 uint64 = v3
    var v14 uint32 = 2147483647
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int16 = 32767
        _ = fi15
        _ = v16
    }
    var v17 bool = v5
    v18 := GenS1{F0: 0, F1: (v14 - 637), F2: 42}
    v19 := GenS1{F0: 9223372036854775807, F1: v14, F2: 921}
    fmt.Println(v5, v14)
    switch v6 {
    case "hello":
        var v20 int16 = 42
        _ = v20
    case "\n\t":
        fmt.Println(v5)
    case "":
        v21 := GenS2{F0: 0, F1: 4294967295, F2: 32767}
        _ = v21
    }
    v2 -= 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
