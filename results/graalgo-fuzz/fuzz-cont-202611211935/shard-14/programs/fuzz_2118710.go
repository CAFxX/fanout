package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
type GenS2 struct { F0 int16 }
func genFn0() int64 {
    var v4 uint64 = 752
    v5 := []int{4294967295}
    var v6 int64 = 1
    var v7 int = 2147483647
    fmt.Println(v6)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := []uint32{193, 1, 2147483648}
    var v2 uint32 = 2
    var v3 int16 = 167
    var v8 int32 = -7
    if (128 < 128) {
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: fi9, F1: 0, F2: v8}
        fmt.Println(v8, fi9, v2)
        _ = fi9
        _ = v10
    }
    fmt.Println(v2, v2)
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int8 = 127
        v3 -= (v3 ^ -32768)
        _ = fi11
        _ = v12
    }
    if false {
        v13 := GenS2{F0: (v3 * (v3 - 42))}
        v3 = (^(v3))
        var v14 int8 = -1
        _ = v13
        _ = v14
    } else {
        v8 = (-(v8))
        switch 42 {
        case 971:
            v3 = v3
        case 4294967295:
            if true {
                _ = len(v1)
            } else {
                fmt.Println(v8, v8, v8)
            }
        }
    }
    var v15 int8 = 127
    var v16 float64 = (-2.25 / -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
