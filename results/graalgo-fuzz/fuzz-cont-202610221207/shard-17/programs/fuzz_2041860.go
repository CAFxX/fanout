package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
func genFn0(p0 int16) uint64 {
    if false {
        var v5 int64 = -793
        _ = v5
    } else {
        v6 := GenS1{F0: 4294967295, F1: 2147483647}
        p0 = (p0 & -1)
        _ = v6
    }
    if false {
        p0 += (p0 - p0)
    }
    v7 := []int{}
    for fi8 := 0; fi8 < 13; fi8++ {
        if true {
            var v9 uint16 = 1
            _ = v9
        }
        if len(v7) > 0 {
            v7[0] = fi8
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 702
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{0, 2147483647, 0}
    var v3 float32 = ((-2.25 - 1e10) + 1e10)
    var v4 int32 = 42
    fmt.Println(v3, v4)
    if (!false) {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 uint = 144
            _ = fi10
            _ = v11
        }
        var v12 bool = true
        var v13 uint16 = 283
        _ = v12
        _ = v13
    } else {
        v4 = -2147483648
    }
    v14 := GenS1{F0: 42, F1: -882}
    var v15 int8 = -1
    v16 := GenS1{F0: 674, F1: 9223372036854775807}
    fmt.Println(v15, v15)
    v17 := v1["foo bar"]
    var v18 uint32 = 328
    var v19 int16 = -1
    var v20 int32 = 111
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
