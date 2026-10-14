package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    var v3 uint16 = 303
    if false {
        fmt.Println(v3)
    } else {
        _ = len(v2)
        v4 := GenS1{F0: 1, F1: 1}
        _ = v4
    }
    v3 -= (v3 | v3)
    v5 := GenS1{F0: 1, F1: 168}
    v6 := GenS2{F0: 33}
    for fi7 := 0; fi7 < 11; fi7++ {
        if (-1 > 0) {
            var v8 int = fi7
            var v9 bool = (true || (884 < 4294967295))
            _ = v8
            _ = v9
        } else {
            fmt.Println(v3)
        }
        switch "\n\t" {
        case "a":
            v10 := GenS3{F0: (("hello" + "") + "\n\t"), F1: 0, F2: -1}
            _ = v10
        case "foo bar":
            var v11 uint = 42
            _ = v11
        default:
            v12 := GenS2{F0: -2147483648}
            _ = v12
        }
        _ = fi7
    }
    v3 = (v3 + (^(v3)))
    var v13 uint32 = 0
    var v14 float64 = 1e10
    var v15 uint = 4294967295
    var v16 int = -1
    var v17 uint64 = 10
    var v18 bool = true
    var v19 float32 = 3.14159
    v18 = v18
    v20 := []int{32767, (v16 << (((v15 % (((v15 << ((v15) & 63))) | 1))) & 63)), (-(v16))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v5
    _ = v6
}
