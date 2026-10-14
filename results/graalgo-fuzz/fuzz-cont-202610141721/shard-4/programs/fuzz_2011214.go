package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint64 }
func genFn0(p0 int64) uint {
    p0 -= -223
    v5 := []int{-1, -292, 32767}
    _ = p0
    _ = v5
    return 4294967295
}

func main() {
    v1 := []string{}
    var v2 bool = true
    var v3 float64 = float64(float32(0.1))
    var v4 int32 = -786
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int64 = 1
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(v2, v4, v3)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    v3 = ((v3 - v3) - 1e10)
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v9 int64 = -693
    switch 0 {
    case 32768:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 1:
        v10 := GenS1{F0: 2}
        _ = v10
    case -630:
        fmt.Println(v9, v2, v3)
    }
    fmt.Println(v4, v2, v4)
    if true {
        _ = len(v1)
    }
    v11 := GenS1{F0: 2}
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v12 int16 = -594
        _ = v12
    case "a":
        if (false && (false || v2)) {
            v4 = v4
        }
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v9 -= (v9 * (v9 * v9))
        if v2 {
            fmt.Println(v9, v3, fi13)
        }
        _ = fi13
    }
    switch 2147483647 {
    case -678:
        v14 := GenS2{F0: 1, F1: -32768, F2: 956}
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
