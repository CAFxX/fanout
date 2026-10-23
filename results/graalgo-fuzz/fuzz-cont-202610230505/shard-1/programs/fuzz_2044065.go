package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int; F2 int8 }
func genFn0(p0 int16) int64 {
    var v4 float64 = (0.1 / 1e10)
    p0 = p0
    if true {
        p0 += p0
        v5 := []int{42}
        _ = len(v5)
        _ = v5
    }
    var v6 uint8 = 41
    _ = p0
    _ = v4
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 float32 = 1e10
    var v2 int8 = 127
    var v3 bool = (132 != 2)
    v7 := []int{128}
    v2 = (v2 - v2)
    if ((^(v2)) >= (v2 ^ (v2 << ((0) & 63)))) {
        v8 := GenS1{F0: 666}
        v2 -= v2
        v3 = v3
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: 2}
        switch "" {
        case "\n\t":
            fmt.Println(fi9, v1, v1)
        case "a":
            v7 = append(v7, fi9)
        case "":
            v2 += v2
        }
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: ((490 > 970) || (!((true && ((v2 % (((v2 - v2)) | 1)) <= v2)) || true))), F1: 717, F2: -1}
    _ = len(v7)
    v1 = (v1 + 0.0)
    var v12 float64 = 3.14159
    var v13 uint8 = 0
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
}
