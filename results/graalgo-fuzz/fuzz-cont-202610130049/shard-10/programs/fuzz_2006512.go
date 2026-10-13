package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int8, p1 uint) uint {
    var v6 int16 = 32767
    v6 += (v6 - (v6 | v6))
    var v7 int64 = 528
    p1 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 ^ p1)
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 42
    var v3 uint32 = v2
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    var v8 uint64 = 2
    for fi9 := 0; fi9 < 4; fi9++ {
        v8 -= v8
        fmt.Println(v3, fi9, fi9)
        _ = fi9
    }
    if (false && false) {
        _ = len(v4)
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: 42, F1: (v8 & (^(v8)))}
            v1 = 255
            fmt.Println(fi10, v8)
            _ = fi10
            _ = v11
        }
        var v12 int8 = 1
        v3 = 4294967295
        _ = v12
    }
    v13 := []int{}
    switch "a" {
    case "a":
        v14 := v4["hello"]
        _ = v14
    case "hello":
        if (32 == 42) {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
    }
    v8 = (v8 ^ uint64(v1))
    if (1 != -9223372036854775808) {
        fmt.Println(v1, v8)
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                fmt.Println(v3, v8, v3)
                var v17 int64 = 9223372036854775807
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        v18 := GenS1{F0: 2, F1: 246}
        _ = v18
    }
    v3 += v3
    var v19 int = -434
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
