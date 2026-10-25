package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int64 }

func main() {
    var v1 uint = 121
    var v2 float32 = 0.1
    v3 := []int{11, 1, -750}
    v2 += (v2 - (3.14159 / v2))
    v2 = 1.5
    for fi4 := 0; fi4 < 13; fi4++ {
        v2 += v2
        _ = fi4
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS1{F0: -128}
        var v7 int32 = 439
        v7 -= (v7 % (((v7 | -2147483648)) | 1))
        _ = fi5
        _ = v6
        _ = v7
    }
    v1 = v1
    if false {
        for fi8 := 0; fi8 < 10; fi8++ {
            v3 = append(v3, (^(fi8)))
            _ = fi8
        }
    } else {
        v9 := GenS2{F0: 1, F1: 42, F2: 9223372036854775807}
        _ = v9
    }
    var v10 float64 = (0.0 * 3.14159)
    switch "hello" {
    case "":
        v2 += v2
    case "foo bar":
        var v11 uint64 = 18446744073709551615
        _ = v11
    }
    v12 := GenS2{F0: 4294967295, F1: 0, F2: 370}
    v2 = (1e10 + v2)
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
}
