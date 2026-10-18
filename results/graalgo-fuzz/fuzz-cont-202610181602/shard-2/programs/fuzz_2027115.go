package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }

func main() {
    var v1 int64 = 467
    v2 := make(map[string]int)
    v3 := []int{1}
    var v4 int16 = 0
    _ = len(v3)
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(v4)
        v1 = v1
        v1 -= v1
        _ = fi5
    }
    v1 = -915
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            delete(v2, "a")
            _ = fi6
        }
        v1 = (v1 % ((v1) | 1))
    } else {
        v7 := GenS1{F0: v4, F1: (v4 - (v4 * (v4 / (((v4 << ((515) & 63))) | 1))))}
        var v8 int8 = -46
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint = 541
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v2, "")
    var v11 uint32 = 65535
    var v12 uint32 = (v11 & (v11 + (v11 + v11)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
