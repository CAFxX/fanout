package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint64 = 481
    var v2 int = 32767
    v3 := []string{"\n\t", "", "\n\t", ("" + "hello")}
    var v4 int = (v2 - int(v1))
    switch (v4 + (v4 - v4)) {
    case -746:
        for fi5 := 0; fi5 < 13; fi5++ {
            fmt.Println(v4, fi5, fi5)
            _ = fi5
        }
    default:
        fmt.Println(v1)
    }
    v4 += (^(v4))
    v6 := GenS1{F0: 42}
    var v7 uint32 = 4294967295
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 := GenS2{F0: -1, F1: 1}
    v9 := GenS3{F0: "hello", F1: v4}
    v10 := GenS4{F0: 0}
    _ = len(v3)
    fmt.Println(v2)
    if true {
        var v11 int = (v2 & (^(v4)))
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(v7)
            v4 -= (fi12 & (v2 / ((-9223372036854775808) | 1)))
            _ = fi12
        }
        _ = v11
    } else {
        fmt.Println(v7, v2, v4)
        v13 := GenS1{F0: 42}
        _ = v13
    }
    if (2 <= 2) {
        v14 := GenS5{F0: 32, F1: (v7 | v7)}
        _ = len(v3)
        var v15 int32 = 2147483647
        _ = v14
        _ = v15
    } else {
        var v16 uint = 2
        _ = v16
    }
    var v17 uint64 = v1
    v18 := GenS3{F0: "\n\t", F1: (v2 % (((v2 >> ((42) & 63))) | 1))}
    if (false || false) {
        fmt.Println(v2)
        var v19 bool = ((^(v7)) != v7)
        _ = v19
    } else {
        var v20 int32 = -1
        fmt.Println(v2, v1)
        _ = v20
    }
    var v21 float64 = 0.0
    var v22 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
