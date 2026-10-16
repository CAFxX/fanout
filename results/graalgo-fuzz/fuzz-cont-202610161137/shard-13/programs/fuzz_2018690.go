package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint64, p1 int16) uint {
    var v6 int8 = 1
    var v7 int64 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 518
}

func main() {
    v1 := []int{-776}
    v2 := []int{}
    var v3 uint = 125
    var v4 uint = 0
    var v5 int32 = 726
    var v8 int = 594
    _ = len(v1)
    fmt.Println(v4)
    var v9 float64 = (3.14159 / 1.0)
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 bool = (1 <= -317)
        _ = fi10
        _ = v11
    }
    v9 -= v9
    switch 42 {
    case -2147483648:
        var v12 int32 = (v5 | v5)
        _ = v12
    default:
        var v13 int32 = v5
        _ = v13
    }
    v14 := GenS1{F0: (-1 <= 125), F1: 4294967295, F2: 4294967295}
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(fi15)
        switch "a" {
        case "":
            var v16 bool = (100 == 100)
            _ = v16
        case "a":
            v5 += (v5 / (((v5 | (v5 % (((v5 << (((^(v4))) & 63))) | 1)))) | 1))
        case "\n\t":
            fmt.Println(v5, v8, v9)
        default:
            fmt.Println(v4, v3, fi15)
        }
        _ = fi15
    }
    if false {
        switch "" {
        case "":
            v1 = append(v1, v8)
        case "hello":
            v17 := GenS2{F0: 4294967295, F1: 36}
            _ = v17
        case "\n\t":
            var v18 bool = ((3.14159 == float32(v9)) || (true || true))
            _ = v18
        }
    } else {
        var v19 uint64 = 1
        _ = v19
    }
    var v20 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
