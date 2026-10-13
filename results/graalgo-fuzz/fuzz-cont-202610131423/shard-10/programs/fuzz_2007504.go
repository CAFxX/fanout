package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int16 = 1000
    var v2 string = "a"
    var v3 uint8 = 2
    v4 := []string{}
    var v5 int64 = -9223372036854775808
    if true {
        v6 := GenS1{F0: v1, F1: true}
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS2{F0: 4294967295, F1: v2, F2: v2}
            v2 = (v2 + (v2 + v2))
            _ = len(v4)
            fmt.Println(fi7, fi7)
            _ = fi8
            _ = v9
        }
        v4 = append(v4, v2)
        if (!true) {
            v10 := GenS3{F0: 4294967295}
            _ = v10
        }
        _ = fi7
    }
    fmt.Println(v5)
    for fi11 := 0; fi11 < 1; fi11++ {
        v1 = (-(v1))
        _ = fi11
    }
    var v12 uint32 = 487
    _ = len(v4)
    _ = len(v4)
    fmt.Println(v12, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
