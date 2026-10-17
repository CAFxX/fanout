package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := []int{562, 42, 2147483647, -288}
    var v4 uint = 306
    var v5 uint16 = 65535
    v3 = append(v3, 88)
    v4 -= 0
    fmt.Println(v5, v4)
    v5 += v5
    var v6 int = 475
    var v7 uint8 = 177
    v8 := GenS1{F0: int16(v6), F1: 0, F2: 950}
    if true {
        for fi9 := 0; fi9 < 16; fi9++ {
            v7 += (^(v7))
            v4 = 66
            _ = fi9
        }
        var v10 uint8 = v7
        _ = v10
    } else {
        var v11 string = "\n\t"
        var v12 int8 = 127
        _ = v11
        _ = v12
    }
    var v13 uint = 4294967295
    _ = len(v1)
    var v14 int16 = 252
    var v15 int8 = -128
    fmt.Println(v6)
    v5 = (v5 - (v5 ^ (v5 ^ 1)))
    _ = len(v3)
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 int = 128
        var v18 string = ""
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
