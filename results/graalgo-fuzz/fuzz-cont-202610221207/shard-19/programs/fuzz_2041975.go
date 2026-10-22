package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }

func main() {
    var v1 int8 = -3
    var v2 int16 = 0
    var v3 int64 = 319
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{}
        fmt.Println(v3, fi6, v3)
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    switch "" {
    case "":
        var v9 uint = 255
        _ = v9
    case "foo bar":
        var v10 float64 = 0.1
        _ = v10
    default:
        if (18446744073709551615 >= 465) {
            var v11 int16 = int16(v3)
            v11 = 533
            var v12 int64 = v3
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                v3 = -9223372036854775808
                delete(v5, 13)
                _ = fi13
            }
            v3 = (v3 ^ (v3 | 9223372036854775807))
        }
    }
    if false {
        if (((336 != -182) && (!(((196 >= 1) && true) || ((false || false) && true)))) && (!false)) {
            _ = len(v8)
            var v14 uint32 = 1
            v2 += 0
            _ = v14
        } else {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 uint = 860
                _ = fi15
                _ = v16
            }
        }
        v17 := GenS1{F0: 42, F1: v1}
        _ = v17
    } else {
        var v18 int16 = -32768
        v3 = (v3 & 42)
        var v19 int8 = (v1 * v1)
        _ = v18
        _ = v19
    }
    fmt.Println(v1, v3, v1)
    var v20 uint = 541
    var v21 int64 = 9223372036854775807
    if ((0 < 42) || (!(((1.5 + 0.0) < 1.5) || true))) {
        v22 := GenS1{F0: (v20 >> ((v20) & 63)), F1: v1}
        v23 := GenS1{F0: v20, F1: 1}
        switch 0 {
        case -1:
            var v24 uint = 1
            _ = v24
        }
        _ = v22
        _ = v23
    } else {
        var v25 int8 = v1
        _ = v25
    }
    v1 -= -1
    v3 -= 1
    var v26 bool = true
    v21 += v3
    v27 := v4["foo bar"]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
