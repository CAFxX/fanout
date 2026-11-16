package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 string; F1 uint; F2 int }

func main() {
    var v1 int8 = 0
    v2 := []string{("foo bar" + "\n\t")}
    var v3 uint32 = 0
    var v4 uint8 = 183
    var v5 float64 = -2.25
    for fi6 := 0; fi6 < 2; fi6++ {
        v1 -= (^(v1))
        var v7 uint = 1
        _ = fi6
        _ = v7
    }
    if ((^(v3)) > (v3 / ((65535) | 1))) {
        v8 := GenS1{F0: false}
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS2{F0: 655, F1: -1}
            v1 = v1
            fmt.Println(v5, fi9)
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    v4 -= v4
    v11 := GenS3{F0: 2}
    v5 = 1e10
    if (-287 != -602) {
        var v12 uint32 = 0
        v13 := GenS1{F0: false}
        _ = v12
        _ = v13
    } else {
        v1 = v1
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v14 := GenS4{F0: (^(v4)), F1: 255}
    v15 := GenS5{F0: "hello", F1: 0, F2: 799}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
